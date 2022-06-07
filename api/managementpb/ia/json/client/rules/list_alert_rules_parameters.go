// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListAlertRulesParams creates a new ListAlertRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAlertRulesParams() *ListAlertRulesParams {
	return &ListAlertRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAlertRulesParamsWithTimeout creates a new ListAlertRulesParams object
// with the ability to set a timeout on a request.
func NewListAlertRulesParamsWithTimeout(timeout time.Duration) *ListAlertRulesParams {
	return &ListAlertRulesParams{
		timeout: timeout,
	}
}

// NewListAlertRulesParamsWithContext creates a new ListAlertRulesParams object
// with the ability to set a context for a request.
func NewListAlertRulesParamsWithContext(ctx context.Context) *ListAlertRulesParams {
	return &ListAlertRulesParams{
		Context: ctx,
	}
}

// NewListAlertRulesParamsWithHTTPClient creates a new ListAlertRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAlertRulesParamsWithHTTPClient(client *http.Client) *ListAlertRulesParams {
	return &ListAlertRulesParams{
		HTTPClient: client,
	}
}

/* ListAlertRulesParams contains all the parameters to send to the API endpoint
   for the list alert rules operation.

   Typically these are written to a http.Request.
*/
type ListAlertRulesParams struct {
	// Body.
	Body ListAlertRulesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list alert rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAlertRulesParams) WithDefaults() *ListAlertRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list alert rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAlertRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list alert rules params
func (o *ListAlertRulesParams) WithTimeout(timeout time.Duration) *ListAlertRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list alert rules params
func (o *ListAlertRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list alert rules params
func (o *ListAlertRulesParams) WithContext(ctx context.Context) *ListAlertRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list alert rules params
func (o *ListAlertRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list alert rules params
func (o *ListAlertRulesParams) WithHTTPClient(client *http.Client) *ListAlertRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list alert rules params
func (o *ListAlertRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list alert rules params
func (o *ListAlertRulesParams) WithBody(body ListAlertRulesBody) *ListAlertRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list alert rules params
func (o *ListAlertRulesParams) SetBody(body ListAlertRulesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListAlertRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
