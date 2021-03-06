package binance

import "github.com/ImPedro29/exchange-sdk/lib"

type binance struct {
	*lib.UnimplementedExchange
	Api    string `json:"api"`
	Secret string `json:"secret"`
	ApiKey string `json:"apiKey"`
}

type symbol struct {
	Symbol                     string   `json:"symbol"`
	Status                     string   `json:"status"`
	BaseAsset                  string   `json:"baseAsset"`
	BaseAssetPrecision         int      `json:"baseAssetPrecision"`
	QuoteAsset                 string   `json:"quoteAsset"`
	QuotePrecision             int      `json:"quotePrecision"`
	QuoteAssetPrecision        int      `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int      `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int      `json:"quoteCommissionPrecision"`
	OrderTypes                 []string `json:"orderTypes"`
	IcebergAllowed             bool     `json:"icebergAllowed"`
	OcoAllowed                 bool     `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed bool     `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop          bool     `json:"allowTrailingStop"`
	IsSpotTradingAllowed       bool     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool     `json:"isMarginTradingAllowed"`
	Filters                    []struct {
		FilterType            string `json:"filterType"`
		MinPrice              string `json:"minPrice,omitempty"`
		MaxPrice              string `json:"maxPrice,omitempty"`
		TickSize              string `json:"tickSize,omitempty"`
		MultiplierUp          string `json:"multiplierUp,omitempty"`
		MultiplierDown        string `json:"multiplierDown,omitempty"`
		AvgPriceMins          int    `json:"avgPriceMins,omitempty"`
		MinQty                string `json:"minQty,omitempty"`
		MaxQty                string `json:"maxQty,omitempty"`
		StepSize              string `json:"stepSize,omitempty"`
		MinNotional           string `json:"minNotional,omitempty"`
		ApplyToMarket         bool   `json:"applyToMarket,omitempty"`
		Limit                 int    `json:"limit,omitempty"`
		MinTrailingAboveDelta int    `json:"minTrailingAboveDelta,omitempty"`
		MaxTrailingAboveDelta int    `json:"maxTrailingAboveDelta,omitempty"`
		MinTrailingBelowDelta int    `json:"minTrailingBelowDelta,omitempty"`
		MaxTrailingBelowDelta int    `json:"maxTrailingBelowDelta,omitempty"`
		MaxNumOrders          int    `json:"maxNumOrders,omitempty"`
		MaxNumAlgoOrders      int    `json:"maxNumAlgoOrders,omitempty"`
	} `json:"filters"`
	Permissions []string `json:"permissions"`
}

type supportedPairsResponse struct {
	Symbols []symbol `json:"symbols"`
}

type depositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

type serverTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}
