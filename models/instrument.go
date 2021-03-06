// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Instrument Full specification of an Instrument.
// swagger:model Instrument
type Instrument struct {

	// commission
	Commission *InstrumentCommission `json:"commission,omitempty"`

	// The display name of the Instrument
	DisplayName string `json:"displayName,omitempty"`

	// The number of decimal places that should be used to display prices for this instrument. (e.g. a displayPrecision of 5 would result in a price of "1" being displayed as "1.00000")
	DisplayPrecision int64 `json:"displayPrecision,omitempty"`

	// The margin rate for this instrument.
	MarginRate string `json:"marginRate,omitempty"`

	// The maximum units allowed for an Order placed for this instrument. Specified in units.
	MaximumOrderUnits string `json:"maximumOrderUnits,omitempty"`

	// The maximum position size allowed for this instrument. Specified in units.
	MaximumPositionSize string `json:"maximumPositionSize,omitempty"`

	// The maximum trailing stop distance allowed for a trailing stop loss created for this instrument. Specified in price units.
	MaximumTrailingStopDistance string `json:"maximumTrailingStopDistance,omitempty"`

	// The smallest number of units allowed to be traded for this instrument.
	MinimumTradeSize string `json:"minimumTradeSize,omitempty"`

	// The minimum trailing stop distance allowed for a trailing stop loss created for this instrument. Specified in price units.
	MinimumTrailingStopDistance string `json:"minimumTrailingStopDistance,omitempty"`

	// The name of the Instrument
	Name string `json:"name,omitempty"`

	// The location of the "pip" for this instrument. The decimal position of the pip in this Instrument's price can be found at 10 ^ pipLocation (e.g. -4 pipLocation results in a decimal pip position of 10 ^ -4 = 0.0001).
	PipLocation int64 `json:"pipLocation,omitempty"`

	// The amount of decimal places that may be provided when specifying the number of units traded for this instrument.
	TradeUnitsPrecision int64 `json:"tradeUnitsPrecision,omitempty"`

	// The type of the Instrument
	// Enum: [CURRENCY CFD METAL]
	Type string `json:"type,omitempty"`
}

// Validate validates this instrument
func (m *Instrument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Instrument) validateCommission(formats strfmt.Registry) error {

	if swag.IsZero(m.Commission) { // not required
		return nil
	}

	if m.Commission != nil {
		if err := m.Commission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commission")
			}
			return err
		}
	}

	return nil
}

var instrumentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CURRENCY","CFD","METAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instrumentTypeTypePropEnum = append(instrumentTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *Instrument) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instrumentTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Instrument) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Instrument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Instrument) UnmarshalBinary(b []byte) error {
	var res Instrument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
