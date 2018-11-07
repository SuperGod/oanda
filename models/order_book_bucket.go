// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OrderBookBucket The order book data for a partition of the instrument's prices.
// swagger:model OrderBookBucket
type OrderBookBucket struct {

	// The percentage of the total number of orders represented by the long orders found in this bucket.
	LongCountPercent string `json:"longCountPercent,omitempty"`

	// The lowest price (inclusive) covered by the bucket. The bucket covers the price range from the price to price + the order book's bucketWidth.
	Price string `json:"price,omitempty"`

	// The percentage of the total number of orders represented by the short orders found in this bucket.
	ShortCountPercent string `json:"shortCountPercent,omitempty"`
}

// Validate validates this order book bucket
func (m *OrderBookBucket) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderBookBucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderBookBucket) UnmarshalBinary(b []byte) error {
	var res OrderBookBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
