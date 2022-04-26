package models

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
