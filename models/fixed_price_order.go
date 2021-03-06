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

// FixedPriceOrder A FixedPriceOrder is an order that is filled immediately upon creation using a fixed price.
// swagger:model FixedPriceOrder
type FixedPriceOrder struct {

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime string `json:"cancelledTime,omitempty"`

	// ID of the Transaction that cancelled the Order (only provided when the Order's state is CANCELLED)
	CancellingTransactionID string `json:"cancellingTransactionID,omitempty"`

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The time when the Order was created.
	CreateTime string `json:"createTime,omitempty"`

	// Date/time when the Order was filled (only provided when the Order's state is FILLED)
	FilledTime string `json:"filledTime,omitempty"`

	// ID of the Transaction that filled this Order (only provided when the Order's state is FILLED)
	FillingTransactionID string `json:"fillingTransactionID,omitempty"`

	// The Order's identifier, unique within the Order's Account.
	ID string `json:"id,omitempty"`

	// The Fixed Price Order's Instrument.
	Instrument string `json:"instrument,omitempty"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	// Enum: [OPEN_ONLY REDUCE_FIRST REDUCE_ONLY DEFAULT]
	PositionFill string `json:"positionFill,omitempty"`

	// The price specified for the Fixed Price Order. This price is the exact price that the Fixed Price Order will be filled at.
	Price string `json:"price,omitempty"`

	// The current state of the Order.
	// Enum: [PENDING FILLED TRIGGERED CANCELLED]
	State string `json:"state,omitempty"`

	// stop loss on fill
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill,omitempty"`

	// take profit on fill
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill,omitempty"`

	// trade client extensions
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions,omitempty"`

	// Trade IDs of Trades closed when the Order was filled (only provided when the Order's state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIds []string `json:"tradeClosedIDs"`

	// Trade ID of Trade opened when the Order was filled (only provided when the Order's state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID string `json:"tradeOpenedID,omitempty"`

	// Trade ID of Trade reduced when the Order was filled (only provided when the Order's state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID string `json:"tradeReducedID,omitempty"`

	// The state that the trade resulting from the Fixed Price Order should be set to.
	TradeState string `json:"tradeState,omitempty"`

	// trailing stop loss on fill
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill,omitempty"`

	// The type of the Order. Always set to "FIXED_PRICE" for Fixed Price Orders.
	// Enum: [MARKET LIMIT STOP MARKET_IF_TOUCHED TAKE_PROFIT STOP_LOSS TRAILING_STOP_LOSS FIXED_PRICE]
	Type string `json:"type,omitempty"`

	// The quantity requested to be filled by the Fixed Price Order. A posititive number of units results in a long Order, and a negative number of units results in a short Order.
	Units string `json:"units,omitempty"`
}

// Validate validates this fixed price order
func (m *FixedPriceOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositionFill(formats); err != nil {
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

	if err := m.validateTradeClientExtensions(formats); err != nil {
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

func (m *FixedPriceOrder) validateClientExtensions(formats strfmt.Registry) error {

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

var fixedPriceOrderTypePositionFillPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN_ONLY","REDUCE_FIRST","REDUCE_ONLY","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fixedPriceOrderTypePositionFillPropEnum = append(fixedPriceOrderTypePositionFillPropEnum, v)
	}
}

const (

	// FixedPriceOrderPositionFillOPENONLY captures enum value "OPEN_ONLY"
	FixedPriceOrderPositionFillOPENONLY string = "OPEN_ONLY"

	// FixedPriceOrderPositionFillREDUCEFIRST captures enum value "REDUCE_FIRST"
	FixedPriceOrderPositionFillREDUCEFIRST string = "REDUCE_FIRST"

	// FixedPriceOrderPositionFillREDUCEONLY captures enum value "REDUCE_ONLY"
	FixedPriceOrderPositionFillREDUCEONLY string = "REDUCE_ONLY"

	// FixedPriceOrderPositionFillDEFAULT captures enum value "DEFAULT"
	FixedPriceOrderPositionFillDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *FixedPriceOrder) validatePositionFillEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fixedPriceOrderTypePositionFillPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FixedPriceOrder) validatePositionFill(formats strfmt.Registry) error {

	if swag.IsZero(m.PositionFill) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionFillEnum("positionFill", "body", m.PositionFill); err != nil {
		return err
	}

	return nil
}

var fixedPriceOrderTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","FILLED","TRIGGERED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fixedPriceOrderTypeStatePropEnum = append(fixedPriceOrderTypeStatePropEnum, v)
	}
}

const (

	// FixedPriceOrderStatePENDING captures enum value "PENDING"
	FixedPriceOrderStatePENDING string = "PENDING"

	// FixedPriceOrderStateFILLED captures enum value "FILLED"
	FixedPriceOrderStateFILLED string = "FILLED"

	// FixedPriceOrderStateTRIGGERED captures enum value "TRIGGERED"
	FixedPriceOrderStateTRIGGERED string = "TRIGGERED"

	// FixedPriceOrderStateCANCELLED captures enum value "CANCELLED"
	FixedPriceOrderStateCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *FixedPriceOrder) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fixedPriceOrderTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FixedPriceOrder) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *FixedPriceOrder) validateStopLossOnFill(formats strfmt.Registry) error {

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

func (m *FixedPriceOrder) validateTakeProfitOnFill(formats strfmt.Registry) error {

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

func (m *FixedPriceOrder) validateTradeClientExtensions(formats strfmt.Registry) error {

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

func (m *FixedPriceOrder) validateTrailingStopLossOnFill(formats strfmt.Registry) error {

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

var fixedPriceOrderTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MARKET","LIMIT","STOP","MARKET_IF_TOUCHED","TAKE_PROFIT","STOP_LOSS","TRAILING_STOP_LOSS","FIXED_PRICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fixedPriceOrderTypeTypePropEnum = append(fixedPriceOrderTypeTypePropEnum, v)
	}
}

const (

	// FixedPriceOrderTypeMARKET captures enum value "MARKET"
	FixedPriceOrderTypeMARKET string = "MARKET"

	// FixedPriceOrderTypeLIMIT captures enum value "LIMIT"
	FixedPriceOrderTypeLIMIT string = "LIMIT"

	// FixedPriceOrderTypeSTOP captures enum value "STOP"
	FixedPriceOrderTypeSTOP string = "STOP"

	// FixedPriceOrderTypeMARKETIFTOUCHED captures enum value "MARKET_IF_TOUCHED"
	FixedPriceOrderTypeMARKETIFTOUCHED string = "MARKET_IF_TOUCHED"

	// FixedPriceOrderTypeTAKEPROFIT captures enum value "TAKE_PROFIT"
	FixedPriceOrderTypeTAKEPROFIT string = "TAKE_PROFIT"

	// FixedPriceOrderTypeSTOPLOSS captures enum value "STOP_LOSS"
	FixedPriceOrderTypeSTOPLOSS string = "STOP_LOSS"

	// FixedPriceOrderTypeTRAILINGSTOPLOSS captures enum value "TRAILING_STOP_LOSS"
	FixedPriceOrderTypeTRAILINGSTOPLOSS string = "TRAILING_STOP_LOSS"

	// FixedPriceOrderTypeFIXEDPRICE captures enum value "FIXED_PRICE"
	FixedPriceOrderTypeFIXEDPRICE string = "FIXED_PRICE"
)

// prop value enum
func (m *FixedPriceOrder) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fixedPriceOrderTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FixedPriceOrder) validateType(formats strfmt.Registry) error {

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
func (m *FixedPriceOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FixedPriceOrder) UnmarshalBinary(b []byte) error {
	var res FixedPriceOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
