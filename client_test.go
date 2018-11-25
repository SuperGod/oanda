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
	testApiClient.SetSymbol("EUR/USD")
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
