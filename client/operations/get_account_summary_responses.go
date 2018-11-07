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

// GetAccountSummaryReader is a Reader for the GetAccountSummary structure.
type GetAccountSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAccountSummaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAccountSummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetAccountSummaryMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountSummaryOK creates a GetAccountSummaryOK with default headers values
func NewGetAccountSummaryOK() *GetAccountSummaryOK {
	return &GetAccountSummaryOK{}
}

/*GetAccountSummaryOK handles this case with default header values.

The Account summary  are provided
*/
type GetAccountSummaryOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetAccountSummaryOKBody
}

func (o *GetAccountSummaryOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/summary][%d] getAccountSummaryOK  %+v", 200, o.Payload)
}

func (o *GetAccountSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetAccountSummaryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountSummaryBadRequest creates a GetAccountSummaryBadRequest with default headers values
func NewGetAccountSummaryBadRequest() *GetAccountSummaryBadRequest {
	return &GetAccountSummaryBadRequest{}
}

/*GetAccountSummaryBadRequest handles this case with default header values.

Bad Request. The client has provided invalid data to be processed by the server.
*/
type GetAccountSummaryBadRequest struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetAccountSummaryBadRequestBody
}

func (o *GetAccountSummaryBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/summary][%d] getAccountSummaryBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountSummaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetAccountSummaryBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountSummaryUnauthorized creates a GetAccountSummaryUnauthorized with default headers values
func NewGetAccountSummaryUnauthorized() *GetAccountSummaryUnauthorized {
	return &GetAccountSummaryUnauthorized{}
}

/*GetAccountSummaryUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type GetAccountSummaryUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetAccountSummaryUnauthorizedBody
}

func (o *GetAccountSummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/summary][%d] getAccountSummaryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccountSummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetAccountSummaryUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountSummaryMethodNotAllowed creates a GetAccountSummaryMethodNotAllowed with default headers values
func NewGetAccountSummaryMethodNotAllowed() *GetAccountSummaryMethodNotAllowed {
	return &GetAccountSummaryMethodNotAllowed{}
}

/*GetAccountSummaryMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type GetAccountSummaryMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetAccountSummaryMethodNotAllowedBody
}

func (o *GetAccountSummaryMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/summary][%d] getAccountSummaryMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetAccountSummaryMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetAccountSummaryMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAccountSummaryBadRequestBody get account summary bad request body
swagger:model GetAccountSummaryBadRequestBody
*/
type GetAccountSummaryBadRequestBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get account summary bad request body
func (o *GetAccountSummaryBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAccountSummaryBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAccountSummaryBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetAccountSummaryBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAccountSummaryMethodNotAllowedBody get account summary method not allowed body
swagger:model GetAccountSummaryMethodNotAllowedBody
*/
type GetAccountSummaryMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get account summary method not allowed body
func (o *GetAccountSummaryMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAccountSummaryMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAccountSummaryMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res GetAccountSummaryMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAccountSummaryOKBody get account summary o k body
swagger:model GetAccountSummaryOKBody
*/
type GetAccountSummaryOKBody struct {

	// account
	Account *models.AccountSummary `json:"account,omitempty"`

	// The ID of the most recent Transaction created for the Account.
	LastTransactionID string `json:"lastTransactionID,omitempty"`
}

// Validate validates this get account summary o k body
func (o *GetAccountSummaryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAccountSummaryOKBody) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(o.Account) { // not required
		return nil
	}

	if o.Account != nil {
		if err := o.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAccountSummaryOK" + "." + "account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAccountSummaryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAccountSummaryOKBody) UnmarshalBinary(b []byte) error {
	var res GetAccountSummaryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAccountSummaryUnauthorizedBody get account summary unauthorized body
swagger:model GetAccountSummaryUnauthorizedBody
*/
type GetAccountSummaryUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get account summary unauthorized body
func (o *GetAccountSummaryUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAccountSummaryUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAccountSummaryUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetAccountSummaryUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
