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

// ClosePositionReader is a Reader for the ClosePosition structure.
type ClosePositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClosePositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewClosePositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewClosePositionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewClosePositionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewClosePositionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewClosePositionMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClosePositionOK creates a ClosePositionOK with default headers values
func NewClosePositionOK() *ClosePositionOK {
	return &ClosePositionOK{}
}

/*ClosePositionOK handles this case with default header values.

The Position closeout request has been successfully processed.
*/
type ClosePositionOK struct {
	/*A link to the Position that was just closed out
	 */
	Location string
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ClosePositionOKBody
}

func (o *ClosePositionOK) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionOK  %+v", 200, o.Payload)
}

func (o *ClosePositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ClosePositionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClosePositionBadRequest creates a ClosePositionBadRequest with default headers values
func NewClosePositionBadRequest() *ClosePositionBadRequest {
	return &ClosePositionBadRequest{}
}

/*ClosePositionBadRequest handles this case with default header values.

The Parameters provided that describe the Position closeout are invalid.
*/
type ClosePositionBadRequest struct {
	Payload *ClosePositionBadRequestBody
}

func (o *ClosePositionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionBadRequest  %+v", 400, o.Payload)
}

func (o *ClosePositionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ClosePositionBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClosePositionUnauthorized creates a ClosePositionUnauthorized with default headers values
func NewClosePositionUnauthorized() *ClosePositionUnauthorized {
	return &ClosePositionUnauthorized{}
}

/*ClosePositionUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type ClosePositionUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ClosePositionUnauthorizedBody
}

func (o *ClosePositionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionUnauthorized  %+v", 401, o.Payload)
}

func (o *ClosePositionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ClosePositionUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClosePositionNotFound creates a ClosePositionNotFound with default headers values
func NewClosePositionNotFound() *ClosePositionNotFound {
	return &ClosePositionNotFound{}
}

/*ClosePositionNotFound handles this case with default header values.

The Account or one or more of the Positions specified does not exist.
*/
type ClosePositionNotFound struct {
	Payload *ClosePositionNotFoundBody
}

func (o *ClosePositionNotFound) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionNotFound  %+v", 404, o.Payload)
}

func (o *ClosePositionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ClosePositionNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClosePositionMethodNotAllowed creates a ClosePositionMethodNotAllowed with default headers values
func NewClosePositionMethodNotAllowed() *ClosePositionMethodNotAllowed {
	return &ClosePositionMethodNotAllowed{}
}

/*ClosePositionMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type ClosePositionMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ClosePositionMethodNotAllowedBody
}

func (o *ClosePositionMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ClosePositionMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ClosePositionMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ClosePositionBadRequestBody close position bad request body
swagger:model ClosePositionBadRequestBody
*/
type ClosePositionBadRequestBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// long order reject transaction
	LongOrderRejectTransaction *models.MarketOrderRejectTransaction `json:"longOrderRejectTransaction,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`

	// short order reject transaction
	ShortOrderRejectTransaction *models.MarketOrderRejectTransaction `json:"shortOrderRejectTransaction,omitempty"`
}

// Validate validates this close position bad request body
func (o *ClosePositionBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLongOrderRejectTransaction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShortOrderRejectTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClosePositionBadRequestBody) validateLongOrderRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.LongOrderRejectTransaction) { // not required
		return nil
	}

	if o.LongOrderRejectTransaction != nil {
		if err := o.LongOrderRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionBadRequest" + "." + "longOrderRejectTransaction")
			}
			return err
		}
	}

	return nil
}

func (o *ClosePositionBadRequestBody) validateShortOrderRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.ShortOrderRejectTransaction) { // not required
		return nil
	}

	if o.ShortOrderRejectTransaction != nil {
		if err := o.ShortOrderRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionBadRequest" + "." + "shortOrderRejectTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClosePositionBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClosePositionBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ClosePositionBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClosePositionBody close position body
swagger:model ClosePositionBody
*/
type ClosePositionBody struct {

	// long client extensions
	LongClientExtensions *models.ClientExtensions `json:"longClientExtensions,omitempty"`

	// Indication of how much of the long Position to closeout. Either the string "ALL", the string "NONE", or a DecimalNumber representing how many units of the long position to close using a PositionCloseout MarketOrder. The units specified must always be positive.
	LongUnits string `json:"longUnits,omitempty"`

	// short client extensions
	ShortClientExtensions *models.ClientExtensions `json:"shortClientExtensions,omitempty"`

	// Indication of how much of the short Position to closeout. Either the string "ALL", the string "NONE", or a DecimalNumber representing how many units of the short position to close using a PositionCloseout MarketOrder. The units specified must always be positive.
	ShortUnits string `json:"shortUnits,omitempty"`
}

// Validate validates this close position body
func (o *ClosePositionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLongClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShortClientExtensions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClosePositionBody) validateLongClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(o.LongClientExtensions) { // not required
		return nil
	}

	if o.LongClientExtensions != nil {
		if err := o.LongClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionBody" + "." + "longClientExtensions")
			}
			return err
		}
	}

	return nil
}

func (o *ClosePositionBody) validateShortClientExtensions(formats strfmt.Registry) error {

	if swag.IsZero(o.ShortClientExtensions) { // not required
		return nil
	}

	if o.ShortClientExtensions != nil {
		if err := o.ShortClientExtensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionBody" + "." + "shortClientExtensions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClosePositionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClosePositionBody) UnmarshalBinary(b []byte) error {
	var res ClosePositionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClosePositionMethodNotAllowedBody close position method not allowed body
swagger:model ClosePositionMethodNotAllowedBody
*/
type ClosePositionMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this close position method not allowed body
func (o *ClosePositionMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClosePositionMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClosePositionMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ClosePositionMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClosePositionNotFoundBody close position not found body
swagger:model ClosePositionNotFoundBody
*/
type ClosePositionNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The ID of the most recent Transaction created for the Account. Only present if the Account exists.
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// long order reject transaction
	LongOrderRejectTransaction *models.MarketOrderRejectTransaction `json:"longOrderRejectTransaction,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request. Only present if the Account exists.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`

	// short order reject transaction
	ShortOrderRejectTransaction *models.MarketOrderRejectTransaction `json:"shortOrderRejectTransaction,omitempty"`
}

// Validate validates this close position not found body
func (o *ClosePositionNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLongOrderRejectTransaction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShortOrderRejectTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClosePositionNotFoundBody) validateLongOrderRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.LongOrderRejectTransaction) { // not required
		return nil
	}

	if o.LongOrderRejectTransaction != nil {
		if err := o.LongOrderRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionNotFound" + "." + "longOrderRejectTransaction")
			}
			return err
		}
	}

	return nil
}

func (o *ClosePositionNotFoundBody) validateShortOrderRejectTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.ShortOrderRejectTransaction) { // not required
		return nil
	}

	if o.ShortOrderRejectTransaction != nil {
		if err := o.ShortOrderRejectTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionNotFound" + "." + "shortOrderRejectTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClosePositionNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClosePositionNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ClosePositionNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClosePositionOKBody close position o k body
swagger:model ClosePositionOKBody
*/
type ClosePositionOKBody struct {

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// long order cancel transaction
	LongOrderCancelTransaction *models.OrderCancelTransaction `json:"longOrderCancelTransaction,omitempty"`

	// long order create transaction
	LongOrderCreateTransaction *models.MarketOrderTransaction `json:"longOrderCreateTransaction,omitempty"`

	// long order fill transaction
	LongOrderFillTransaction *models.OrderFillTransaction `json:"longOrderFillTransaction,omitempty"`

	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIds []string `json:"relatedTransactionIDs"`

	// short order cancel transaction
	ShortOrderCancelTransaction *models.OrderCancelTransaction `json:"shortOrderCancelTransaction,omitempty"`

	// short order create transaction
	ShortOrderCreateTransaction *models.MarketOrderTransaction `json:"shortOrderCreateTransaction,omitempty"`

	// short order fill transaction
	ShortOrderFillTransaction *models.OrderFillTransaction `json:"shortOrderFillTransaction,omitempty"`
}

// Validate validates this close position o k body
func (o *ClosePositionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLongOrderCancelTransaction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLongOrderCreateTransaction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLongOrderFillTransaction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShortOrderCancelTransaction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShortOrderCreateTransaction(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShortOrderFillTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClosePositionOKBody) validateLongOrderCancelTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.LongOrderCancelTransaction) { // not required
		return nil
	}

	if o.LongOrderCancelTransaction != nil {
		if err := o.LongOrderCancelTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionOK" + "." + "longOrderCancelTransaction")
			}
			return err
		}
	}

	return nil
}

func (o *ClosePositionOKBody) validateLongOrderCreateTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.LongOrderCreateTransaction) { // not required
		return nil
	}

	if o.LongOrderCreateTransaction != nil {
		if err := o.LongOrderCreateTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionOK" + "." + "longOrderCreateTransaction")
			}
			return err
		}
	}

	return nil
}

func (o *ClosePositionOKBody) validateLongOrderFillTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.LongOrderFillTransaction) { // not required
		return nil
	}

	if o.LongOrderFillTransaction != nil {
		if err := o.LongOrderFillTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionOK" + "." + "longOrderFillTransaction")
			}
			return err
		}
	}

	return nil
}

func (o *ClosePositionOKBody) validateShortOrderCancelTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.ShortOrderCancelTransaction) { // not required
		return nil
	}

	if o.ShortOrderCancelTransaction != nil {
		if err := o.ShortOrderCancelTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionOK" + "." + "shortOrderCancelTransaction")
			}
			return err
		}
	}

	return nil
}

func (o *ClosePositionOKBody) validateShortOrderCreateTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.ShortOrderCreateTransaction) { // not required
		return nil
	}

	if o.ShortOrderCreateTransaction != nil {
		if err := o.ShortOrderCreateTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionOK" + "." + "shortOrderCreateTransaction")
			}
			return err
		}
	}

	return nil
}

func (o *ClosePositionOKBody) validateShortOrderFillTransaction(formats strfmt.Registry) error {

	if swag.IsZero(o.ShortOrderFillTransaction) { // not required
		return nil
	}

	if o.ShortOrderFillTransaction != nil {
		if err := o.ShortOrderFillTransaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closePositionOK" + "." + "shortOrderFillTransaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClosePositionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClosePositionOKBody) UnmarshalBinary(b []byte) error {
	var res ClosePositionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClosePositionUnauthorizedBody close position unauthorized body
swagger:model ClosePositionUnauthorizedBody
*/
type ClosePositionUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this close position unauthorized body
func (o *ClosePositionUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClosePositionUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClosePositionUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res ClosePositionUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
