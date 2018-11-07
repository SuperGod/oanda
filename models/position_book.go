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

// PositionBook The representation of an instrument's position book at a point in time
// swagger:model PositionBook
type PositionBook struct {

	// The price width for each bucket. Each bucket covers the price range from the bucket's price to the bucket's price + bucketWidth.
	BucketWidth string `json:"bucketWidth,omitempty"`

	// The partitioned position book, divided into buckets using a default bucket width. These buckets are only provided for price ranges which actually contain order or position data.
	Buckets []*PositionBookBucket `json:"buckets"`

	// The position book's instrument
	Instrument string `json:"instrument,omitempty"`

	// The price (midpoint) for the position book's instrument at the time of the position book snapshot
	Price string `json:"price,omitempty"`

	// The time when the position book snapshot was created
	Time string `json:"time,omitempty"`
}

// Validate validates this position book
func (m *PositionBook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuckets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PositionBook) validateBuckets(formats strfmt.Registry) error {

	if swag.IsZero(m.Buckets) { // not required
		return nil
	}

	for i := 0; i < len(m.Buckets); i++ {
		if swag.IsZero(m.Buckets[i]) { // not required
			continue
		}

		if m.Buckets[i] != nil {
			if err := m.Buckets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("buckets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PositionBook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PositionBook) UnmarshalBinary(b []byte) error {
	var res PositionBook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
