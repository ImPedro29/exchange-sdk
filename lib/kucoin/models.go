package kucoin

import (
	"github.com/ImPedro29/exchange-sdk/lib"
	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/gorilla/websocket"
)

type kucoin struct {
	*lib.UnimplementedExchange
	Api    string `json:"api"`
	Secret string `json:"secret"`
	Key    string `json:"key"`
	events *kucoinEvents
}

type kucoinEvents struct {
	Started   bool            `json:"started"`
	Api       string          `json:"api"`
	Token     string          `json:"token"`
	ConnectID string          `json:"connectID"`
	Conn      *websocket.Conn `json:"conn"`
	handlers  map[string]models.EventHandler
	close     chan bool
}

type data struct {
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	BaseCurrency    string `json:"baseCurrency"`
	QuoteCurrency   string `json:"quoteCurrency"`
	FeeCurrency     string `json:"feeCurrency"`
	Market          string `json:"market"`
	BaseMinSize     string `json:"baseMinSize"`
	QuoteMinSize    string `json:"quoteMinSize"`
	BaseMaxSize     string `json:"baseMaxSize"`
	QuoteMaxSize    string `json:"quoteMaxSize"`
	BaseIncrement   string `json:"baseIncrement"`
	QuoteIncrement  string `json:"quoteIncrement"`
	PriceIncrement  string `json:"priceIncrement"`
	PriceLimitRate  string `json:"priceLimitRate"`
	MinFunds        string `json:"minFunds"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
	EnableTrading   bool   `json:"enableTrading"`
}

type supportedPairsResponse struct {
	Data []data `json:"data"`
	Code string `json:"code"`
}

type marketResponse struct {
	Code string `json:"code"`
	Data struct {
		Time   int64            `json:"time"`
		Ticker []tickerResponse `json:"ticker"`
	} `json:"data"`
}

type tickerResponse struct {
	Symbol           string `json:"symbol"`
	SymbolName       string `json:"symbolName"`
	Buy              string `json:"buy"`
	Sell             string `json:"sell"`
	ChangeRate       string `json:"changeRate"`
	ChangePrice      string `json:"changePrice"`
	High             string `json:"high"`
	Low              string `json:"low"`
	Vol              string `json:"vol"`
	VolValue         string `json:"volValue"`
	Last             string `json:"last"`
	AveragePrice     string `json:"averagePrice"`
	TakerFeeRate     string `json:"takerFeeRate"`
	MakerFeeRate     string `json:"makerFeeRate"`
	TakerCoefficient string `json:"takerCoefficient"`
	MakerCoefficient string `json:"makerCoefficient"`
}

// Websocket
type websocketStartResponse struct {
	Code string `json:"code"`
	Data struct {
		Token           string `json:"token"`
		InstanceServers []struct {
			Endpoint     string `json:"endpoint"`
			Encrypt      bool   `json:"encrypt"`
			Protocol     string `json:"protocol"`
			PingInterval int    `json:"pingInterval"`
			PingTimeout  int    `json:"pingTimeout"`
		} `json:"instanceServers"`
	} `json:"data"`
}

type websocketWelcomeResponse struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type websocketRequest struct {
	Id             int64  `json:"id"`
	Type           string `json:"type"`
	Topic          string `json:"topic"`
	PrivateChannel bool   `json:"privateChannel"`
	Response       bool   `json:"response"`
}
