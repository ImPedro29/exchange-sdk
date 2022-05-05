package models

type KucoinAssetMarket struct {
	BestAsk     string `json:"bestAsk"`
	BestAskSize string `json:"bestAskSize"`
	BestBid     string `json:"bestBid"`
	BestBidSize string `json:"bestBidSize"`
	Price       string `json:"price"`
	Sequence    string `json:"sequence"`
	Size        string `json:"size"`
	Time        int64  `json:"time"`
}

type KucoinMarket struct {
	Data    KucoinAssetMarket `json:"data"`
	Subject string            `json:"subject"`
	Topic   string            `json:"topic"`
	Type    string            `json:"type"`
}
