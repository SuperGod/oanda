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

// StreamPricingReader is a Reader for the StreamPricing structure.
type StreamPricingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamPricingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStreamPricingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewStreamPricingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewStreamPricingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStreamPricingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewStreamPricingMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStreamPricingOK creates a StreamPricingOK with default headers values
func NewStreamPricingOK() *StreamPricingOK {
	return &StreamPricingOK{}
}

/*StreamPricingOK handles this case with default header values.

Connecting to the Price Stream was successful.
*/
type StreamPricingOK struct {
	/*A link to the next/previous order book snapshot.
	 */
	Link string
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *StreamPricingOKBody
}

func (o *StreamPricingOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing/stream][%d] streamPricingOK  %+v", 200, o.Payload)
}

func (o *StreamPricingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(StreamPricingOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamPricingBadRequest creates a StreamPricingBadRequest with default headers values
func NewStreamPricingBadRequest() *StreamPricingBadRequest {
	return &StreamPricingBadRequest{}
}

/*StreamPricingBadRequest handles this case with default header values.

Bad Request. The client has provided invalid data to be processed by the server.
*/
type StreamPricingBadRequest struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *StreamPricingBadRequestBody
}

func (o *StreamPricingBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing/stream][%d] streamPricingBadRequest  %+v", 400, o.Payload)
}

func (o *StreamPricingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(StreamPricingBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamPricingUnauthorized creates a StreamPricingUnauthorized with default headers values
func NewStreamPricingUnauthorized() *StreamPricingUnauthorized {
	return &StreamPricingUnauthorized{}
}

/*StreamPricingUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type StreamPricingUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *StreamPricingUnauthorizedBody
}

func (o *StreamPricingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing/stream][%d] streamPricingUnauthorized  %+v", 401, o.Payload)
}

func (o *StreamPricingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(StreamPricingUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamPricingNotFound creates a StreamPricingNotFound with default headers values
func NewStreamPricingNotFound() *StreamPricingNotFound {
	return &StreamPricingNotFound{}
}

/*StreamPricingNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type StreamPricingNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *StreamPricingNotFoundBody
}

func (o *StreamPricingNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing/stream][%d] streamPricingNotFound  %+v", 404, o.Payload)
}

func (o *StreamPricingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(StreamPricingNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamPricingMethodNotAllowed creates a StreamPricingMethodNotAllowed with default headers values
func NewStreamPricingMethodNotAllowed() *StreamPricingMethodNotAllowed {
	return &StreamPricingMethodNotAllowed{}
}

/*StreamPricingMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type StreamPricingMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *StreamPricingMethodNotAllowedBody
}

func (o *StreamPricingMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing/stream][%d] streamPricingMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *StreamPricingMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(StreamPricingMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StreamPricingBadRequestBody stream pricing bad request body
swagger:model StreamPricingBadRequestBody
*/
type StreamPricingBadRequestBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this stream pricing bad request body
func (o *StreamPricingBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StreamPricingBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StreamPricingBadRequestBody) UnmarshalBinary(b []byte) error {
	var res StreamPricingBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StreamPricingMethodNotAllowedBody stream pricing method not allowed body
swagger:model StreamPricingMethodNotAllowedBody
*/
type StreamPricingMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this stream pricing method not allowed body
func (o *StreamPricingMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StreamPricingMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StreamPricingMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res StreamPricingMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StreamPricingNotFoundBody stream pricing not found body
swagger:model StreamPricingNotFoundBody
*/
type StreamPricingNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this stream pricing not found body
func (o *StreamPricingNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StreamPricingNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StreamPricingNotFoundBody) UnmarshalBinary(b []byte) error {
	var res StreamPricingNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StreamPricingOKBody The response body for the Pricing Stream uses chunked transfer encoding.  Each chunk contains Price and/or PricingHeartbeat objects encoded as JSON.  Each JSON object is serialized into a single line of text, and multiple objects found in the same chunk are separated by newlines.
// Heartbeats are sent every 5 seconds.
swagger:model StreamPricingOKBody
*/
type StreamPricingOKBody struct {

	// heartbeat
	Heartbeat *models.PricingHeartbeat `json:"heartbeat,omitempty"`

	// price
	Price *models.ClientPrice `json:"price,omitempty"`
}

// Validate validates this stream pricing o k body
func (o *StreamPricingOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHeartbeat(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StreamPricingOKBody) validateHeartbeat(formats strfmt.Registry) error {

	if swag.IsZero(o.Heartbeat) { // not required
		return nil
	}

	if o.Heartbeat != nil {
		if err := o.Heartbeat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streamPricingOK" + "." + "heartbeat")
			}
			return err
		}
	}

	return nil
}

func (o *StreamPricingOKBody) validatePrice(formats strfmt.Registry) error {

	if swag.IsZero(o.Price) { // not required
		return nil
	}

	if o.Price != nil {
		if err := o.Price.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streamPricingOK" + "." + "price")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StreamPricingOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StreamPricingOKBody) UnmarshalBinary(b []byte) error {
	var res StreamPricingOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StreamPricingUnauthorizedBody stream pricing unauthorized body
swagger:model StreamPricingUnauthorizedBody
*/
type StreamPricingUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this stream pricing unauthorized body
func (o *StreamPricingUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StreamPricingUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StreamPricingUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res StreamPricingUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
