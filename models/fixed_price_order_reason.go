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

// FixedPriceOrderReason The reason that the Fixed Price Order was created
// swagger:model FixedPriceOrderReason
type FixedPriceOrderReason string

const (

	// FixedPriceOrderReasonPLATFORMACCOUNTMIGRATION captures enum value "PLATFORM_ACCOUNT_MIGRATION"
	FixedPriceOrderReasonPLATFORMACCOUNTMIGRATION FixedPriceOrderReason = "PLATFORM_ACCOUNT_MIGRATION"
)

// for schema
var fixedPriceOrderReasonEnum []interface{}

func init() {
	var res []FixedPriceOrderReason
	if err := json.Unmarshal([]byte(`["PLATFORM_ACCOUNT_MIGRATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fixedPriceOrderReasonEnum = append(fixedPriceOrderReasonEnum, v)
	}
}

func (m FixedPriceOrderReason) validateFixedPriceOrderReasonEnum(path, location string, value FixedPriceOrderReason) error {
	if err := validate.Enum(path, location, value, fixedPriceOrderReasonEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this fixed price order reason
func (m FixedPriceOrderReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFixedPriceOrderReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
