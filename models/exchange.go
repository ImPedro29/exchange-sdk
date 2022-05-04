package models

import "github.com/shopspring/decimal"

type PairStatus int

const (
	Enabled PairStatus = 1 + iota
	Disabled
)

type Account struct {
	Address string `json:"address"`
	Network string `json:"network"`
}

type Asset struct {
	Symbol    string `json:"symbol"`
	Precision int64  `json:"precision"`
	Network   string `json:"network,omitempty"`
}

type Pair struct {
	Status PairStatus `json:"status"`
	Base   Asset      `json:"base"`
	Quote  Asset      `json:"quote"`
}

type MarketAsset struct {
	Symbol      string           `json:"symbol,omitempty"`
	Buy         *decimal.Decimal `json:"buy,omitempty"`
	Sell        *decimal.Decimal `json:"sell,omitempty"`
	ChangeRate  *decimal.Decimal `json:"changeRate,omitempty"`
	High        *decimal.Decimal `json:"high,omitempty"`
	Low         *decimal.Decimal `json:"low,omitempty"`
	Volume      *decimal.Decimal `json:"volume,omitempty"`
	VolumeValue *decimal.Decimal `json:"volumeValue,omitempty"`
	Last        *decimal.Decimal `json:"last,omitempty"`
}
