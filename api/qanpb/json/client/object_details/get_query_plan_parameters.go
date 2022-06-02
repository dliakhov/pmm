// Code generated by go-swagger; DO NOT EDIT.

package object_details

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

// NewGetQueryPlanParams creates a new GetQueryPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQueryPlanParams() *GetQueryPlanParams {
	return &GetQueryPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQueryPlanParamsWithTimeout creates a new GetQueryPlanParams object
// with the ability to set a timeout on a request.
func NewGetQueryPlanParamsWithTimeout(timeout time.Duration) *GetQueryPlanParams {
	return &GetQueryPlanParams{
		timeout: timeout,
	}
}

// NewGetQueryPlanParamsWithContext creates a new GetQueryPlanParams object
// with the ability to set a context for a request.
func NewGetQueryPlanParamsWithContext(ctx context.Context) *GetQueryPlanParams {
	return &GetQueryPlanParams{
		Context: ctx,
	}
}

// NewGetQueryPlanParamsWithHTTPClient creates a new GetQueryPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQueryPlanParamsWithHTTPClient(client *http.Client) *GetQueryPlanParams {
	return &GetQueryPlanParams{
		HTTPClient: client,
	}
}

/* GetQueryPlanParams contains all the parameters to send to the API endpoint
   for the get query plan operation.

   Typically these are written to a http.Request.
*/
type GetQueryPlanParams struct {
	// Body.
	Body GetQueryPlanBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get query plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueryPlanParams) WithDefaults() *GetQueryPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get query plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueryPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get query plan params
func (o *GetQueryPlanParams) WithTimeout(timeout time.Duration) *GetQueryPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get query plan params
func (o *GetQueryPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get query plan params
func (o *GetQueryPlanParams) WithContext(ctx context.Context) *GetQueryPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get query plan params
func (o *GetQueryPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get query plan params
func (o *GetQueryPlanParams) WithHTTPClient(client *http.Client) *GetQueryPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get query plan params
func (o *GetQueryPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get query plan params
func (o *GetQueryPlanParams) WithBody(body GetQueryPlanBody) *GetQueryPlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get query plan params
func (o *GetQueryPlanParams) SetBody(body GetQueryPlanBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetQueryPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
