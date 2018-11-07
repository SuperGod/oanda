// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PositionFinancing OpenTradeFinancing is used to pay/collect daily financing charge for a Position within an Account
// swagger:model PositionFinancing
type PositionFinancing struct {

	// The amount of financing paid/collected for the Position.
	Financing string `json:"financing,omitempty"`

	// The instrument of the Position that financing is being paid/collected for.
	Instrument string `json:"instrument,omitempty"`

	// The financing paid/collecte for each open Trade within the Position.
	OpenTradeFinancings []*OpenTradeFinancing `json:"openTradeFinancings"`
}

// Validate validates this position financing
func (m *PositionFinancing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOpenTradeFinancings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PositionFinancing) validateOpenTradeFinancings(formats strfmt.Registry) error {

	if swag.IsZero(m.OpenTradeFinancings) { // not required
		return nil
	}

	for i := 0; i < len(m.OpenTradeFinancings); i++ {
		if swag.IsZero(m.OpenTradeFinancings[i]) { // not required
			continue
		}

		if m.OpenTradeFinancings[i] != nil {
			if err := m.OpenTradeFinancings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("openTradeFinancings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PositionFinancing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PositionFinancing) UnmarshalBinary(b []byte) error {
	var res PositionFinancing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}