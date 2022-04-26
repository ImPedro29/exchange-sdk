package mexc

import (
	"fmt"
	"time"

	"github.com/ImPedro29/exchange-sdk/utils"
)

func (s *mexc) getServerTime() (*time.Time, error) {
	var serverTimeRes serverTimeResponse
	if err := utils.GetURL(fmt.Sprintf("%s/open/api/v2/common/timestamp", s.Api), &serverTimeRes, nil); err != nil {
		return nil, err
	}

	timeParsed := time.UnixMilli(serverTimeRes.Data)
	return &timeParsed, nil
}
