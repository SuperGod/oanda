package oanda

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SuperGod/coinex"
	"github.com/SuperGod/oanda/client/operations"
	"github.com/SuperGod/oanda/models"
	. "github.com/SuperGod/trademodel"
)

type Position = coinex.Position
type Contract = coinex.Contract

func (o *Oanda) OpenLong(price float64, amount float64) (order *Order, err error) {
	order, err = o.createLimitOrder(price, amount, "long", "OpenLong by api")
	return
}

func (o *Oanda) OpenLongMarket(amount float64) (order *Order, err error) {
	order, err = o.createMarketOrder(amount, "long", "OpenLongMarket by api")
	return
}

func (o *Oanda) CloseLongMarket(amount float64) (order *Order, err error) {
	param := operations.NewClosePositionParams()
	o.initParam(param)
	param.ClosePositionBody = operations.ClosePositionBody{
		LongUnits:            fmt.Sprintf("%f", amount),
		LongClientExtensions: &models.ClientExtensions{},
	}
	ret, err := o.api.Operations.ClosePosition(param)
	if err != nil {
		return
	}
	order = &Order{}
	order.OrderID = ret.Payload.ShortOrderFillTransaction.ID
	order.Amount = amount
	// order.Price = price
	order.Side = "short"
	order.Time, err = parseTime(ret.Payload.ShortOrderFillTransaction.Time)
	order.Currency = o.symbol
	if err != nil {
		return
	}
	var _ = ret
	return
}

func (o *Oanda) CloseLongPrice(price, amount float64) (order *Order, err error) {
	order, err = o.createLimitOrder(price, 0-amount, "short", "CloseLongPrice by api")
	return
}

func (o *Oanda) OpenShort(price float64, amount float64) (order *Order, err error) {
	order, err = o.createLimitOrder(price, 0-amount, "short", "OpenShort by api")
	return
}

func (o *Oanda) OpenShortMarket(amount float64) (order *Order, err error) {
	order, err = o.createMarketOrder(0-amount, "short", "OpenLongMarket by api")
	return
}

func (o *Oanda) CloseShortMarket(amount float64) (order *Order, err error) {
	order, err = o.createMarketOrder(amount, "long", "CloseShortMarket by api")
	return
}

func (o *Oanda) CloseShortPrice(price, amount float64) (order *Order, err error) {
	order, err = o.createLimitOrder(price, amount, "long", "CloseShortPrice by api")
	return
}

func (o *Oanda) createLimitOrder(price, amount float64, side, comment string) (order *Order, err error) {
	param := operations.NewCreateOrderParams()
	param.CreateOrderBody.Order = &models.LimitOrderRequest{
		Instrument:   o.symbol,
		Price:        fmt.Sprintf("%f", price),
		Type:         string(models.OrderTypeLIMIT),
		Units:        fmt.Sprintf("%d", int32(amount)),
		PositionFill: string(models.OrderPositionFillDEFAULT),
		ClientExtensions: &models.ClientExtensions{
			Comment: comment,
			Tag:     side,
			ID:      ""},
	}
	o.initParam(param)
	newOrder, err := o.api.Operations.CreateOrder(param)
	if err == nil {
		order = transCreateOrder(o.symbol, price, amount, side, newOrder)
	}
	return
}

func (o *Oanda) createMarketOrder(amount float64, side, comment string) (order *Order, err error) {
	param := operations.NewCreateOrderParams()
	param.CreateOrderBody.Order = &models.MarketOrderRequest{
		Instrument:   o.symbol,
		Type:         string(models.OrderTypeMARKET),
		Units:        fmt.Sprintf("%d", int32(amount)),
		PositionFill: string(models.OrderPositionFillOPENONLY),
		ClientExtensions: &models.ClientExtensions{
			Comment: comment,
			Tag:     side,
			ID:      ""},
	}
	o.initParam(param)
	var price float64
	newOrder, err := o.api.Operations.CreateOrder(param)
	if newOrder.Payload.OrderFillTransaction != nil {
		priceStr := newOrder.Payload.OrderFillTransaction.TradeOpened.Price
		price, err = parseFloat(priceStr)
		if err != nil {
			err = fmt.Errorf("createMarketOrder: price is not float64 %s", priceStr)
			return
		}
	}
	if err == nil {
		order = transCreateOrder(o.symbol, price, amount, side, newOrder)
	}
	return
}

func (o *Oanda) Buy(price float64, amount float64) (*Order, error) {
	return o.OpenLong(price, amount)
}
func (o *Oanda) Sell(price float64, amount float64) (*Order, error) {
	return o.OpenShort(price, amount)
}

func transCreateOrder(symbol string, price, amount float64, side string, rawOrder *operations.CreateOrderCreated) (order *Order) {
	order = &Order{}
	order.OrderID = rawOrder.Payload.OrderCreateTransaction.ID
	order.Amount = float64(amount)
	order.Price = price
	order.Side = side
	var err error
	order.Time, err = parseTime(rawOrder.Payload.OrderCreateTransaction.Time)
	if err != nil {
		log.Println("transCreateOrder parse time failed:", rawOrder.Payload.OrderCreateTransaction.Time)
	}
	order.Currency = symbol
	return
}

func parseTime(str string) (t time.Time, err error) {
	t, err = time.Parse(TimeLayout, str)
	return
}

func parseFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}
