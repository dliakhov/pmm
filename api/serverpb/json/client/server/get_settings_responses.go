// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSettingsReader is a Reader for the GetSettings structure.
type GetSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSettingsOK creates a GetSettingsOK with default headers values
func NewGetSettingsOK() *GetSettingsOK {
	return &GetSettingsOK{}
}

/*GetSettingsOK handles this case with default header values.

A successful response.
*/
type GetSettingsOK struct {
	Payload *GetSettingsOKBody
}

func (o *GetSettingsOK) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Get][%d] getSettingsOk  %+v", 200, o.Payload)
}

func (o *GetSettingsOK) GetPayload() *GetSettingsOKBody {
	return o.Payload
}

func (o *GetSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSettingsDefault creates a GetSettingsDefault with default headers values
func NewGetSettingsDefault(code int) *GetSettingsDefault {
	return &GetSettingsDefault{
		_statusCode: code,
	}
}

/*GetSettingsDefault handles this case with default header values.

An unexpected error response.
*/
type GetSettingsDefault struct {
	_statusCode int

	Payload *GetSettingsDefaultBody
}

// Code gets the status code for the get settings default response
func (o *GetSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetSettingsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Get][%d] GetSettings default  %+v", o._statusCode, o.Payload)
}

func (o *GetSettingsDefault) GetPayload() *GetSettingsDefaultBody {
	return o.Payload
}

func (o *GetSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSettingsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSettingsDefaultBody get settings default body
swagger:model GetSettingsDefaultBody
*/
type GetSettingsDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get settings default body
func (o *GetSettingsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("GetSettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetSettingsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBody get settings OK body
swagger:model GetSettingsOKBody
*/
type GetSettingsOKBody struct {

	// settings
	Settings *GetSettingsOKBodySettings `json:"settings,omitempty"`
}

// Validate validates this get settings OK body
func (o *GetSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsOKBody) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.Settings) { // not required
		return nil
	}

	if o.Settings != nil {
		if err := o.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettings Settings represents PMM Server settings.
swagger:model GetSettingsOKBodySettings
*/
type GetSettingsOKBodySettings struct {

	// True if updates are disabled.
	UpdatesDisabled bool `json:"updates_disabled,omitempty"`

	// True if telemetry is enabled.
	TelemetryEnabled bool `json:"telemetry_enabled,omitempty"`

	// data retention
	DataRetention string `json:"data_retention,omitempty"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// aws partitions
	AWSPartitions []string `json:"aws_partitions"`

	// External AlertManager URL (e.g., https://username:password@1.2.3.4/path).
	AlertManagerURL string `json:"alert_manager_url,omitempty"`

	// Custom alerting or recording rules.
	AlertManagerRules string `json:"alert_manager_rules,omitempty"`

	// True if Security Threat Tool is enabled.
	SttEnabled bool `json:"stt_enabled,omitempty"`

	// Percona Platform user's email, if this PMM instance is linked to the Platform.
	PlatformEmail string `json:"platform_email,omitempty"`

	// True if DBaaS is enabled.
	DbaasEnabled bool `json:"dbaas_enabled,omitempty"`

	// Alerting is always enabled and cannot be disabled.
	AlertingEnabled bool `json:"alerting_enabled,omitempty"`

	// PMM Server public address.
	PMMPublicAddress string `json:"pmm_public_address,omitempty"`

	// True if Backup Management is enabled.
	BackupManagementEnabled bool `json:"backup_management_enabled,omitempty"`

	// True if Azure Discover is enabled.
	AzurediscoverEnabled bool `json:"azurediscover_enabled,omitempty"`

	// True if the PMM instance is connected to Platform
	ConnectedToPlatform bool `json:"connected_to_platform,omitempty"`

	// email alerting settings
	EmailAlertingSettings *GetSettingsOKBodySettingsEmailAlertingSettings `json:"email_alerting_settings,omitempty"`

	// metrics resolutions
	MetricsResolutions *GetSettingsOKBodySettingsMetricsResolutions `json:"metrics_resolutions,omitempty"`

	// slack alerting settings
	SlackAlertingSettings *GetSettingsOKBodySettingsSlackAlertingSettings `json:"slack_alerting_settings,omitempty"`

	// stt check intervals
	SttCheckIntervals *GetSettingsOKBodySettingsSttCheckIntervals `json:"stt_check_intervals,omitempty"`
}

// Validate validates this get settings OK body settings
func (o *GetSettingsOKBodySettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmailAlertingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetricsResolutions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlackAlertingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSttCheckIntervals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsOKBodySettings) validateEmailAlertingSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.EmailAlertingSettings) { // not required
		return nil
	}

	if o.EmailAlertingSettings != nil {
		if err := o.EmailAlertingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "email_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *GetSettingsOKBodySettings) validateMetricsResolutions(formats strfmt.Registry) error {

	if swag.IsZero(o.MetricsResolutions) { // not required
		return nil
	}

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

func (o *GetSettingsOKBodySettings) validateSlackAlertingSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.SlackAlertingSettings) { // not required
		return nil
	}

	if o.SlackAlertingSettings != nil {
		if err := o.SlackAlertingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "slack_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *GetSettingsOKBodySettings) validateSttCheckIntervals(formats strfmt.Registry) error {

	if swag.IsZero(o.SttCheckIntervals) { // not required
		return nil
	}

	if o.SttCheckIntervals != nil {
		if err := o.SttCheckIntervals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "stt_check_intervals")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettings) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettingsEmailAlertingSettings EmailAlertingSettings represents email (SMTP) configuration for Alerting.
swagger:model GetSettingsOKBodySettingsEmailAlertingSettings
*/
type GetSettingsOKBodySettingsEmailAlertingSettings struct {

	// SMTP From header field.
	From string `json:"from,omitempty"`

	// SMTP host and port.
	Smarthost string `json:"smarthost,omitempty"`

	// Hostname to identify to the SMTP server.
	Hello string `json:"hello,omitempty"`

	// Auth using CRAM-MD5, LOGIN and PLAIN.
	Username string `json:"username,omitempty"`

	// Auth using LOGIN and PLAIN.
	Password string `json:"password,omitempty"`

	// Auth using PLAIN.
	Identity string `json:"identity,omitempty"`

	// Auth using CRAM-MD5.
	Secret string `json:"secret,omitempty"`

	// Require TLS.
	RequireTLS bool `json:"require_tls,omitempty"`
}

// Validate validates this get settings OK body settings email alerting settings
func (o *GetSettingsOKBodySettingsEmailAlertingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsEmailAlertingSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsEmailAlertingSettings) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettingsEmailAlertingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettingsMetricsResolutions MetricsResolutions represents Prometheus exporters metrics resolutions.
swagger:model GetSettingsOKBodySettingsMetricsResolutions
*/
type GetSettingsOKBodySettingsMetricsResolutions struct {

	// High resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Hr string `json:"hr,omitempty"`

	// Medium resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Mr string `json:"mr,omitempty"`

	// Low resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Lr string `json:"lr,omitempty"`
}

// Validate validates this get settings OK body settings metrics resolutions
func (o *GetSettingsOKBodySettingsMetricsResolutions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsMetricsResolutions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsMetricsResolutions) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettingsMetricsResolutions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettingsSlackAlertingSettings SlackAlertingSettings represents Slack configuration for Alerting.
swagger:model GetSettingsOKBodySettingsSlackAlertingSettings
*/
type GetSettingsOKBodySettingsSlackAlertingSettings struct {

	// Slack API (webhook) URL.
	URL string `json:"url,omitempty"`
}

// Validate validates this get settings OK body settings slack alerting settings
func (o *GetSettingsOKBodySettingsSlackAlertingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsSlackAlertingSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsSlackAlertingSettings) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettingsSlackAlertingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettingsSttCheckIntervals STTCheckIntervals represents intervals between STT checks.
swagger:model GetSettingsOKBodySettingsSttCheckIntervals
*/
type GetSettingsOKBodySettingsSttCheckIntervals struct {

	// Standard check interval.
	StandardInterval string `json:"standard_interval,omitempty"`

	// Interval for rare check runs.
	RareInterval string `json:"rare_interval,omitempty"`

	// Interval for frequent check runs.
	FrequentInterval string `json:"frequent_interval,omitempty"`
}

// Validate validates this get settings OK body settings stt check intervals
func (o *GetSettingsOKBodySettingsSttCheckIntervals) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsSttCheckIntervals) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsSttCheckIntervals) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettingsSttCheckIntervals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
