package oanda

import (
	"testing"
)

var (
	testApiClient *Oanda
)

func getTestClient() *Oanda {
	if testApiClient != nil {
		return testApiClient
	}
	testApiClient = GetDebugClient()
	testApiClient.FetchAccount()
	testApiClient.SetSymbol("EUR_USD")
	return testApiClient
}

func TestListAccounts(t *testing.T) {
	api := getTestClient()
	accounts, err := api.ListAccounts()
	if err != nil {
		t.Fatal(err.Error())
	}
	for acc := range accounts {
		t.Log(acc)
	}
}

func TestProxy(t *testing.T) {
	api := getTestClient()

	err := api.SetProxy("socks5://127.0.0.1:1080")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = api.FetchAccount()
	if err != nil {
		t.Fatal(err.Error())
	}
	rets, err := api.GetInstruments()
	if err != nil {
		t.Fatal(err.Error())
	}
	for ret := range rets {
		t.Log(ret)
	}
}
