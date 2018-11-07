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

// NewGetExternalUserInfoParams creates a new GetExternalUserInfoParams object
// with the default values initialized.
func NewGetExternalUserInfoParams() *GetExternalUserInfoParams {
	var ()
	return &GetExternalUserInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalUserInfoParamsWithTimeout creates a new GetExternalUserInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExternalUserInfoParamsWithTimeout(timeout time.Duration) *GetExternalUserInfoParams {
	var ()
	return &GetExternalUserInfoParams{

		timeout: timeout,
	}
}

// NewGetExternalUserInfoParamsWithContext creates a new GetExternalUserInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExternalUserInfoParamsWithContext(ctx context.Context) *GetExternalUserInfoParams {
	var ()
	return &GetExternalUserInfoParams{

		Context: ctx,
	}
}

// NewGetExternalUserInfoParamsWithHTTPClient creates a new GetExternalUserInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExternalUserInfoParamsWithHTTPClient(client *http.Client) *GetExternalUserInfoParams {
	var ()
	return &GetExternalUserInfoParams{
		HTTPClient: client,
	}
}

/*GetExternalUserInfoParams contains all the parameters to send to the API endpoint
for the get external user info operation typically these are written to a http.Request
*/
type GetExternalUserInfoParams struct {

	/*Authorization
	  The authorization bearer token previously obtained by the client

	*/
	Authorization string
	/*UserSpecifier
	  The User Specifier

	*/
	UserSpecifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get external user info params
func (o *GetExternalUserInfoParams) WithTimeout(timeout time.Duration) *GetExternalUserInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external user info params
func (o *GetExternalUserInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external user info params
func (o *GetExternalUserInfoParams) WithContext(ctx context.Context) *GetExternalUserInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external user info params
func (o *GetExternalUserInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external user info params
func (o *GetExternalUserInfoParams) WithHTTPClient(client *http.Client) *GetExternalUserInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external user info params
func (o *GetExternalUserInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get external user info params
func (o *GetExternalUserInfoParams) WithAuthorization(authorization string) *GetExternalUserInfoParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get external user info params
func (o *GetExternalUserInfoParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithUserSpecifier adds the userSpecifier to the get external user info params
func (o *GetExternalUserInfoParams) WithUserSpecifier(userSpecifier string) *GetExternalUserInfoParams {
	o.SetUserSpecifier(userSpecifier)
	return o
}

// SetUserSpecifier adds the userSpecifier to the get external user info params
func (o *GetExternalUserInfoParams) SetUserSpecifier(userSpecifier string) {
	o.UserSpecifier = userSpecifier
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalUserInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param userSpecifier
	if err := r.SetPathParam("userSpecifier", o.UserSpecifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}