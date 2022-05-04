package interfaces

import "github.com/ImPedro29/exchange-sdk/models"

type Exchange interface {
	GetPairs() (map[string]models.Pair, error)
	DepositAddress(pair models.Asset) (string, error)
	GetMarket() (map[string]models.MarketAsset, error)
}
