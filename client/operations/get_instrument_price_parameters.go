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

// NewGetInstrumentPriceParams creates a new GetInstrumentPriceParams object
// with the default values initialized.
func NewGetInstrumentPriceParams() *GetInstrumentPriceParams {
	var ()
	return &GetInstrumentPriceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstrumentPriceParamsWithTimeout creates a new GetInstrumentPriceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstrumentPriceParamsWithTimeout(timeout time.Duration) *GetInstrumentPriceParams {
	var ()
	return &GetInstrumentPriceParams{

		timeout: timeout,
	}
}

// NewGetInstrumentPriceParamsWithContext creates a new GetInstrumentPriceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstrumentPriceParamsWithContext(ctx context.Context) *GetInstrumentPriceParams {
	var ()
	return &GetInstrumentPriceParams{

		Context: ctx,
	}
}

// NewGetInstrumentPriceParamsWithHTTPClient creates a new GetInstrumentPriceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstrumentPriceParamsWithHTTPClient(client *http.Client) *GetInstrumentPriceParams {
	var ()
	return &GetInstrumentPriceParams{
		HTTPClient: client,
	}
}

/*GetInstrumentPriceParams contains all the parameters to send to the API endpoint
for the get instrument price operation typically these are written to a http.Request
*/
type GetInstrumentPriceParams struct {

	/*AcceptDatetimeFormat
	  Format of DateTime fields in the request and response.

	*/
	AcceptDatetimeFormat *string
	/*Authorization
	  The authorization bearer token previously obtained by the client

	*/
	Authorization string
	/*Instrument
	  Name of the Instrument

	*/
	Instrument string
	/*Time
	  The time at which the desired price is in effect. The current price is returned if no time is provided.

	*/
	Time *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instrument price params
func (o *GetInstrumentPriceParams) WithTimeout(timeout time.Duration) *GetInstrumentPriceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instrument price params
func (o *GetInstrumentPriceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instrument price params
func (o *GetInstrumentPriceParams) WithContext(ctx context.Context) *GetInstrumentPriceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instrument price params
func (o *GetInstrumentPriceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instrument price params
func (o *GetInstrumentPriceParams) WithHTTPClient(client *http.Client) *GetInstrumentPriceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instrument price params
func (o *GetInstrumentPriceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the get instrument price params
func (o *GetInstrumentPriceParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *GetInstrumentPriceParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the get instrument price params
func (o *GetInstrumentPriceParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the get instrument price params
func (o *GetInstrumentPriceParams) WithAuthorization(authorization string) *GetInstrumentPriceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get instrument price params
func (o *GetInstrumentPriceParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithInstrument adds the instrument to the get instrument price params
func (o *GetInstrumentPriceParams) WithInstrument(instrument string) *GetInstrumentPriceParams {
	o.SetInstrument(instrument)
	return o
}

// SetInstrument adds the instrument to the get instrument price params
func (o *GetInstrumentPriceParams) SetInstrument(instrument string) {
	o.Instrument = instrument
}

// WithTime adds the time to the get instrument price params
func (o *GetInstrumentPriceParams) WithTime(time *string) *GetInstrumentPriceParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the get instrument price params
func (o *GetInstrumentPriceParams) SetTime(time *string) {
	o.Time = time
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstrumentPriceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param instrument
	if err := r.SetPathParam("instrument", o.Instrument); err != nil {
		return err
	}

	if o.Time != nil {

		// query param time
		var qrTime string
		if o.Time != nil {
			qrTime = *o.Time
		}
		qTime := qrTime
		if qTime != "" {
			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}