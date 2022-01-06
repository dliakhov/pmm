// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewTestEmailAlertingSettingsParams creates a new TestEmailAlertingSettingsParams object
// with the default values initialized.
func NewTestEmailAlertingSettingsParams() *TestEmailAlertingSettingsParams {
	var ()
	return &TestEmailAlertingSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestEmailAlertingSettingsParamsWithTimeout creates a new TestEmailAlertingSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestEmailAlertingSettingsParamsWithTimeout(timeout time.Duration) *TestEmailAlertingSettingsParams {
	var ()
	return &TestEmailAlertingSettingsParams{

		timeout: timeout,
	}
}

// NewTestEmailAlertingSettingsParamsWithContext creates a new TestEmailAlertingSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestEmailAlertingSettingsParamsWithContext(ctx context.Context) *TestEmailAlertingSettingsParams {
	var ()
	return &TestEmailAlertingSettingsParams{

		Context: ctx,
	}
}

// NewTestEmailAlertingSettingsParamsWithHTTPClient creates a new TestEmailAlertingSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestEmailAlertingSettingsParamsWithHTTPClient(client *http.Client) *TestEmailAlertingSettingsParams {
	var ()
	return &TestEmailAlertingSettingsParams{
		HTTPClient: client,
	}
}

/*TestEmailAlertingSettingsParams contains all the parameters to send to the API endpoint
for the test email alerting settings operation typically these are written to a http.Request
*/
type TestEmailAlertingSettingsParams struct {

	/*Body*/
	Body TestEmailAlertingSettingsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test email alerting settings params
func (o *TestEmailAlertingSettingsParams) WithTimeout(timeout time.Duration) *TestEmailAlertingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test email alerting settings params
func (o *TestEmailAlertingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test email alerting settings params
func (o *TestEmailAlertingSettingsParams) WithContext(ctx context.Context) *TestEmailAlertingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test email alerting settings params
func (o *TestEmailAlertingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test email alerting settings params
func (o *TestEmailAlertingSettingsParams) WithHTTPClient(client *http.Client) *TestEmailAlertingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test email alerting settings params
func (o *TestEmailAlertingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test email alerting settings params
func (o *TestEmailAlertingSettingsParams) WithBody(body TestEmailAlertingSettingsBody) *TestEmailAlertingSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test email alerting settings params
func (o *TestEmailAlertingSettingsParams) SetBody(body TestEmailAlertingSettingsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestEmailAlertingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
