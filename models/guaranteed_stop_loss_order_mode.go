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

// GuaranteedStopLossOrderMode The overall behaviour of the Account regarding guaranteed Stop Loss Orders.
// swagger:model GuaranteedStopLossOrderMode
type GuaranteedStopLossOrderMode string

const (

	// GuaranteedStopLossOrderModeDISABLED captures enum value "DISABLED"
	GuaranteedStopLossOrderModeDISABLED GuaranteedStopLossOrderMode = "DISABLED"

	// GuaranteedStopLossOrderModeALLOWED captures enum value "ALLOWED"
	GuaranteedStopLossOrderModeALLOWED GuaranteedStopLossOrderMode = "ALLOWED"

	// GuaranteedStopLossOrderModeREQUIRED captures enum value "REQUIRED"
	GuaranteedStopLossOrderModeREQUIRED GuaranteedStopLossOrderMode = "REQUIRED"
)

// for schema
var guaranteedStopLossOrderModeEnum []interface{}

func init() {
	var res []GuaranteedStopLossOrderMode
	if err := json.Unmarshal([]byte(`["DISABLED","ALLOWED","REQUIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		guaranteedStopLossOrderModeEnum = append(guaranteedStopLossOrderModeEnum, v)
	}
}

func (m GuaranteedStopLossOrderMode) validateGuaranteedStopLossOrderModeEnum(path, location string, value GuaranteedStopLossOrderMode) error {
	if err := validate.Enum(path, location, value, guaranteedStopLossOrderModeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this guaranteed stop loss order mode
func (m GuaranteedStopLossOrderMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGuaranteedStopLossOrderModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}