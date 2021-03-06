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

// NewSetTradeDependentOrdersParams creates a new SetTradeDependentOrdersParams object
// with the default values initialized.
func NewSetTradeDependentOrdersParams() *SetTradeDependentOrdersParams {
	var ()
	return &SetTradeDependentOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetTradeDependentOrdersParamsWithTimeout creates a new SetTradeDependentOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetTradeDependentOrdersParamsWithTimeout(timeout time.Duration) *SetTradeDependentOrdersParams {
	var ()
	return &SetTradeDependentOrdersParams{

		timeout: timeout,
	}
}

// NewSetTradeDependentOrdersParamsWithContext creates a new SetTradeDependentOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetTradeDependentOrdersParamsWithContext(ctx context.Context) *SetTradeDependentOrdersParams {
	var ()
	return &SetTradeDependentOrdersParams{

		Context: ctx,
	}
}

// NewSetTradeDependentOrdersParamsWithHTTPClient creates a new SetTradeDependentOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetTradeDependentOrdersParamsWithHTTPClient(client *http.Client) *SetTradeDependentOrdersParams {
	var ()
	return &SetTradeDependentOrdersParams{
		HTTPClient: client,
	}
}

/*SetTradeDependentOrdersParams contains all the parameters to send to the API endpoint
for the set trade dependent orders operation typically these are written to a http.Request
*/
type SetTradeDependentOrdersParams struct {

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
	/*SetTradeDependentOrdersBody
	  Details of how to modify the Trade's dependent Orders.

	*/
	SetTradeDependentOrdersBody SetTradeDependentOrdersBody
	/*TradeSpecifier
	  Specifier for the Trade

	*/
	TradeSpecifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) WithTimeout(timeout time.Duration) *SetTradeDependentOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) WithContext(ctx context.Context) *SetTradeDependentOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) WithHTTPClient(client *http.Client) *SetTradeDependentOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *SetTradeDependentOrdersParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) WithAuthorization(authorization string) *SetTradeDependentOrdersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) WithAccountID(accountID string) *SetTradeDependentOrdersParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithSetTradeDependentOrdersBody adds the setTradeDependentOrdersBody to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) WithSetTradeDependentOrdersBody(setTradeDependentOrdersBody SetTradeDependentOrdersBody) *SetTradeDependentOrdersParams {
	o.SetSetTradeDependentOrdersBody(setTradeDependentOrdersBody)
	return o
}

// SetSetTradeDependentOrdersBody adds the setTradeDependentOrdersBody to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) SetSetTradeDependentOrdersBody(setTradeDependentOrdersBody SetTradeDependentOrdersBody) {
	o.SetTradeDependentOrdersBody = setTradeDependentOrdersBody
}

// WithTradeSpecifier adds the tradeSpecifier to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) WithTradeSpecifier(tradeSpecifier string) *SetTradeDependentOrdersParams {
	o.SetTradeSpecifier(tradeSpecifier)
	return o
}

// SetTradeSpecifier adds the tradeSpecifier to the set trade dependent orders params
func (o *SetTradeDependentOrdersParams) SetTradeSpecifier(tradeSpecifier string) {
	o.TradeSpecifier = tradeSpecifier
}

// WriteToRequest writes these params to a swagger request
func (o *SetTradeDependentOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.SetTradeDependentOrdersBody); err != nil {
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
