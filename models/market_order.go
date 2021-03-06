// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MarketOrder A MarketOrder is an order that is filled immediately upon creation using the current market price.
// swagger:model MarketOrder
type MarketOrder struct {

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime string `json:"cancelledTime,omitempty"`

	// ID of the Transaction that cancelled the Order (only provided when the Order's state is CANCELLED)
	CancellingTransactionID string `json:"cancellingTransactionID,omitempty"`

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The time when the Order was created.
	CreateTime string `json:"createTime,omitempty"`

	// delayed trade close
	DelayedTradeClose *MarketOrderDelayedTradeClose `json:"delayedTradeClose,omitempty"`

	// Date/time when the Order was filled (only provided when the Order's state is FILLED)
	FilledTime string `json:"filledTime,omitempty"`

	// ID of the Transaction that filled this Order (only provided when the Order's state is FILLED)
	FillingTransactionID string `json:"fillingTransactionID,omitempty"`

	// The Order's identifier, unique within the Order's Account.
	ID string `json:"id,omitempty"`

	// The Market Order's Instrument.
	Instrument string `json:"instrument,omitempty"`

	// long position closeout
	LongPositionCloseout *MarketOrderPositionCloseout `json:"longPositionCloseout,omitempty"`

	// margin closeout
	MarginCloseout *MarketOrderMarginCloseout `json:"marginCloseout,omitempty"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	// Enum: [OPEN_ONLY REDUCE_FIRST REDUCE_ONLY DEFAULT]
	PositionFill string `json:"positionFill,omitempty"`

	// The worst price that the client is willing to have the Market Order filled at.
	PriceBound string `json:"priceBound,omitempty"`

	// short position closeout
	ShortPositionCloseout *MarketOrderPositionCloseout `json:"shortPositionCloseout,omitempty"`

	// The current state of the Order.
	// Enum: [PENDING FILLED TRIGGERED CANCELLED]
	State string `json:"state,omitempty"`

	// stop loss on fill
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill,omitempty"`

	// take profit on fill
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill,omitempty"`

	// The time-in-force requested for the Market Order. Restricted to FOK or IOC for a MarketOrder.
	// Enum: [GTC GTD GFD FOK IOC]
	TimeInForce string `json:"timeInForce,omitempty"`

	// trade client extensions
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions,omitempty"`

	// trade close
	TradeClose *MarketOrderTradeClose `json:"tradeClose,omitempty"`

	// Trade IDs of Trades closed when the Order was filled (only provided when the Order's state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIds []string `json:"tradeClosedIDs"`

	// Trade ID of Trade opened when the Order was filled (only provided when the Order's state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID string `json:"tradeOpenedID,omitempty"`

	// Trade ID of Trade reduced when the Order was filled (only provided when the Order's state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID string `json:"tradeReducedID,omitempty"`

	// trailing stop loss on fill
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill,omitempty"`

	// The type of the Order. Always set to "MARKET" for Market Orders.
	// Enum: [MARKET LIMIT STOP MARKET_IF_TOUCHED TAKE_PROFIT STOP_LOSS TRAILING_STOP_LOSS FIXED_PRICE]
	Type string `json:"type,omitempty"`

	// The quantity requested to be filled by the Market Order. A posititive number of units results in a long Order, and a negative number of units results in a short Order.
	Units string `json:"units,omitempty"`
}

// Validate validates this market order
func (m *MarketOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelayedTradeClose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLongPositionCloseout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarginCloseout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositionFill(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortPositionCloseout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopLossOnFill(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTakeProfitOnFill(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeInForce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTradeClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTradeClose(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrailingStopLossOnFill(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketOrder) validateClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientExtensions) { // not required
		return nil
	}

	if m.ClientExtensions != nil {
		if err := m.ClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientExtensions")
			}
			return err
		}
	}

	return nil
}

func (m *MarketOrder) validateDelayedTradeClose(formats strfmt.Registry) error {

	if swag.IsZero(m.DelayedTradeClose) { // not required
		return nil
	}

	if m.DelayedTradeClose != nil {
		if err := m.DelayedTradeClose.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delayedTradeClose")
			}
			return err
		}
	}

	return nil
}

func (m *MarketOrder) validateLongPositionCloseout(formats strfmt.Registry) error {

	if swag.IsZero(m.LongPositionCloseout) { // not required
		return nil
	}

	if m.LongPositionCloseout != nil {
		if err := m.LongPositionCloseout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("longPositionCloseout")
			}
			return err
		}
	}

	return nil
}

func (m *MarketOrder) validateMarginCloseout(formats strfmt.Registry) error {

	if swag.IsZero(m.MarginCloseout) { // not required
		return nil
	}

	if m.MarginCloseout != nil {
		if err := m.MarginCloseout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("marginCloseout")
			}
			return err
		}
	}

	return nil
}

var marketOrderTypePositionFillPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN_ONLY","REDUCE_FIRST","REDUCE_ONLY","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketOrderTypePositionFillPropEnum = append(marketOrderTypePositionFillPropEnum, v)
	}
}

const (

	// MarketOrderPositionFillOPENONLY captures enum value "OPEN_ONLY"
	MarketOrderPositionFillOPENONLY string = "OPEN_ONLY"

	// MarketOrderPositionFillREDUCEFIRST captures enum value "REDUCE_FIRST"
	MarketOrderPositionFillREDUCEFIRST string = "REDUCE_FIRST"

	// MarketOrderPositionFillREDUCEONLY captures enum value "REDUCE_ONLY"
	MarketOrderPositionFillREDUCEONLY string = "REDUCE_ONLY"

	// MarketOrderPositionFillDEFAULT captures enum value "DEFAULT"
	MarketOrderPositionFillDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *MarketOrder) validatePositionFillEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketOrderTypePositionFillPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketOrder) validatePositionFill(formats strfmt.Registry) error {

	if swag.IsZero(m.PositionFill) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionFillEnum("positionFill", "body", m.PositionFill); err != nil {
		return err
	}

	return nil
}

func (m *MarketOrder) validateShortPositionCloseout(formats strfmt.Registry) error {

	if swag.IsZero(m.ShortPositionCloseout) { // not required
		return nil
	}

	if m.ShortPositionCloseout != nil {
		if err := m.ShortPositionCloseout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortPositionCloseout")
			}
			return err
		}
	}

	return nil
}

var marketOrderTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","FILLED","TRIGGERED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketOrderTypeStatePropEnum = append(marketOrderTypeStatePropEnum, v)
	}
}

const (

	// MarketOrderStatePENDING captures enum value "PENDING"
	MarketOrderStatePENDING string = "PENDING"

	// MarketOrderStateFILLED captures enum value "FILLED"
	MarketOrderStateFILLED string = "FILLED"

	// MarketOrderStateTRIGGERED captures enum value "TRIGGERED"
	MarketOrderStateTRIGGERED string = "TRIGGERED"

	// MarketOrderStateCANCELLED captures enum value "CANCELLED"
	MarketOrderStateCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *MarketOrder) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketOrderTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketOrder) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *MarketOrder) validateStopLossOnFill(formats strfmt.Registry) error {

	if swag.IsZero(m.StopLossOnFill) { // not required
		return nil
	}

	if m.StopLossOnFill != nil {
		if err := m.StopLossOnFill.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stopLossOnFill")
			}
			return err
		}
	}

	return nil
}

func (m *MarketOrder) validateTakeProfitOnFill(formats strfmt.Registry) error {

	if swag.IsZero(m.TakeProfitOnFill) { // not required
		return nil
	}

	if m.TakeProfitOnFill != nil {
		if err := m.TakeProfitOnFill.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("takeProfitOnFill")
			}
			return err
		}
	}

	return nil
}

var marketOrderTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketOrderTypeTimeInForcePropEnum = append(marketOrderTypeTimeInForcePropEnum, v)
	}
}

const (

	// MarketOrderTimeInForceGTC captures enum value "GTC"
	MarketOrderTimeInForceGTC string = "GTC"

	// MarketOrderTimeInForceGTD captures enum value "GTD"
	MarketOrderTimeInForceGTD string = "GTD"

	// MarketOrderTimeInForceGFD captures enum value "GFD"
	MarketOrderTimeInForceGFD string = "GFD"

	// MarketOrderTimeInForceFOK captures enum value "FOK"
	MarketOrderTimeInForceFOK string = "FOK"

	// MarketOrderTimeInForceIOC captures enum value "IOC"
	MarketOrderTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *MarketOrder) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketOrderTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketOrder) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

func (m *MarketOrder) validateTradeClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.TradeClientExtensions) { // not required
		return nil
	}

	if m.TradeClientExtensions != nil {
		if err := m.TradeClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tradeClientExtensions")
			}
			return err
		}
	}

	return nil
}

func (m *MarketOrder) validateTradeClose(formats strfmt.Registry) error {

	if swag.IsZero(m.TradeClose) { // not required
		return nil
	}

	if m.TradeClose != nil {
		if err := m.TradeClose.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tradeClose")
			}
			return err
		}
	}

	return nil
}

func (m *MarketOrder) validateTrailingStopLossOnFill(formats strfmt.Registry) error {

	if swag.IsZero(m.TrailingStopLossOnFill) { // not required
		return nil
	}

	if m.TrailingStopLossOnFill != nil {
		if err := m.TrailingStopLossOnFill.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trailingStopLossOnFill")
			}
			return err
		}
	}

	return nil
}

var marketOrderTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MARKET","LIMIT","STOP","MARKET_IF_TOUCHED","TAKE_PROFIT","STOP_LOSS","TRAILING_STOP_LOSS","FIXED_PRICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketOrderTypeTypePropEnum = append(marketOrderTypeTypePropEnum, v)
	}
}

const (

	// MarketOrderTypeMARKET captures enum value "MARKET"
	MarketOrderTypeMARKET string = "MARKET"

	// MarketOrderTypeLIMIT captures enum value "LIMIT"
	MarketOrderTypeLIMIT string = "LIMIT"

	// MarketOrderTypeSTOP captures enum value "STOP"
	MarketOrderTypeSTOP string = "STOP"

	// MarketOrderTypeMARKETIFTOUCHED captures enum value "MARKET_IF_TOUCHED"
	MarketOrderTypeMARKETIFTOUCHED string = "MARKET_IF_TOUCHED"

	// MarketOrderTypeTAKEPROFIT captures enum value "TAKE_PROFIT"
	MarketOrderTypeTAKEPROFIT string = "TAKE_PROFIT"

	// MarketOrderTypeSTOPLOSS captures enum value "STOP_LOSS"
	MarketOrderTypeSTOPLOSS string = "STOP_LOSS"

	// MarketOrderTypeTRAILINGSTOPLOSS captures enum value "TRAILING_STOP_LOSS"
	MarketOrderTypeTRAILINGSTOPLOSS string = "TRAILING_STOP_LOSS"

	// MarketOrderTypeFIXEDPRICE captures enum value "FIXED_PRICE"
	MarketOrderTypeFIXEDPRICE string = "FIXED_PRICE"
)

// prop value enum
func (m *MarketOrder) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketOrderTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketOrder) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarketOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketOrder) UnmarshalBinary(b []byte) error {
	var res MarketOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
