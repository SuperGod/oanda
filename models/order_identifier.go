// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OrderIdentifier An OrderIdentifier is used to refer to an Order, and contains both the OrderID and the ClientOrderID.
// swagger:model OrderIdentifier
type OrderIdentifier struct {

	// The client-provided client Order ID
	ClientOrderID string `json:"clientOrderID,omitempty"`

	// The OANDA-assigned Order ID
	OrderID string `json:"orderID,omitempty"`
}

// Validate validates this order identifier
func (m *OrderIdentifier) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderIdentifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderIdentifier) UnmarshalBinary(b []byte) error {
	var res OrderIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
