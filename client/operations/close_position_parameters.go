// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewClosePositionParams creates a new ClosePositionParams object
// with the default values initialized.
func NewClosePositionParams() *ClosePositionParams {
	var ()
	return &ClosePositionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewClosePositionParamsWithTimeout creates a new ClosePositionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewClosePositionParamsWithTimeout(timeout time.Duration) *ClosePositionParams {
	var ()
	return &ClosePositionParams{

		timeout: timeout,
	}
}

// NewClosePositionParamsWithContext creates a new ClosePositionParams object
// with the default values initialized, and the ability to set a context for a request
func NewClosePositionParamsWithContext(ctx context.Context) *ClosePositionParams {
	var ()
	return &ClosePositionParams{

		Context: ctx,
	}
}

// NewClosePositionParamsWithHTTPClient creates a new ClosePositionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewClosePositionParamsWithHTTPClient(client *http.Client) *ClosePositionParams {
	var ()
	return &ClosePositionParams{
		HTTPClient: client,
	}
}

/*ClosePositionParams contains all the parameters to send to the API endpoint
for the close position operation typically these are written to a http.Request
*/
type ClosePositionParams struct {

	/*AcceptDatetimeFormat
	  Format of DateTime fields in the request and response.

	*/
	AcceptDatetimeFormat *string
	/*Authorization
	  The authorization bearer token previously obtained by the client

	*/
	Authorization string
	/*AccountID
	  Account Identifier

	*/
	AccountID string
	/*ClosePositionBody
	  Representation of how to close the position

	*/
	ClosePositionBody ClosePositionBody
	/*Instrument
	  Name of the Instrument

	*/
	Instrument string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the close position params
func (o *ClosePositionParams) WithTimeout(timeout time.Duration) *ClosePositionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the close position params
func (o *ClosePositionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the close position params
func (o *ClosePositionParams) WithContext(ctx context.Context) *ClosePositionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the close position params
func (o *ClosePositionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the close position params
func (o *ClosePositionParams) WithHTTPClient(client *http.Client) *ClosePositionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the close position params
func (o *ClosePositionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the close position params
func (o *ClosePositionParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *ClosePositionParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the close position params
func (o *ClosePositionParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the close position params
func (o *ClosePositionParams) WithAuthorization(authorization string) *ClosePositionParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the close position params
func (o *ClosePositionParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the close position params
func (o *ClosePositionParams) WithAccountID(accountID string) *ClosePositionParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the close position params
func (o *ClosePositionParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithClosePositionBody adds the closePositionBody to the close position params
func (o *ClosePositionParams) WithClosePositionBody(closePositionBody ClosePositionBody) *ClosePositionParams {
	o.SetClosePositionBody(closePositionBody)
	return o
}

// SetClosePositionBody adds the closePositionBody to the close position params
func (o *ClosePositionParams) SetClosePositionBody(closePositionBody ClosePositionBody) {
	o.ClosePositionBody = closePositionBody
}

// WithInstrument adds the instrument to the close position params
func (o *ClosePositionParams) WithInstrument(instrument string) *ClosePositionParams {
	o.SetInstrument(instrument)
	return o
}

// SetInstrument adds the instrument to the close position params
func (o *ClosePositionParams) SetInstrument(instrument string) {
	o.Instrument = instrument
}

// WriteToRequest writes these params to a swagger request
func (o *ClosePositionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptDatetimeFormat != nil {

		// header param Accept-Datetime-Format
		if err := r.SetHeaderParam("Accept-Datetime-Format", *o.AcceptDatetimeFormat); err != nil {
			return err
		}

	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param accountID
	if err := r.SetPathParam("accountID", o.AccountID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.ClosePositionBody); err != nil {
		return err
	}

	// path param instrument
	if err := r.SetPathParam("instrument", o.Instrument); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
