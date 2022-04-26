package mexc

import (
	"fmt"
	"strings"

	"github.com/ImPedro29/exchange-sdk/models"
)

const (
	Enabled = "ENABLED"
)

func ParseStatus(statusForParse string) models.PairStatus {
	switch statusForParse {
	case Enabled:
		return models.Enabled
	default:
		return models.Disabled
	}
}

func ParsePair(pairForParse data) (string, *models.Pair) {
	symbolSplit := strings.Split(pairForParse.Symbol, "_")
	var quoteSymbol string
	if len(symbolSplit) > 1 {
		quoteSymbol = symbolSplit[1]
	}

	key := fmt.Sprintf("%s%s", pairForParse.VcoinName, quoteSymbol)

	return key, &models.Pair{
		Status: ParseStatus(pairForParse.State),
		Base: models.Asset{
			Symbol:    pairForParse.VcoinName,
			Precision: pairForParse.QuantityScale,
		},
		Quote: models.Asset{
			Symbol:    quoteSymbol,
			Precision: pairForParse.PriceScale,
		},
	}
}
