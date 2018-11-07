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

// ListAccountsReader is a Reader for the ListAccounts structure.
type ListAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListAccountsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAccountsOK creates a ListAccountsOK with default headers values
func NewListAccountsOK() *ListAccountsOK {
	return &ListAccountsOK{}
}

/*ListAccountsOK handles this case with default header values.

The list of authorized Accounts has been provided.
*/
type ListAccountsOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListAccountsOKBody
}

func (o *ListAccountsOK) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] listAccountsOK  %+v", 200, o.Payload)
}

func (o *ListAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListAccountsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountsUnauthorized creates a ListAccountsUnauthorized with default headers values
func NewListAccountsUnauthorized() *ListAccountsUnauthorized {
	return &ListAccountsUnauthorized{}
}

/*ListAccountsUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type ListAccountsUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListAccountsUnauthorizedBody
}

func (o *ListAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] listAccountsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListAccountsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountsMethodNotAllowed creates a ListAccountsMethodNotAllowed with default headers values
func NewListAccountsMethodNotAllowed() *ListAccountsMethodNotAllowed {
	return &ListAccountsMethodNotAllowed{}
}

/*ListAccountsMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type ListAccountsMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListAccountsMethodNotAllowedBody
}

func (o *ListAccountsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] listAccountsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListAccountsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListAccountsMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListAccountsMethodNotAllowedBody list accounts method not allowed body
swagger:model ListAccountsMethodNotAllowedBody
*/
type ListAccountsMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list accounts method not allowed body
func (o *ListAccountsMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAccountsMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAccountsMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ListAccountsMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListAccountsOKBody list accounts o k body
swagger:model ListAccountsOKBody
*/
type ListAccountsOKBody struct {

	// The list of Accounts the client is authorized to access and their associated properties.
	Accounts []*models.AccountProperties `json:"accounts"`
}

// Validate validates this list accounts o k body
func (o *ListAccountsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListAccountsOKBody) validateAccounts(formats strfmt.Registry) error {

	if swag.IsZero(o.Accounts) { // not required
		return nil
	}

	for i := 0; i < len(o.Accounts); i++ {
		if swag.IsZero(o.Accounts[i]) { // not required
			continue
		}

		if o.Accounts[i] != nil {
			if err := o.Accounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listAccountsOK" + "." + "accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListAccountsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAccountsOKBody) UnmarshalBinary(b []byte) error {
	var res ListAccountsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListAccountsUnauthorizedBody list accounts unauthorized body
swagger:model ListAccountsUnauthorizedBody
*/
type ListAccountsUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list accounts unauthorized body
func (o *ListAccountsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListAccountsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListAccountsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ListAccountsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
