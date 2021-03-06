// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SuperGod/oanda/models"
)

// GetInstrumentCandlesReader is a Reader for the GetInstrumentCandles structure.
type GetInstrumentCandlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstrumentCandlesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstrumentCandlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstrumentCandlesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstrumentCandlesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstrumentCandlesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstrumentCandlesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstrumentCandlesOK creates a GetInstrumentCandlesOK with default headers values
func NewGetInstrumentCandlesOK() *GetInstrumentCandlesOK {
	return &GetInstrumentCandlesOK{}
}

/*GetInstrumentCandlesOK handles this case with default header values.

Pricing information has been successfully provided.
*/
type GetInstrumentCandlesOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentCandlesOKBody
}

func (o *GetInstrumentCandlesOK) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/candles][%d] getInstrumentCandlesOK  %+v", 200, o.Payload)
}

func (o *GetInstrumentCandlesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentCandlesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentCandlesBadRequest creates a GetInstrumentCandlesBadRequest with default headers values
func NewGetInstrumentCandlesBadRequest() *GetInstrumentCandlesBadRequest {
	return &GetInstrumentCandlesBadRequest{}
}

/*GetInstrumentCandlesBadRequest handles this case with default header values.

Bad Request. The client has provided invalid data to be processed by the server.
*/
type GetInstrumentCandlesBadRequest struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentCandlesBadRequestBody
}

func (o *GetInstrumentCandlesBadRequest) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/candles][%d] getInstrumentCandlesBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstrumentCandlesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentCandlesBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentCandlesUnauthorized creates a GetInstrumentCandlesUnauthorized with default headers values
func NewGetInstrumentCandlesUnauthorized() *GetInstrumentCandlesUnauthorized {
	return &GetInstrumentCandlesUnauthorized{}
}

/*GetInstrumentCandlesUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type GetInstrumentCandlesUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentCandlesUnauthorizedBody
}

func (o *GetInstrumentCandlesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/candles][%d] getInstrumentCandlesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstrumentCandlesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentCandlesUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentCandlesNotFound creates a GetInstrumentCandlesNotFound with default headers values
func NewGetInstrumentCandlesNotFound() *GetInstrumentCandlesNotFound {
	return &GetInstrumentCandlesNotFound{}
}

/*GetInstrumentCandlesNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type GetInstrumentCandlesNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentCandlesNotFoundBody
}

func (o *GetInstrumentCandlesNotFound) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/candles][%d] getInstrumentCandlesNotFound  %+v", 404, o.Payload)
}

func (o *GetInstrumentCandlesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentCandlesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentCandlesMethodNotAllowed creates a GetInstrumentCandlesMethodNotAllowed with default headers values
func NewGetInstrumentCandlesMethodNotAllowed() *GetInstrumentCandlesMethodNotAllowed {
	return &GetInstrumentCandlesMethodNotAllowed{}
}

/*GetInstrumentCandlesMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type GetInstrumentCandlesMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentCandlesMethodNotAllowedBody
}

func (o *GetInstrumentCandlesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/candles][%d] getInstrumentCandlesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetInstrumentCandlesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentCandlesMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInstrumentCandlesBadRequestBody get instrument candles bad request body
swagger:model GetInstrumentCandlesBadRequestBody
*/
type GetInstrumentCandlesBadRequestBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instrument candles bad request body
func (o *GetInstrumentCandlesBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentCandlesBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentCandlesBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentCandlesBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentCandlesMethodNotAllowedBody get instrument candles method not allowed body
swagger:model GetInstrumentCandlesMethodNotAllowedBody
*/
type GetInstrumentCandlesMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instrument candles method not allowed body
func (o *GetInstrumentCandlesMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentCandlesMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentCandlesMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentCandlesMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentCandlesNotFoundBody get instrument candles not found body
swagger:model GetInstrumentCandlesNotFoundBody
*/
type GetInstrumentCandlesNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instrument candles not found body
func (o *GetInstrumentCandlesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentCandlesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentCandlesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentCandlesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentCandlesOKBody get instrument candles o k body
swagger:model GetInstrumentCandlesOKBody
*/
type GetInstrumentCandlesOKBody struct {

	// The list of candlesticks that satisfy the request.
	Candles []*models.Candlestick `json:"candles"`

	// The granularity of the candlesticks provided.
	// Enum: [S5 S10 S15 S30 M1 M2 M4 M5 M10 M15 M30 H1 H2 H3 H4 H6 H8 H12 D W M]
	Granularity string `json:"granularity,omitempty"`

	// The instrument whose Prices are represented by the candlesticks.
	Instrument string `json:"instrument,omitempty"`
}

// Validate validates this get instrument candles o k body
func (o *GetInstrumentCandlesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCandles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGranularity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInstrumentCandlesOKBody) validateCandles(formats strfmt.Registry) error {

	if swag.IsZero(o.Candles) { // not required
		return nil
	}

	for i := 0; i < len(o.Candles); i++ {
		if swag.IsZero(o.Candles[i]) { // not required
			continue
		}

		if o.Candles[i] != nil {
			if err := o.Candles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getInstrumentCandlesOK" + "." + "candles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getInstrumentCandlesOKBodyTypeGranularityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["S5","S10","S15","S30","M1","M2","M4","M5","M10","M15","M30","H1","H2","H3","H4","H6","H8","H12","D","W","M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getInstrumentCandlesOKBodyTypeGranularityPropEnum = append(getInstrumentCandlesOKBodyTypeGranularityPropEnum, v)
	}
}

const (

	// GetInstrumentCandlesOKBodyGranularityS5 captures enum value "S5"
	GetInstrumentCandlesOKBodyGranularityS5 string = "S5"

	// GetInstrumentCandlesOKBodyGranularityS10 captures enum value "S10"
	GetInstrumentCandlesOKBodyGranularityS10 string = "S10"

	// GetInstrumentCandlesOKBodyGranularityS15 captures enum value "S15"
	GetInstrumentCandlesOKBodyGranularityS15 string = "S15"

	// GetInstrumentCandlesOKBodyGranularityS30 captures enum value "S30"
	GetInstrumentCandlesOKBodyGranularityS30 string = "S30"

	// GetInstrumentCandlesOKBodyGranularityM1 captures enum value "M1"
	GetInstrumentCandlesOKBodyGranularityM1 string = "M1"

	// GetInstrumentCandlesOKBodyGranularityM2 captures enum value "M2"
	GetInstrumentCandlesOKBodyGranularityM2 string = "M2"

	// GetInstrumentCandlesOKBodyGranularityM4 captures enum value "M4"
	GetInstrumentCandlesOKBodyGranularityM4 string = "M4"

	// GetInstrumentCandlesOKBodyGranularityM5 captures enum value "M5"
	GetInstrumentCandlesOKBodyGranularityM5 string = "M5"

	// GetInstrumentCandlesOKBodyGranularityM10 captures enum value "M10"
	GetInstrumentCandlesOKBodyGranularityM10 string = "M10"

	// GetInstrumentCandlesOKBodyGranularityM15 captures enum value "M15"
	GetInstrumentCandlesOKBodyGranularityM15 string = "M15"

	// GetInstrumentCandlesOKBodyGranularityM30 captures enum value "M30"
	GetInstrumentCandlesOKBodyGranularityM30 string = "M30"

	// GetInstrumentCandlesOKBodyGranularityH1 captures enum value "H1"
	GetInstrumentCandlesOKBodyGranularityH1 string = "H1"

	// GetInstrumentCandlesOKBodyGranularityH2 captures enum value "H2"
	GetInstrumentCandlesOKBodyGranularityH2 string = "H2"

	// GetInstrumentCandlesOKBodyGranularityH3 captures enum value "H3"
	GetInstrumentCandlesOKBodyGranularityH3 string = "H3"

	// GetInstrumentCandlesOKBodyGranularityH4 captures enum value "H4"
	GetInstrumentCandlesOKBodyGranularityH4 string = "H4"

	// GetInstrumentCandlesOKBodyGranularityH6 captures enum value "H6"
	GetInstrumentCandlesOKBodyGranularityH6 string = "H6"

	// GetInstrumentCandlesOKBodyGranularityH8 captures enum value "H8"
	GetInstrumentCandlesOKBodyGranularityH8 string = "H8"

	// GetInstrumentCandlesOKBodyGranularityH12 captures enum value "H12"
	GetInstrumentCandlesOKBodyGranularityH12 string = "H12"

	// GetInstrumentCandlesOKBodyGranularityD captures enum value "D"
	GetInstrumentCandlesOKBodyGranularityD string = "D"

	// GetInstrumentCandlesOKBodyGranularityW captures enum value "W"
	GetInstrumentCandlesOKBodyGranularityW string = "W"

	// GetInstrumentCandlesOKBodyGranularityM captures enum value "M"
	GetInstrumentCandlesOKBodyGranularityM string = "M"
)

// prop value enum
func (o *GetInstrumentCandlesOKBody) validateGranularityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getInstrumentCandlesOKBodyTypeGranularityPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetInstrumentCandlesOKBody) validateGranularity(formats strfmt.Registry) error {

	if swag.IsZero(o.Granularity) { // not required
		return nil
	}

	// value enum
	if err := o.validateGranularityEnum("getInstrumentCandlesOK"+"."+"granularity", "body", o.Granularity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentCandlesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentCandlesOKBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentCandlesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentCandlesUnauthorizedBody get instrument candles unauthorized body
swagger:model GetInstrumentCandlesUnauthorizedBody
*/
type GetInstrumentCandlesUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instrument candles unauthorized body
func (o *GetInstrumentCandlesUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentCandlesUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentCandlesUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentCandlesUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
