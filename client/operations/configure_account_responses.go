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

// ConfigureAccountReader is a Reader for the ConfigureAccount structure.
type ConfigureAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigureAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConfigureAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewConfigureAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewConfigureAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewConfigureAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewConfigureAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewConfigureAccountMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConfigureAccountOK creates a ConfigureAccountOK with default headers values
func NewConfigureAccountOK() *ConfigureAccountOK {
	return &ConfigureAccountOK{}
}

/*ConfigureAccountOK handles this case with default header values.

The Account was configured successfully.
*/
type ConfigureAccountOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ConfigureAccountOKBody
}

func (o *ConfigureAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{accountID}/configuration][%d] configureAccountOK  %+v", 200, o.Payload)
}

func (o *ConfigureAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ConfigureAccountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureAccountBadRequest creates a ConfigureAccountBadRequest with default headers values
func NewConfigureAccountBadRequest() *ConfigureAccountBadRequest {
	return &ConfigureAccountBadRequest{}
}

/*ConfigureAccountBadRequest handles this case with default header values.

The configuration specification was invalid.
*/
type ConfigureAccountBadRequest struct {
	Payload *ConfigureAccountBadRequestBody
}

func (o *ConfigureAccountBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{accountID}/configuration][%d] configureAccountBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigureAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ConfigureAccountBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureAccountUnauthorized creates a ConfigureAccountUnauthorized with default headers values
func NewConfigureAccountUnauthorized() *ConfigureAccountUnauthorized {
	return &ConfigureAccountUnauthorized{}
}

/*ConfigureAccountUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type ConfigureAccountUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ConfigureAccountUnauthorizedBody
}

func (o *ConfigureAccountUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{accountID}/configuration][%d] configureAccountUnauthorized  %+v", 401, o.Payload)
}

func (o *ConfigureAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ConfigureAccountUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureAccountForbidden creates a ConfigureAccountForbidden with default headers values
func NewConfigureAccountForbidden() *ConfigureAccountForbidden {
	return &ConfigureAccountForbidden{}
}

/*ConfigureAccountForbidden handles this case with default header values.

The configuration operation was forbidden on the Account.
*/
type ConfigureAccountForbidden struct {
	Payload *ConfigureAccountForbiddenBody
}

func (o *ConfigureAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{accountID}/configuration][%d] configureAccountForbidden  %+v", 403, o.Payload)
}

func (o *ConfigureAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ConfigureAccountForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureAccountNotFound creates a ConfigureAccountNotFound with default headers values
func NewConfigureAccountNotFound() *ConfigureAccountNotFound {
	return &ConfigureAccountNotFound{}
}

/*ConfigureAccountNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type ConfigureAccountNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ConfigureAccountNotFoundBody
}

func (o *ConfigureAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{accountID}/configuration][%d] configureAccountNotFound  %+v", 404, o.Payload)
}

func (o *ConfigureAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ConfigureAccountNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureAccountMethodNotAllowed creates a ConfigureAccountMethodNotAllowed with default headers values
func NewConfigureAccountMethodNotAllowed() *ConfigureAccountMethodNotAllowed {
	return &ConfigureAccountMethodNotAllowed{}
}

/*ConfigureAccountMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type ConfigureAccountMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ConfigureAccountMethodNotAllowedBody
}

func (o *ConfigureAccountMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{accountID}/configuration][%d] configureAccountMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ConfigureAccountMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ConfigureAccountMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ConfigureAccountBadRequestBody configure account bad request body
swagger:model ConfigureAccountBadRequestBody
*/
type ConfigureAccountBadRequestBody struct {

	// client configure reject transaction
	ClientConfigureRejectTransaction *models.ClientConfigureRejectTransaction `json:"clientConfigureRejectTransaction,omitempty"`

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The ID of the last Transaction created for the Account.
	LastTransactionID string `json:"lastTransactionID,omitempty"`
}

// Validate validates this configure account bad request body
func (o *ConfigureAccountBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientConfigureRejectTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConfigureAccountBadRequestBody) validateClientConfigureRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientConfigureRejectTransaction) { // not required
		return nil
	}

	if o.ClientConfigureRejectTransaction != nil {
		if err := o.ClientConfigureRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configureAccountBadRequest" + "." + "clientConfigureRejectTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConfigureAccountBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfigureAccountBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ConfigureAccountBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConfigureAccountBody configure account body
swagger:model ConfigureAccountBody
*/
type ConfigureAccountBody struct {

	// Client-defined alias (name) for the Account
	Alias string `json:"alias,omitempty"`

	// The string representation of a decimal number.
	MarginRate string `json:"marginRate,omitempty"`
}

// Validate validates this configure account body
func (o *ConfigureAccountBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConfigureAccountBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfigureAccountBody) UnmarshalBinary(b []byte) error {
	var res ConfigureAccountBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConfigureAccountForbiddenBody configure account forbidden body
swagger:model ConfigureAccountForbiddenBody
*/
type ConfigureAccountForbiddenBody struct {

	// client configure reject transaction
	ClientConfigureRejectTransaction *models.ClientConfigureRejectTransaction `json:"clientConfigureRejectTransaction,omitempty"`

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The ID of the last Transaction created for the Account.
	LastTransactionID string `json:"lastTransactionID,omitempty"`
}

// Validate validates this configure account forbidden body
func (o *ConfigureAccountForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientConfigureRejectTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConfigureAccountForbiddenBody) validateClientConfigureRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientConfigureRejectTransaction) { // not required
		return nil
	}

	if o.ClientConfigureRejectTransaction != nil {
		if err := o.ClientConfigureRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configureAccountForbidden" + "." + "clientConfigureRejectTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConfigureAccountForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfigureAccountForbiddenBody) UnmarshalBinary(b []byte) error {
	var res ConfigureAccountForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConfigureAccountMethodNotAllowedBody configure account method not allowed body
swagger:model ConfigureAccountMethodNotAllowedBody
*/
type ConfigureAccountMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this configure account method not allowed body
func (o *ConfigureAccountMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConfigureAccountMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfigureAccountMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ConfigureAccountMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConfigureAccountNotFoundBody configure account not found body
swagger:model ConfigureAccountNotFoundBody
*/
type ConfigureAccountNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this configure account not found body
func (o *ConfigureAccountNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConfigureAccountNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfigureAccountNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ConfigureAccountNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConfigureAccountOKBody configure account o k body
swagger:model ConfigureAccountOKBody
*/
type ConfigureAccountOKBody struct {

	// client configure transaction
	ClientConfigureTransaction *models.ClientConfigureTransaction `json:"clientConfigureTransaction,omitempty"`

	// The ID of the last Transaction created for the Account.
	LastTransactionID string `json:"lastTransactionID,omitempty"`
}

// Validate validates this configure account o k body
func (o *ConfigureAccountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientConfigureTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConfigureAccountOKBody) validateClientConfigureTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientConfigureTransaction) { // not required
		return nil
	}

	if o.ClientConfigureTransaction != nil {
		if err := o.ClientConfigureTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configureAccountOK" + "." + "clientConfigureTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConfigureAccountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfigureAccountOKBody) UnmarshalBinary(b []byte) error {
	var res ConfigureAccountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConfigureAccountUnauthorizedBody configure account unauthorized body
swagger:model ConfigureAccountUnauthorizedBody
*/
type ConfigureAccountUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this configure account unauthorized body
func (o *ConfigureAccountUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConfigureAccountUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConfigureAccountUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ConfigureAccountUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
