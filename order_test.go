package oanda

import "testing"

func TestOpenLong(t *testing.T) {
	api := getTestClient()
	order, err := api.OpenLong(1.0, 1.0)
	if err != nil {
		return
	}
	t.Log("order:", order)
	order, err = api.OpenLongMarket(1.0)
	if err != nil {
		return
	}
	t.Log("order2:", order)
}

func TestOpenShort(t *testing.T) {
	api := getTestClient()
	order, err := api.OpenShort(1.0, 1.0)
	if err != nil {
		return
	}
	t.Log("order:", order)
	order, err = api.OpenShortMarket(1.0)
	if err != nil {
		return
	}
	t.Log("order2:", order)
}
