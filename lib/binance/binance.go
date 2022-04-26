package binance

import (
	"fmt"

	"github.com/ImPedro29/exchange-sdk/interfaces"
	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/ImPedro29/exchange-sdk/utils"
)

func NewBinance(api, apiKey, secret string) interfaces.Exchange {
	return &binance{
		Api:    api,
		Secret: secret,
		ApiKey: apiKey,
	}
}

func (s *binance) GetPairs() (map[string]models.Pair, error) {
	var supportedPairsRes supportedPairsResponse
	if err := utils.GetURL(fmt.Sprintf("%s/api/v1/exchangeInfo", s.Api), &supportedPairsRes, nil); err != nil {
		return nil, err
	}

	pairs := make(map[string]models.Pair)
	for _, pair := range supportedPairsRes.Symbols {
		key, parsed := ParsePair(pair)
		pairs[key] = *parsed
	}

	return pairs, nil
}

func (s *binance) DepositAddress(asset models.Asset) (string, error) {
	timestamp, err := s.getServerTime()
	if err != nil {
		return "", err
	}

	params := fmt.Sprintf("coin=%s&network=%s&timestamp=%d", asset.Symbol, asset.Network, timestamp)

	signed := utils.SignSha256(params, s.Secret)
	headers := [][2]string{{"X-MBX-APIKEY", s.ApiKey}}
	params = fmt.Sprintf("%s&signature=%s", params, signed)

	var depositAddressRes depositAddressResponse
	if err := utils.GetURL(fmt.Sprintf("%s/sapi/v1/capital/deposit/address?%s", s.Api, params), &depositAddressRes, headers); err != nil {
		return "", err
	}

	return depositAddressRes.Address, nil
}
