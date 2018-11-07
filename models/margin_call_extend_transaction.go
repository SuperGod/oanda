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

// MarginCallExtendTransaction A MarginCallExtendTransaction is created when the margin call state for an Account has been extended.
// swagger:model MarginCallExtendTransaction
type MarginCallExtendTransaction struct {

	// The ID of the Account the Transaction was created for.
	AccountID string `json:"accountID,omitempty"`

	// The ID of the "batch" that the Transaction belongs to. Transactions in the same batch are applied to the Account simultaneously.
	BatchID string `json:"batchID,omitempty"`

	// The number of the extensions to the Account's current margin call that have been applied. This value will be set to 1 for the first MarginCallExtend Transaction
	ExtensionNumber int64 `json:"extensionNumber,omitempty"`

	// The Transaction's Identifier.
	ID string `json:"id,omitempty"`

	// The Request ID of the request which generated the transaction.
	RequestID string `json:"requestID,omitempty"`

	// The date/time when the Transaction was created.
	Time string `json:"time,omitempty"`

	// The Type of the Transaction. Always set to "MARGIN_CALL_EXTEND" for an MarginCallExtendTransaction.
	// Enum: [CREATE CLOSE REOPEN CLIENT_CONFIGURE CLIENT_CONFIGURE_REJECT TRANSFER_FUNDS TRANSFER_FUNDS_REJECT MARKET_ORDER MARKET_ORDER_REJECT FIXED_PRICE_ORDER LIMIT_ORDER LIMIT_ORDER_REJECT STOP_ORDER STOP_ORDER_REJECT MARKET_IF_TOUCHED_ORDER MARKET_IF_TOUCHED_ORDER_REJECT TAKE_PROFIT_ORDER TAKE_PROFIT_ORDER_REJECT STOP_LOSS_ORDER STOP_LOSS_ORDER_REJECT TRAILING_STOP_LOSS_ORDER TRAILING_STOP_LOSS_ORDER_REJECT ORDER_FILL ORDER_CANCEL ORDER_CANCEL_REJECT ORDER_CLIENT_EXTENSIONS_MODIFY ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT TRADE_CLIENT_EXTENSIONS_MODIFY TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT MARGIN_CALL_ENTER MARGIN_CALL_EXTEND MARGIN_CALL_EXIT DELAYED_TRADE_CLOSURE DAILY_FINANCING RESET_RESETTABLE_PL]
	Type string `json:"type,omitempty"`

	// The ID of the user that initiated the creation of the Transaction.
	UserID int64 `json:"userID,omitempty"`
}

// Validate validates this margin call extend transaction
func (m *MarginCallExtendTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var marginCallExtendTransactionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","CLOSE","REOPEN","CLIENT_CONFIGURE","CLIENT_CONFIGURE_REJECT","TRANSFER_FUNDS","TRANSFER_FUNDS_REJECT","MARKET_ORDER","MARKET_ORDER_REJECT","FIXED_PRICE_ORDER","LIMIT_ORDER","LIMIT_ORDER_REJECT","STOP_ORDER","STOP_ORDER_REJECT","MARKET_IF_TOUCHED_ORDER","MARKET_IF_TOUCHED_ORDER_REJECT","TAKE_PROFIT_ORDER","TAKE_PROFIT_ORDER_REJECT","STOP_LOSS_ORDER","STOP_LOSS_ORDER_REJECT","TRAILING_STOP_LOSS_ORDER","TRAILING_STOP_LOSS_ORDER_REJECT","ORDER_FILL","ORDER_CANCEL","ORDER_CANCEL_REJECT","ORDER_CLIENT_EXTENSIONS_MODIFY","ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT","TRADE_CLIENT_EXTENSIONS_MODIFY","TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT","MARGIN_CALL_ENTER","MARGIN_CALL_EXTEND","MARGIN_CALL_EXIT","DELAYED_TRADE_CLOSURE","DAILY_FINANCING","RESET_RESETTABLE_PL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marginCallExtendTransactionTypeTypePropEnum = append(marginCallExtendTransactionTypeTypePropEnum, v)
	}
}

const (

	// MarginCallExtendTransactionTypeCREATE captures enum value "CREATE"
	MarginCallExtendTransactionTypeCREATE string = "CREATE"

	// MarginCallExtendTransactionTypeCLOSE captures enum value "CLOSE"
	MarginCallExtendTransactionTypeCLOSE string = "CLOSE"

	// MarginCallExtendTransactionTypeREOPEN captures enum value "REOPEN"
	MarginCallExtendTransactionTypeREOPEN string = "REOPEN"

	// MarginCallExtendTransactionTypeCLIENTCONFIGURE captures enum value "CLIENT_CONFIGURE"
	MarginCallExtendTransactionTypeCLIENTCONFIGURE string = "CLIENT_CONFIGURE"

	// MarginCallExtendTransactionTypeCLIENTCONFIGUREREJECT captures enum value "CLIENT_CONFIGURE_REJECT"
	MarginCallExtendTransactionTypeCLIENTCONFIGUREREJECT string = "CLIENT_CONFIGURE_REJECT"

	// MarginCallExtendTransactionTypeTRANSFERFUNDS captures enum value "TRANSFER_FUNDS"
	MarginCallExtendTransactionTypeTRANSFERFUNDS string = "TRANSFER_FUNDS"

	// MarginCallExtendTransactionTypeTRANSFERFUNDSREJECT captures enum value "TRANSFER_FUNDS_REJECT"
	MarginCallExtendTransactionTypeTRANSFERFUNDSREJECT string = "TRANSFER_FUNDS_REJECT"

	// MarginCallExtendTransactionTypeMARKETORDER captures enum value "MARKET_ORDER"
	MarginCallExtendTransactionTypeMARKETORDER string = "MARKET_ORDER"

	// MarginCallExtendTransactionTypeMARKETORDERREJECT captures enum value "MARKET_ORDER_REJECT"
	MarginCallExtendTransactionTypeMARKETORDERREJECT string = "MARKET_ORDER_REJECT"

	// MarginCallExtendTransactionTypeFIXEDPRICEORDER captures enum value "FIXED_PRICE_ORDER"
	MarginCallExtendTransactionTypeFIXEDPRICEORDER string = "FIXED_PRICE_ORDER"

	// MarginCallExtendTransactionTypeLIMITORDER captures enum value "LIMIT_ORDER"
	MarginCallExtendTransactionTypeLIMITORDER string = "LIMIT_ORDER"

	// MarginCallExtendTransactionTypeLIMITORDERREJECT captures enum value "LIMIT_ORDER_REJECT"
	MarginCallExtendTransactionTypeLIMITORDERREJECT string = "LIMIT_ORDER_REJECT"

	// MarginCallExtendTransactionTypeSTOPORDER captures enum value "STOP_ORDER"
	MarginCallExtendTransactionTypeSTOPORDER string = "STOP_ORDER"

	// MarginCallExtendTransactionTypeSTOPORDERREJECT captures enum value "STOP_ORDER_REJECT"
	MarginCallExtendTransactionTypeSTOPORDERREJECT string = "STOP_ORDER_REJECT"

	// MarginCallExtendTransactionTypeMARKETIFTOUCHEDORDER captures enum value "MARKET_IF_TOUCHED_ORDER"
	MarginCallExtendTransactionTypeMARKETIFTOUCHEDORDER string = "MARKET_IF_TOUCHED_ORDER"

	// MarginCallExtendTransactionTypeMARKETIFTOUCHEDORDERREJECT captures enum value "MARKET_IF_TOUCHED_ORDER_REJECT"
	MarginCallExtendTransactionTypeMARKETIFTOUCHEDORDERREJECT string = "MARKET_IF_TOUCHED_ORDER_REJECT"

	// MarginCallExtendTransactionTypeTAKEPROFITORDER captures enum value "TAKE_PROFIT_ORDER"
	MarginCallExtendTransactionTypeTAKEPROFITORDER string = "TAKE_PROFIT_ORDER"

	// MarginCallExtendTransactionTypeTAKEPROFITORDERREJECT captures enum value "TAKE_PROFIT_ORDER_REJECT"
	MarginCallExtendTransactionTypeTAKEPROFITORDERREJECT string = "TAKE_PROFIT_ORDER_REJECT"

	// MarginCallExtendTransactionTypeSTOPLOSSORDER captures enum value "STOP_LOSS_ORDER"
	MarginCallExtendTransactionTypeSTOPLOSSORDER string = "STOP_LOSS_ORDER"

	// MarginCallExtendTransactionTypeSTOPLOSSORDERREJECT captures enum value "STOP_LOSS_ORDER_REJECT"
	MarginCallExtendTransactionTypeSTOPLOSSORDERREJECT string = "STOP_LOSS_ORDER_REJECT"

	// MarginCallExtendTransactionTypeTRAILINGSTOPLOSSORDER captures enum value "TRAILING_STOP_LOSS_ORDER"
	MarginCallExtendTransactionTypeTRAILINGSTOPLOSSORDER string = "TRAILING_STOP_LOSS_ORDER"

	// MarginCallExtendTransactionTypeTRAILINGSTOPLOSSORDERREJECT captures enum value "TRAILING_STOP_LOSS_ORDER_REJECT"
	MarginCallExtendTransactionTypeTRAILINGSTOPLOSSORDERREJECT string = "TRAILING_STOP_LOSS_ORDER_REJECT"

	// MarginCallExtendTransactionTypeORDERFILL captures enum value "ORDER_FILL"
	MarginCallExtendTransactionTypeORDERFILL string = "ORDER_FILL"

	// MarginCallExtendTransactionTypeORDERCANCEL captures enum value "ORDER_CANCEL"
	MarginCallExtendTransactionTypeORDERCANCEL string = "ORDER_CANCEL"

	// MarginCallExtendTransactionTypeORDERCANCELREJECT captures enum value "ORDER_CANCEL_REJECT"
	MarginCallExtendTransactionTypeORDERCANCELREJECT string = "ORDER_CANCEL_REJECT"

	// MarginCallExtendTransactionTypeORDERCLIENTEXTENSIONSMODIFY captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY"
	MarginCallExtendTransactionTypeORDERCLIENTEXTENSIONSMODIFY string = "ORDER_CLIENT_EXTENSIONS_MODIFY"

	// MarginCallExtendTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	MarginCallExtendTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT string = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// MarginCallExtendTransactionTypeTRADECLIENTEXTENSIONSMODIFY captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY"
	MarginCallExtendTransactionTypeTRADECLIENTEXTENSIONSMODIFY string = "TRADE_CLIENT_EXTENSIONS_MODIFY"

	// MarginCallExtendTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	MarginCallExtendTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT string = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// MarginCallExtendTransactionTypeMARGINCALLENTER captures enum value "MARGIN_CALL_ENTER"
	MarginCallExtendTransactionTypeMARGINCALLENTER string = "MARGIN_CALL_ENTER"

	// MarginCallExtendTransactionTypeMARGINCALLEXTEND captures enum value "MARGIN_CALL_EXTEND"
	MarginCallExtendTransactionTypeMARGINCALLEXTEND string = "MARGIN_CALL_EXTEND"

	// MarginCallExtendTransactionTypeMARGINCALLEXIT captures enum value "MARGIN_CALL_EXIT"
	MarginCallExtendTransactionTypeMARGINCALLEXIT string = "MARGIN_CALL_EXIT"

	// MarginCallExtendTransactionTypeDELAYEDTRADECLOSURE captures enum value "DELAYED_TRADE_CLOSURE"
	MarginCallExtendTransactionTypeDELAYEDTRADECLOSURE string = "DELAYED_TRADE_CLOSURE"

	// MarginCallExtendTransactionTypeDAILYFINANCING captures enum value "DAILY_FINANCING"
	MarginCallExtendTransactionTypeDAILYFINANCING string = "DAILY_FINANCING"

	// MarginCallExtendTransactionTypeRESETRESETTABLEPL captures enum value "RESET_RESETTABLE_PL"
	MarginCallExtendTransactionTypeRESETRESETTABLEPL string = "RESET_RESETTABLE_PL"
)

// prop value enum
func (m *MarginCallExtendTransaction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marginCallExtendTransactionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarginCallExtendTransaction) validateType(formats strfmt.Registry) error {

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
func (m *MarginCallExtendTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarginCallExtendTransaction) UnmarshalBinary(b []byte) error {
	var res MarginCallExtendTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}