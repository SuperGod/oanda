// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HomeConversions HomeConversions represents the factors to use to convert quantities of a given currency into the Account's home currency. The conversion factor depends on the scenario the conversion is required for.
// swagger:model HomeConversions
type HomeConversions struct {

	// The factor used to convert any gains for an Account in the specified currency into the Account's home currency. This would include positive realized P/L and positive financing amounts. Conversion is performed by multiplying the positive P/L by the conversion factor.
	AccountGain string `json:"accountGain,omitempty"`

	// The string representation of a decimal number.
	AccountLoss string `json:"accountLoss,omitempty"`

	// The currency to be converted into the home currency.
	Currency string `json:"currency,omitempty"`

	// The factor used to convert a Position or Trade Value in the specified currency into the Account's home currency. Conversion is performed by multiplying the Position or Trade Value by the conversion factor.
	PositionValue string `json:"positionValue,omitempty"`
}

// Validate validates this home conversions
func (m *HomeConversions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HomeConversions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HomeConversions) UnmarshalBinary(b []byte) error {
	var res HomeConversions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}