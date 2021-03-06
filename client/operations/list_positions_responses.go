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

// ListPositionsReader is a Reader for the ListPositions structure.
type ListPositionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPositionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListPositionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListPositionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListPositionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListPositionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPositionsOK creates a ListPositionsOK with default headers values
func NewListPositionsOK() *ListPositionsOK {
	return &ListPositionsOK{}
}

/*ListPositionsOK handles this case with default header values.

The Account's Positions are provided.
*/
type ListPositionsOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListPositionsOKBody
}

func (o *ListPositionsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/positions][%d] listPositionsOK  %+v", 200, o.Payload)
}

func (o *ListPositionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListPositionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPositionsUnauthorized creates a ListPositionsUnauthorized with default headers values
func NewListPositionsUnauthorized() *ListPositionsUnauthorized {
	return &ListPositionsUnauthorized{}
}

/*ListPositionsUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type ListPositionsUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListPositionsUnauthorizedBody
}

func (o *ListPositionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/positions][%d] listPositionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListPositionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListPositionsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPositionsNotFound creates a ListPositionsNotFound with default headers values
func NewListPositionsNotFound() *ListPositionsNotFound {
	return &ListPositionsNotFound{}
}

/*ListPositionsNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type ListPositionsNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListPositionsNotFoundBody
}

func (o *ListPositionsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/positions][%d] listPositionsNotFound  %+v", 404, o.Payload)
}

func (o *ListPositionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListPositionsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPositionsMethodNotAllowed creates a ListPositionsMethodNotAllowed with default headers values
func NewListPositionsMethodNotAllowed() *ListPositionsMethodNotAllowed {
	return &ListPositionsMethodNotAllowed{}
}

/*ListPositionsMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type ListPositionsMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListPositionsMethodNotAllowedBody
}

func (o *ListPositionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/positions][%d] listPositionsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListPositionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListPositionsMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListPositionsMethodNotAllowedBody list positions method not allowed body
swagger:model ListPositionsMethodNotAllowedBody
*/
type ListPositionsMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list positions method not allowed body
func (o *ListPositionsMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListPositionsMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPositionsMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ListPositionsMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListPositionsNotFoundBody list positions not found body
swagger:model ListPositionsNotFoundBody
*/
type ListPositionsNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list positions not found body
func (o *ListPositionsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListPositionsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPositionsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ListPositionsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListPositionsOKBody list positions o k body
swagger:model ListPositionsOKBody
*/
type ListPositionsOKBody struct {

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// The list of Account Positions.
	Positions []*models.Position `json:"positions"`
}

// Validate validates this list positions o k body
func (o *ListPositionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPositionsOKBody) validatePositions(formats strfmt.Registry) error {

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
					return ve.ValidateName("listPositionsOK" + "." + "positions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListPositionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPositionsOKBody) UnmarshalBinary(b []byte) error {
	var res ListPositionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListPositionsUnauthorizedBody list positions unauthorized body
swagger:model ListPositionsUnauthorizedBody
*/
type ListPositionsUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list positions unauthorized body
func (o *ListPositionsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListPositionsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPositionsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ListPositionsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
