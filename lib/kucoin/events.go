package kucoin

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/ImPedro29/exchange-sdk/common"
	"github.com/ImPedro29/exchange-sdk/constraints"
	"github.com/ImPedro29/exchange-sdk/interfaces"
	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/ImPedro29/exchange-sdk/utils"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (s *kucoin) Events() (interfaces.Events, error) {
	if s.events.Conn != nil {
		return s.events, nil
	}

	url := fmt.Sprintf("%s/api/v1/bullet-public", s.Api)
	var response websocketStartResponse
	if err := utils.PostURL(url, nil, &response, nil); err != nil {
		return nil, err
	}

	if len(response.Data.InstanceServers) < 1 {
		return nil, common.ErrReturnedLen0
	}

	s.events.Token = response.Data.Token
	s.events.Api = response.Data.InstanceServers[0].Endpoint

	ctx, cancel := context.WithTimeout(context.Background(), constraints.Timeout)
	defer cancel()
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	url = fmt.Sprintf("%s?token=%s&connectId=%s", s.events.Api, s.events.Token, id.String())
	c, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	var wsResponse websocketWelcomeResponse
	if err := c.ReadJSON(&wsResponse); err != nil {
		return nil, err
	}

	if id.String() != wsResponse.Id {
		return nil, common.ErrIdReturnedWrong
	}

	s.events.Conn = c
	go s.events.handler()

	return s.events, nil
}

func (s *kucoinEvents) handler() {
	for {
		select {
		case <-s.close:
			return
		default:
			var object websocketResponse
			if err := s.Conn.ReadJSON(&object); err != nil {
				zap.L().Error("failed trying to read json from kucoin server", zap.Error(err))
				continue
			}

			if object.Topic == "" {
				zap.L().Debug("server not returned topic", zap.Any("object", object))
				continue
			}

			function, ok := s.handlers[object.Topic]
			if ok {
				err := function(object.Data)
				if err != nil {
					zap.L().Error("failed trying to handle function", zap.Error(err))
					continue
				}
			}
		}
	}
}

func (s *kucoinEvents) Listen(event models.EventType, handler models.EventHandler) error {
	switch event {
	case constraints.MarketAll:
		return s.MarketAll(handler)
	}

	return common.ErrNotSupported
}

func (s *kucoinEvents) MarketAll(handler models.EventHandler) error {
	topic := "/market/ticker:all"
	if _, ok := s.handlers[topic]; ok {
		s.handlers[topic] = handler
		zap.L().Warn("marketAll handler registered already, now replaced...")
		return nil
	}

	request := websocketRequest{
		Id:       rand.Int63(),
		Type:     "subscribe",
		Topic:    topic,
		Response: true,
	}

	if err := s.Conn.WriteJSON(request); err != nil {
		return err
	}

	s.handlers[topic] = handler
	return nil
}

func (s *kucoinEvents) Close() error {
	if s.Conn != nil {
		if err := s.Conn.Close(); err != nil {
			return err
		}
		s.Conn = nil
		s.close <- true
	}

	return common.ErrConnectionClosed
}
