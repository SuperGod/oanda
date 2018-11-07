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

// ListOrdersReader is a Reader for the ListOrders structure.
type ListOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListOrdersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListOrdersOK creates a ListOrdersOK with default headers values
func NewListOrdersOK() *ListOrdersOK {
	return &ListOrdersOK{}
}

/*ListOrdersOK handles this case with default header values.

The list of Orders requested
*/
type ListOrdersOK struct {
	/*A link to the next page of results if the results were paginated
	 */
	Link string
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListOrdersOKBody
}

func (o *ListOrdersOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/orders][%d] listOrdersOK  %+v", 200, o.Payload)
}

func (o *ListOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListOrdersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrdersBadRequest creates a ListOrdersBadRequest with default headers values
func NewListOrdersBadRequest() *ListOrdersBadRequest {
	return &ListOrdersBadRequest{}
}

/*ListOrdersBadRequest handles this case with default header values.

Bad Request. The client has provided invalid data to be processed by the server.
*/
type ListOrdersBadRequest struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListOrdersBadRequestBody
}

func (o *ListOrdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/orders][%d] listOrdersBadRequest  %+v", 400, o.Payload)
}

func (o *ListOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListOrdersBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrdersNotFound creates a ListOrdersNotFound with default headers values
func NewListOrdersNotFound() *ListOrdersNotFound {
	return &ListOrdersNotFound{}
}

/*ListOrdersNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type ListOrdersNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListOrdersNotFoundBody
}

func (o *ListOrdersNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/orders][%d] listOrdersNotFound  %+v", 404, o.Payload)
}

func (o *ListOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListOrdersNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrdersMethodNotAllowed creates a ListOrdersMethodNotAllowed with default headers values
func NewListOrdersMethodNotAllowed() *ListOrdersMethodNotAllowed {
	return &ListOrdersMethodNotAllowed{}
}

/*ListOrdersMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type ListOrdersMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *ListOrdersMethodNotAllowedBody
}

func (o *ListOrdersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/orders][%d] listOrdersMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListOrdersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(ListOrdersMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListOrdersBadRequestBody list orders bad request body
swagger:model ListOrdersBadRequestBody
*/
type ListOrdersBadRequestBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list orders bad request body
func (o *ListOrdersBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListOrdersBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOrdersBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ListOrdersBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListOrdersMethodNotAllowedBody list orders method not allowed body
swagger:model ListOrdersMethodNotAllowedBody
*/
type ListOrdersMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list orders method not allowed body
func (o *ListOrdersMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListOrdersMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOrdersMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ListOrdersMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListOrdersNotFoundBody list orders not found body
swagger:model ListOrdersNotFoundBody
*/
type ListOrdersNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this list orders not found body
func (o *ListOrdersNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListOrdersNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOrdersNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ListOrdersNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListOrdersOKBody list orders o k body
swagger:model ListOrdersOKBody
*/
type ListOrdersOKBody struct {

	// The ID of the most recent Transaction created for the Account
	LastTransactionID string `json:"lastTransactionID,omitempty"`

	// The list of Order detail objects
	Orders []*models.Order `json:"orders"`
}

// Validate validates this list orders o k body
func (o *ListOrdersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListOrdersOKBody) validateOrders(formats strfmt.Registry) error {

	if swag.IsZero(o.Orders) { // not required
		return nil
	}

	for i := 0; i < len(o.Orders); i++ {
		if swag.IsZero(o.Orders[i]) { // not required
			continue
		}

		if o.Orders[i] != nil {
			if err := o.Orders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listOrdersOK" + "." + "orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListOrdersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListOrdersOKBody) UnmarshalBinary(b []byte) error {
	var res ListOrdersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
