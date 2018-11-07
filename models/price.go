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

// Price The Price representation
// swagger:model Price
type Price struct {

	// The list of prices and liquidity available on the Instrument's ask side. It is possible for this list to be empty if there is no ask liquidity currently available for the Instrument in the Account.
	Asks []*PriceBucket `json:"asks"`

	// The base ask price as calculated by pricing.
	BaseAsk string `json:"baseAsk,omitempty"`

	// The base bid price as calculated by pricing.
	BaseBid string `json:"baseBid,omitempty"`

	// The list of prices and liquidity available on the Instrument's bid side. It is possible for this list to be empty if there is no bid liquidity currently available for the Instrument in the Account.
	Bids []*PriceBucket `json:"bids"`

	// The closeout ask price. This price is used when an ask is required to closeout a Position (margin closeout or manual) yet there is no ask liquidity. The closeout ask is never used to open a new position.
	CloseoutAsk string `json:"closeoutAsk,omitempty"`

	// The closeout bid price. This price is used when a bid is required to closeout a Position (margin closeout or manual) yet there is no bid liquidity. The closeout bid is never used to open a new position.
	CloseoutBid string `json:"closeoutBid,omitempty"`

	// The Price's Instrument.
	Instrument string `json:"instrument,omitempty"`

	// The date/time when the Price was created.
	Timestamp string `json:"timestamp,omitempty"`

	// Flag indicating if the Price is tradeable or not
	Tradeable bool `json:"tradeable,omitempty"`
}

// Validate validates this price
func (m *Price) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBids(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Price) validateAsks(formats strfmt.Registry) error {

	if swag.IsZero(m.Asks) { // not required
		return nil
	}

	for i := 0; i < len(m.Asks); i++ {
		if swag.IsZero(m.Asks[i]) { // not required
			continue
		}

		if m.Asks[i] != nil {
			if err := m.Asks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("asks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Price) validateBids(formats strfmt.Registry) error {

	if swag.IsZero(m.Bids) { // not required
		return nil
	}

	for i := 0; i < len(m.Bids); i++ {
		if swag.IsZero(m.Bids[i]) { // not required
			continue
		}

		if m.Bids[i] != nil {
			if err := m.Bids[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Price) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Price) UnmarshalBinary(b []byte) error {
	var res Price
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}