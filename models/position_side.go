// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PositionSide The representation of a Position for a single direction (long or short).
// swagger:model PositionSide
type PositionSide struct {

	// Volume-weighted average of the underlying Trade open prices for the Position.
	AveragePrice string `json:"averagePrice,omitempty"`

	// The total amount of financing paid/collected for this PositionSide over the lifetime of the Account.
	Financing string `json:"financing,omitempty"`

	// The total amount of fees charged over the lifetime of the Account for the execution of guaranteed Stop Loss Orders attached to Trades for this PositionSide.
	GuaranteedExecutionFees string `json:"guaranteedExecutionFees,omitempty"`

	// Profit/loss realized by the PositionSide over the lifetime of the Account.
	Pl string `json:"pl,omitempty"`

	// Profit/loss realized by the PositionSide since the Account's resettablePL was last reset by the client.
	ResettablePL string `json:"resettablePL,omitempty"`

	// List of the open Trade IDs which contribute to the open Position.
	TradeIds []string `json:"tradeIDs"`

	// Number of units in the position (negative value indicates short position, positive indicates long position).
	Units string `json:"units,omitempty"`

	// The unrealized profit/loss of all open Trades that contribute to this PositionSide.
	UnrealizedPL string `json:"unrealizedPL,omitempty"`
}

// Validate validates this position side
func (m *PositionSide) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PositionSide) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PositionSide) UnmarshalBinary(b []byte) error {
	var res PositionSide
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}