// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// OrderCancelReason The reason that an Order was cancelled.
// swagger:model OrderCancelReason
type OrderCancelReason string

const (

	// OrderCancelReasonINTERNALSERVERERROR captures enum value "INTERNAL_SERVER_ERROR"
	OrderCancelReasonINTERNALSERVERERROR OrderCancelReason = "INTERNAL_SERVER_ERROR"

	// OrderCancelReasonACCOUNTLOCKED captures enum value "ACCOUNT_LOCKED"
	OrderCancelReasonACCOUNTLOCKED OrderCancelReason = "ACCOUNT_LOCKED"

	// OrderCancelReasonACCOUNTNEWPOSITIONSLOCKED captures enum value "ACCOUNT_NEW_POSITIONS_LOCKED"
	OrderCancelReasonACCOUNTNEWPOSITIONSLOCKED OrderCancelReason = "ACCOUNT_NEW_POSITIONS_LOCKED"

	// OrderCancelReasonACCOUNTORDERCREATIONLOCKED captures enum value "ACCOUNT_ORDER_CREATION_LOCKED"
	OrderCancelReasonACCOUNTORDERCREATIONLOCKED OrderCancelReason = "ACCOUNT_ORDER_CREATION_LOCKED"

	// OrderCancelReasonACCOUNTORDERFILLLOCKED captures enum value "ACCOUNT_ORDER_FILL_LOCKED"
	OrderCancelReasonACCOUNTORDERFILLLOCKED OrderCancelReason = "ACCOUNT_ORDER_FILL_LOCKED"

	// OrderCancelReasonCLIENTREQUEST captures enum value "CLIENT_REQUEST"
	OrderCancelReasonCLIENTREQUEST OrderCancelReason = "CLIENT_REQUEST"

	// OrderCancelReasonMIGRATION captures enum value "MIGRATION"
	OrderCancelReasonMIGRATION OrderCancelReason = "MIGRATION"

	// OrderCancelReasonMARKETHALTED captures enum value "MARKET_HALTED"
	OrderCancelReasonMARKETHALTED OrderCancelReason = "MARKET_HALTED"

	// OrderCancelReasonLINKEDTRADECLOSED captures enum value "LINKED_TRADE_CLOSED"
	OrderCancelReasonLINKEDTRADECLOSED OrderCancelReason = "LINKED_TRADE_CLOSED"

	// OrderCancelReasonTIMEINFORCEEXPIRED captures enum value "TIME_IN_FORCE_EXPIRED"
	OrderCancelReasonTIMEINFORCEEXPIRED OrderCancelReason = "TIME_IN_FORCE_EXPIRED"

	// OrderCancelReasonINSUFFICIENTMARGIN captures enum value "INSUFFICIENT_MARGIN"
	OrderCancelReasonINSUFFICIENTMARGIN OrderCancelReason = "INSUFFICIENT_MARGIN"

	// OrderCancelReasonFIFOVIOLATION captures enum value "FIFO_VIOLATION"
	OrderCancelReasonFIFOVIOLATION OrderCancelReason = "FIFO_VIOLATION"

	// OrderCancelReasonBOUNDSVIOLATION captures enum value "BOUNDS_VIOLATION"
	OrderCancelReasonBOUNDSVIOLATION OrderCancelReason = "BOUNDS_VIOLATION"

	// OrderCancelReasonCLIENTREQUESTREPLACED captures enum value "CLIENT_REQUEST_REPLACED"
	OrderCancelReasonCLIENTREQUESTREPLACED OrderCancelReason = "CLIENT_REQUEST_REPLACED"

	// OrderCancelReasonINSUFFICIENTLIQUIDITY captures enum value "INSUFFICIENT_LIQUIDITY"
	OrderCancelReasonINSUFFICIENTLIQUIDITY OrderCancelReason = "INSUFFICIENT_LIQUIDITY"

	// OrderCancelReasonTAKEPROFITONFILLGTDTIMESTAMPINPAST captures enum value "TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	OrderCancelReasonTAKEPROFITONFILLGTDTIMESTAMPINPAST OrderCancelReason = "TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST"

	// OrderCancelReasonTAKEPROFITONFILLLOSS captures enum value "TAKE_PROFIT_ON_FILL_LOSS"
	OrderCancelReasonTAKEPROFITONFILLLOSS OrderCancelReason = "TAKE_PROFIT_ON_FILL_LOSS"

	// OrderCancelReasonLOSINGTAKEPROFIT captures enum value "LOSING_TAKE_PROFIT"
	OrderCancelReasonLOSINGTAKEPROFIT OrderCancelReason = "LOSING_TAKE_PROFIT"

	// OrderCancelReasonSTOPLOSSONFILLGTDTIMESTAMPINPAST captures enum value "STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	OrderCancelReasonSTOPLOSSONFILLGTDTIMESTAMPINPAST OrderCancelReason = "STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"

	// OrderCancelReasonSTOPLOSSONFILLLOSS captures enum value "STOP_LOSS_ON_FILL_LOSS"
	OrderCancelReasonSTOPLOSSONFILLLOSS OrderCancelReason = "STOP_LOSS_ON_FILL_LOSS"

	// OrderCancelReasonSTOPLOSSONFILLPRICEDISTANCEMAXIMUMEXCEEDED captures enum value "STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"
	OrderCancelReasonSTOPLOSSONFILLPRICEDISTANCEMAXIMUMEXCEEDED OrderCancelReason = "STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"

	// OrderCancelReasonSTOPLOSSONFILLREQUIRED captures enum value "STOP_LOSS_ON_FILL_REQUIRED"
	OrderCancelReasonSTOPLOSSONFILLREQUIRED OrderCancelReason = "STOP_LOSS_ON_FILL_REQUIRED"

	// OrderCancelReasonSTOPLOSSONFILLGUARANTEEDREQUIRED captures enum value "STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED"
	OrderCancelReasonSTOPLOSSONFILLGUARANTEEDREQUIRED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED"

	// OrderCancelReasonSTOPLOSSONFILLGUARANTEEDNOTALLOWED captures enum value "STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED"
	OrderCancelReasonSTOPLOSSONFILLGUARANTEEDNOTALLOWED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED"

	// OrderCancelReasonSTOPLOSSONFILLGUARANTEEDMINIMUMDISTANCENOTMET captures enum value "STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET"
	OrderCancelReasonSTOPLOSSONFILLGUARANTEEDMINIMUMDISTANCENOTMET OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET"

	// OrderCancelReasonSTOPLOSSONFILLGUARANTEEDLEVELRESTRICTIONEXCEEDED captures enum value "STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED"
	OrderCancelReasonSTOPLOSSONFILLGUARANTEEDLEVELRESTRICTIONEXCEEDED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED"

	// OrderCancelReasonSTOPLOSSONFILLGUARANTEEDHEDGINGNOTALLOWED captures enum value "STOP_LOSS_ON_FILL_GUARANTEED_HEDGING_NOT_ALLOWED"
	OrderCancelReasonSTOPLOSSONFILLGUARANTEEDHEDGINGNOTALLOWED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_HEDGING_NOT_ALLOWED"

	// OrderCancelReasonSTOPLOSSONFILLTIMEINFORCEINVALID captures enum value "STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID"
	OrderCancelReasonSTOPLOSSONFILLTIMEINFORCEINVALID OrderCancelReason = "STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID"

	// OrderCancelReasonSTOPLOSSONFILLTRIGGERCONDITIONINVALID captures enum value "STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID"
	OrderCancelReasonSTOPLOSSONFILLTRIGGERCONDITIONINVALID OrderCancelReason = "STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID"

	// OrderCancelReasonTAKEPROFITONFILLPRICEDISTANCEMAXIMUMEXCEEDED captures enum value "TAKE_PROFIT_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"
	OrderCancelReasonTAKEPROFITONFILLPRICEDISTANCEMAXIMUMEXCEEDED OrderCancelReason = "TAKE_PROFIT_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"

	// OrderCancelReasonTRAILINGSTOPLOSSONFILLGTDTIMESTAMPINPAST captures enum value "TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	OrderCancelReasonTRAILINGSTOPLOSSONFILLGTDTIMESTAMPINPAST OrderCancelReason = "TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"

	// OrderCancelReasonCLIENTTRADEIDALREADYEXISTS captures enum value "CLIENT_TRADE_ID_ALREADY_EXISTS"
	OrderCancelReasonCLIENTTRADEIDALREADYEXISTS OrderCancelReason = "CLIENT_TRADE_ID_ALREADY_EXISTS"

	// OrderCancelReasonPOSITIONCLOSEOUTFAILED captures enum value "POSITION_CLOSEOUT_FAILED"
	OrderCancelReasonPOSITIONCLOSEOUTFAILED OrderCancelReason = "POSITION_CLOSEOUT_FAILED"

	// OrderCancelReasonOPENTRADESALLOWEDEXCEEDED captures enum value "OPEN_TRADES_ALLOWED_EXCEEDED"
	OrderCancelReasonOPENTRADESALLOWEDEXCEEDED OrderCancelReason = "OPEN_TRADES_ALLOWED_EXCEEDED"

	// OrderCancelReasonPENDINGORDERSALLOWEDEXCEEDED captures enum value "PENDING_ORDERS_ALLOWED_EXCEEDED"
	OrderCancelReasonPENDINGORDERSALLOWEDEXCEEDED OrderCancelReason = "PENDING_ORDERS_ALLOWED_EXCEEDED"

	// OrderCancelReasonTAKEPROFITONFILLCLIENTORDERIDALREADYEXISTS captures enum value "TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"
	OrderCancelReasonTAKEPROFITONFILLCLIENTORDERIDALREADYEXISTS OrderCancelReason = "TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"

	// OrderCancelReasonSTOPLOSSONFILLCLIENTORDERIDALREADYEXISTS captures enum value "STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"
	OrderCancelReasonSTOPLOSSONFILLCLIENTORDERIDALREADYEXISTS OrderCancelReason = "STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"

	// OrderCancelReasonTRAILINGSTOPLOSSONFILLCLIENTORDERIDALREADYEXISTS captures enum value "TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"
	OrderCancelReasonTRAILINGSTOPLOSSONFILLCLIENTORDERIDALREADYEXISTS OrderCancelReason = "TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"

	// OrderCancelReasonPOSITIONSIZEEXCEEDED captures enum value "POSITION_SIZE_EXCEEDED"
	OrderCancelReasonPOSITIONSIZEEXCEEDED OrderCancelReason = "POSITION_SIZE_EXCEEDED"

	// OrderCancelReasonHEDGINGGSLOVIOLATION captures enum value "HEDGING_GSLO_VIOLATION"
	OrderCancelReasonHEDGINGGSLOVIOLATION OrderCancelReason = "HEDGING_GSLO_VIOLATION"

	// OrderCancelReasonACCOUNTPOSITIONVALUELIMITEXCEEDED captures enum value "ACCOUNT_POSITION_VALUE_LIMIT_EXCEEDED"
	OrderCancelReasonACCOUNTPOSITIONVALUELIMITEXCEEDED OrderCancelReason = "ACCOUNT_POSITION_VALUE_LIMIT_EXCEEDED"

	// OrderCancelReasonINSTRUMENTBIDREDUCEONLY captures enum value "INSTRUMENT_BID_REDUCE_ONLY"
	OrderCancelReasonINSTRUMENTBIDREDUCEONLY OrderCancelReason = "INSTRUMENT_BID_REDUCE_ONLY"

	// OrderCancelReasonINSTRUMENTASKREDUCEONLY captures enum value "INSTRUMENT_ASK_REDUCE_ONLY"
	OrderCancelReasonINSTRUMENTASKREDUCEONLY OrderCancelReason = "INSTRUMENT_ASK_REDUCE_ONLY"

	// OrderCancelReasonINSTRUMENTBIDHALTED captures enum value "INSTRUMENT_BID_HALTED"
	OrderCancelReasonINSTRUMENTBIDHALTED OrderCancelReason = "INSTRUMENT_BID_HALTED"

	// OrderCancelReasonINSTRUMENTASKHALTED captures enum value "INSTRUMENT_ASK_HALTED"
	OrderCancelReasonINSTRUMENTASKHALTED OrderCancelReason = "INSTRUMENT_ASK_HALTED"

	// OrderCancelReasonSTOPLOSSONFILLGUARANTEEDBIDHALTED captures enum value "STOP_LOSS_ON_FILL_GUARANTEED_BID_HALTED"
	OrderCancelReasonSTOPLOSSONFILLGUARANTEEDBIDHALTED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_BID_HALTED"

	// OrderCancelReasonSTOPLOSSONFILLGUARANTEEDASKHALTED captures enum value "STOP_LOSS_ON_FILL_GUARANTEED_ASK_HALTED"
	OrderCancelReasonSTOPLOSSONFILLGUARANTEEDASKHALTED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_ASK_HALTED"
)

// for schema
var orderCancelReasonEnum []interface{}

func init() {
	var res []OrderCancelReason
	if err := json.Unmarshal([]byte(`["INTERNAL_SERVER_ERROR","ACCOUNT_LOCKED","ACCOUNT_NEW_POSITIONS_LOCKED","ACCOUNT_ORDER_CREATION_LOCKED","ACCOUNT_ORDER_FILL_LOCKED","CLIENT_REQUEST","MIGRATION","MARKET_HALTED","LINKED_TRADE_CLOSED","TIME_IN_FORCE_EXPIRED","INSUFFICIENT_MARGIN","FIFO_VIOLATION","BOUNDS_VIOLATION","CLIENT_REQUEST_REPLACED","INSUFFICIENT_LIQUIDITY","TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST","TAKE_PROFIT_ON_FILL_LOSS","LOSING_TAKE_PROFIT","STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST","STOP_LOSS_ON_FILL_LOSS","STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED","STOP_LOSS_ON_FILL_REQUIRED","STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED","STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED","STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET","STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED","STOP_LOSS_ON_FILL_GUARANTEED_HEDGING_NOT_ALLOWED","STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID","STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID","TAKE_PROFIT_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED","TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST","CLIENT_TRADE_ID_ALREADY_EXISTS","POSITION_CLOSEOUT_FAILED","OPEN_TRADES_ALLOWED_EXCEEDED","PENDING_ORDERS_ALLOWED_EXCEEDED","TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS","STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS","TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS","POSITION_SIZE_EXCEEDED","HEDGING_GSLO_VIOLATION","ACCOUNT_POSITION_VALUE_LIMIT_EXCEEDED","INSTRUMENT_BID_REDUCE_ONLY","INSTRUMENT_ASK_REDUCE_ONLY","INSTRUMENT_BID_HALTED","INSTRUMENT_ASK_HALTED","STOP_LOSS_ON_FILL_GUARANTEED_BID_HALTED","STOP_LOSS_ON_FILL_GUARANTEED_ASK_HALTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderCancelReasonEnum = append(orderCancelReasonEnum, v)
	}
}

func (m OrderCancelReason) validateOrderCancelReasonEnum(path, location string, value OrderCancelReason) error {
	if err := validate.Enum(path, location, value, orderCancelReasonEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this order cancel reason
func (m OrderCancelReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrderCancelReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
