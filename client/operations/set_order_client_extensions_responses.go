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

// SetOrderClientExtensionsReader is a Reader for the SetOrderClientExtensions structure.
type SetOrderClientExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetOrderClientExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetOrderClientExtensionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSetOrderClientExtensionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSetOrderClientExtensionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSetOrderClientExtensionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSetOrderClientExtensionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetOrderClientExtensionsOK creates a SetOrderClientExtensionsOK with default headers values
func NewSetOrderClientExtensionsOK() *SetOrderClientExtensionsOK {
	return &SetOrderClientExtensionsOK{}
}

/*SetOrderClientExtensionsOK handles this case with default header values.

The Order's Client Extensions were successfully modified
*/
type SetOrderClientExtensionsOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *SetOrderClientExtensionsOKBody
}

func (o *SetOrderClientExtensionsOK) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}/clientExtensions][%d] setOrderClientExtensionsOK  %+v", 200, o.Payload)
}

func (o *SetOrderClientExtensionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(SetOrderClientExtensionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrderClientExtensionsBadRequest creates a SetOrderClientExtensionsBadRequest with default headers values
func NewSetOrderClientExtensionsBadRequest() *SetOrderClientExtensionsBadRequest {
	return &SetOrderClientExtensionsBadRequest{}
}

/*SetOrderClientExtensionsBadRequest handles this case with default header values.

The Order Client Extensions specification was invalid
*/
type SetOrderClientExtensionsBadRequest struct {
	Payload *SetOrderClientExtensionsBadRequestBody
}

func (o *SetOrderClientExtensionsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}/clientExtensions][%d] setOrderClientExtensionsBadRequest  %+v", 400, o.Payload)
}

func (o *SetOrderClientExtensionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetOrderClientExtensionsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrderClientExtensionsUnauthorized creates a SetOrderClientExtensionsUnauthorized with default headers values
func NewSetOrderClientExtensionsUnauthorized() *SetOrderClientExtensionsUnauthorized {
	return &SetOrderClientExtensionsUnauthorized{}
}

/*SetOrderClientExtensionsUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type SetOrderClientExtensionsUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *SetOrderClientExtensionsUnauthorizedBody
}

func (o *SetOrderClientExtensionsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}/clientExtensions][%d] setOrderClientExtensionsUnauthorized  %+v", 401, o.Payload)
}

func (o *SetOrderClientExtensionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(SetOrderClientExtensionsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrderClientExtensionsNotFound creates a SetOrderClientExtensionsNotFound with default headers values
func NewSetOrderClientExtensionsNotFound() *SetOrderClientExtensionsNotFound {
	return &SetOrderClientExtensionsNotFound{}
}

/*SetOrderClientExtensionsNotFound handles this case with default header values.

The Account or Order specified does not exist.
*/
type SetOrderClientExtensionsNotFound struct {
	Payload *SetOrderClientExtensionsNotFoundBody
}

func (o *SetOrderClientExtensionsNotFound) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}/clientExtensions][%d] setOrderClientExtensionsNotFound  %+v", 404, o.Payload)
}

func (o *SetOrderClientExtensionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetOrderClientExtensionsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetOrderClientExtensionsMethodNotAllowed creates a SetOrderClientExtensionsMethodNotAllowed with default headers values
func NewSetOrderClientExtensionsMethodNotAllowed() *SetOrderClientExtensionsMethodNotAllowed {
	return &SetOrderClientExtensionsMethodNotAllowed{}
}

/*SetOrderClientExtensionsMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type SetOrderClientExtensionsMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *SetOrderClientExtensionsMethodNotAllowedBody
}

func (o *SetOrderClientExtensionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}/clientExtensions][%d] setOrderClientExtensionsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SetOrderClientExtensionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(SetOrderClientExtensionsMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SetOrderClientExtensionsBadRequestBody set order client extensions bad request body
swagger:model SetOrderClientExtensionsBadRequestBody
*/
type SetOrderClientExtensionsBadRequestBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// order client extensions modify reject transaction
	OrderClientExtensionsModifyRejectTransaction *models.OrderClientExtensionsModifyRejectTransaction `json:"orderClientExtensionsModifyRejectTransaction,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`
}

// Validate validates this set order client extensions bad request body
func (o *SetOrderClientExtensionsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrderClientExtensionsModifyRejectTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetOrderClientExtensionsBadRequestBody) validateOrderClientExtensionsModifyRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.OrderClientExtensionsModifyRejectTransaction) { // not required
		return nil
	}

	if o.OrderClientExtensionsModifyRejectTransaction != nil {
		if err := o.OrderClientExtensionsModifyRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("setOrderClientExtensionsBadRequest" + "." + "orderClientExtensionsModifyRejectTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetOrderClientExtensionsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetOrderClientExtensionsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res SetOrderClientExtensionsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SetOrderClientExtensionsBody set order client extensions body
swagger:model SetOrderClientExtensionsBody
*/
type SetOrderClientExtensionsBody struct {

	// client extensions
	ClientExtensions *models.ClientExtensions `json:"clientExtensions,omitempty"`

	// trade client extensions
	TradeClientExtensions *models.ClientExtensions `json:"tradeClientExtensions,omitempty"`
}

// Validate validates this set order client extensions body
func (o *SetOrderClientExtensionsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTradeClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetOrderClientExtensionsBody) validateClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientExtensions) { // not required
		return nil
	}

	if o.ClientExtensions != nil {
		if err := o.ClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("setOrderClientExtensionsBody" + "." + "clientExtensions")
			}
			return err
		}
	}

	return nil
}

func (o *SetOrderClientExtensionsBody) validateTradeClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(o.TradeClientExtensions) { // not required
		return nil
	}

	if o.TradeClientExtensions != nil {
		if err := o.TradeClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("setOrderClientExtensionsBody" + "." + "tradeClientExtensions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetOrderClientExtensionsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetOrderClientExtensionsBody) UnmarshalBinary(b []byte) error {
	var res SetOrderClientExtensionsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SetOrderClientExtensionsMethodNotAllowedBody set order client extensions method not allowed body
swagger:model SetOrderClientExtensionsMethodNotAllowedBody
*/
type SetOrderClientExtensionsMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this set order client extensions method not allowed body
func (o *SetOrderClientExtensionsMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetOrderClientExtensionsMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetOrderClientExtensionsMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res SetOrderClientExtensionsMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SetOrderClientExtensionsNotFoundBody set order client extensions not found body
swagger:model SetOrderClientExtensionsNotFoundBody
*/
type SetOrderClientExtensionsNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The ID of the most recent Transaction created for the Account. Only present if the Account exists.
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// order client extensions modify reject transaction
	OrderClientExtensionsModifyRejectTransaction *models.OrderClientExtensionsModifyRejectTransaction `json:"orderClientExtensionsModifyRejectTransaction,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request. Only present if the Account exists.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`
}

// Validate validates this set order client extensions not found body
func (o *SetOrderClientExtensionsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrderClientExtensionsModifyRejectTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetOrderClientExtensionsNotFoundBody) validateOrderClientExtensionsModifyRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.OrderClientExtensionsModifyRejectTransaction) { // not required
		return nil
	}

	if o.OrderClientExtensionsModifyRejectTransaction != nil {
		if err := o.OrderClientExtensionsModifyRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("setOrderClientExtensionsNotFound" + "." + "orderClientExtensionsModifyRejectTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetOrderClientExtensionsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetOrderClientExtensionsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SetOrderClientExtensionsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SetOrderClientExtensionsOKBody set order client extensions o k body
swagger:model SetOrderClientExtensionsOKBody
*/
type SetOrderClientExtensionsOKBody struct {

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// order client extensions modify transaction
	OrderClientExtensionsModifyTransaction *models.OrderClientExtensionsModifyTransaction `json:"orderClientExtensionsModifyTransaction,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`
}

// Validate validates this set order client extensions o k body
func (o *SetOrderClientExtensionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrderClientExtensionsModifyTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetOrderClientExtensionsOKBody) validateOrderClientExtensionsModifyTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.OrderClientExtensionsModifyTransaction) { // not required
		return nil
	}

	if o.OrderClientExtensionsModifyTransaction != nil {
		if err := o.OrderClientExtensionsModifyTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("setOrderClientExtensionsOK" + "." + "orderClientExtensionsModifyTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetOrderClientExtensionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetOrderClientExtensionsOKBody) UnmarshalBinary(b []byte) error {
	var res SetOrderClientExtensionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SetOrderClientExtensionsUnauthorizedBody set order client extensions unauthorized body
swagger:model SetOrderClientExtensionsUnauthorizedBody
*/
type SetOrderClientExtensionsUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this set order client extensions unauthorized body
func (o *SetOrderClientExtensionsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetOrderClientExtensionsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetOrderClientExtensionsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res SetOrderClientExtensionsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
