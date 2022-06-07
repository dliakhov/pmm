// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewAddAzureDatabaseExporterParams creates a new AddAzureDatabaseExporterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAzureDatabaseExporterParams() *AddAzureDatabaseExporterParams {
	return &AddAzureDatabaseExporterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAzureDatabaseExporterParamsWithTimeout creates a new AddAzureDatabaseExporterParams object
// with the ability to set a timeout on a request.
func NewAddAzureDatabaseExporterParamsWithTimeout(timeout time.Duration) *AddAzureDatabaseExporterParams {
	return &AddAzureDatabaseExporterParams{
		timeout: timeout,
	}
}

// NewAddAzureDatabaseExporterParamsWithContext creates a new AddAzureDatabaseExporterParams object
// with the ability to set a context for a request.
func NewAddAzureDatabaseExporterParamsWithContext(ctx context.Context) *AddAzureDatabaseExporterParams {
	return &AddAzureDatabaseExporterParams{
		Context: ctx,
	}
}

// NewAddAzureDatabaseExporterParamsWithHTTPClient creates a new AddAzureDatabaseExporterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAzureDatabaseExporterParamsWithHTTPClient(client *http.Client) *AddAzureDatabaseExporterParams {
	return &AddAzureDatabaseExporterParams{
		HTTPClient: client,
	}
}

/* AddAzureDatabaseExporterParams contains all the parameters to send to the API endpoint
   for the add azure database exporter operation.

   Typically these are written to a http.Request.
*/
type AddAzureDatabaseExporterParams struct {
	// Body.
	Body AddAzureDatabaseExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add azure database exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAzureDatabaseExporterParams) WithDefaults() *AddAzureDatabaseExporterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add azure database exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAzureDatabaseExporterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add azure database exporter params
func (o *AddAzureDatabaseExporterParams) WithTimeout(timeout time.Duration) *AddAzureDatabaseExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add azure database exporter params
func (o *AddAzureDatabaseExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add azure database exporter params
func (o *AddAzureDatabaseExporterParams) WithContext(ctx context.Context) *AddAzureDatabaseExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add azure database exporter params
func (o *AddAzureDatabaseExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add azure database exporter params
func (o *AddAzureDatabaseExporterParams) WithHTTPClient(client *http.Client) *AddAzureDatabaseExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add azure database exporter params
func (o *AddAzureDatabaseExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add azure database exporter params
func (o *AddAzureDatabaseExporterParams) WithBody(body AddAzureDatabaseExporterBody) *AddAzureDatabaseExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add azure database exporter params
func (o *AddAzureDatabaseExporterParams) SetBody(body AddAzureDatabaseExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAzureDatabaseExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
