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

// MarketIfTouchedOrderTransaction A MarketIfTouchedOrderTransaction represents the creation of a MarketIfTouched Order in the user's Account.
// swagger:model MarketIfTouchedOrderTransaction
type MarketIfTouchedOrderTransaction struct {

	// The ID of the Account the Transaction was created for.
	AccountID string `json:"accountID,omitempty"`

	// The ID of the "batch" that the Transaction belongs to. Transactions in the same batch are applied to the Account simultaneously.
	BatchID string `json:"batchID,omitempty"`

	// The ID of the Transaction that cancels the replaced Order (only provided if this Order replaces an existing Order).
	CancellingTransactionID string `json:"cancellingTransactionID,omitempty"`

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The date/time when the MarketIfTouched Order will be cancelled if its timeInForce is "GTD".
	GtdTime string `json:"gtdTime,omitempty"`

	// The Transaction's Identifier.
	ID string `json:"id,omitempty"`

	// The MarketIfTouched Order's Instrument.
	Instrument string `json:"instrument,omitempty"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	// Enum: [OPEN_ONLY REDUCE_FIRST REDUCE_ONLY DEFAULT]
	PositionFill string `json:"positionFill,omitempty"`

	// The price threshold specified for the MarketIfTouched Order. The MarketIfTouched Order will only be filled by a market price that crosses this price from the direction of the market price at the time when the Order was created (the initialMarketPrice). Depending on the value of the Order's price and initialMarketPrice, the MarketIfTouchedOrder will behave like a Limit or a Stop Order.
	Price string `json:"price,omitempty"`

	// The worst market price that may be used to fill this MarketIfTouched Order.
	PriceBound string `json:"priceBound,omitempty"`

	// The reason that the Market-if-touched Order was initiated
	// Enum: [CLIENT_ORDER REPLACEMENT]
	Reason string `json:"reason,omitempty"`

	// The ID of the Order that this Order replaces (only provided if this Order replaces an existing Order).
	ReplacesOrderID string `json:"replacesOrderID,omitempty"`

	// The Request ID of the request which generated the transaction.
	RequestID string `json:"requestID,omitempty"`

	// stop loss on fill
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill,omitempty"`

	// take profit on fill
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill,omitempty"`

	// The date/time when the Transaction was created.
	Time string `json:"time,omitempty"`

	// The time-in-force requested for the MarketIfTouched Order. Restricted to "GTC", "GFD" and "GTD" for MarketIfTouched Orders.
	// Enum: [GTC GTD GFD FOK IOC]
	TimeInForce string `json:"timeInForce,omitempty"`

	// trade client extensions
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions,omitempty"`

	// trailing stop loss on fill
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill,omitempty"`

	// Specification of which price component should be used when determining if an Order should be triggered and filled. This allows Orders to be triggered based on the bid, ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy) price depending on the desired behaviour. Orders are always filled using their default price component.
	// This feature is only provided through the REST API. Clients who choose to specify a non-default trigger condition will not see it reflected in any of OANDA's proprietary or partner trading platforms, their transaction history or their account statements. OANDA platforms always assume that an Order's trigger condition is set to the default value when indicating the distance from an Order's trigger price, and will always provide the default trigger condition when creating or modifying an Order.
	// A special restriction applies when creating a guaranteed Stop Loss Order. In this case the TriggerCondition value must either be "DEFAULT", or the "natural" trigger side "DEFAULT" results in. So for a Stop Loss Order for a long trade valid values are "DEFAULT" and "BID", and for short trades "DEFAULT" and "ASK" are valid.
	// Enum: [DEFAULT INVERSE BID ASK MID]
	TriggerCondition string `json:"triggerCondition,omitempty"`

	// The Type of the Transaction. Always set to "MARKET_IF_TOUCHED_ORDER" in a MarketIfTouchedOrderTransaction.
	// Enum: [CREATE CLOSE REOPEN CLIENT_CONFIGURE CLIENT_CONFIGURE_REJECT TRANSFER_FUNDS TRANSFER_FUNDS_REJECT MARKET_ORDER MARKET_ORDER_REJECT FIXED_PRICE_ORDER LIMIT_ORDER LIMIT_ORDER_REJECT STOP_ORDER STOP_ORDER_REJECT MARKET_IF_TOUCHED_ORDER MARKET_IF_TOUCHED_ORDER_REJECT TAKE_PROFIT_ORDER TAKE_PROFIT_ORDER_REJECT STOP_LOSS_ORDER STOP_LOSS_ORDER_REJECT TRAILING_STOP_LOSS_ORDER TRAILING_STOP_LOSS_ORDER_REJECT ORDER_FILL ORDER_CANCEL ORDER_CANCEL_REJECT ORDER_CLIENT_EXTENSIONS_MODIFY ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT TRADE_CLIENT_EXTENSIONS_MODIFY TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT MARGIN_CALL_ENTER MARGIN_CALL_EXTEND MARGIN_CALL_EXIT DELAYED_TRADE_CLOSURE DAILY_FINANCING RESET_RESETTABLE_PL]
	Type string `json:"type,omitempty"`

	// The quantity requested to be filled by the MarketIfTouched Order. A posititive number of units results in a long Order, and a negative number of units results in a short Order.
	Units string `json:"units,omitempty"`

	// The ID of the user that initiated the creation of the Transaction.
	UserID int64 `json:"userID,omitempty"`
}

// Validate validates this market if touched order transaction
func (m *MarketIfTouchedOrderTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositionFill(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
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

	if err := m.validateTrailingStopLossOnFill(formats); err != nil {
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

func (m *MarketIfTouchedOrderTransaction) validateClientExtensions(formats strfmt.Registry) error {

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

var marketIfTouchedOrderTransactionTypePositionFillPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN_ONLY","REDUCE_FIRST","REDUCE_ONLY","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketIfTouchedOrderTransactionTypePositionFillPropEnum = append(marketIfTouchedOrderTransactionTypePositionFillPropEnum, v)
	}
}

const (

	// MarketIfTouchedOrderTransactionPositionFillOPENONLY captures enum value "OPEN_ONLY"
	MarketIfTouchedOrderTransactionPositionFillOPENONLY string = "OPEN_ONLY"

	// MarketIfTouchedOrderTransactionPositionFillREDUCEFIRST captures enum value "REDUCE_FIRST"
	MarketIfTouchedOrderTransactionPositionFillREDUCEFIRST string = "REDUCE_FIRST"

	// MarketIfTouchedOrderTransactionPositionFillREDUCEONLY captures enum value "REDUCE_ONLY"
	MarketIfTouchedOrderTransactionPositionFillREDUCEONLY string = "REDUCE_ONLY"

	// MarketIfTouchedOrderTransactionPositionFillDEFAULT captures enum value "DEFAULT"
	MarketIfTouchedOrderTransactionPositionFillDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *MarketIfTouchedOrderTransaction) validatePositionFillEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketIfTouchedOrderTransactionTypePositionFillPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketIfTouchedOrderTransaction) validatePositionFill(formats strfmt.Registry) error {

	if swag.IsZero(m.PositionFill) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionFillEnum("positionFill", "body", m.PositionFill); err != nil {
		return err
	}

	return nil
}

var marketIfTouchedOrderTransactionTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLIENT_ORDER","REPLACEMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketIfTouchedOrderTransactionTypeReasonPropEnum = append(marketIfTouchedOrderTransactionTypeReasonPropEnum, v)
	}
}

const (

	// MarketIfTouchedOrderTransactionReasonCLIENTORDER captures enum value "CLIENT_ORDER"
	MarketIfTouchedOrderTransactionReasonCLIENTORDER string = "CLIENT_ORDER"

	// MarketIfTouchedOrderTransactionReasonREPLACEMENT captures enum value "REPLACEMENT"
	MarketIfTouchedOrderTransactionReasonREPLACEMENT string = "REPLACEMENT"
)

// prop value enum
func (m *MarketIfTouchedOrderTransaction) validateReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketIfTouchedOrderTransactionTypeReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketIfTouchedOrderTransaction) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	// value enum
	if err := m.validateReasonEnum("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *MarketIfTouchedOrderTransaction) validateStopLossOnFill(formats strfmt.Registry) error {

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

func (m *MarketIfTouchedOrderTransaction) validateTakeProfitOnFill(formats strfmt.Registry) error {

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

var marketIfTouchedOrderTransactionTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketIfTouchedOrderTransactionTypeTimeInForcePropEnum = append(marketIfTouchedOrderTransactionTypeTimeInForcePropEnum, v)
	}
}

const (

	// MarketIfTouchedOrderTransactionTimeInForceGTC captures enum value "GTC"
	MarketIfTouchedOrderTransactionTimeInForceGTC string = "GTC"

	// MarketIfTouchedOrderTransactionTimeInForceGTD captures enum value "GTD"
	MarketIfTouchedOrderTransactionTimeInForceGTD string = "GTD"

	// MarketIfTouchedOrderTransactionTimeInForceGFD captures enum value "GFD"
	MarketIfTouchedOrderTransactionTimeInForceGFD string = "GFD"

	// MarketIfTouchedOrderTransactionTimeInForceFOK captures enum value "FOK"
	MarketIfTouchedOrderTransactionTimeInForceFOK string = "FOK"

	// MarketIfTouchedOrderTransactionTimeInForceIOC captures enum value "IOC"
	MarketIfTouchedOrderTransactionTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *MarketIfTouchedOrderTransaction) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketIfTouchedOrderTransactionTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketIfTouchedOrderTransaction) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

func (m *MarketIfTouchedOrderTransaction) validateTradeClientExtensions(formats strfmt.Registry) error {

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

func (m *MarketIfTouchedOrderTransaction) validateTrailingStopLossOnFill(formats strfmt.Registry) error {

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

var marketIfTouchedOrderTransactionTypeTriggerConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","INVERSE","BID","ASK","MID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketIfTouchedOrderTransactionTypeTriggerConditionPropEnum = append(marketIfTouchedOrderTransactionTypeTriggerConditionPropEnum, v)
	}
}

const (

	// MarketIfTouchedOrderTransactionTriggerConditionDEFAULT captures enum value "DEFAULT"
	MarketIfTouchedOrderTransactionTriggerConditionDEFAULT string = "DEFAULT"

	// MarketIfTouchedOrderTransactionTriggerConditionINVERSE captures enum value "INVERSE"
	MarketIfTouchedOrderTransactionTriggerConditionINVERSE string = "INVERSE"

	// MarketIfTouchedOrderTransactionTriggerConditionBID captures enum value "BID"
	MarketIfTouchedOrderTransactionTriggerConditionBID string = "BID"

	// MarketIfTouchedOrderTransactionTriggerConditionASK captures enum value "ASK"
	MarketIfTouchedOrderTransactionTriggerConditionASK string = "ASK"

	// MarketIfTouchedOrderTransactionTriggerConditionMID captures enum value "MID"
	MarketIfTouchedOrderTransactionTriggerConditionMID string = "MID"
)

// prop value enum
func (m *MarketIfTouchedOrderTransaction) validateTriggerConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketIfTouchedOrderTransactionTypeTriggerConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketIfTouchedOrderTransaction) validateTriggerCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerCondition) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerConditionEnum("triggerCondition", "body", m.TriggerCondition); err != nil {
		return err
	}

	return nil
}

var marketIfTouchedOrderTransactionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","CLOSE","REOPEN","CLIENT_CONFIGURE","CLIENT_CONFIGURE_REJECT","TRANSFER_FUNDS","TRANSFER_FUNDS_REJECT","MARKET_ORDER","MARKET_ORDER_REJECT","FIXED_PRICE_ORDER","LIMIT_ORDER","LIMIT_ORDER_REJECT","STOP_ORDER","STOP_ORDER_REJECT","MARKET_IF_TOUCHED_ORDER","MARKET_IF_TOUCHED_ORDER_REJECT","TAKE_PROFIT_ORDER","TAKE_PROFIT_ORDER_REJECT","STOP_LOSS_ORDER","STOP_LOSS_ORDER_REJECT","TRAILING_STOP_LOSS_ORDER","TRAILING_STOP_LOSS_ORDER_REJECT","ORDER_FILL","ORDER_CANCEL","ORDER_CANCEL_REJECT","ORDER_CLIENT_EXTENSIONS_MODIFY","ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT","TRADE_CLIENT_EXTENSIONS_MODIFY","TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT","MARGIN_CALL_ENTER","MARGIN_CALL_EXTEND","MARGIN_CALL_EXIT","DELAYED_TRADE_CLOSURE","DAILY_FINANCING","RESET_RESETTABLE_PL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketIfTouchedOrderTransactionTypeTypePropEnum = append(marketIfTouchedOrderTransactionTypeTypePropEnum, v)
	}
}

const (

	// MarketIfTouchedOrderTransactionTypeCREATE captures enum value "CREATE"
	MarketIfTouchedOrderTransactionTypeCREATE string = "CREATE"

	// MarketIfTouchedOrderTransactionTypeCLOSE captures enum value "CLOSE"
	MarketIfTouchedOrderTransactionTypeCLOSE string = "CLOSE"

	// MarketIfTouchedOrderTransactionTypeREOPEN captures enum value "REOPEN"
	MarketIfTouchedOrderTransactionTypeREOPEN string = "REOPEN"

	// MarketIfTouchedOrderTransactionTypeCLIENTCONFIGURE captures enum value "CLIENT_CONFIGURE"
	MarketIfTouchedOrderTransactionTypeCLIENTCONFIGURE string = "CLIENT_CONFIGURE"

	// MarketIfTouchedOrderTransactionTypeCLIENTCONFIGUREREJECT captures enum value "CLIENT_CONFIGURE_REJECT"
	MarketIfTouchedOrderTransactionTypeCLIENTCONFIGUREREJECT string = "CLIENT_CONFIGURE_REJECT"

	// MarketIfTouchedOrderTransactionTypeTRANSFERFUNDS captures enum value "TRANSFER_FUNDS"
	MarketIfTouchedOrderTransactionTypeTRANSFERFUNDS string = "TRANSFER_FUNDS"

	// MarketIfTouchedOrderTransactionTypeTRANSFERFUNDSREJECT captures enum value "TRANSFER_FUNDS_REJECT"
	MarketIfTouchedOrderTransactionTypeTRANSFERFUNDSREJECT string = "TRANSFER_FUNDS_REJECT"

	// MarketIfTouchedOrderTransactionTypeMARKETORDER captures enum value "MARKET_ORDER"
	MarketIfTouchedOrderTransactionTypeMARKETORDER string = "MARKET_ORDER"

	// MarketIfTouchedOrderTransactionTypeMARKETORDERREJECT captures enum value "MARKET_ORDER_REJECT"
	MarketIfTouchedOrderTransactionTypeMARKETORDERREJECT string = "MARKET_ORDER_REJECT"

	// MarketIfTouchedOrderTransactionTypeFIXEDPRICEORDER captures enum value "FIXED_PRICE_ORDER"
	MarketIfTouchedOrderTransactionTypeFIXEDPRICEORDER string = "FIXED_PRICE_ORDER"

	// MarketIfTouchedOrderTransactionTypeLIMITORDER captures enum value "LIMIT_ORDER"
	MarketIfTouchedOrderTransactionTypeLIMITORDER string = "LIMIT_ORDER"

	// MarketIfTouchedOrderTransactionTypeLIMITORDERREJECT captures enum value "LIMIT_ORDER_REJECT"
	MarketIfTouchedOrderTransactionTypeLIMITORDERREJECT string = "LIMIT_ORDER_REJECT"

	// MarketIfTouchedOrderTransactionTypeSTOPORDER captures enum value "STOP_ORDER"
	MarketIfTouchedOrderTransactionTypeSTOPORDER string = "STOP_ORDER"

	// MarketIfTouchedOrderTransactionTypeSTOPORDERREJECT captures enum value "STOP_ORDER_REJECT"
	MarketIfTouchedOrderTransactionTypeSTOPORDERREJECT string = "STOP_ORDER_REJECT"

	// MarketIfTouchedOrderTransactionTypeMARKETIFTOUCHEDORDER captures enum value "MARKET_IF_TOUCHED_ORDER"
	MarketIfTouchedOrderTransactionTypeMARKETIFTOUCHEDORDER string = "MARKET_IF_TOUCHED_ORDER"

	// MarketIfTouchedOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT captures enum value "MARKET_IF_TOUCHED_ORDER_REJECT"
	MarketIfTouchedOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT string = "MARKET_IF_TOUCHED_ORDER_REJECT"

	// MarketIfTouchedOrderTransactionTypeTAKEPROFITORDER captures enum value "TAKE_PROFIT_ORDER"
	MarketIfTouchedOrderTransactionTypeTAKEPROFITORDER string = "TAKE_PROFIT_ORDER"

	// MarketIfTouchedOrderTransactionTypeTAKEPROFITORDERREJECT captures enum value "TAKE_PROFIT_ORDER_REJECT"
	MarketIfTouchedOrderTransactionTypeTAKEPROFITORDERREJECT string = "TAKE_PROFIT_ORDER_REJECT"

	// MarketIfTouchedOrderTransactionTypeSTOPLOSSORDER captures enum value "STOP_LOSS_ORDER"
	MarketIfTouchedOrderTransactionTypeSTOPLOSSORDER string = "STOP_LOSS_ORDER"

	// MarketIfTouchedOrderTransactionTypeSTOPLOSSORDERREJECT captures enum value "STOP_LOSS_ORDER_REJECT"
	MarketIfTouchedOrderTransactionTypeSTOPLOSSORDERREJECT string = "STOP_LOSS_ORDER_REJECT"

	// MarketIfTouchedOrderTransactionTypeTRAILINGSTOPLOSSORDER captures enum value "TRAILING_STOP_LOSS_ORDER"
	MarketIfTouchedOrderTransactionTypeTRAILINGSTOPLOSSORDER string = "TRAILING_STOP_LOSS_ORDER"

	// MarketIfTouchedOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT captures enum value "TRAILING_STOP_LOSS_ORDER_REJECT"
	MarketIfTouchedOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT string = "TRAILING_STOP_LOSS_ORDER_REJECT"

	// MarketIfTouchedOrderTransactionTypeORDERFILL captures enum value "ORDER_FILL"
	MarketIfTouchedOrderTransactionTypeORDERFILL string = "ORDER_FILL"

	// MarketIfTouchedOrderTransactionTypeORDERCANCEL captures enum value "ORDER_CANCEL"
	MarketIfTouchedOrderTransactionTypeORDERCANCEL string = "ORDER_CANCEL"

	// MarketIfTouchedOrderTransactionTypeORDERCANCELREJECT captures enum value "ORDER_CANCEL_REJECT"
	MarketIfTouchedOrderTransactionTypeORDERCANCELREJECT string = "ORDER_CANCEL_REJECT"

	// MarketIfTouchedOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY"
	MarketIfTouchedOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY string = "ORDER_CLIENT_EXTENSIONS_MODIFY"

	// MarketIfTouchedOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	MarketIfTouchedOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT string = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// MarketIfTouchedOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY"
	MarketIfTouchedOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY string = "TRADE_CLIENT_EXTENSIONS_MODIFY"

	// MarketIfTouchedOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	MarketIfTouchedOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT string = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// MarketIfTouchedOrderTransactionTypeMARGINCALLENTER captures enum value "MARGIN_CALL_ENTER"
	MarketIfTouchedOrderTransactionTypeMARGINCALLENTER string = "MARGIN_CALL_ENTER"

	// MarketIfTouchedOrderTransactionTypeMARGINCALLEXTEND captures enum value "MARGIN_CALL_EXTEND"
	MarketIfTouchedOrderTransactionTypeMARGINCALLEXTEND string = "MARGIN_CALL_EXTEND"

	// MarketIfTouchedOrderTransactionTypeMARGINCALLEXIT captures enum value "MARGIN_CALL_EXIT"
	MarketIfTouchedOrderTransactionTypeMARGINCALLEXIT string = "MARGIN_CALL_EXIT"

	// MarketIfTouchedOrderTransactionTypeDELAYEDTRADECLOSURE captures enum value "DELAYED_TRADE_CLOSURE"
	MarketIfTouchedOrderTransactionTypeDELAYEDTRADECLOSURE string = "DELAYED_TRADE_CLOSURE"

	// MarketIfTouchedOrderTransactionTypeDAILYFINANCING captures enum value "DAILY_FINANCING"
	MarketIfTouchedOrderTransactionTypeDAILYFINANCING string = "DAILY_FINANCING"

	// MarketIfTouchedOrderTransactionTypeRESETRESETTABLEPL captures enum value "RESET_RESETTABLE_PL"
	MarketIfTouchedOrderTransactionTypeRESETRESETTABLEPL string = "RESET_RESETTABLE_PL"
)

// prop value enum
func (m *MarketIfTouchedOrderTransaction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketIfTouchedOrderTransactionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketIfTouchedOrderTransaction) validateType(formats strfmt.Registry) error {

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
func (m *MarketIfTouchedOrderTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketIfTouchedOrderTransaction) UnmarshalBinary(b []byte) error {
	var res MarketIfTouchedOrderTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}