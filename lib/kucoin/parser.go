package kucoin

import (
	"fmt"
	"strings"

	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/shopspring/decimal"
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

func ParseMarket(ticker tickerResponse) (*models.MarketAsset, error) {
	buy, err := decimal.NewFromString(ticker.Buy)
	if err != nil {
		return nil, err
	}

	sell, err := decimal.NewFromString(ticker.Sell)
	if err != nil {
		return nil, err
	}

	changeRate, err := decimal.NewFromString(ticker.ChangeRate)
	if err != nil {
		return nil, err
	}

	high, err := decimal.NewFromString(ticker.High)
	if err != nil {
		return nil, err
	}

	low, err := decimal.NewFromString(ticker.Low)
	if err != nil {
		return nil, err
	}

	volume, err := decimal.NewFromString(ticker.Vol)
	if err != nil {
		return nil, err
	}

	volumeValue, err := decimal.NewFromString(ticker.VolValue)
	if err != nil {
		return nil, err
	}

	last, err := decimal.NewFromString(ticker.VolValue)
	if err != nil {
		return nil, err
	}

	return &models.MarketAsset{
		Symbol:      ticker.Symbol,
		Buy:         &buy,
		Sell:        &sell,
		ChangeRate:  &changeRate,
		High:        &high,
		Low:         &low,
		Volume:      &volume,
		VolumeValue: &volumeValue,
		Last:        &last,
	}, nil
}
