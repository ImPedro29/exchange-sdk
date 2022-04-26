package kucoin

import (
	"fmt"
	"strings"

	"github.com/ImPedro29/exchange-sdk/models"
)

const (
	Trading = "TRADING"
)

func ParseStatus(isEnabled bool) models.PairStatus {
	if isEnabled {
		return models.Enabled
	}

	return models.Disabled
}

func ParsePair(pairForParse data) (string, *models.Pair) {
	key := fmt.Sprintf("%s%s", pairForParse.BaseCurrency, pairForParse.QuoteCurrency)

	minSizeSplitBase := strings.Split(pairForParse.BaseIncrement, ".")
	minSizeSplitQuote := strings.Split(pairForParse.QuoteIncrement, ".")
	var basePrecision, quotePrecision int64
	if len(minSizeSplitBase) > 1 {
		basePrecision = int64(len(minSizeSplitBase[1]))
	}
	if len(minSizeSplitQuote) > 1 {
		quotePrecision = int64(len(minSizeSplitQuote[1]))
	}

	return key, &models.Pair{
		Status: ParseStatus(pairForParse.EnableTrading),
		Base: models.Asset{
			Symbol:    pairForParse.BaseCurrency,
			Precision: basePrecision,
		},
		Quote: models.Asset{
			Symbol:    pairForParse.QuoteCurrency,
			Precision: quotePrecision,
		},
	}
}
