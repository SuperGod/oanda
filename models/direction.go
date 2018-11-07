// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Direction In the context of an Order or a Trade, defines whether the units are positive or negative.
// swagger:model Direction
type Direction string

const (

	// DirectionLONG captures enum value "LONG"
	DirectionLONG Direction = "LONG"

	// DirectionSHORT captures enum value "SHORT"
	DirectionSHORT Direction = "SHORT"
)

// for schema
var directionEnum []interface{}

func init() {
	var res []Direction
	if err := json.Unmarshal([]byte(`["LONG","SHORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directionEnum = append(directionEnum, v)
	}
}

func (m Direction) validateDirectionEnum(path, location string, value Direction) error {
	if err := validate.Enum(path, location, value, directionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this direction
func (m Direction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
