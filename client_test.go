package oanda

import (
	"testing"
)

func TestListAccounts(t *testing.T) {
	name, token := GetToken()
	api := NewTestOanda(name, token)
	api.SetDebug(true)
	accounts, err := api.ListAccounts()
	if err != nil {
		t.Fatal(err.Error())
	}
	for acc := range accounts {
		t.Log(acc)
	}
}
