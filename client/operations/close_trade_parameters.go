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

// NewCloseTradeParams creates a new CloseTradeParams object
// with the default values initialized.
func NewCloseTradeParams() *CloseTradeParams {
	var ()
	return &CloseTradeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCloseTradeParamsWithTimeout creates a new CloseTradeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloseTradeParamsWithTimeout(timeout time.Duration) *CloseTradeParams {
	var ()
	return &CloseTradeParams{

		timeout: timeout,
	}
}

// NewCloseTradeParamsWithContext creates a new CloseTradeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloseTradeParamsWithContext(ctx context.Context) *CloseTradeParams {
	var ()
	return &CloseTradeParams{

		Context: ctx,
	}
}

// NewCloseTradeParamsWithHTTPClient creates a new CloseTradeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloseTradeParamsWithHTTPClient(client *http.Client) *CloseTradeParams {
	var ()
	return &CloseTradeParams{
		HTTPClient: client,
	}
}

/*CloseTradeParams contains all the parameters to send to the API endpoint
for the close trade operation typically these are written to a http.Request
*/
type CloseTradeParams struct {

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
	/*CloseTradeBody
	  Details of how much of the open Trade to close.

	*/
	CloseTradeBody CloseTradeBody
	/*TradeSpecifier
	  Specifier for the Trade

	*/
	TradeSpecifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the close trade params
func (o *CloseTradeParams) WithTimeout(timeout time.Duration) *CloseTradeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the close trade params
func (o *CloseTradeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the close trade params
func (o *CloseTradeParams) WithContext(ctx context.Context) *CloseTradeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the close trade params
func (o *CloseTradeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the close trade params
func (o *CloseTradeParams) WithHTTPClient(client *http.Client) *CloseTradeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the close trade params
func (o *CloseTradeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the close trade params
func (o *CloseTradeParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *CloseTradeParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the close trade params
func (o *CloseTradeParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the close trade params
func (o *CloseTradeParams) WithAuthorization(authorization string) *CloseTradeParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the close trade params
func (o *CloseTradeParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the close trade params
func (o *CloseTradeParams) WithAccountID(accountID string) *CloseTradeParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the close trade params
func (o *CloseTradeParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithCloseTradeBody adds the closeTradeBody to the close trade params
func (o *CloseTradeParams) WithCloseTradeBody(closeTradeBody CloseTradeBody) *CloseTradeParams {
	o.SetCloseTradeBody(closeTradeBody)
	return o
}

// SetCloseTradeBody adds the closeTradeBody to the close trade params
func (o *CloseTradeParams) SetCloseTradeBody(closeTradeBody CloseTradeBody) {
	o.CloseTradeBody = closeTradeBody
}

// WithTradeSpecifier adds the tradeSpecifier to the close trade params
func (o *CloseTradeParams) WithTradeSpecifier(tradeSpecifier string) *CloseTradeParams {
	o.SetTradeSpecifier(tradeSpecifier)
	return o
}

// SetTradeSpecifier adds the tradeSpecifier to the close trade params
func (o *CloseTradeParams) SetTradeSpecifier(tradeSpecifier string) {
	o.TradeSpecifier = tradeSpecifier
}

// WriteToRequest writes these params to a swagger request
func (o *CloseTradeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.CloseTradeBody); err != nil {
		return err
	}

	// path param tradeSpecifier
	if err := r.SetPathParam("tradeSpecifier", o.TradeSpecifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
