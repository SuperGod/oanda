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

// NewGetOrderParams creates a new GetOrderParams object
// with the default values initialized.
func NewGetOrderParams() *GetOrderParams {
	var ()
	return &GetOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrderParamsWithTimeout creates a new GetOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrderParamsWithTimeout(timeout time.Duration) *GetOrderParams {
	var ()
	return &GetOrderParams{

		timeout: timeout,
	}
}

// NewGetOrderParamsWithContext creates a new GetOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrderParamsWithContext(ctx context.Context) *GetOrderParams {
	var ()
	return &GetOrderParams{

		Context: ctx,
	}
}

// NewGetOrderParamsWithHTTPClient creates a new GetOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrderParamsWithHTTPClient(client *http.Client) *GetOrderParams {
	var ()
	return &GetOrderParams{
		HTTPClient: client,
	}
}

/*GetOrderParams contains all the parameters to send to the API endpoint
for the get order operation typically these are written to a http.Request
*/
type GetOrderParams struct {

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
	/*OrderSpecifier
	  The Order Specifier

	*/
	OrderSpecifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get order params
func (o *GetOrderParams) WithTimeout(timeout time.Duration) *GetOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get order params
func (o *GetOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get order params
func (o *GetOrderParams) WithContext(ctx context.Context) *GetOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get order params
func (o *GetOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get order params
func (o *GetOrderParams) WithHTTPClient(client *http.Client) *GetOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get order params
func (o *GetOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the get order params
func (o *GetOrderParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *GetOrderParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the get order params
func (o *GetOrderParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the get order params
func (o *GetOrderParams) WithAuthorization(authorization string) *GetOrderParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get order params
func (o *GetOrderParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the get order params
func (o *GetOrderParams) WithAccountID(accountID string) *GetOrderParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get order params
func (o *GetOrderParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithOrderSpecifier adds the orderSpecifier to the get order params
func (o *GetOrderParams) WithOrderSpecifier(orderSpecifier string) *GetOrderParams {
	o.SetOrderSpecifier(orderSpecifier)
	return o
}

// SetOrderSpecifier adds the orderSpecifier to the get order params
func (o *GetOrderParams) SetOrderSpecifier(orderSpecifier string) {
	o.OrderSpecifier = orderSpecifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param orderSpecifier
	if err := r.SetPathParam("orderSpecifier", o.OrderSpecifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
