package oanda

import (
	"testing"
	"time"
)

func TestKline(t *testing.T) {
	api := getTestClient()
	end := time.Now().Add(-time.Hour * 24)
	start := end.Add(-time.Hour * 24 * 7)
	rets, err := api.Kline(start, end, 100, "1h")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range rets {
		t.Log("kline:", v)
	}

}
