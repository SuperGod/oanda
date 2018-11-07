// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SuperGod/oanda/models"
)

// GetInstrumentPriceReader is a Reader for the GetInstrumentPrice structure.
type GetInstrumentPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstrumentPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstrumentPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstrumentPriceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstrumentPriceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstrumentPriceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstrumentPriceMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstrumentPriceOK creates a GetInstrumentPriceOK with default headers values
func NewGetInstrumentPriceOK() *GetInstrumentPriceOK {
	return &GetInstrumentPriceOK{}
}

/*GetInstrumentPriceOK handles this case with default header values.

Pricing information has been successfully provided.
*/
type GetInstrumentPriceOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentPriceOKBody
}

func (o *GetInstrumentPriceOK) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/price][%d] getInstrumentPriceOK  %+v", 200, o.Payload)
}

func (o *GetInstrumentPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentPriceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentPriceBadRequest creates a GetInstrumentPriceBadRequest with default headers values
func NewGetInstrumentPriceBadRequest() *GetInstrumentPriceBadRequest {
	return &GetInstrumentPriceBadRequest{}
}

/*GetInstrumentPriceBadRequest handles this case with default header values.

Bad Request. The client has provided invalid data to be processed by the server.
*/
type GetInstrumentPriceBadRequest struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentPriceBadRequestBody
}

func (o *GetInstrumentPriceBadRequest) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/price][%d] getInstrumentPriceBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstrumentPriceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentPriceBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentPriceUnauthorized creates a GetInstrumentPriceUnauthorized with default headers values
func NewGetInstrumentPriceUnauthorized() *GetInstrumentPriceUnauthorized {
	return &GetInstrumentPriceUnauthorized{}
}

/*GetInstrumentPriceUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type GetInstrumentPriceUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentPriceUnauthorizedBody
}

func (o *GetInstrumentPriceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/price][%d] getInstrumentPriceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstrumentPriceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentPriceUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentPriceNotFound creates a GetInstrumentPriceNotFound with default headers values
func NewGetInstrumentPriceNotFound() *GetInstrumentPriceNotFound {
	return &GetInstrumentPriceNotFound{}
}

/*GetInstrumentPriceNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type GetInstrumentPriceNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentPriceNotFoundBody
}

func (o *GetInstrumentPriceNotFound) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/price][%d] getInstrumentPriceNotFound  %+v", 404, o.Payload)
}

func (o *GetInstrumentPriceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentPriceNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentPriceMethodNotAllowed creates a GetInstrumentPriceMethodNotAllowed with default headers values
func NewGetInstrumentPriceMethodNotAllowed() *GetInstrumentPriceMethodNotAllowed {
	return &GetInstrumentPriceMethodNotAllowed{}
}

/*GetInstrumentPriceMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type GetInstrumentPriceMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentPriceMethodNotAllowedBody
}

func (o *GetInstrumentPriceMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/price][%d] getInstrumentPriceMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetInstrumentPriceMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentPriceMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInstrumentPriceBadRequestBody get instrument price bad request body
swagger:model GetInstrumentPriceBadRequestBody
*/
type GetInstrumentPriceBadRequestBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instrument price bad request body
func (o *GetInstrumentPriceBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentPriceBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentPriceBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentPriceBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentPriceMethodNotAllowedBody get instrument price method not allowed body
swagger:model GetInstrumentPriceMethodNotAllowedBody
*/
type GetInstrumentPriceMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instrument price method not allowed body
func (o *GetInstrumentPriceMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentPriceMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentPriceMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentPriceMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentPriceNotFoundBody get instrument price not found body
swagger:model GetInstrumentPriceNotFoundBody
*/
type GetInstrumentPriceNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instrument price not found body
func (o *GetInstrumentPriceNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentPriceNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentPriceNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentPriceNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentPriceOKBody get instrument price o k body
swagger:model GetInstrumentPriceOKBody
*/
type GetInstrumentPriceOKBody struct {

	// price
	Price *models.Price `json:"price,omitempty"`
}

// Validate validates this get instrument price o k body
func (o *GetInstrumentPriceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInstrumentPriceOKBody) validatePrice(formats strfmt.Registry) error {

	if swag.IsZero(o.Price) { // not required
		return nil
	}

	if o.Price != nil {
		if err := o.Price.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInstrumentPriceOK" + "." + "price")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentPriceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentPriceOKBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentPriceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentPriceUnauthorizedBody get instrument price unauthorized body
swagger:model GetInstrumentPriceUnauthorizedBody
*/
type GetInstrumentPriceUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instrument price unauthorized body
func (o *GetInstrumentPriceUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentPriceUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentPriceUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentPriceUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}