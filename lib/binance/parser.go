package binance

import (
	"fmt"

	"github.com/ImPedro29/exchange-sdk/models"
)

const (
	Trading = "TRADING"
)

func ParseStatus(statusForParse string) models.PairStatus {
	switch statusForParse {
	case Trading:
		return models.Enabled
	default:
		return models.Disabled
	}
}

func ParsePair(pairForParse symbol) (string, *models.Pair) {
	key := fmt.Sprintf("%s%s", pairForParse.BaseAsset, pairForParse.QuoteAsset)

	return key, &models.Pair{
		Status: ParseStatus(pairForParse.Status),
		Base: models.Asset{
			Symbol:    pairForParse.BaseAsset,
			Precision: int64(pairForParse.BaseAssetPrecision),
		},
		Quote: models.Asset{
			Symbol:    pairForParse.QuoteAsset,
			Precision: int64(pairForParse.QuoteAssetPrecision),
		},
	}
}
