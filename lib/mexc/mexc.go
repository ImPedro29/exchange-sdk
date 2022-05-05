package mexc

import (
	"fmt"

	"github.com/ImPedro29/exchange-sdk/constraints"
	"github.com/ImPedro29/exchange-sdk/interfaces"
	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/ImPedro29/exchange-sdk/utils"
)

func NewMexc(accessKey, secret string) interfaces.Exchange {
	return &mexc{
		Api:       fmt.Sprintf("https://%s", constraints.Mexc),
		AccessKey: accessKey,
		Secret:    secret,
	}
}

func (s *mexc) GetPairs() (map[string]models.Pair, error) {
	var supportedPairsRes supportedPairsResponse
	if err := utils.GetURL(fmt.Sprintf("%s/open/api/v2/market/symbols", s.Api), &supportedPairsRes, nil); err != nil {
		return nil, err
	}

	pairs := make(map[string]models.Pair)
	for _, pair := range supportedPairsRes.Data {
		key, parsed := ParsePair(pair)
		pairs[key] = *parsed
	}

	return pairs, nil
}
