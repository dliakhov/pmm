// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CheckUpdatesReader is a Reader for the CheckUpdates structure.
type CheckUpdatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckUpdatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckUpdatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCheckUpdatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckUpdatesOK creates a CheckUpdatesOK with default headers values
func NewCheckUpdatesOK() *CheckUpdatesOK {
	return &CheckUpdatesOK{}
}

/* CheckUpdatesOK describes a response with status code 200, with default header values.

A successful response.
*/
type CheckUpdatesOK struct {
	Payload *CheckUpdatesOKBody
}

func (o *CheckUpdatesOK) Error() string {
	return fmt.Sprintf("[POST /v1/Updates/Check][%d] checkUpdatesOk  %+v", 200, o.Payload)
}

func (o *CheckUpdatesOK) GetPayload() *CheckUpdatesOKBody {
	return o.Payload
}

func (o *CheckUpdatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(CheckUpdatesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckUpdatesDefault creates a CheckUpdatesDefault with default headers values
func NewCheckUpdatesDefault(code int) *CheckUpdatesDefault {
	return &CheckUpdatesDefault{
		_statusCode: code,
	}
}

/* CheckUpdatesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CheckUpdatesDefault struct {
	_statusCode int

	Payload *CheckUpdatesDefaultBody
}

// Code gets the status code for the check updates default response
func (o *CheckUpdatesDefault) Code() int {
	return o._statusCode
}

func (o *CheckUpdatesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Updates/Check][%d] CheckUpdates default  %+v", o._statusCode, o.Payload)
}

func (o *CheckUpdatesDefault) GetPayload() *CheckUpdatesDefaultBody {
	return o.Payload
}

func (o *CheckUpdatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(CheckUpdatesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckUpdatesBody check updates body
swagger:model CheckUpdatesBody
*/
type CheckUpdatesBody struct {
	// If false, cached information may be returned.
	Force bool `json:"force,omitempty"`

	// If true, only installed version will be in response.
	OnlyInstalledVersion bool `json:"only_installed_version,omitempty"`
}

// Validate validates this check updates body
func (o *CheckUpdatesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this check updates body based on context it is used
func (o *CheckUpdatesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesBody) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesDefaultBody check updates default body
swagger:model CheckUpdatesDefaultBody
*/
type CheckUpdatesDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*CheckUpdatesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this check updates default body
func (o *CheckUpdatesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CheckUpdates default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CheckUpdates default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this check updates default body based on the context it is used
func (o *CheckUpdatesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CheckUpdates default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CheckUpdates default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesDefaultBody) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesDefaultBodyDetailsItems0 check updates default body details items0
swagger:model CheckUpdatesDefaultBodyDetailsItems0
*/
type CheckUpdatesDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this check updates default body details items0
func (o *CheckUpdatesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this check updates default body details items0 based on context it is used
func (o *CheckUpdatesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesOKBody check updates OK body
swagger:model CheckUpdatesOKBody
*/
type CheckUpdatesOKBody struct {
	// True if there is a PMM Server update available.
	UpdateAvailable bool `json:"update_available,omitempty"`

	// Latest available PMM Server release announcement URL.
	LatestNewsURL string `json:"latest_news_url,omitempty"`

	// Last check time.
	// Format: date-time
	LastCheck strfmt.DateTime `json:"last_check,omitempty"`

	// installed
	Installed *CheckUpdatesOKBodyInstalled `json:"installed,omitempty"`

	// latest
	Latest *CheckUpdatesOKBodyLatest `json:"latest,omitempty"`
}

// Validate validates this check updates OK body
func (o *CheckUpdatesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLastCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInstalled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLatest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesOKBody) validateLastCheck(formats strfmt.Registry) error {
	if swag.IsZero(o.LastCheck) { // not required
		return nil
	}

	if err := validate.FormatOf("checkUpdatesOk"+"."+"last_check", "body", "date-time", o.LastCheck.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CheckUpdatesOKBody) validateInstalled(formats strfmt.Registry) error {
	if swag.IsZero(o.Installed) { // not required
		return nil
	}

	if o.Installed != nil {
		if err := o.Installed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkUpdatesOk" + "." + "installed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkUpdatesOk" + "." + "installed")
			}
			return err
		}
	}

	return nil
}

func (o *CheckUpdatesOKBody) validateLatest(formats strfmt.Registry) error {
	if swag.IsZero(o.Latest) { // not required
		return nil
	}

	if o.Latest != nil {
		if err := o.Latest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkUpdatesOk" + "." + "latest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkUpdatesOk" + "." + "latest")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this check updates OK body based on the context it is used
func (o *CheckUpdatesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateInstalled(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLatest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesOKBody) contextValidateInstalled(ctx context.Context, formats strfmt.Registry) error {
	if o.Installed != nil {
		if err := o.Installed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkUpdatesOk" + "." + "installed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkUpdatesOk" + "." + "installed")
			}
			return err
		}
	}

	return nil
}

func (o *CheckUpdatesOKBody) contextValidateLatest(ctx context.Context, formats strfmt.Registry) error {
	if o.Latest != nil {
		if err := o.Latest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checkUpdatesOk" + "." + "latest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checkUpdatesOk" + "." + "latest")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesOKBody) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesOKBodyInstalled VersionInfo describes component version, or PMM Server as a whole.
swagger:model CheckUpdatesOKBodyInstalled
*/
type CheckUpdatesOKBodyInstalled struct {
	// User-visible version.
	Version string `json:"version,omitempty"`

	// Full version for debugging.
	FullVersion string `json:"full_version,omitempty"`

	// Build or release date.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this check updates OK body installed
func (o *CheckUpdatesOKBodyInstalled) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesOKBodyInstalled) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("checkUpdatesOk"+"."+"installed"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this check updates OK body installed based on context it is used
func (o *CheckUpdatesOKBodyInstalled) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesOKBodyInstalled) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesOKBodyInstalled) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesOKBodyInstalled
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckUpdatesOKBodyLatest VersionInfo describes component version, or PMM Server as a whole.
swagger:model CheckUpdatesOKBodyLatest
*/
type CheckUpdatesOKBodyLatest struct {
	// User-visible version.
	Version string `json:"version,omitempty"`

	// Full version for debugging.
	FullVersion string `json:"full_version,omitempty"`

	// Build or release date.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this check updates OK body latest
func (o *CheckUpdatesOKBodyLatest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckUpdatesOKBodyLatest) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("checkUpdatesOk"+"."+"latest"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this check updates OK body latest based on context it is used
func (o *CheckUpdatesOKBodyLatest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckUpdatesOKBodyLatest) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckUpdatesOKBodyLatest) UnmarshalBinary(b []byte) error {
	var res CheckUpdatesOKBodyLatest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
