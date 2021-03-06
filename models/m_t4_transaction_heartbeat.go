// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MT4TransactionHeartbeat A TransactionHeartbeat object is injected into the Transaction stream to ensure that the HTTP connection remains active.
// swagger:model MT4TransactionHeartbeat
type MT4TransactionHeartbeat struct {

	// The date/time when the TransactionHeartbeat was created.
	Time string `json:"time,omitempty"`

	// The string "HEARTBEAT"
	Type string `json:"type,omitempty"`
}

// Validate validates this m t4 transaction heartbeat
func (m *MT4TransactionHeartbeat) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MT4TransactionHeartbeat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MT4TransactionHeartbeat) UnmarshalBinary(b []byte) error {
	var res MT4TransactionHeartbeat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
