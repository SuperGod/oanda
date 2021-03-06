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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListOrdersParams creates a new ListOrdersParams object
// with the default values initialized.
func NewListOrdersParams() *ListOrdersParams {
	var ()
	return &ListOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOrdersParamsWithTimeout creates a new ListOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOrdersParamsWithTimeout(timeout time.Duration) *ListOrdersParams {
	var ()
	return &ListOrdersParams{

		timeout: timeout,
	}
}

// NewListOrdersParamsWithContext creates a new ListOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOrdersParamsWithContext(ctx context.Context) *ListOrdersParams {
	var ()
	return &ListOrdersParams{

		Context: ctx,
	}
}

// NewListOrdersParamsWithHTTPClient creates a new ListOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOrdersParamsWithHTTPClient(client *http.Client) *ListOrdersParams {
	var ()
	return &ListOrdersParams{
		HTTPClient: client,
	}
}

/*ListOrdersParams contains all the parameters to send to the API endpoint
for the list orders operation typically these are written to a http.Request
*/
type ListOrdersParams struct {

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
	/*BeforeID
	  The maximum Order ID to return. If not provided the most recent Orders in the Account are returned

	*/
	BeforeID *string
	/*Count
	  The maximum number of Orders to return

	*/
	Count *int64
	/*Ids
	  List of Order IDs to retrieve

	*/
	Ids []string
	/*Instrument
	  The instrument to filter the requested orders by

	*/
	Instrument *string
	/*State
	  The state to filter the requested Orders by

	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list orders params
func (o *ListOrdersParams) WithTimeout(timeout time.Duration) *ListOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list orders params
func (o *ListOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list orders params
func (o *ListOrdersParams) WithContext(ctx context.Context) *ListOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list orders params
func (o *ListOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list orders params
func (o *ListOrdersParams) WithHTTPClient(client *http.Client) *ListOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list orders params
func (o *ListOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptDatetimeFormat adds the acceptDatetimeFormat to the list orders params
func (o *ListOrdersParams) WithAcceptDatetimeFormat(acceptDatetimeFormat *string) *ListOrdersParams {
	o.SetAcceptDatetimeFormat(acceptDatetimeFormat)
	return o
}

// SetAcceptDatetimeFormat adds the acceptDatetimeFormat to the list orders params
func (o *ListOrdersParams) SetAcceptDatetimeFormat(acceptDatetimeFormat *string) {
	o.AcceptDatetimeFormat = acceptDatetimeFormat
}

// WithAuthorization adds the authorization to the list orders params
func (o *ListOrdersParams) WithAuthorization(authorization string) *ListOrdersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list orders params
func (o *ListOrdersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAccountID adds the accountID to the list orders params
func (o *ListOrdersParams) WithAccountID(accountID string) *ListOrdersParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list orders params
func (o *ListOrdersParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithBeforeID adds the beforeID to the list orders params
func (o *ListOrdersParams) WithBeforeID(beforeID *string) *ListOrdersParams {
	o.SetBeforeID(beforeID)
	return o
}

// SetBeforeID adds the beforeId to the list orders params
func (o *ListOrdersParams) SetBeforeID(beforeID *string) {
	o.BeforeID = beforeID
}

// WithCount adds the count to the list orders params
func (o *ListOrdersParams) WithCount(count *int64) *ListOrdersParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the list orders params
func (o *ListOrdersParams) SetCount(count *int64) {
	o.Count = count
}

// WithIds adds the ids to the list orders params
func (o *ListOrdersParams) WithIds(ids []string) *ListOrdersParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the list orders params
func (o *ListOrdersParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithInstrument adds the instrument to the list orders params
func (o *ListOrdersParams) WithInstrument(instrument *string) *ListOrdersParams {
	o.SetInstrument(instrument)
	return o
}

// SetInstrument adds the instrument to the list orders params
func (o *ListOrdersParams) SetInstrument(instrument *string) {
	o.Instrument = instrument
}

// WithState adds the state to the list orders params
func (o *ListOrdersParams) WithState(state *string) *ListOrdersParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the list orders params
func (o *ListOrdersParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *ListOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.BeforeID != nil {

		// query param beforeID
		var qrBeforeID string
		if o.BeforeID != nil {
			qrBeforeID = *o.BeforeID
		}
		qBeforeID := qrBeforeID
		if qBeforeID != "" {
			if err := r.SetQueryParam("beforeID", qBeforeID); err != nil {
				return err
			}
		}

	}

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "csv")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if o.Instrument != nil {

		// query param instrument
		var qrInstrument string
		if o.Instrument != nil {
			qrInstrument = *o.Instrument
		}
		qInstrument := qrInstrument
		if qInstrument != "" {
			if err := r.SetQueryParam("instrument", qInstrument); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
