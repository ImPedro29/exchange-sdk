package mexc

import "github.com/ImPedro29/exchange-sdk/lib"

type mexc struct {
	*lib.UnimplementedExchange
	Api       string `json:"api"`
	AccessKey string `json:"accessKey"`
	Secret    string `json:"secret"`
}

type data struct {
	Symbol          string `json:"symbol"`
	State           string `json:"state"`
	CountDownMark   int    `json:"countDownMark"`
	TimeZone        string `json:"timeZone"`
	FullName        string `json:"fullName"`
	SymbolStatus    int    `json:"symbolStatus"`
	VcoinName       string `json:"vcoinName"`
	VcoinStatus     int    `json:"vcoinStatus"`
	PriceScale      int64  `json:"price_scale"`
	QuantityScale   int64  `json:"quantity_scale"`
	MinAmount       string `json:"min_amount"`
	MaxAmount       string `json:"max_amount"`
	MakerFeeRate    string `json:"maker_fee_rate"`
	TakerFeeRate    string `json:"taker_fee_rate"`
	Limited         bool   `json:"limited"`
	EtfMark         int64  `json:"etf_mark"`
	SymbolPartition string `json:"symbol_partition"`
}

type supportedPairsResponse struct {
	Data []data `json:"data"`
}

type serverTimeResponse struct {
	Code int   `json:"code"`
	Data int64 `json:"data"`
}
