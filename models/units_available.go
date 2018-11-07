// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UnitsAvailable Representation of how many units of an Instrument are available to be traded by an Order depending on its postionFill option.
// swagger:model UnitsAvailable
type UnitsAvailable struct {

	// default
	Default *UnitsAvailableDetails `json:"default,omitempty"`

	// open only
	OpenOnly *UnitsAvailableDetails `json:"openOnly,omitempty"`

	// reduce first
	ReduceFirst *UnitsAvailableDetails `json:"reduceFirst,omitempty"`

	// reduce only
	ReduceOnly *UnitsAvailableDetails `json:"reduceOnly,omitempty"`
}

// Validate validates this units available
func (m *UnitsAvailable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReduceFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReduceOnly(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnitsAvailable) validateDefault(formats strfmt.Registry) error {

	if swag.IsZero(m.Default) { // not required
		return nil
	}

	if m.Default != nil {
		if err := m.Default.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default")
			}
			return err
		}
	}

	return nil
}

func (m *UnitsAvailable) validateOpenOnly(formats strfmt.Registry) error {

	if swag.IsZero(m.OpenOnly) { // not required
		return nil
	}

	if m.OpenOnly != nil {
		if err := m.OpenOnly.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openOnly")
			}
			return err
		}
	}

	return nil
}

func (m *UnitsAvailable) validateReduceFirst(formats strfmt.Registry) error {

	if swag.IsZero(m.ReduceFirst) { // not required
		return nil
	}

	if m.ReduceFirst != nil {
		if err := m.ReduceFirst.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reduceFirst")
			}
			return err
		}
	}

	return nil
}

func (m *UnitsAvailable) validateReduceOnly(formats strfmt.Registry) error {

	if swag.IsZero(m.ReduceOnly) { // not required
		return nil
	}

	if m.ReduceOnly != nil {
		if err := m.ReduceOnly.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reduceOnly")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnitsAvailable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnitsAvailable) UnmarshalBinary(b []byte) error {
	var res UnitsAvailable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
