package kucoin

import (
	"fmt"

	"github.com/ImPedro29/exchange-sdk/common"
	"github.com/ImPedro29/exchange-sdk/constraints"
	"github.com/ImPedro29/exchange-sdk/interfaces"
	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/ImPedro29/exchange-sdk/utils"
)

func NewKucoin(apiKey, secret string) interfaces.Exchange {
	return &kucoin{
		Api:    fmt.Sprintf("https://%s", constraints.KuCoin),
		Secret: secret,
		Key:    apiKey,
		events: &kucoinEvents{
			Api:      fmt.Sprintf("wss://%s", constraints.KuCoin),
			handlers: make(map[string]models.EventHandler),
			close:    make(chan bool),
		},
	}
}

func (s *kucoin) GetPairs() (map[string]models.Pair, error) {
	var supportedPairsRes supportedPairsResponse
	if err := utils.GetURL(fmt.Sprintf("%s/api/v1/symbols", s.Api), &supportedPairsRes, nil); err != nil {
		return nil, err
	}

	pairs := make(map[string]models.Pair)
	for _, pair := range supportedPairsRes.Data {
		key, parsed := ParsePair(pair)
		pairs[key] = *parsed
	}

	return pairs, nil
}

func (s *kucoin) GetMarket() (map[string]models.MarketAsset, error) {
	var supportedPairsRes marketResponse
	if err := utils.GetURL(fmt.Sprintf("%s/api/v1/market/allTickers", s.Api), &supportedPairsRes, nil); err != nil {
		return nil, err
	}

	if len(supportedPairsRes.Data.Ticker) < 1 {
		return nil, common.ErrReturnedLen0
	}

	market := make(map[string]models.MarketAsset)
	for _, marketPair := range supportedPairsRes.Data.Ticker {
		marketPairParsed, err := ParseMarket(marketPair)
		if err != nil {
			return nil, err
		}
		market[marketPairParsed.Symbol] = *marketPairParsed
	}

	return market, nil
}
