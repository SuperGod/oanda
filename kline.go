package oanda

import (
	"strings"
	"time"

	"github.com/SuperGod/oanda/client/operations"
	. "github.com/SuperGod/trademodel"
)

func (o *Oanda) KlineChan(start, end time.Time, bSize string) (datas chan []interface{}, err chan error) {
	datas = make(chan []interface{}, 1024)
	err = make(chan error, 1)
	klines, err1 := o.Kline(start, end, 0, bSize)
	if err1 != nil {
		close(datas)
		err <- err1
		close(err)
		return
	}
	var rets []interface{}
	for _, v := range klines {
		rets = append(rets, v)
	}
	datas <- rets
	close(datas)
	close(err)
	return
}

func (o *Oanda) Kline(start, end time.Time, nLimit int, binSize string) (klines []*Candle, err error) {
	return o.KlineSymbol(o.symbol, start, end, nLimit, binSize)
}

func (o *Oanda) KlineSymbol(symbol string, start, end time.Time, nLimit int, binSize string) (klines []*Candle, err error) {
	param := operations.NewGetInstrumentCandlesParams()
	o.initParam(param)
	endChar := binSize[len(binSize)-1]
	bSize := strings.ToUpper(string(endChar) + binSize[0:len(binSize)-1])
	param.Granularity = &bSize
	var strStart, strEnd string
	strStart = start.Format(TimeLayout)
	strEnd = end.Format(TimeLayout)
	param.From = &strStart
	param.To = &strEnd
	param.Instrument = symbol
	ret, err := o.api.Operations.GetInstrumentCandles(param)
	if err != nil {
		return
	}

	for _, v := range ret.Payload.Candles {
		c := Candle{}
		c.Open, _ = parseFloat(v.Mid.O)
		c.Close, _ = parseFloat(v.Mid.C)
		c.High, _ = parseFloat(v.Mid.H)
		c.Low, _ = parseFloat(v.Mid.L)
		t, _ := parseTime(v.Time)
		c.Start = t.Unix()
		c.Volume = float64(v.Volume)
		klines = append(klines, &c)
	}
	return
}
