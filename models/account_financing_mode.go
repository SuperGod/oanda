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

// AccountFinancingMode The financing mode of an Account
// swagger:model AccountFinancingMode
type AccountFinancingMode string

const (

	// AccountFinancingModeNOFINANCING captures enum value "NO_FINANCING"
	AccountFinancingModeNOFINANCING AccountFinancingMode = "NO_FINANCING"

	// AccountFinancingModeSECONDBYSECOND captures enum value "SECOND_BY_SECOND"
	AccountFinancingModeSECONDBYSECOND AccountFinancingMode = "SECOND_BY_SECOND"

	// AccountFinancingModeDAILY captures enum value "DAILY"
	AccountFinancingModeDAILY AccountFinancingMode = "DAILY"
)

// for schema
var accountFinancingModeEnum []interface{}

func init() {
	var res []AccountFinancingMode
	if err := json.Unmarshal([]byte(`["NO_FINANCING","SECOND_BY_SECOND","DAILY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountFinancingModeEnum = append(accountFinancingModeEnum, v)
	}
}

func (m AccountFinancingMode) validateAccountFinancingModeEnum(path, location string, value AccountFinancingMode) error {
	if err := validate.Enum(path, location, value, accountFinancingModeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this account financing mode
func (m AccountFinancingMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountFinancingModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
