package kucoin

import (
	"fmt"

	"github.com/ImPedro29/exchange-sdk/common"
	"github.com/ImPedro29/exchange-sdk/interfaces"
	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/ImPedro29/exchange-sdk/utils"
)

func NewKucoin(api, apiKey, secret string) interfaces.Exchange {
	return &kucoin{
		Api:    api,
		Secret: secret,
		Key:    apiKey,
	}
}

func (s *kucoin) GetPairs() (map[string]models.Pair, error) {
	var supportedPairsRes supportedPairsResponse
	if err := utils.GetURL(fmt.Sprintf("%s/api/v1/symbols", s.Api), &supportedPairsRes, nil); err != nil {
		return nil, err
	}

	pairs := make(map[string]models.Pair)
	for _, pair := range supportedPairsRes.Data {
		key, parsed := ParsePair(pair)
		pairs[key] = *parsed
	}

	return pairs, nil
}

func (s *kucoin) DepositAddress(asset models.Asset) (string, error) {
	return "", common.ErrNotSupported
}
