// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SuperGod/oanda/models"
)

// ListOpenPositionsReader is a Reader for the ListOpenPositions structure.
type ListOpenPositionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenPositionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListOpenPositionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListOpenPositionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListOpenPositionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListOpenPositionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListOpenPositionsOK creates a ListOpenPositionsOK with default headers values
func NewListOpenPositionsOK() *ListOpenPositionsOK {
	return &ListOpenPositionsOK{}
}

/*ListOpenPositionsOK handles this case with default header values.

The Account's open Positions are provided.
*/
type ListOpenPositionsOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListOpenPositionsOKBody
}

func (o *ListOpenPositionsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/openPositions][%d] listOpenPositionsOK  %+v", 200, o.Payload)
}

func (o *ListOpenPositionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListOpenPositionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenPositionsUnauthorized creates a ListOpenPositionsUnauthorized with default headers values
func NewListOpenPositionsUnauthorized() *ListOpenPositionsUnauthorized {
	return &ListOpenPositionsUnauthorized{}
}

/*ListOpenPositionsUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type ListOpenPositionsUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListOpenPositionsUnauthorizedBody
}

func (o *ListOpenPositionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/openPositions][%d] listOpenPositionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListOpenPositionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListOpenPositionsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenPositionsNotFound creates a ListOpenPositionsNotFound with default headers values
func NewListOpenPositionsNotFound() *ListOpenPositionsNotFound {
	return &ListOpenPositionsNotFound{}
}

/*ListOpenPositionsNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type ListOpenPositionsNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListOpenPositionsNotFoundBody
}

func (o *ListOpenPositionsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/openPositions][%d] listOpenPositionsNotFound  %+v", 404, o.Payload)
}

func (o *ListOpenPositionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListOpenPositionsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenPositionsMethodNotAllowed creates a ListOpenPositionsMethodNotAllowed with default headers values
func NewListOpenPositionsMethodNotAllowed() *ListOpenPositionsMethodNotAllowed {
	return &ListOpenPositionsMethodNotAllowed{}
}

/*ListOpenPositionsMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type ListOpenPositionsMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListOpenPositionsMethodNotAllowedBody
}

func (o *ListOpenPositionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/openPositions][%d] listOpenPositionsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListOpenPositionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListOpenPositionsMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListOpenPositionsMethodNotAllowedBody list open positions method not allowed body
swagger:model ListOpenPositionsMethodNotAllowedBody
*/
type ListOpenPositionsMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list open positions method not allowed body
func (o *ListOpenPositionsMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListOpenPositionsMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOpenPositionsMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ListOpenPositionsMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListOpenPositionsNotFoundBody list open positions not found body
swagger:model ListOpenPositionsNotFoundBody
*/
type ListOpenPositionsNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list open positions not found body
func (o *ListOpenPositionsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListOpenPositionsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOpenPositionsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ListOpenPositionsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListOpenPositionsOKBody list open positions o k body
swagger:model ListOpenPositionsOKBody
*/
type ListOpenPositionsOKBody struct {

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// The list of open Positions in the Account.
	Positions []*models.Position `json:"positions"`
}

// Validate validates this list open positions o k body
func (o *ListOpenPositionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListOpenPositionsOKBody) validatePositions(formats strfmt.Registry) error {

	if swag.IsZero(o.Positions) { // not required
		return nil
	}

	for i := 0; i < len(o.Positions); i++ {
		if swag.IsZero(o.Positions[i]) { // not required
			continue
		}

		if o.Positions[i] != nil {
			if err := o.Positions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listOpenPositionsOK" + "." + "positions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListOpenPositionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOpenPositionsOKBody) UnmarshalBinary(b []byte) error {
	var res ListOpenPositionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListOpenPositionsUnauthorizedBody list open positions unauthorized body
swagger:model ListOpenPositionsUnauthorizedBody
*/
type ListOpenPositionsUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list open positions unauthorized body
func (o *ListOpenPositionsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListOpenPositionsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOpenPositionsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ListOpenPositionsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
