package oanda

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/SuperGod/oanda/client"
	"github.com/SuperGod/oanda/client/operations"
	httptransport "github.com/go-openapi/runtime/client"
)

const (
	BaseURL     = ""
	TestBaseURL = "api-fxpractice.oanda.com"
)

var (
	AcceptDateTimeFormat = "RFC3339"
	DtPtr                = &AcceptDateTimeFormat
	TimeLayout           = time.RFC3339Nano
)

type AuthAccount interface {
	SetAccountID(string)
}

type AuthDatetimeFormat interface {
	SetAcceptDatetimeFormat(*string)
}

type AuthParam interface {
	SetAuthorization(string)
	SetTimeout(time.Duration)
}

type Oanda struct {
	name         string // just a name to identify the api client
	api          *client.Oanda
	base         string
	token        string
	transport    *httptransport.Runtime
	httpClient   *http.Client
	accountAlias string // which account to use
	accountID    string // fetch account use alias
	symbol       string
	timeout      time.Duration
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
	o.httpClient = &http.Client{Transport: &http.Transport{}}
	cfg := client.DefaultTransportConfig()
	cfg.Host = baseURL
	o.transport = httptransport.NewWithClient(cfg.Host, cfg.BasePath, cfg.Schemes, o.httpClient)
	o.api = client.New(o.transport, nil)
	o.token = fmt.Sprintf("Bearer %s", token)
	o.timeout = time.Second * 10
	o.symbol = "EUR_USD"
	return o
}

func (o *Oanda) SetProxy(proxy string) (err error) {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return
	}
	httpTransport := o.httpClient.Transport.(*http.Transport)
	httpTransport.Proxy = http.ProxyURL(proxyURL)
	return
}

// SetSymbol set symbol
func (o *Oanda) SetSymbol(symbol string) {
	o.symbol = symbol
}

// SetAccountAlias set which account to use
// if not set, use the first account
func (o *Oanda) SetAccountAlias(alias string) {
	o.accountAlias = alias
}

func (o *Oanda) SetDebug(debug bool) {
	o.transport.Debug = debug
}

// FetchAccount fetch account id from server
// must call first
func (o *Oanda) FetchAccount() (err error) {
	accounts, err := o.ListAccounts()
	if err != nil {
		return
	}
	if len(accounts) == 0 {
		err = errors.New("no account found")
		return
	}
	if o.accountAlias == "" {
		o.accountID = accounts[0].ID
		o.accountAlias = accounts[0].Alias
	}
	for _, v := range accounts {
		if v.Alias == o.accountAlias {
			o.accountID = v.ID
			break
		}
	}
	if o.accountID == "" {
		err = fmt.Errorf("can't account of alias %s", o.accountAlias)
	}
	return
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
	o.initParam(params)
	params.SetAccountID(id)
	ret, err := o.api.Operations.GetAccount(params)
	if err != nil {
		return
	}
	account = ret.Payload.Account
	return
}

func (o *Oanda) GetInstruments() (instruments []*Instrument, err error) {
	params := operations.NewGetAccountInstrumentsParams()
	o.initParam(params)
	params.AccountID = o.accountID
	fmt.Println("param:", o.accountID)
	ret, err := o.api.Operations.GetAccountInstruments(params)
	if err != nil {
		return
	}
	instruments = ret.Payload.Instruments
	return
}

func (o *Oanda) initParam(param interface{}) {
	authParam, ok := param.(AuthParam)
	if !ok {
		log.Println("param not impl AuthParam", param)
		return
	}
	authParam.SetAuthorization(o.token)
	authAccount, ok := param.(AuthAccount)
	if ok {
		authAccount.SetAccountID(o.accountID)
	}
	authParam.SetTimeout(o.timeout)
	authDatetimeFormat, ok := param.(AuthDatetimeFormat)
	if ok {
		authDatetimeFormat.SetAcceptDatetimeFormat(DtPtr)
	}
}
