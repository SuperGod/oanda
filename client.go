package oanda

import (
	"fmt"
	"time"

	"github.com/SuperGod/oanda/client"
	"github.com/SuperGod/oanda/client/operations"
	httptransport "github.com/go-openapi/runtime/client"
)

const (
	BaseURL     = ""
	TestBaseURL = "api-fxpractice.oanda.com"
)

type Oanda struct {
	name      string // just a name to identify the api client
	api       *client.Oanda
	base      string
	token     string
	transport *httptransport.Runtime
}

// NewOanda create onada with token
func NewOanda(name, token string) *Oanda {
	return NewOandaWithURL(BaseURL, name, token)
}

// NewOanda create test onada with token
func NewTestOanda(name, token string) *Oanda {
	return NewOandaWithURL(TestBaseURL, name, token)
}

// NewOandaWithURL create oanda with url and token
func NewOandaWithURL(baseURL, name, token string) *Oanda {
	o := new(Oanda)
	o.name = name
	o.base = baseURL

	cfg := client.DefaultTransportConfig()
	cfg.Host = baseURL
	o.transport = httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	o.api = client.New(o.transport, nil)
	o.token = fmt.Sprintf("Bearer %s", token)
	return o
}

func (o *Oanda) SetDebug(debug bool) {
	o.transport.Debug = debug
}

func (o *Oanda) ListAccounts() (accounts []*Account, err error) {
	params := operations.NewListAccountsParams()
	params.SetAuthorization(o.token)
	ret, err := o.api.Operations.ListAccounts(params)
	if err != nil {
		return
	}
	time.Sleep(time.Second * 5)
	var account *Account
	for _, v := range ret.Payload.Accounts {
		account, err = o.GetAccountInfo(v.ID)
		if err != nil {
			fmt.Println("get accountinfo error:", err.Error())
			return
		}
		accounts = append(accounts, account)
		time.Sleep(time.Second * 5)
	}
	return
}

func (o *Oanda) GetAccountInfo(id string) (account *Account, err error) {
	params := operations.NewGetAccountParams()
	params.SetAuthorization(o.token)
	params.SetAccountID(id)
	params.SetTimeout(time.Second * 10)
	strFormat := "RFC3339"
	params.SetAcceptDatetimeFormat(&strFormat)
	ret, err := o.api.Operations.GetAccount(params)
	if err != nil {
		return
	}
	account = ret.Payload.Account
	return
}
