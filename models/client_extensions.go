// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClientExtensions A ClientExtensions object allows a client to attach a clientID, tag and comment to Orders and Trades in their Account.  Do not set, modify, or delete this field if your account is associated with MT4.
// swagger:model ClientExtensions
type ClientExtensions struct {

	// A comment associated with the Order/Trade
	Comment string `json:"comment,omitempty"`

	// The Client ID of the Order/Trade
	ID string `json:"id,omitempty"`

	// A tag associated with the Order/Trade
	Tag string `json:"tag,omitempty"`
}

// Validate validates this client extensions
func (m *ClientExtensions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientExtensions) UnmarshalBinary(b []byte) error {
	var res ClientExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
