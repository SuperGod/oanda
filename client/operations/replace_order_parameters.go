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

// NewReplaceOrderParams creates a new ReplaceOrderParams object
// with the default values initialized.
func NewReplaceOrderParams() *ReplaceOrderParams {
	var ()
	return &ReplaceOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceOrderParamsWithTimeout creates a new ReplaceOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceOrderParamsWithTimeout(timeout time.Duration) *ReplaceOrderParams {
	var ()
	return &ReplaceOrderParams{

		timeout: timeout,
	}
}

// NewReplaceOrderParamsWithContext creates a new ReplaceOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceOrderParamsWithContext(ctx context.Context) *ReplaceOrderParams {
	var ()
	return &ReplaceOrderParams{

		Context: ctx,
	}
}

// NewReplaceOrderParamsWithHTTPClient creates a new ReplaceOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceOrderParamsWithHTTPClient(client *http.Client) *ReplaceOrderParams {
	var ()
	return &ReplaceOrderParams{
		HTTPClient: client,
	}
}

/*ReplaceOrderParams contains all the parameters to send to the API endpoint
for the replace order operation typically these are written to a http.Request
*/
type ReplaceOrderParams struct {

	/*AcceptDatetimeFormat
	  Format of DateTime fields in the request and response.

	*/
	AcceptDatetimeFormat *string
	/*Authorization
	  The authorization bearer token previously obtained by the client

	*/
	Authorization string
	/*ClientRequestID
	  Client specified RequestID to be sent with request.

	*/
	ClientRequestID *string
	/*AccountID
	  Account Identifier

	*/
	AccountID string
	/*OrderSpecifier
	  The Order Specifier

	*/
	OrderSpecifier string
	/*ReplaceOrderBody
	  Specification of the replacing Order. The replacing order must have the same type as the replaced Order.

	*/
	ReplaceOrderBody ReplaceOrderBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace order params
func (o *ReplaceOrderParams) WithTimeout(timeout time.Duration) *ReplaceOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace order params
func (o *ReplaceOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace order params
func (o *ReplaceOrderParams) WithContext(ctx context.Context) *ReplaceOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace order params
func (o *ReplaceOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace order params
func (o *ReplaceOrderParams) WithHTTPClient(client *http.Client) *ReplaceOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace order params
func (o *ReplaceOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the replace order params
func (o *ReplaceOrderParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *ReplaceOrderParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the replace order params
func (o *ReplaceOrderParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the replace order params
func (o *ReplaceOrderParams) WithAuthorization(authorization string) *ReplaceOrderParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the replace order params
func (o *ReplaceOrderParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithClientRequestID adds the clientRequestID to the replace order params
func (o *ReplaceOrderParams) WithClientRequestID(clientRequestID *string) *ReplaceOrderParams {
	o.SetClientRequestID(clientRequestID)
	return o
}

// SetClientRequestID adds the clientRequestId to the replace order params
func (o *ReplaceOrderParams) SetClientRequestID(clientRequestID *string) {
	o.ClientRequestID = clientRequestID
}

// WithAccountID adds the accountID to the replace order params
func (o *ReplaceOrderParams) WithAccountID(accountID string) *ReplaceOrderParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the replace order params
func (o *ReplaceOrderParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithOrderSpecifier adds the orderSpecifier to the replace order params
func (o *ReplaceOrderParams) WithOrderSpecifier(orderSpecifier string) *ReplaceOrderParams {
	o.SetOrderSpecifier(orderSpecifier)
	return o
}

// SetOrderSpecifier adds the orderSpecifier to the replace order params
func (o *ReplaceOrderParams) SetOrderSpecifier(orderSpecifier string) {
	o.OrderSpecifier = orderSpecifier
}

// WithReplaceOrderBody adds the replaceOrderBody to the replace order params
func (o *ReplaceOrderParams) WithReplaceOrderBody(replaceOrderBody ReplaceOrderBody) *ReplaceOrderParams {
	o.SetReplaceOrderBody(replaceOrderBody)
	return o
}

// SetReplaceOrderBody adds the replaceOrderBody to the replace order params
func (o *ReplaceOrderParams) SetReplaceOrderBody(replaceOrderBody ReplaceOrderBody) {
	o.ReplaceOrderBody = replaceOrderBody
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ClientRequestID != nil {

		// header param ClientRequestID
		if err := r.SetHeaderParam("ClientRequestID", *o.ClientRequestID); err != nil {
			return err
		}

	}

	// path param accountID
	if err := r.SetPathParam("accountID", o.AccountID); err != nil {
		return err
	}

	// path param orderSpecifier
	if err := r.SetPathParam("orderSpecifier", o.OrderSpecifier); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.ReplaceOrderBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
