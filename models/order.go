package models

import "github.com/shopspring/decimal"

type OrderCreation struct {
	Type        OrderType        `json:"type"`
	OrderMarket OrderMarket      `json:"orderMarket"`
	Amount      *decimal.Decimal `json:"amount"`
}

type OrderType int

const (
	MarketOrder OrderType = 1 + iota
	Limit
)

type OrderMarket int

const (
	Buy OrderMarket = 1 + iota
	Sell
)
