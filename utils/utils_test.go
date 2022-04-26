package utils_test

import (
	"testing"

	"github.com/ImPedro29/exchange-sdk/utils"
	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {
	dataToSign := "symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559"

	t.Run("Success sign test", func(t *testing.T) {
		res := utils.SignSha256(dataToSign, "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j")
		require.Equal(t, "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71", res)
	})

	t.Run("Empty sign test", func(t *testing.T) {
		res := utils.SignSha256(dataToSign, "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j")
		require.NotEqual(t, "", res)
	})
}
