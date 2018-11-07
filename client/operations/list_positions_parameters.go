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

// NewListPositionsParams creates a new ListPositionsParams object
// with the default values initialized.
func NewListPositionsParams() *ListPositionsParams {
	var ()
	return &ListPositionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPositionsParamsWithTimeout creates a new ListPositionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPositionsParamsWithTimeout(timeout time.Duration) *ListPositionsParams {
	var ()
	return &ListPositionsParams{

		timeout: timeout,
	}
}

// NewListPositionsParamsWithContext creates a new ListPositionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPositionsParamsWithContext(ctx context.Context) *ListPositionsParams {
	var ()
	return &ListPositionsParams{

		Context: ctx,
	}
}

// NewListPositionsParamsWithHTTPClient creates a new ListPositionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPositionsParamsWithHTTPClient(client *http.Client) *ListPositionsParams {
	var ()
	return &ListPositionsParams{
		HTTPClient: client,
	}
}

/*ListPositionsParams contains all the parameters to send to the API endpoint
for the list positions operation typically these are written to a http.Request
*/
type ListPositionsParams struct {

	/*Authorization
	  The authorization bearer token previously obtained by the client

	*/
	Authorization string
	/*AccountID
	  Account Identifier

	*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list positions params
func (o *ListPositionsParams) WithTimeout(timeout time.Duration) *ListPositionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list positions params
func (o *ListPositionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list positions params
func (o *ListPositionsParams) WithContext(ctx context.Context) *ListPositionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list positions params
func (o *ListPositionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list positions params
func (o *ListPositionsParams) WithHTTPClient(client *http.Client) *ListPositionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list positions params
func (o *ListPositionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the list positions params
func (o *ListPositionsParams) WithAuthorization(authorization string) *ListPositionsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list positions params
func (o *ListPositionsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the list positions params
func (o *ListPositionsParams) WithAccountID(accountID string) *ListPositionsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list positions params
func (o *ListPositionsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPositionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param accountID
	if err := r.SetPathParam("accountID", o.AccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
