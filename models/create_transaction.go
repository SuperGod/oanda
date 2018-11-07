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

// CreateTransaction A CreateTransaction represents the creation of an Account.
// swagger:model CreateTransaction
type CreateTransaction struct {

	// The ID of the Account the Transaction was created for.
	AccountID string `json:"accountID,omitempty"`

	// The number of the Account within the site/division/user
	AccountNumber int64 `json:"accountNumber,omitempty"`

	// The ID of the user that the Account was created for
	AccountUserID int64 `json:"accountUserID,omitempty"`

	// The ID of the "batch" that the Transaction belongs to. Transactions in the same batch are applied to the Account simultaneously.
	BatchID string `json:"batchID,omitempty"`

	// The ID of the Division that the Account is in
	DivisionID int64 `json:"divisionID,omitempty"`

	// The home currency of the Account
	HomeCurrency string `json:"homeCurrency,omitempty"`

	// The Transaction's Identifier.
	ID string `json:"id,omitempty"`

	// The Request ID of the request which generated the transaction.
	RequestID string `json:"requestID,omitempty"`

	// The ID of the Site that the Account was created at
	SiteID int64 `json:"siteID,omitempty"`

	// The date/time when the Transaction was created.
	Time string `json:"time,omitempty"`

	// The Type of the Transaction. Always set to "CREATE" in a CreateTransaction.
	// Enum: [CREATE CLOSE REOPEN CLIENT_CONFIGURE CLIENT_CONFIGURE_REJECT TRANSFER_FUNDS TRANSFER_FUNDS_REJECT MARKET_ORDER MARKET_ORDER_REJECT FIXED_PRICE_ORDER LIMIT_ORDER LIMIT_ORDER_REJECT STOP_ORDER STOP_ORDER_REJECT MARKET_IF_TOUCHED_ORDER MARKET_IF_TOUCHED_ORDER_REJECT TAKE_PROFIT_ORDER TAKE_PROFIT_ORDER_REJECT STOP_LOSS_ORDER STOP_LOSS_ORDER_REJECT TRAILING_STOP_LOSS_ORDER TRAILING_STOP_LOSS_ORDER_REJECT ORDER_FILL ORDER_CANCEL ORDER_CANCEL_REJECT ORDER_CLIENT_EXTENSIONS_MODIFY ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT TRADE_CLIENT_EXTENSIONS_MODIFY TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT MARGIN_CALL_ENTER MARGIN_CALL_EXTEND MARGIN_CALL_EXIT DELAYED_TRADE_CLOSURE DAILY_FINANCING RESET_RESETTABLE_PL]
	Type string `json:"type,omitempty"`

	// The ID of the user that initiated the creation of the Transaction.
	UserID int64 `json:"userID,omitempty"`
}

// Validate validates this create transaction
func (m *CreateTransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createTransactionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","CLOSE","REOPEN","CLIENT_CONFIGURE","CLIENT_CONFIGURE_REJECT","TRANSFER_FUNDS","TRANSFER_FUNDS_REJECT","MARKET_ORDER","MARKET_ORDER_REJECT","FIXED_PRICE_ORDER","LIMIT_ORDER","LIMIT_ORDER_REJECT","STOP_ORDER","STOP_ORDER_REJECT","MARKET_IF_TOUCHED_ORDER","MARKET_IF_TOUCHED_ORDER_REJECT","TAKE_PROFIT_ORDER","TAKE_PROFIT_ORDER_REJECT","STOP_LOSS_ORDER","STOP_LOSS_ORDER_REJECT","TRAILING_STOP_LOSS_ORDER","TRAILING_STOP_LOSS_ORDER_REJECT","ORDER_FILL","ORDER_CANCEL","ORDER_CANCEL_REJECT","ORDER_CLIENT_EXTENSIONS_MODIFY","ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT","TRADE_CLIENT_EXTENSIONS_MODIFY","TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT","MARGIN_CALL_ENTER","MARGIN_CALL_EXTEND","MARGIN_CALL_EXIT","DELAYED_TRADE_CLOSURE","DAILY_FINANCING","RESET_RESETTABLE_PL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createTransactionTypeTypePropEnum = append(createTransactionTypeTypePropEnum, v)
	}
}

const (

	// CreateTransactionTypeCREATE captures enum value "CREATE"
	CreateTransactionTypeCREATE string = "CREATE"

	// CreateTransactionTypeCLOSE captures enum value "CLOSE"
	CreateTransactionTypeCLOSE string = "CLOSE"

	// CreateTransactionTypeREOPEN captures enum value "REOPEN"
	CreateTransactionTypeREOPEN string = "REOPEN"

	// CreateTransactionTypeCLIENTCONFIGURE captures enum value "CLIENT_CONFIGURE"
	CreateTransactionTypeCLIENTCONFIGURE string = "CLIENT_CONFIGURE"

	// CreateTransactionTypeCLIENTCONFIGUREREJECT captures enum value "CLIENT_CONFIGURE_REJECT"
	CreateTransactionTypeCLIENTCONFIGUREREJECT string = "CLIENT_CONFIGURE_REJECT"

	// CreateTransactionTypeTRANSFERFUNDS captures enum value "TRANSFER_FUNDS"
	CreateTransactionTypeTRANSFERFUNDS string = "TRANSFER_FUNDS"

	// CreateTransactionTypeTRANSFERFUNDSREJECT captures enum value "TRANSFER_FUNDS_REJECT"
	CreateTransactionTypeTRANSFERFUNDSREJECT string = "TRANSFER_FUNDS_REJECT"

	// CreateTransactionTypeMARKETORDER captures enum value "MARKET_ORDER"
	CreateTransactionTypeMARKETORDER string = "MARKET_ORDER"

	// CreateTransactionTypeMARKETORDERREJECT captures enum value "MARKET_ORDER_REJECT"
	CreateTransactionTypeMARKETORDERREJECT string = "MARKET_ORDER_REJECT"

	// CreateTransactionTypeFIXEDPRICEORDER captures enum value "FIXED_PRICE_ORDER"
	CreateTransactionTypeFIXEDPRICEORDER string = "FIXED_PRICE_ORDER"

	// CreateTransactionTypeLIMITORDER captures enum value "LIMIT_ORDER"
	CreateTransactionTypeLIMITORDER string = "LIMIT_ORDER"

	// CreateTransactionTypeLIMITORDERREJECT captures enum value "LIMIT_ORDER_REJECT"
	CreateTransactionTypeLIMITORDERREJECT string = "LIMIT_ORDER_REJECT"

	// CreateTransactionTypeSTOPORDER captures enum value "STOP_ORDER"
	CreateTransactionTypeSTOPORDER string = "STOP_ORDER"

	// CreateTransactionTypeSTOPORDERREJECT captures enum value "STOP_ORDER_REJECT"
	CreateTransactionTypeSTOPORDERREJECT string = "STOP_ORDER_REJECT"

	// CreateTransactionTypeMARKETIFTOUCHEDORDER captures enum value "MARKET_IF_TOUCHED_ORDER"
	CreateTransactionTypeMARKETIFTOUCHEDORDER string = "MARKET_IF_TOUCHED_ORDER"

	// CreateTransactionTypeMARKETIFTOUCHEDORDERREJECT captures enum value "MARKET_IF_TOUCHED_ORDER_REJECT"
	CreateTransactionTypeMARKETIFTOUCHEDORDERREJECT string = "MARKET_IF_TOUCHED_ORDER_REJECT"

	// CreateTransactionTypeTAKEPROFITORDER captures enum value "TAKE_PROFIT_ORDER"
	CreateTransactionTypeTAKEPROFITORDER string = "TAKE_PROFIT_ORDER"

	// CreateTransactionTypeTAKEPROFITORDERREJECT captures enum value "TAKE_PROFIT_ORDER_REJECT"
	CreateTransactionTypeTAKEPROFITORDERREJECT string = "TAKE_PROFIT_ORDER_REJECT"

	// CreateTransactionTypeSTOPLOSSORDER captures enum value "STOP_LOSS_ORDER"
	CreateTransactionTypeSTOPLOSSORDER string = "STOP_LOSS_ORDER"

	// CreateTransactionTypeSTOPLOSSORDERREJECT captures enum value "STOP_LOSS_ORDER_REJECT"
	CreateTransactionTypeSTOPLOSSORDERREJECT string = "STOP_LOSS_ORDER_REJECT"

	// CreateTransactionTypeTRAILINGSTOPLOSSORDER captures enum value "TRAILING_STOP_LOSS_ORDER"
	CreateTransactionTypeTRAILINGSTOPLOSSORDER string = "TRAILING_STOP_LOSS_ORDER"

	// CreateTransactionTypeTRAILINGSTOPLOSSORDERREJECT captures enum value "TRAILING_STOP_LOSS_ORDER_REJECT"
	CreateTransactionTypeTRAILINGSTOPLOSSORDERREJECT string = "TRAILING_STOP_LOSS_ORDER_REJECT"

	// CreateTransactionTypeORDERFILL captures enum value "ORDER_FILL"
	CreateTransactionTypeORDERFILL string = "ORDER_FILL"

	// CreateTransactionTypeORDERCANCEL captures enum value "ORDER_CANCEL"
	CreateTransactionTypeORDERCANCEL string = "ORDER_CANCEL"

	// CreateTransactionTypeORDERCANCELREJECT captures enum value "ORDER_CANCEL_REJECT"
	CreateTransactionTypeORDERCANCELREJECT string = "ORDER_CANCEL_REJECT"

	// CreateTransactionTypeORDERCLIENTEXTENSIONSMODIFY captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY"
	CreateTransactionTypeORDERCLIENTEXTENSIONSMODIFY string = "ORDER_CLIENT_EXTENSIONS_MODIFY"

	// CreateTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT captures enum value "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	CreateTransactionTypeORDERCLIENTEXTENSIONSMODIFYREJECT string = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// CreateTransactionTypeTRADECLIENTEXTENSIONSMODIFY captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY"
	CreateTransactionTypeTRADECLIENTEXTENSIONSMODIFY string = "TRADE_CLIENT_EXTENSIONS_MODIFY"

	// CreateTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT captures enum value "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	CreateTransactionTypeTRADECLIENTEXTENSIONSMODIFYREJECT string = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"

	// CreateTransactionTypeMARGINCALLENTER captures enum value "MARGIN_CALL_ENTER"
	CreateTransactionTypeMARGINCALLENTER string = "MARGIN_CALL_ENTER"

	// CreateTransactionTypeMARGINCALLEXTEND captures enum value "MARGIN_CALL_EXTEND"
	CreateTransactionTypeMARGINCALLEXTEND string = "MARGIN_CALL_EXTEND"

	// CreateTransactionTypeMARGINCALLEXIT captures enum value "MARGIN_CALL_EXIT"
	CreateTransactionTypeMARGINCALLEXIT string = "MARGIN_CALL_EXIT"

	// CreateTransactionTypeDELAYEDTRADECLOSURE captures enum value "DELAYED_TRADE_CLOSURE"
	CreateTransactionTypeDELAYEDTRADECLOSURE string = "DELAYED_TRADE_CLOSURE"

	// CreateTransactionTypeDAILYFINANCING captures enum value "DAILY_FINANCING"
	CreateTransactionTypeDAILYFINANCING string = "DAILY_FINANCING"

	// CreateTransactionTypeRESETRESETTABLEPL captures enum value "RESET_RESETTABLE_PL"
	CreateTransactionTypeRESETRESETTABLEPL string = "RESET_RESETTABLE_PL"
)

// prop value enum
func (m *CreateTransaction) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createTransactionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateTransaction) validateType(formats strfmt.Registry) error {

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
func (m *CreateTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTransaction) UnmarshalBinary(b []byte) error {
	var res CreateTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
