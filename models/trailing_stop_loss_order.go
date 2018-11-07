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

// TrailingStopLossOrder A TrailingStopLossOrder is an order that is linked to an open Trade and created with a price distance. The price distance is used to calculate a trailing stop value for the order that is in the losing direction from the market price at the time of the order's creation. The trailing stop value will follow the market price as it moves in the winning direction, and the order will filled (closing the Trade) by the first price that is equal to or worse than the trailing stop value. A TrailingStopLossOrder cannot be used to open a new Position.
// swagger:model TrailingStopLossOrder
type TrailingStopLossOrder struct {

	// Date/time when the Order was cancelled (only provided when the state of the Order is CANCELLED)
	CancelledTime string `json:"cancelledTime,omitempty"`

	// ID of the Transaction that cancelled the Order (only provided when the Order's state is CANCELLED)
	CancellingTransactionID string `json:"cancellingTransactionID,omitempty"`

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID string `json:"clientTradeID,omitempty"`

	// The time when the Order was created.
	CreateTime string `json:"createTime,omitempty"`

	// The price distance (in price units) specified for the TrailingStopLoss Order.
	Distance string `json:"distance,omitempty"`

	// Date/time when the Order was filled (only provided when the Order's state is FILLED)
	FilledTime string `json:"filledTime,omitempty"`

	// ID of the Transaction that filled this Order (only provided when the Order's state is FILLED)
	FillingTransactionID string `json:"fillingTransactionID,omitempty"`

	// The date/time when the StopLoss Order will be cancelled if its timeInForce is "GTD".
	GtdTime string `json:"gtdTime,omitempty"`

	// The Order's identifier, unique within the Order's Account.
	ID string `json:"id,omitempty"`

	// The ID of the Order that replaced this Order (only provided if this Order was cancelled as part of a cancel/replace).
	ReplacedByOrderID string `json:"replacedByOrderID,omitempty"`

	// The ID of the Order that was replaced by this Order (only provided if this Order was created as part of a cancel/replace).
	ReplacesOrderID string `json:"replacesOrderID,omitempty"`

	// The current state of the Order.
	// Enum: [PENDING FILLED TRIGGERED CANCELLED]
	State string `json:"state,omitempty"`

	// The time-in-force requested for the TrailingStopLoss Order. Restricted to "GTC", "GFD" and "GTD" for TrailingStopLoss Orders.
	// Enum: [GTC GTD GFD FOK IOC]
	TimeInForce string `json:"timeInForce,omitempty"`

	// Trade IDs of Trades closed when the Order was filled (only provided when the Order's state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIds []string `json:"tradeClosedIDs"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID string `json:"tradeID,omitempty"`

	// Trade ID of Trade opened when the Order was filled (only provided when the Order's state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID string `json:"tradeOpenedID,omitempty"`

	// Trade ID of Trade reduced when the Order was filled (only provided when the Order's state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID string `json:"tradeReducedID,omitempty"`

	// The trigger price for the Trailing Stop Loss Order. The trailing stop value will trail (follow) the market price by the TSL order's configured "distance" as the market price moves in the winning direction. If the market price moves to a level that is equal to or worse than the trailing stop value, the order will be filled and the Trade will be closed.
	TrailingStopValue string `json:"trailingStopValue,omitempty"`

	// Specification of which price component should be used when determining if an Order should be triggered and filled. This allows Orders to be triggered based on the bid, ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy) price depending on the desired behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to specify a non-default trigger condition will not see it reflected in any of OANDA's proprietary or partner trading platforms, their transaction history or their account statements. OANDA platforms always assume that an Order's trigger condition is set to the default value when indicating the distance from an Order's trigger price, and will always provide the default trigger condition when creating or modifying an Order.
	// A special restriction applies when creating a guaranteed Stop Loss Order. In this case the TriggerCondition value must either be "DEFAULT", or the "natural" trigger side "DEFAULT" results in. So for a Stop Loss Order for a long trade valid values are "DEFAULT" and "BID", and for short trades "DEFAULT" and "ASK" are valid.
	// Enum: [DEFAULT INVERSE BID ASK MID]
	TriggerCondition string `json:"triggerCondition,omitempty"`

	// The type of the Order. Always set to "TRAILING_STOP_LOSS" for Trailing Stop Loss Orders.
	// Enum: [MARKET LIMIT STOP MARKET_IF_TOUCHED TAKE_PROFIT STOP_LOSS TRAILING_STOP_LOSS FIXED_PRICE]
	Type string `json:"type,omitempty"`
}

// Validate validates this trailing stop loss order
func (m *TrailingStopLossOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeInForce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggerCondition(formats); err != nil {
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

func (m *TrailingStopLossOrder) validateClientExtensions(formats strfmt.Registry) error {

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

var trailingStopLossOrderTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","FILLED","TRIGGERED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderTypeStatePropEnum = append(trailingStopLossOrderTypeStatePropEnum, v)
	}
}

const (

	// TrailingStopLossOrderStatePENDING captures enum value "PENDING"
	TrailingStopLossOrderStatePENDING string = "PENDING"

	// TrailingStopLossOrderStateFILLED captures enum value "FILLED"
	TrailingStopLossOrderStateFILLED string = "FILLED"

	// TrailingStopLossOrderStateTRIGGERED captures enum value "TRIGGERED"
	TrailingStopLossOrderStateTRIGGERED string = "TRIGGERED"

	// TrailingStopLossOrderStateCANCELLED captures enum value "CANCELLED"
	TrailingStopLossOrderStateCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *TrailingStopLossOrder) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrder) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var trailingStopLossOrderTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderTypeTimeInForcePropEnum = append(trailingStopLossOrderTypeTimeInForcePropEnum, v)
	}
}

const (

	// TrailingStopLossOrderTimeInForceGTC captures enum value "GTC"
	TrailingStopLossOrderTimeInForceGTC string = "GTC"

	// TrailingStopLossOrderTimeInForceGTD captures enum value "GTD"
	TrailingStopLossOrderTimeInForceGTD string = "GTD"

	// TrailingStopLossOrderTimeInForceGFD captures enum value "GFD"
	TrailingStopLossOrderTimeInForceGFD string = "GFD"

	// TrailingStopLossOrderTimeInForceFOK captures enum value "FOK"
	TrailingStopLossOrderTimeInForceFOK string = "FOK"

	// TrailingStopLossOrderTimeInForceIOC captures enum value "IOC"
	TrailingStopLossOrderTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *TrailingStopLossOrder) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrder) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

var trailingStopLossOrderTypeTriggerConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","INVERSE","BID","ASK","MID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderTypeTriggerConditionPropEnum = append(trailingStopLossOrderTypeTriggerConditionPropEnum, v)
	}
}

const (

	// TrailingStopLossOrderTriggerConditionDEFAULT captures enum value "DEFAULT"
	TrailingStopLossOrderTriggerConditionDEFAULT string = "DEFAULT"

	// TrailingStopLossOrderTriggerConditionINVERSE captures enum value "INVERSE"
	TrailingStopLossOrderTriggerConditionINVERSE string = "INVERSE"

	// TrailingStopLossOrderTriggerConditionBID captures enum value "BID"
	TrailingStopLossOrderTriggerConditionBID string = "BID"

	// TrailingStopLossOrderTriggerConditionASK captures enum value "ASK"
	TrailingStopLossOrderTriggerConditionASK string = "ASK"

	// TrailingStopLossOrderTriggerConditionMID captures enum value "MID"
	TrailingStopLossOrderTriggerConditionMID string = "MID"
)

// prop value enum
func (m *TrailingStopLossOrder) validateTriggerConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderTypeTriggerConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrder) validateTriggerCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerConditionEnum("triggerCondition", "body", m.TriggerCondition); err != nil {
		return err
	}

	return nil
}

var trailingStopLossOrderTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MARKET","LIMIT","STOP","MARKET_IF_TOUCHED","TAKE_PROFIT","STOP_LOSS","TRAILING_STOP_LOSS","FIXED_PRICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderTypeTypePropEnum = append(trailingStopLossOrderTypeTypePropEnum, v)
	}
}

const (

	// TrailingStopLossOrderTypeMARKET captures enum value "MARKET"
	TrailingStopLossOrderTypeMARKET string = "MARKET"

	// TrailingStopLossOrderTypeLIMIT captures enum value "LIMIT"
	TrailingStopLossOrderTypeLIMIT string = "LIMIT"

	// TrailingStopLossOrderTypeSTOP captures enum value "STOP"
	TrailingStopLossOrderTypeSTOP string = "STOP"

	// TrailingStopLossOrderTypeMARKETIFTOUCHED captures enum value "MARKET_IF_TOUCHED"
	TrailingStopLossOrderTypeMARKETIFTOUCHED string = "MARKET_IF_TOUCHED"

	// TrailingStopLossOrderTypeTAKEPROFIT captures enum value "TAKE_PROFIT"
	TrailingStopLossOrderTypeTAKEPROFIT string = "TAKE_PROFIT"

	// TrailingStopLossOrderTypeSTOPLOSS captures enum value "STOP_LOSS"
	TrailingStopLossOrderTypeSTOPLOSS string = "STOP_LOSS"

	// TrailingStopLossOrderTypeTRAILINGSTOPLOSS captures enum value "TRAILING_STOP_LOSS"
	TrailingStopLossOrderTypeTRAILINGSTOPLOSS string = "TRAILING_STOP_LOSS"

	// TrailingStopLossOrderTypeFIXEDPRICE captures enum value "FIXED_PRICE"
	TrailingStopLossOrderTypeFIXEDPRICE string = "FIXED_PRICE"
)

// prop value enum
func (m *TrailingStopLossOrder) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrder) validateType(formats strfmt.Registry) error {

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
func (m *TrailingStopLossOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrailingStopLossOrder) UnmarshalBinary(b []byte) error {
	var res TrailingStopLossOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}