// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DailyFinancingTransaction A DailyFinancingTransaction represents the daily payment/collection of financing for an Account.
// swagger:model DailyFinancingTransaction
type DailyFinancingTransaction struct {

	// The Account's balance after daily financing.
	AccountBalance string `json:"accountBalance,omitempty"`

	// The account financing mode at the time of the daily financing.
	// Enum: [NO_FINANCING SECOND_BY_SECOND DAILY]
	AccountFinancingMode string `json:"accountFinancingMode,omitempty"`

	// The ID of the Account the Transaction was created for.
	AccountID string `json:"accountID,omitempty"`

	// The ID of the "batch" that the Transaction belongs to. Transactions in the same batch are applied to the Account simultaneously.
	BatchID string `json:"batchID,omitempty"`

	// The amount of financing paid/collected for the Account.
	Financing string `json:"financing,omitempty"`

	// The Transaction's Identifier.
	ID string `json:"id,omitempty"`

	// The financing paid/collected for each Position in the Account.
	PositionFinancings []*PositionFinancing `json:"positionFinancings"`

	// The Request ID of the request which generated the transaction.
	RequestID string `json:"requestID,omitempty"`

	// The date/time when the Transaction was created.
	Time string `json:"time,omitempty"`

	// The Type of the Transaction. Always set to "DAILY_FINANCING" for a DailyFinancingTransaction.
	// Enum: [CREATE CLOSE REOPEN CLIENT_CONFIGURE CLIENT_CONFIGURE_REJECT TRANSFER_FUNDS TRANSFER_FUNDS_REJECT MARKET_ORDER MARKET_ORDER_REJECT FIXED_PRICE_ORDER LIMIT_ORDER LIMIT_ORDER_REJECT STOP_ORDER STOP_ORDER_REJECT MARKET_IF_TOUCHED_ORDER MARKET_IF_TOUCHED_ORDER_REJECT TAKE_PROFIT_ORDER TAKE_PROFIT_ORDER_REJECT STOP_LOSS_ORDER STOP_LOSS_ORDER_REJECT TRAILING_STOP_LOSS_ORDER TRAILING_STOP_LOSS_ORDER_REJECT ORDER_FILL ORDER_CANCEL ORDER_CANCEL_REJECT ORDER_CLIENT_EXTENSIONS_MODIFY ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT TRADE_CLIENT_EXTENSIONS_MODIFY TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT MARGIN_CALL_ENTER MARGIN_CALL_EXTEND MARGIN_CALL_EXIT DELAYED_TRADE_CLOSURE DAILY_FINANCING RESET_RESETTABLE_PL]
	Type string `json:"type,omitempty"`

	// The ID of the user that initiated the creation of the Transaction.
	UserID int64 `json:"userID,omitempty"`
}

// Validate validates this daily financing transaction
func (m *DailyFinancingTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountFinancingMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositionFinancings(formats); err != nil {
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

var dailyFinancingTransactionTypeAccountFinancingModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NO_FINANCING","SECOND_BY_SECOND","DAILY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dailyFinancingTransactionTypeAccountFinancingModePropEnum = append(dailyFinancingTransactionTypeAccountFinancingModePropEnum, v)
	}
}

const (

	// DailyFinancingTransactionAccountFinancingModeNOFINANCING captures enum value "NO_FINANCING"
	DailyFinancingTransactionAccountFinancingModeNOFINANCING string = "NO_FINANCING"

	// DailyFinancingTransactionAccountFinancingModeSECONDBYSECOND captures enum value "SECOND_BY_SECOND"
	DailyFinancingTransactionAccountFinancingModeSECONDBYSECOND string = "SECOND_BY_SECOND"

	// DailyFinancingTransactionAccountFinancingModeDAILY captures enum value "DAILY"
	DailyFinancingTransactionAccountFinancingModeDAILY string = "DAILY"
)

// prop value enum
func (m *DailyFinancingTransaction) validateAccountFinancingModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dailyFinancingTransactionTypeAccountFinancingModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DailyFinancingTransaction) validateAccountFinancingMode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountFinancingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccountFinancingModeEnum("accountFinancingMode", "body", m.AccountFinancingMode); err != nil {
		return err
	}

	return nil
}

func (m *DailyFinancingTransaction) validatePositionFinancings(formats strfmt.Registry) error {

	if swag.IsZero(m.PositionFinancings) { // not required
		return nil
	}

	for i := 0; i < len(m.PositionFinancings); i++ {
		if swag.IsZero(m.PositionFinancings[i]) { // not required
			continue
		}

		if m.PositionFinancings[i] != nil {
			if err := m.PositionFinancings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("positionFinancings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var dailyFinancingTransactionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","CLOSE","REOPEN","CLIENT_CONFIGURE","CLIENT_CONFIGURE_REJECT","TRANSFER_FUNDS","TRANSFER_FUNDS_REJECT","MARKET_ORDER","MARKET_ORDER_REJECT","FIXED_PRICE_ORDER","LIMIT_ORDER","LIMIT_ORDER_REJECT","STOP_ORDER","STOP_ORDER_REJECT","MARKET_IF_TOUCHED_ORDER","MARKET_IF_TOUCHED_ORDER_REJECT","TAKE_PROFIT_ORDER","TAKE_PROFIT_ORDER_REJECT","STOP_LOSS_ORDER","STOP_LOSS_ORDER_REJECT","TRAILING_STOP_LOSS_ORDER","TRAILING_STOP_LOSS_ORDER_REJECT","ORDER_FILL","ORDER_CANCEL","ORDER_CANCEL_REJECT","ORDER_CLIENT_EXTENSIONS_MODIFY","ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT","TRADE_CLIENT_EXTENSIONS_MODIFY","TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT","MARGIN_CALL_ENTER","MARGIN_CALL_EXTEND","MARGIN_CALL_EXIT","DELAYED_TRADE_CLOSURE","DAILY_FINANCING","RESET_RESETTABLE_PL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dailyFinancingTransactionTypeTypePropEnum = append(dailyFinancingTransactionTypeTypePropEnum, v)
	}
}

const (

	// DailyFinancingTransactionTypeCREATE captures enum value "CREATE"
	DailyFinancingTransactionTypeCREATE string = "CREATE"

	// DailyFinancingTransactionTypeCLOSE captures enum value "CLOSE"
	DailyFinancingTransactionTypeCLOSE string = "CLOSE"

	// DailyFinancingTransactionTypeREOPEN captures enum value "REOPEN"
	DailyFinancingTransactionTypeREOPEN string = "REOPEN"

	// DailyFinancingTransactionTypeCLIENTCONFIGURE captures enum value "CLIENT_CONFIGURE"
	DailyFinancingTransactionTypeCLIENTCONFIGURE string = "CLIENT_CONFIGURE"

	// DailyFinancingTransactionTypeCLIENTCONFIGUREREJECT captures enum value "CLIENT_CONFIGURE_REJECT"
	DailyFinancingTransactionTypeCLIENTCONFIGUREREJECT string = "CLIENT_CONFIGURE_REJECT"

	// DailyFinancingTransactionTypeTRANSFERFUNDS captures enum value "TRANSFER_FUNDS"
	DailyFinancingTransactionTypeTRANSFERFUNDS string = "TRANSFER_FUNDS"

	// DailyFinancingTransactionTypeTRANSFERFUNDSREJECT captures enum value "TRANSFER_FUNDS_REJECT"
	DailyFinancingTransactionTypeTRANSFERFUNDSREJECT string = "TRANSFER_FUNDS_REJECT"

	// DailyFinancingTransactionTypeMARKETORDER captures enum value "MARKET_ORDER"
	DailyFinancingTransactionTypeMARKETORDER string = "MARKET_ORDER"

	// DailyFinancingTransactionTypeMARKETORDERREJECT captures enum value "MARKET_ORDER_REJECT"
	DailyFinancingTransactionTypeMARKETORDERREJECT string = "MARKET_ORDER_REJECT"

	// DailyFinancingTransactionTypeFIXEDPRICEORDER captures enum value "FIXED_PRICE_ORDER"
	DailyFinancingTransactionTypeFIXEDPRICEORDER string = "FIXED_PRICE_ORDER"

	// DailyFinancingTransactionTypeLIMITORDER captures enum value "LIMIT_ORDER"
	DailyFinancingTransactionTypeLIMITORDER string = "LIMIT_ORDER"

	// DailyFinancingTransactionTypeLIMITORDERREJECT captures enum value "LIMIT_ORDER_REJECT"
	DailyFinancingTransactionTypeLIMITORDERREJECT string = "LIMIT_ORDER_REJECT"

	// DailyFinancingTransactionTypeSTOPORDER captures enum value "STOP_ORDER"
	DailyFinancingTransactionTypeSTOPORDER string = "STOP_ORDER"

	// DailyFinancingTransactionTypeSTOPORDERREJECT captures enum value "STOP_ORDER_REJECT"
	DailyFinancingTransactionTypeSTOPORDERREJECT string = "STOP_ORDER_REJECT"

	// DailyFinancingTransactionTypeMARKETIFTOUCHEDORDER captures enum value "MARKET_IF_TOUCHED_ORDER"
	DailyFinancingTransactionTypeMARKETIFTOUCHEDORDER string = "MARKET_IF_TOUCHED_ORDER"

	// DailyFinancingTransactionTypeMARKETIFTOUCHEDORDERREJECT captures enum value "MARKET_IF_TOUCHED_ORDER_REJECT"
	DailyFinancingTransactionTypeMARKETIFTOUCHEDORDERREJECT string = "MARKET_IF_TOUCHED_ORDER_REJECT"

	// DailyFinancingTransactionTypeTAKEPROFITORDER captures enum value "TAKE_PROFIT_ORDER"
	DailyFinancingTransactionTypeTAKEPROFITORDER string = "TAKE_PROFIT_ORDER"

	// DailyFinancingTransactionTypeTAKEPROFITORDERREJECT captures enum value "TAKE_PROFIT_ORDER_REJECT"
	DailyFinancingTransactionTypeTAKEPROFITORDERREJECT string = "TAKE_PROFIT_ORDER_REJECT"

	// DailyFinancingTransactionTypeSTOPLOSSORDER captures enum value "STOP_LOSS_ORDER"
	DailyFinancingTransactionTypeSTOPLOSSORDER string = "STOP_LOSS_ORDER"

	// DailyFinancingTransactionTypeSTOPLOSSORDERREJECT captures enum value "STOP_LOSS_ORDER_REJECT"
	DailyFinancingTransactionTypeSTOPLOSSORDERREJECT string = "STOP_LOSS_ORDER_REJECT"

	// DailyFinancingTransactionTypeTRAILINGSTOPLOSSORDER captures enum value "TRAILING_STOP_LOSS_ORDER"
	DailyFinancingTransactionTypeTRAILINGSTOPLOSSORDER string = "TRAILING_STOP_LOSS_ORDER"

	// DailyFinancingTransactionTypeTRAILINGSTOPLOSSORDERREJECT captures enum value "TRAILING_STOP_LOSS_ORDER_REJECT"
	DailyFinancingTransactionTypeTRAILINGSTOPLOSSORDERREJECT string = "TRAILING_STOP_LOSS_ORDER_REJECT"

	// DailyFinancingTransactionTypeORDERFILL captures enum value "ORDER_FILL"
	DailyFinancingTransactionTypeORDERFILL string = "ORDER_FILL"

	// DailyFinancingTransactionTypeORDERCANCEL captures enum value "ORDER_CANCEL"
	DailyFinancingTransactionTypeORDERCANCEL string = "ORDER_CANCEL"

	// DailyFinancingTransactionTypeORDERCANCELREJECT captures enum value "ORDER_CANCEL_REJECT"
	DailyFinancingTransactionTypeORDERCANCELREJECT string = "ORDER_CANCEL_REJECT"

	// DailyFinancingTransactionTypeORDERCLIENTEXTENSIONSMODIFY captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY"
	DailyFinancingTransactionTypeORDERCLIENTEXTENSIONSMODIFY string = "ORDER_CLIENT_EXTENSIONS_MODIFY"

	// DailyFinancingTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	DailyFinancingTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT string = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// DailyFinancingTransactionTypeTRADECLIENTEXTENSIONSMODIFY captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY"
	DailyFinancingTransactionTypeTRADECLIENTEXTENSIONSMODIFY string = "TRADE_CLIENT_EXTENSIONS_MODIFY"

	// DailyFinancingTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	DailyFinancingTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT string = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// DailyFinancingTransactionTypeMARGINCALLENTER captures enum value "MARGIN_CALL_ENTER"
	DailyFinancingTransactionTypeMARGINCALLENTER string = "MARGIN_CALL_ENTER"

	// DailyFinancingTransactionTypeMARGINCALLEXTEND captures enum value "MARGIN_CALL_EXTEND"
	DailyFinancingTransactionTypeMARGINCALLEXTEND string = "MARGIN_CALL_EXTEND"

	// DailyFinancingTransactionTypeMARGINCALLEXIT captures enum value "MARGIN_CALL_EXIT"
	DailyFinancingTransactionTypeMARGINCALLEXIT string = "MARGIN_CALL_EXIT"

	// DailyFinancingTransactionTypeDELAYEDTRADECLOSURE captures enum value "DELAYED_TRADE_CLOSURE"
	DailyFinancingTransactionTypeDELAYEDTRADECLOSURE string = "DELAYED_TRADE_CLOSURE"

	// DailyFinancingTransactionTypeDAILYFINANCING captures enum value "DAILY_FINANCING"
	DailyFinancingTransactionTypeDAILYFINANCING string = "DAILY_FINANCING"

	// DailyFinancingTransactionTypeRESETRESETTABLEPL captures enum value "RESET_RESETTABLE_PL"
	DailyFinancingTransactionTypeRESETRESETTABLEPL string = "RESET_RESETTABLE_PL"
)

// prop value enum
func (m *DailyFinancingTransaction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dailyFinancingTransactionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DailyFinancingTransaction) validateType(formats strfmt.Registry) error {

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
func (m *DailyFinancingTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DailyFinancingTransaction) UnmarshalBinary(b []byte) error {
	var res DailyFinancingTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
