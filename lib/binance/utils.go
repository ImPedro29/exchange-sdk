package binance

import (
	"fmt"

	"github.com/ImPedro29/exchange-sdk/utils"
)

func (s *binance) getServerTime() (int64, error) {
	var serverTimeRes serverTimeResponse
	if err := utils.GetURL(fmt.Sprintf("%s/api/v3/time", s.Api), &serverTimeRes, nil); err != nil {
		return 0, err
	}

	return serverTimeRes.ServerTime, nil
}
