package kucoin_test

import (
	"testing"
	"time"

	"github.com/ImPedro29/exchange-sdk/constraints"
	"github.com/stretchr/testify/require"
)

func TestKucoinEvents_Start(t *testing.T) {
	vt, err := exchange.Events()
	require.NoError(t, err)
	require.NotEmpty(t, vt)
	defer func() {
		err = vt.Close()
		require.Error(t, err)
	}()

	err = vt.Listen(constraints.MarketAll, func(data interface{}) error {
		t.Log(data)
		return nil
	})

	require.NoError(t, err)
	time.Sleep(time.Second * 5)
}
