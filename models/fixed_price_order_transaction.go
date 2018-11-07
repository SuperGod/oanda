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

// FixedPriceOrderTransaction A FixedPriceOrderTransaction represents the creation of a Fixed Price Order in the user's account. A Fixed Price Order is an Order that is filled immediately at a specified price.
// swagger:model FixedPriceOrderTransaction
type FixedPriceOrderTransaction struct {

	// The ID of the Account the Transaction was created for.
	AccountID string `json:"accountID,omitempty"`

	// The ID of the "batch" that the Transaction belongs to. Transactions in the same batch are applied to the Account simultaneously.
	BatchID string `json:"batchID,omitempty"`

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The Transaction's Identifier.
	ID string `json:"id,omitempty"`

	// The Fixed Price Order's Instrument.
	Instrument string `json:"instrument,omitempty"`

	// Specification of how Positions in the Account are modified when the Order is filled.
	// Enum: [OPEN_ONLY REDUCE_FIRST REDUCE_ONLY DEFAULT]
	PositionFill string `json:"positionFill,omitempty"`

	// The price specified for the Fixed Price Order. This price is the exact price that the Fixed Price Order will be filled at.
	Price string `json:"price,omitempty"`

	// The reason that the Fixed Price Order was created
	// Enum: [PLATFORM_ACCOUNT_MIGRATION]
	Reason string `json:"reason,omitempty"`

	// The Request ID of the request which generated the transaction.
	RequestID string `json:"requestID,omitempty"`

	// stop loss on fill
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill,omitempty"`

	// take profit on fill
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill,omitempty"`

	// The date/time when the Transaction was created.
	Time string `json:"time,omitempty"`

	// trade client extensions
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions,omitempty"`

	// The state that the trade resulting from the Fixed Price Order should be set to.
	TradeState string `json:"tradeState,omitempty"`

	// trailing stop loss on fill
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill,omitempty"`

	// The Type of the Transaction. Always set to "FIXED_PRICE_ORDER" in a FixedPriceOrderTransaction.
	// Enum: [CREATE CLOSE REOPEN CLIENT_CONFIGURE CLIENT_CONFIGURE_REJECT TRANSFER_FUNDS TRANSFER_FUNDS_REJECT MARKET_ORDER MARKET_ORDER_REJECT FIXED_PRICE_ORDER LIMIT_ORDER LIMIT_ORDER_REJECT STOP_ORDER STOP_ORDER_REJECT MARKET_IF_TOUCHED_ORDER MARKET_IF_TOUCHED_ORDER_REJECT TAKE_PROFIT_ORDER TAKE_PROFIT_ORDER_REJECT STOP_LOSS_ORDER STOP_LOSS_ORDER_REJECT TRAILING_STOP_LOSS_ORDER TRAILING_STOP_LOSS_ORDER_REJECT ORDER_FILL ORDER_CANCEL ORDER_CANCEL_REJECT ORDER_CLIENT_EXTENSIONS_MODIFY ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT TRADE_CLIENT_EXTENSIONS_MODIFY TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT MARGIN_CALL_ENTER MARGIN_CALL_EXTEND MARGIN_CALL_EXIT DELAYED_TRADE_CLOSURE DAILY_FINANCING RESET_RESETTABLE_PL]
	Type string `json:"type,omitempty"`

	// The quantity requested to be filled by the Fixed Price Order. A posititive number of units results in a long Order, and a negative number of units results in a short Order.
	Units string `json:"units,omitempty"`

	// The ID of the user that initiated the creation of the Transaction.
	UserID int64 `json:"userID,omitempty"`
}

// Validate validates this fixed price order transaction
func (m *FixedPriceOrderTransaction) Validate(formats strfmt.Registry) error {
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

func (m *FixedPriceOrderTransaction) validateClientExtensions(formats strfmt.Registry) error {

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

var fixedPriceOrderTransactionTypePositionFillPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN_ONLY","REDUCE_FIRST","REDUCE_ONLY","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fixedPriceOrderTransactionTypePositionFillPropEnum = append(fixedPriceOrderTransactionTypePositionFillPropEnum, v)
	}
}

const (

	// FixedPriceOrderTransactionPositionFillOPENONLY captures enum value "OPEN_ONLY"
	FixedPriceOrderTransactionPositionFillOPENONLY string = "OPEN_ONLY"

	// FixedPriceOrderTransactionPositionFillREDUCEFIRST captures enum value "REDUCE_FIRST"
	FixedPriceOrderTransactionPositionFillREDUCEFIRST string = "REDUCE_FIRST"

	// FixedPriceOrderTransactionPositionFillREDUCEONLY captures enum value "REDUCE_ONLY"
	FixedPriceOrderTransactionPositionFillREDUCEONLY string = "REDUCE_ONLY"

	// FixedPriceOrderTransactionPositionFillDEFAULT captures enum value "DEFAULT"
	FixedPriceOrderTransactionPositionFillDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *FixedPriceOrderTransaction) validatePositionFillEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fixedPriceOrderTransactionTypePositionFillPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FixedPriceOrderTransaction) validatePositionFill(formats strfmt.Registry) error {

	if swag.IsZero(m.PositionFill) { // not required
		return nil
	}

	// value enum
	if err := m.validatePositionFillEnum("positionFill", "body", m.PositionFill); err != nil {
		return err
	}

	return nil
}

var fixedPriceOrderTransactionTypeReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PLATFORM_ACCOUNT_MIGRATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fixedPriceOrderTransactionTypeReasonPropEnum = append(fixedPriceOrderTransactionTypeReasonPropEnum, v)
	}
}

const (

	// FixedPriceOrderTransactionReasonPLATFORMACCOUNTMIGRATION captures enum value "PLATFORM_ACCOUNT_MIGRATION"
	FixedPriceOrderTransactionReasonPLATFORMACCOUNTMIGRATION string = "PLATFORM_ACCOUNT_MIGRATION"
)

// prop value enum
func (m *FixedPriceOrderTransaction) validateReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fixedPriceOrderTransactionTypeReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FixedPriceOrderTransaction) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	// value enum
	if err := m.validateReasonEnum("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *FixedPriceOrderTransaction) validateStopLossOnFill(formats strfmt.Registry) error {

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

func (m *FixedPriceOrderTransaction) validateTakeProfitOnFill(formats strfmt.Registry) error {

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

func (m *FixedPriceOrderTransaction) validateTradeClientExtensions(formats strfmt.Registry) error {

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

func (m *FixedPriceOrderTransaction) validateTrailingStopLossOnFill(formats strfmt.Registry) error {

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

var fixedPriceOrderTransactionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","CLOSE","REOPEN","CLIENT_CONFIGURE","CLIENT_CONFIGURE_REJECT","TRANSFER_FUNDS","TRANSFER_FUNDS_REJECT","MARKET_ORDER","MARKET_ORDER_REJECT","FIXED_PRICE_ORDER","LIMIT_ORDER","LIMIT_ORDER_REJECT","STOP_ORDER","STOP_ORDER_REJECT","MARKET_IF_TOUCHED_ORDER","MARKET_IF_TOUCHED_ORDER_REJECT","TAKE_PROFIT_ORDER","TAKE_PROFIT_ORDER_REJECT","STOP_LOSS_ORDER","STOP_LOSS_ORDER_REJECT","TRAILING_STOP_LOSS_ORDER","TRAILING_STOP_LOSS_ORDER_REJECT","ORDER_FILL","ORDER_CANCEL","ORDER_CANCEL_REJECT","ORDER_CLIENT_EXTENSIONS_MODIFY","ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT","TRADE_CLIENT_EXTENSIONS_MODIFY","TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT","MARGIN_CALL_ENTER","MARGIN_CALL_EXTEND","MARGIN_CALL_EXIT","DELAYED_TRADE_CLOSURE","DAILY_FINANCING","RESET_RESETTABLE_PL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fixedPriceOrderTransactionTypeTypePropEnum = append(fixedPriceOrderTransactionTypeTypePropEnum, v)
	}
}

const (

	// FixedPriceOrderTransactionTypeCREATE captures enum value "CREATE"
	FixedPriceOrderTransactionTypeCREATE string = "CREATE"

	// FixedPriceOrderTransactionTypeCLOSE captures enum value "CLOSE"
	FixedPriceOrderTransactionTypeCLOSE string = "CLOSE"

	// FixedPriceOrderTransactionTypeREOPEN captures enum value "REOPEN"
	FixedPriceOrderTransactionTypeREOPEN string = "REOPEN"

	// FixedPriceOrderTransactionTypeCLIENTCONFIGURE captures enum value "CLIENT_CONFIGURE"
	FixedPriceOrderTransactionTypeCLIENTCONFIGURE string = "CLIENT_CONFIGURE"

	// FixedPriceOrderTransactionTypeCLIENTCONFIGUREREJECT captures enum value "CLIENT_CONFIGURE_REJECT"
	FixedPriceOrderTransactionTypeCLIENTCONFIGUREREJECT string = "CLIENT_CONFIGURE_REJECT"

	// FixedPriceOrderTransactionTypeTRANSFERFUNDS captures enum value "TRANSFER_FUNDS"
	FixedPriceOrderTransactionTypeTRANSFERFUNDS string = "TRANSFER_FUNDS"

	// FixedPriceOrderTransactionTypeTRANSFERFUNDSREJECT captures enum value "TRANSFER_FUNDS_REJECT"
	FixedPriceOrderTransactionTypeTRANSFERFUNDSREJECT string = "TRANSFER_FUNDS_REJECT"

	// FixedPriceOrderTransactionTypeMARKETORDER captures enum value "MARKET_ORDER"
	FixedPriceOrderTransactionTypeMARKETORDER string = "MARKET_ORDER"

	// FixedPriceOrderTransactionTypeMARKETORDERREJECT captures enum value "MARKET_ORDER_REJECT"
	FixedPriceOrderTransactionTypeMARKETORDERREJECT string = "MARKET_ORDER_REJECT"

	// FixedPriceOrderTransactionTypeFIXEDPRICEORDER captures enum value "FIXED_PRICE_ORDER"
	FixedPriceOrderTransactionTypeFIXEDPRICEORDER string = "FIXED_PRICE_ORDER"

	// FixedPriceOrderTransactionTypeLIMITORDER captures enum value "LIMIT_ORDER"
	FixedPriceOrderTransactionTypeLIMITORDER string = "LIMIT_ORDER"

	// FixedPriceOrderTransactionTypeLIMITORDERREJECT captures enum value "LIMIT_ORDER_REJECT"
	FixedPriceOrderTransactionTypeLIMITORDERREJECT string = "LIMIT_ORDER_REJECT"

	// FixedPriceOrderTransactionTypeSTOPORDER captures enum value "STOP_ORDER"
	FixedPriceOrderTransactionTypeSTOPORDER string = "STOP_ORDER"

	// FixedPriceOrderTransactionTypeSTOPORDERREJECT captures enum value "STOP_ORDER_REJECT"
	FixedPriceOrderTransactionTypeSTOPORDERREJECT string = "STOP_ORDER_REJECT"

	// FixedPriceOrderTransactionTypeMARKETIFTOUCHEDORDER captures enum value "MARKET_IF_TOUCHED_ORDER"
	FixedPriceOrderTransactionTypeMARKETIFTOUCHEDORDER string = "MARKET_IF_TOUCHED_ORDER"

	// FixedPriceOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT captures enum value "MARKET_IF_TOUCHED_ORDER_REJECT"
	FixedPriceOrderTransactionTypeMARKETIFTOUCHEDORDERREJECT string = "MARKET_IF_TOUCHED_ORDER_REJECT"

	// FixedPriceOrderTransactionTypeTAKEPROFITORDER captures enum value "TAKE_PROFIT_ORDER"
	FixedPriceOrderTransactionTypeTAKEPROFITORDER string = "TAKE_PROFIT_ORDER"

	// FixedPriceOrderTransactionTypeTAKEPROFITORDERREJECT captures enum value "TAKE_PROFIT_ORDER_REJECT"
	FixedPriceOrderTransactionTypeTAKEPROFITORDERREJECT string = "TAKE_PROFIT_ORDER_REJECT"

	// FixedPriceOrderTransactionTypeSTOPLOSSORDER captures enum value "STOP_LOSS_ORDER"
	FixedPriceOrderTransactionTypeSTOPLOSSORDER string = "STOP_LOSS_ORDER"

	// FixedPriceOrderTransactionTypeSTOPLOSSORDERREJECT captures enum value "STOP_LOSS_ORDER_REJECT"
	FixedPriceOrderTransactionTypeSTOPLOSSORDERREJECT string = "STOP_LOSS_ORDER_REJECT"

	// FixedPriceOrderTransactionTypeTRAILINGSTOPLOSSORDER captures enum value "TRAILING_STOP_LOSS_ORDER"
	FixedPriceOrderTransactionTypeTRAILINGSTOPLOSSORDER string = "TRAILING_STOP_LOSS_ORDER"

	// FixedPriceOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT captures enum value "TRAILING_STOP_LOSS_ORDER_REJECT"
	FixedPriceOrderTransactionTypeTRAILINGSTOPLOSSORDERREJECT string = "TRAILING_STOP_LOSS_ORDER_REJECT"

	// FixedPriceOrderTransactionTypeORDERFILL captures enum value "ORDER_FILL"
	FixedPriceOrderTransactionTypeORDERFILL string = "ORDER_FILL"

	// FixedPriceOrderTransactionTypeORDERCANCEL captures enum value "ORDER_CANCEL"
	FixedPriceOrderTransactionTypeORDERCANCEL string = "ORDER_CANCEL"

	// FixedPriceOrderTransactionTypeORDERCANCELREJECT captures enum value "ORDER_CANCEL_REJECT"
	FixedPriceOrderTransactionTypeORDERCANCELREJECT string = "ORDER_CANCEL_REJECT"

	// FixedPriceOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY"
	FixedPriceOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFY string = "ORDER_CLIENT_EXTENSIONS_MODIFY"

	// FixedPriceOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	FixedPriceOrderTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT string = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// FixedPriceOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY"
	FixedPriceOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFY string = "TRADE_CLIENT_EXTENSIONS_MODIFY"

	// FixedPriceOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	FixedPriceOrderTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT string = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// FixedPriceOrderTransactionTypeMARGINCALLENTER captures enum value "MARGIN_CALL_ENTER"
	FixedPriceOrderTransactionTypeMARGINCALLENTER string = "MARGIN_CALL_ENTER"

	// FixedPriceOrderTransactionTypeMARGINCALLEXTEND captures enum value "MARGIN_CALL_EXTEND"
	FixedPriceOrderTransactionTypeMARGINCALLEXTEND string = "MARGIN_CALL_EXTEND"

	// FixedPriceOrderTransactionTypeMARGINCALLEXIT captures enum value "MARGIN_CALL_EXIT"
	FixedPriceOrderTransactionTypeMARGINCALLEXIT string = "MARGIN_CALL_EXIT"

	// FixedPriceOrderTransactionTypeDELAYEDTRADECLOSURE captures enum value "DELAYED_TRADE_CLOSURE"
	FixedPriceOrderTransactionTypeDELAYEDTRADECLOSURE string = "DELAYED_TRADE_CLOSURE"

	// FixedPriceOrderTransactionTypeDAILYFINANCING captures enum value "DAILY_FINANCING"
	FixedPriceOrderTransactionTypeDAILYFINANCING string = "DAILY_FINANCING"

	// FixedPriceOrderTransactionTypeRESETRESETTABLEPL captures enum value "RESET_RESETTABLE_PL"
	FixedPriceOrderTransactionTypeRESETRESETTABLEPL string = "RESET_RESETTABLE_PL"
)

// prop value enum
func (m *FixedPriceOrderTransaction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fixedPriceOrderTransactionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FixedPriceOrderTransaction) validateType(formats strfmt.Registry) error {

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
func (m *FixedPriceOrderTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FixedPriceOrderTransaction) UnmarshalBinary(b []byte) error {
	var res FixedPriceOrderTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
