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

// TakeProfitDetails TakeProfitDetails specifies the details of a Take Profit Order to be created on behalf of a client. This may happen when an Order is filled that opens a Trade requiring a Take Profit, or when a Trade's dependent Take Profit Order is modified directly through the Trade.
// swagger:model TakeProfitDetails
type TakeProfitDetails struct {

	// client extensions
	ClientExtensions *ClientExtensions `json:"clientExtensions,omitempty"`

	// The date when the Take Profit Order will be cancelled on if timeInForce is GTD.
	GtdTime string `json:"gtdTime,omitempty"`

	// The price that the Take Profit Order will be triggered at. Only one of the price and distance fields may be specified.
	Price string `json:"price,omitempty"`

	// The time in force for the created Take Profit Order. This may only be GTC, GTD or GFD.
	// Enum: [GTC GTD GFD FOK IOC]
	TimeInForce string `json:"timeInForce,omitempty"`
}

// Validate validates this take profit details
func (m *TakeProfitDetails) Validate(formats strfmt.Registry) error {
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

func (m *TakeProfitDetails) validateClientExtensions(formats strfmt.Registry) error {

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

var takeProfitDetailsTypeTimeInForcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GTC","GTD","GFD","FOK","IOC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		takeProfitDetailsTypeTimeInForcePropEnum = append(takeProfitDetailsTypeTimeInForcePropEnum, v)
	}
}

const (

	// TakeProfitDetailsTimeInForceGTC captures enum value "GTC"
	TakeProfitDetailsTimeInForceGTC string = "GTC"

	// TakeProfitDetailsTimeInForceGTD captures enum value "GTD"
	TakeProfitDetailsTimeInForceGTD string = "GTD"

	// TakeProfitDetailsTimeInForceGFD captures enum value "GFD"
	TakeProfitDetailsTimeInForceGFD string = "GFD"

	// TakeProfitDetailsTimeInForceFOK captures enum value "FOK"
	TakeProfitDetailsTimeInForceFOK string = "FOK"

	// TakeProfitDetailsTimeInForceIOC captures enum value "IOC"
	TakeProfitDetailsTimeInForceIOC string = "IOC"
)

// prop value enum
func (m *TakeProfitDetails) validateTimeInForceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, takeProfitDetailsTypeTimeInForcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TakeProfitDetails) validateTimeInForce(formats strfmt.Registry) error {

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
func (m *TakeProfitDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TakeProfitDetails) UnmarshalBinary(b []byte) error {
	var res TakeProfitDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
