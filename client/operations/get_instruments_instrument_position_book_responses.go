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

// GetInstrumentsInstrumentPositionBookReader is a Reader for the GetInstrumentsInstrumentPositionBook structure.
type GetInstrumentsInstrumentPositionBookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstrumentsInstrumentPositionBookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstrumentsInstrumentPositionBookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstrumentsInstrumentPositionBookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstrumentsInstrumentPositionBookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInstrumentsInstrumentPositionBookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetInstrumentsInstrumentPositionBookMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstrumentsInstrumentPositionBookOK creates a GetInstrumentsInstrumentPositionBookOK with default headers values
func NewGetInstrumentsInstrumentPositionBookOK() *GetInstrumentsInstrumentPositionBookOK {
	return &GetInstrumentsInstrumentPositionBookOK{}
}

/*GetInstrumentsInstrumentPositionBookOK handles this case with default header values.

The position book has been successfully provided.
*/
type GetInstrumentsInstrumentPositionBookOK struct {
	/*Value will be "gzip" regardless of provided Accept-Encoding header
	 */
	ContentEncoding string
	/*A link to the next/previous position book snapshot.
	 */
	Link string
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentsInstrumentPositionBookOKBody
}

func (o *GetInstrumentsInstrumentPositionBookOK) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/positionBook][%d] getInstrumentsInstrumentPositionBookOK  %+v", 200, o.Payload)
}

func (o *GetInstrumentsInstrumentPositionBookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Encoding
	o.ContentEncoding = response.GetHeader("Content-Encoding")

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentsInstrumentPositionBookOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentsInstrumentPositionBookBadRequest creates a GetInstrumentsInstrumentPositionBookBadRequest with default headers values
func NewGetInstrumentsInstrumentPositionBookBadRequest() *GetInstrumentsInstrumentPositionBookBadRequest {
	return &GetInstrumentsInstrumentPositionBookBadRequest{}
}

/*GetInstrumentsInstrumentPositionBookBadRequest handles this case with default header values.

Bad Request. The client has provided invalid data to be processed by the server.
*/
type GetInstrumentsInstrumentPositionBookBadRequest struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentsInstrumentPositionBookBadRequestBody
}

func (o *GetInstrumentsInstrumentPositionBookBadRequest) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/positionBook][%d] getInstrumentsInstrumentPositionBookBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstrumentsInstrumentPositionBookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentsInstrumentPositionBookBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentsInstrumentPositionBookUnauthorized creates a GetInstrumentsInstrumentPositionBookUnauthorized with default headers values
func NewGetInstrumentsInstrumentPositionBookUnauthorized() *GetInstrumentsInstrumentPositionBookUnauthorized {
	return &GetInstrumentsInstrumentPositionBookUnauthorized{}
}

/*GetInstrumentsInstrumentPositionBookUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type GetInstrumentsInstrumentPositionBookUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentsInstrumentPositionBookUnauthorizedBody
}

func (o *GetInstrumentsInstrumentPositionBookUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/positionBook][%d] getInstrumentsInstrumentPositionBookUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstrumentsInstrumentPositionBookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentsInstrumentPositionBookUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentsInstrumentPositionBookNotFound creates a GetInstrumentsInstrumentPositionBookNotFound with default headers values
func NewGetInstrumentsInstrumentPositionBookNotFound() *GetInstrumentsInstrumentPositionBookNotFound {
	return &GetInstrumentsInstrumentPositionBookNotFound{}
}

/*GetInstrumentsInstrumentPositionBookNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type GetInstrumentsInstrumentPositionBookNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentsInstrumentPositionBookNotFoundBody
}

func (o *GetInstrumentsInstrumentPositionBookNotFound) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/positionBook][%d] getInstrumentsInstrumentPositionBookNotFound  %+v", 404, o.Payload)
}

func (o *GetInstrumentsInstrumentPositionBookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentsInstrumentPositionBookNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstrumentsInstrumentPositionBookMethodNotAllowed creates a GetInstrumentsInstrumentPositionBookMethodNotAllowed with default headers values
func NewGetInstrumentsInstrumentPositionBookMethodNotAllowed() *GetInstrumentsInstrumentPositionBookMethodNotAllowed {
	return &GetInstrumentsInstrumentPositionBookMethodNotAllowed{}
}

/*GetInstrumentsInstrumentPositionBookMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type GetInstrumentsInstrumentPositionBookMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *GetInstrumentsInstrumentPositionBookMethodNotAllowedBody
}

func (o *GetInstrumentsInstrumentPositionBookMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /instruments/{instrument}/positionBook][%d] getInstrumentsInstrumentPositionBookMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetInstrumentsInstrumentPositionBookMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(GetInstrumentsInstrumentPositionBookMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetInstrumentsInstrumentPositionBookBadRequestBody get instruments instrument position book bad request body
swagger:model GetInstrumentsInstrumentPositionBookBadRequestBody
*/
type GetInstrumentsInstrumentPositionBookBadRequestBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instruments instrument position book bad request body
func (o *GetInstrumentsInstrumentPositionBookBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentsInstrumentPositionBookBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentsInstrumentPositionBookMethodNotAllowedBody get instruments instrument position book method not allowed body
swagger:model GetInstrumentsInstrumentPositionBookMethodNotAllowedBody
*/
type GetInstrumentsInstrumentPositionBookMethodNotAllowedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instruments instrument position book method not allowed body
func (o *GetInstrumentsInstrumentPositionBookMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentsInstrumentPositionBookMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentsInstrumentPositionBookNotFoundBody get instruments instrument position book not found body
swagger:model GetInstrumentsInstrumentPositionBookNotFoundBody
*/
type GetInstrumentsInstrumentPositionBookNotFoundBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instruments instrument position book not found body
func (o *GetInstrumentsInstrumentPositionBookNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentsInstrumentPositionBookNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentsInstrumentPositionBookOKBody get instruments instrument position book o k body
swagger:model GetInstrumentsInstrumentPositionBookOKBody
*/
type GetInstrumentsInstrumentPositionBookOKBody struct {

	// position book
	PositionBook *models.PositionBook `json:"positionBook,omitempty"`
}

// Validate validates this get instruments instrument position book o k body
func (o *GetInstrumentsInstrumentPositionBookOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePositionBook(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetInstrumentsInstrumentPositionBookOKBody) validatePositionBook(formats strfmt.Registry) error {

	if swag.IsZero(o.PositionBook) { // not required
		return nil
	}

	if o.PositionBook != nil {
		if err := o.PositionBook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getInstrumentsInstrumentPositionBookOK" + "." + "positionBook")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookOKBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentsInstrumentPositionBookOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetInstrumentsInstrumentPositionBookUnauthorizedBody get instruments instrument position book unauthorized body
swagger:model GetInstrumentsInstrumentPositionBookUnauthorizedBody
*/
type GetInstrumentsInstrumentPositionBookUnauthorizedBody struct {

	// The code of the error that has occurred. This field may not be returned for some errors.
	ErrorCode string `json:"errorCode,omitempty"`

	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get instruments instrument position book unauthorized body
func (o *GetInstrumentsInstrumentPositionBookUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetInstrumentsInstrumentPositionBookUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetInstrumentsInstrumentPositionBookUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}