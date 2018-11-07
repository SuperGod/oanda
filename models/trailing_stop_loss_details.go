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

// TrailingStopLossDetails TrailingStopLossDetails specifies the details of a Trailing Stop Loss Order to be created on behalf of a client. This may happen when an Order is filled that opens a Trade requiring a Trailing Stop Loss, or when a Trade's dependent Trailing Stop Loss Order is modified directly through the Trade.
// swagger:model TrailingStopLossDetails
type TrailingStopLossDetails struct {

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The distance (in price units) from the Trade's fill price that the Trailing Stop Loss Order will be triggered at.
	Distance string `json:"distance,omitempty"`

	// The date when the Trailing Stop Loss Order will be cancelled on if timeInForce is GTD.
	GtdTime string `json:"gtdTime,omitempty"`

	// The time in force for the created Trailing Stop Loss Order. This may only be GTC, GTD or GFD.
	// Enum: [GTC GTD GFD FOK IOC]
	TimeInForce string `json:"timeInForce,omitempty"`
}

// Validate validates this trailing stop loss details
func (m *TrailingStopLossDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeInForce(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrailingStopLossDetails) validateClientExtensions(formats strfmt.Registry) error {

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

var trailingStopLossDetailsTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trailingStopLossDetailsTypeTimeInForcePropEnum = append(trailingStopLossDetailsTypeTimeInForcePropEnum, v)
	}
}

const (

	// TrailingStopLossDetailsTimeInForceGTC captures enum value "GTC"
	TrailingStopLossDetailsTimeInForceGTC string = "GTC"

	// TrailingStopLossDetailsTimeInForceGTD captures enum value "GTD"
	TrailingStopLossDetailsTimeInForceGTD string = "GTD"

	// TrailingStopLossDetailsTimeInForceGFD captures enum value "GFD"
	TrailingStopLossDetailsTimeInForceGFD string = "GFD"

	// TrailingStopLossDetailsTimeInForceFOK captures enum value "FOK"
	TrailingStopLossDetailsTimeInForceFOK string = "FOK"

	// TrailingStopLossDetailsTimeInForceIOC captures enum value "IOC"
	TrailingStopLossDetailsTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *TrailingStopLossDetails) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, trailingStopLossDetailsTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TrailingStopLossDetails) validateTimeInForce(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInForce) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimeInForceEnum("timeInForce", "body", m.TimeInForce); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrailingStopLossDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrailingStopLossDetails) UnmarshalBinary(b []byte) error {
	var res TrailingStopLossDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
