// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LiquidityRegenerationScheduleStep A liquidity regeneration schedule Step indicates the amount of bid and ask liquidity that is used by the Account at a certain time. These amounts will only change at the timestamp of the following step.
// swagger:model LiquidityRegenerationScheduleStep
type LiquidityRegenerationScheduleStep struct {

	// The amount of ask liquidity used at this step in the schedule.
	AskLiquidityUsed string `json:"askLiquidityUsed,omitempty"`

	// The amount of bid liquidity used at this step in the schedule.
	BidLiquidityUsed string `json:"bidLiquidityUsed,omitempty"`

	// The timestamp of the schedule step.
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this liquidity regeneration schedule step
func (m *LiquidityRegenerationScheduleStep) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LiquidityRegenerationScheduleStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LiquidityRegenerationScheduleStep) UnmarshalBinary(b []byte) error {
	var res LiquidityRegenerationScheduleStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
