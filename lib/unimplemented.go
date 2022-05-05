package lib

import (
	"github.com/ImPedro29/exchange-sdk/common"
	"github.com/ImPedro29/exchange-sdk/interfaces"
	"github.com/ImPedro29/exchange-sdk/models"
)

var _ interfaces.Exchange = (*UnimplementedExchange)(nil)

type UnimplementedExchange struct{}

func (u UnimplementedExchange) GetPairs() (map[string]models.Pair, error) {
	return nil, common.ErrNotSupported
}

func (u UnimplementedExchange) DepositAddress(_ models.Asset) (string, error) {
	return "", common.ErrNotSupported
}

func (u UnimplementedExchange) GetMarket() (map[string]models.MarketAsset, error) {
	return nil, common.ErrNotSupported
}

func (u UnimplementedExchange) Events() (interfaces.Events, error) {
	return nil, common.ErrNotSupported
}
