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

// InstrumentType The type of an Instrument.
// swagger:model InstrumentType
type InstrumentType string

const (

	// InstrumentTypeCURRENCY captures enum value "CURRENCY"
	InstrumentTypeCURRENCY InstrumentType = "CURRENCY"

	// InstrumentTypeCFD captures enum value "CFD"
	InstrumentTypeCFD InstrumentType = "CFD"

	// InstrumentTypeMETAL captures enum value "METAL"
	InstrumentTypeMETAL InstrumentType = "METAL"
)

// for schema
var instrumentTypeEnum []interface{}

func init() {
	var res []InstrumentType
	if err := json.Unmarshal([]byte(`["CURRENCY","CFD","METAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instrumentTypeEnum = append(instrumentTypeEnum, v)
	}
}

func (m InstrumentType) validateInstrumentTypeEnum(path, location string, value InstrumentType) error {
	if err := validate.Enum(path, location, value, instrumentTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this instrument type
func (m InstrumentType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInstrumentTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
