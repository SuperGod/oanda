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

// TrailingStopLossOrderTransaction A TrailingStopLossOrderTransaction represents the creation of a TrailingStopLoss Order in the user's Account.
// swagger:model TrailingStopLossOrderTransaction
type TrailingStopLossOrderTransaction struct {

	// The ID of the Account the Transaction was created for.
	AccountID string `json:"accountID,omitempty"`

	// The ID of the "batch" that the Transaction belongs to. Transactions in the same batch are applied to the Account simultaneously.
	BatchID string `json:"batchID,omitempty"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID string `json:"cancellingTransactionID,omitempty"`

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID string `json:"clientTradeID,omitempty"`

	// The price distance (in price units) specified for the TrailingStopLoss Order.
	Distance string `json:"distance,omitempty"`

	// The date/time when the StopLoss Order will be cancelled if its timeInForce is "GTD".
	GtdTime string `json:"gtdTime,omitempty"`

	// The Transaction's Identifier.
	ID string `json:"id,omitempty"`

	// The ID of the OrderFill Transaction that caused this Order to be created (only provided if this Order was created automatically when another Order was filled).
	OrderFillTransactionID string `json:"orderFillTransactionID,omitempty"`

	// The reason that the Trailing Stop Loss Order was initiated
	// Enum: [CLIENT_ORDER REPLACEMENT ON_FILL]
	Reason string `json:"reason,omitempty"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID string `json:"replacesOrderID,omitempty"`

	// The Request ID of the request which generated the transaction.
	RequestID string `json:"requestID,omitempty"`

	// The date/time when the Transaction was created.
	Time string `json:"time,omitempty"`

	// The time-in-force requested for the TrailingStopLoss Order. Restricted to "GTC", "GFD" and "GTD" for TrailingStopLoss Orders.
	// Enum: [GTC GTD GFD FOK IOC]
	TimeInForce string `json:"timeInForce,omitempty"`

	// The ID of the Trade to close when the price threshold is breached.
	TradeID string `json:"tradeID,omitempty"`

	// Specification of which price component should be used when determining if an Order should be triggered and filled. This allows Orders to be triggered based on the bid, ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy) price depending on the desired behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to specify a non-default trigger condition will not see it reflected in any of OANDA's proprietary or partner trading platforms, their transaction history or their account statements. OANDA platforms always assume that an Order's trigger condition is set to the default value when indicating the distance from an Order's trigger price, and will always provide the default trigger condition when creating or modifying an Order.
	// A special restriction applies when creating a guaranteed Stop Loss Order. In this case the TriggerCondition value must either be "DEFAULT", or the "natural" trigger side "DEFAULT" results in. So for a Stop Loss Order for a long trade valid values are "DEFAULT" and "BID", and for short trades "DEFAULT" and "ASK" are valid.
	// Enum: [DEFAULT INVERSE BID ASK MID]
	TriggerCondition string `json:"triggerCondition,omitempty"`

	// The Type of the Transaction. Always set to "TRAILING_STOP_LOSS_ORDER" in a TrailingStopLossOrderTransaction.
	// Enum: [CREATE CLOSE REOPEN CLIENT_CONFIGURE CLIENT_CONFIGURE_REJECT TRANSFER_FUNDS TRANSFER_FUNDS_REJECT MARKET_ORDER MARKET_ORDER_REJECT FIXED_PRICE_ORDER LIMIT_ORDER LIMIT_ORDER_REJECT STOP_ORDER STOP_ORDER_REJECT MARKET_IF_TOUCHED_ORDER MARKET_IF_TOUCHED_ORDER_REJECT TAKE_PROFIT_ORDER TAKE_PROFIT_ORDER_REJECT STOP_LOSS_ORDER STOP_LOSS_ORDER_REJECT TRAILING_STOP_LOSS_ORDER TRAILING_STOP_LOSS_ORDER_REJECT ORDER_FILL ORDER_CANCEL ORDER_CANCEL_REJECT ORDER_CLIENT_EXTENSIONS_MODIFY ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT TRADE_CLIENT_EXTENSIONS_MODIFY TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT MARGIN_CALL_ENTER MARGIN_CALL_EXTEND MARGIN_CALL_EXIT DELAYED_TRADE_CLOSURE DAILY_FINANCING RESET_RESETTABLE_PL]
	Type string `json:"type,omitempty"`

	// The ID of the user that initiated the creation of the Transaction.
	UserID int64 `json:"userID,omitempty"`
}

// Validate validates this trailing stop loss order transaction
func (m *TrailingStopLossOrderTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
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

func (m *TrailingStopLossOrderTransaction) validateClientExtensions(formats strfmt.Registry) error {

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

var trailingStopLossOrderTransactionTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLIENT_ORDER","REPLACEMENT","ON_FILL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderTransactionTypeReasonPropEnum = append(trailingStopLossOrderTransactionTypeReasonPropEnum, v)
	}
}

const (

	// TrailingStopLossOrderTransactionReasonCLIENTORDER captures enum value "CLIENT_ORDER"
	TrailingStopLossOrderTransactionReasonCLIENTORDER string = "CLIENT_ORDER"

	// TrailingStopLossOrderTransactionReasonREPLACEMENT captures enum value "REPLACEMENT"
	TrailingStopLossOrderTransactionReasonREPLACEMENT string = "REPLACEMENT"

	// TrailingStopLossOrderTransactionReasonONFILL captures enum value "ON_FILL"
	TrailingStopLossOrderTransactionReasonONFILL string = "ON_FILL"
)

// prop value enum
func (m *TrailingStopLossOrderTransaction) validateReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderTransactionTypeReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrderTransaction) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	// value enum
	if err := m.validateReasonEnum("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

var trailingStopLossOrderTransactionTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderTransactionTypeTimeInForcePropEnum = append(trailingStopLossOrderTransactionTypeTimeInForcePropEnum, v)
	}
}

const (

	// TrailingStopLossOrderTransactionTimeInForceGTC captures enum value "GTC"
	TrailingStopLossOrderTransactionTimeInForceGTC string = "GTC"

	// TrailingStopLossOrderTransactionTimeInForceGTD captures enum value "GTD"
	TrailingStopLossOrderTransactionTimeInForceGTD string = "GTD"

	// TrailingStopLossOrderTransactionTimeInForceGFD captures enum value "GFD"
	TrailingStopLossOrderTransactionTimeInForceGFD string = "GFD"

	// TrailingStopLossOrderTransactionTimeInForceFOK captures enum value "FOK"
	TrailingStopLossOrderTransactionTimeInForceFOK string = "FOK"

	// TrailingStopLossOrderTransactionTimeInForceIOC captures enum value "IOC"
	TrailingStopLossOrderTransactionTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *TrailingStopLossOrderTransaction) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderTransactionTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrderTransaction) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

var trailingStopLossOrderTransactionTypeTriggerConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","INVERSE","BID","ASK","MID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderTransactionTypeTriggerConditionPropEnum = append(trailingStopLossOrderTransactionTypeTriggerConditionPropEnum, v)
	}
}

const (

	// TrailingStopLossOrderTransactionTriggerConditionDEFAULT captures enum value "DEFAULT"
	TrailingStopLossOrderTransactionTriggerConditionDEFAULT string = "DEFAULT"

	// TrailingStopLossOrderTransactionTriggerConditionINVERSE captures enum value "INVERSE"
	TrailingStopLossOrderTransactionTriggerConditionINVERSE string = "INVERSE"

	// TrailingStopLossOrderTransactionTriggerConditionBID captures enum value "BID"
	TrailingStopLossOrderTransactionTriggerConditionBID string = "BID"

	// TrailingStopLossOrderTransactionTriggerConditionASK captures enum value "ASK"
	TrailingStopLossOrderTransactionTriggerConditionASK string = "ASK"

	// TrailingStopLossOrderTransactionTriggerConditionMID captures enum value "MID"
	TrailingStopLossOrderTransactionTriggerConditionMID string = "MID"
)

// prop value enum
func (m *TrailingStopLossOrderTransaction) validateTriggerConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderTransactionTypeTriggerConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrderTransaction) validateTriggerCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerConditionEnum("triggerCondition", "body", m.TriggerCondition); err != nil {
		return err
	}

	return nil
}

var trailingStopLossOrderTransactionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","CLOSE","REOPEN","CLIENT_CONFIGURE","CLIENT_CONFIGURE_REJECT","TRANSFER_FUNDS","TRANSFER_FUNDS_REJECT","MARKET_ORDER","MARKET_ORDER_REJECT","FIXED_PRICE_ORDER","LIMIT_ORDER","LIMIT_ORDER_REJECT","STOP_ORDER","STOP_ORDER_REJECT","MARKET_IF_TOUCHED_ORDER","MARKET_IF_TOUCHED_ORDER_REJECT","TAKE_PROFIT_ORDER","TAKE_PROFIT_ORDER_REJECT","STOP_LOSS_ORDER","STOP_LOSS_ORDER_REJECT","TRAILING_STOP_LOSS_ORDER","TRAILING_STOP_LOSS_ORDER_REJECT","ORDER_FILL","ORDER_CANCEL","ORDER_CANCEL_REJECT","ORDER_CLIENT_EXTENSIONS_MODIFY","ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT","TRADE_CLIENT_EXTENSIONS_MODIFY","TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT","MARGIN_CALL_ENTER","MARGIN_CALL_EXTEND","MARGIN_CALL_EXIT","DELAYED_TRADE_CLOSURE","DAILY_FINANCING","RESET_RESETTABLE_PL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossOrderTransactionTypeTypePropEnum = append(trailingStopLossOrderTransactionTypeTypePropEnum, v)
	}
}

const (

	// TrailingStopLossOrderTransactionTypeCREATE captures enum value "CREATE"
	TrailingStopLossOrderTransactionTypeCREATE string = "CREATE"

	// TrailingStopLossOrderTransactionTypeCLOSE captures enum value "CLOSE"
	TrailingStopLossOrderTransactionTypeCLOSE string = "CLOSE"

	// TrailingStopLossOrderTransactionTypeREOPEN captures enum value "REOPEN"
	TrailingStopLossOrderTransactionTypeREOPEN string = "REOPEN"

	// TrailingStopLossOrderTransactionTypeCLIENTCONFIGURE captures enum value "CLIENT_CONFIGURE"
	TrailingStopLossOrderTransactionTypeCLIENTCONFIGURE string = "CLIENT_CONFIGURE"

	// TrailingStopLossOrderTransactionTypeCLIENTCONFIGUREREJECT captures enum value "CLIENT_CONFIGURE_REJECT"
	TrailingStopLossOrderTransactionTypeCLIENTCONFIGUREREJECT string = "CLIENT_CONFIGURE_REJECT"

	// TrailingStopLossOrderTransactionTypeTRANSFERFUNDS captures enum value "TRANSFER_FUNDS"
	TrailingStopLossOrderTransactionTypeTRANSFERFUNDS string = "TRANSFER_FUNDS"

	// TrailingStopLossOrderTransactionTypeTRANSFERFUNDSREJECT captures enum value "TRANSFER_FUNDS_REJECT"
	TrailingStopLossOrderTransactionTypeTRANSFERFUNDSREJECT string = "TRANSFER_FUNDS_REJECT"

	// TrailingStopLossOrderTransactionTypeMARKETORDER captures enum value "MARKET_ORDER"
	TrailingStopLossOrderTransactionTypeMARKETORDER string = "MARKET_ORDER"

	// TrailingStopLossOrderTransactionTypeMARKETORDERREJECT captures enum value "MARKET_ORDER_REJECT"
	TrailingStopLossOrderTransactionTypeMARKETORDERREJECT string = "MARKET_ORDER_REJECT"

	// TrailingStopLossOrderTransactionTypeFIXEDPRICEORDER captures enum value "FIXED_PRICE_ORDER"
	TrailingStopLossOrderTransactionTypeFIXEDPRICEORDER string = "FIXED_PRICE_ORDER"

	// TrailingStopLossOrderTransactionTypeLIMITORDER captures enum value "LIMIT_ORDER"
	TrailingStopLossOrderTransactionTypeLIMITORDER string = "LIMIT_ORDER"

	// TrailingStopLossOrderTransactionTypeLIMITORDERREJECT captures enum value "LIMIT_ORDER_REJECT"
	TrailingStopLossOrderTransactionTypeLIMITORDERREJECT string = "LIMIT_ORDER_REJECT"

	// TrailingStopLossOrderTransactionTypeSTOPORDER captures enum value "STOP_ORDER"
	TrailingStopLossOrderTransactionTypeSTOPORDER string = "STOP_ORDER"

	// TrailingStopLossOrderTransactionTypeSTOPORDERREJECT captures enum value "STOP_ORDER_REJECT"
	TrailingStopLossOrderTransactionTypeSTOPORDERREJECT string = "STOP_ORDER_REJECT"

	// TrailingStopLossOrderTransactionTypeMARKETIFTOUCHEDORDER captures enum value "MARKET_IF_TOUCHED_ORDER"
	TrailingStopLossOrderTransactionTypeMARKETIFTOUCHEDORDER string = "MARKET_IF_TOUCHED_ORDER"

	// TrailingStopLossOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT captures enum value "MARKET_IF_TOUCHED_ORDER_REJECT"
	TrailingStopLossOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT string = "MARKET_IF_TOUCHED_ORDER_REJECT"

	// TrailingStopLossOrderTransactionTypeTAKEPROFITORDER captures enum value "TAKE_PROFIT_ORDER"
	TrailingStopLossOrderTransactionTypeTAKEPROFITORDER string = "TAKE_PROFIT_ORDER"

	// TrailingStopLossOrderTransactionTypeTAKEPROFITORDERREJECT captures enum value "TAKE_PROFIT_ORDER_REJECT"
	TrailingStopLossOrderTransactionTypeTAKEPROFITORDERREJECT string = "TAKE_PROFIT_ORDER_REJECT"

	// TrailingStopLossOrderTransactionTypeSTOPLOSSORDER captures enum value "STOP_LOSS_ORDER"
	TrailingStopLossOrderTransactionTypeSTOPLOSSORDER string = "STOP_LOSS_ORDER"

	// TrailingStopLossOrderTransactionTypeSTOPLOSSORDERREJECT captures enum value "STOP_LOSS_ORDER_REJECT"
	TrailingStopLossOrderTransactionTypeSTOPLOSSORDERREJECT string = "STOP_LOSS_ORDER_REJECT"

	// TrailingStopLossOrderTransactionTypeTRAILINGSTOPLOSSORDER captures enum value "TRAILING_STOP_LOSS_ORDER"
	TrailingStopLossOrderTransactionTypeTRAILINGSTOPLOSSORDER string = "TRAILING_STOP_LOSS_ORDER"

	// TrailingStopLossOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT captures enum value "TRAILING_STOP_LOSS_ORDER_REJECT"
	TrailingStopLossOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT string = "TRAILING_STOP_LOSS_ORDER_REJECT"

	// TrailingStopLossOrderTransactionTypeORDERFILL captures enum value "ORDER_FILL"
	TrailingStopLossOrderTransactionTypeORDERFILL string = "ORDER_FILL"

	// TrailingStopLossOrderTransactionTypeORDERCANCEL captures enum value "ORDER_CANCEL"
	TrailingStopLossOrderTransactionTypeORDERCANCEL string = "ORDER_CANCEL"

	// TrailingStopLossOrderTransactionTypeORDERCANCELREJECT captures enum value "ORDER_CANCEL_REJECT"
	TrailingStopLossOrderTransactionTypeORDERCANCELREJECT string = "ORDER_CANCEL_REJECT"

	// TrailingStopLossOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY"
	TrailingStopLossOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY string = "ORDER_CLIENT_EXTENSIONS_MODIFY"

	// TrailingStopLossOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	TrailingStopLossOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT string = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// TrailingStopLossOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY"
	TrailingStopLossOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY string = "TRADE_CLIENT_EXTENSIONS_MODIFY"

	// TrailingStopLossOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	TrailingStopLossOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT string = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// TrailingStopLossOrderTransactionTypeMARGINCALLENTER captures enum value "MARGIN_CALL_ENTER"
	TrailingStopLossOrderTransactionTypeMARGINCALLENTER string = "MARGIN_CALL_ENTER"

	// TrailingStopLossOrderTransactionTypeMARGINCALLEXTEND captures enum value "MARGIN_CALL_EXTEND"
	TrailingStopLossOrderTransactionTypeMARGINCALLEXTEND string = "MARGIN_CALL_EXTEND"

	// TrailingStopLossOrderTransactionTypeMARGINCALLEXIT captures enum value "MARGIN_CALL_EXIT"
	TrailingStopLossOrderTransactionTypeMARGINCALLEXIT string = "MARGIN_CALL_EXIT"

	// TrailingStopLossOrderTransactionTypeDELAYEDTRADECLOSURE captures enum value "DELAYED_TRADE_CLOSURE"
	TrailingStopLossOrderTransactionTypeDELAYEDTRADECLOSURE string = "DELAYED_TRADE_CLOSURE"

	// TrailingStopLossOrderTransactionTypeDAILYFINANCING captures enum value "DAILY_FINANCING"
	TrailingStopLossOrderTransactionTypeDAILYFINANCING string = "DAILY_FINANCING"

	// TrailingStopLossOrderTransactionTypeRESETRESETTABLEPL captures enum value "RESET_RESETTABLE_PL"
	TrailingStopLossOrderTransactionTypeRESETRESETTABLEPL string = "RESET_RESETTABLE_PL"
)

// prop value enum
func (m *TrailingStopLossOrderTransaction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossOrderTransactionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossOrderTransaction) validateType(formats strfmt.Registry) error {

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
func (m *TrailingStopLossOrderTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrailingStopLossOrderTransaction) UnmarshalBinary(b []byte) error {
	var res TrailingStopLossOrderTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
