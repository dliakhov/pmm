// Code generated by go-swagger; DO NOT EDIT.

package channels

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
)

// ChangeChannelReader is a Reader for the ChangeChannel structure.
type ChangeChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeChannelOK creates a ChangeChannelOK with default headers values
func NewChangeChannelOK() *ChangeChannelOK {
	return &ChangeChannelOK{}
}

/* ChangeChannelOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeChannelOK struct {
	Payload interface{}
}

func (o *ChangeChannelOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Channels/Change][%d] changeChannelOk  %+v", 200, o.Payload)
}

func (o *ChangeChannelOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeChannelDefault creates a ChangeChannelDefault with default headers values
func NewChangeChannelDefault(code int) *ChangeChannelDefault {
	return &ChangeChannelDefault{
		_statusCode: code,
	}
}

/* ChangeChannelDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeChannelDefault struct {
	_statusCode int

	Payload *ChangeChannelDefaultBody
}

// Code gets the status code for the change channel default response
func (o *ChangeChannelDefault) Code() int {
	return o._statusCode
}

func (o *ChangeChannelDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Channels/Change][%d] ChangeChannel default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeChannelDefault) GetPayload() *ChangeChannelDefaultBody {
	return o.Payload
}

func (o *ChangeChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeChannelDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeChannelBody change channel body
swagger:model ChangeChannelBody
*/
type ChangeChannelBody struct {
	// Machine-readable ID.
	ChannelID string `json:"channel_id,omitempty"`

	// Short human-readable summary. Empty value will not change it.
	Summary string `json:"summary,omitempty"`

	// Enables or disables that channel. Should be set.
	Disabled bool `json:"disabled,omitempty"`

	// email config
	EmailConfig *ChangeChannelParamsBodyEmailConfig `json:"email_config,omitempty"`

	// pagerduty config
	PagerdutyConfig *ChangeChannelParamsBodyPagerdutyConfig `json:"pagerduty_config,omitempty"`

	// slack config
	SlackConfig *ChangeChannelParamsBodySlackConfig `json:"slack_config,omitempty"`

	// webhook config
	WebhookConfig *ChangeChannelParamsBodyWebhookConfig `json:"webhook_config,omitempty"`
}

// Validate validates this change channel body
func (o *ChangeChannelBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmailConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagerdutyConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlackConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWebhookConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeChannelBody) validateEmailConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.EmailConfig) { // not required
		return nil
	}

	if o.EmailConfig != nil {
		if err := o.EmailConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "email_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "email_config")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeChannelBody) validatePagerdutyConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.PagerdutyConfig) { // not required
		return nil
	}

	if o.PagerdutyConfig != nil {
		if err := o.PagerdutyConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pagerduty_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pagerduty_config")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeChannelBody) validateSlackConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.SlackConfig) { // not required
		return nil
	}

	if o.SlackConfig != nil {
		if err := o.SlackConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "slack_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "slack_config")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeChannelBody) validateWebhookConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.WebhookConfig) { // not required
		return nil
	}

	if o.WebhookConfig != nil {
		if err := o.WebhookConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "webhook_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change channel body based on the context it is used
func (o *ChangeChannelBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEmailConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePagerdutyConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSlackConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWebhookConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeChannelBody) contextValidateEmailConfig(ctx context.Context, formats strfmt.Registry) error {
	if o.EmailConfig != nil {
		if err := o.EmailConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "email_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "email_config")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeChannelBody) contextValidatePagerdutyConfig(ctx context.Context, formats strfmt.Registry) error {
	if o.PagerdutyConfig != nil {
		if err := o.PagerdutyConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pagerduty_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "pagerduty_config")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeChannelBody) contextValidateSlackConfig(ctx context.Context, formats strfmt.Registry) error {
	if o.SlackConfig != nil {
		if err := o.SlackConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "slack_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "slack_config")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeChannelBody) contextValidateWebhookConfig(ctx context.Context, formats strfmt.Registry) error {
	if o.WebhookConfig != nil {
		if err := o.WebhookConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "webhook_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelBody) UnmarshalBinary(b []byte) error {
	var res ChangeChannelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeChannelDefaultBody change channel default body
swagger:model ChangeChannelDefaultBody
*/
type ChangeChannelDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeChannelDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change channel default body
func (o *ChangeChannelDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeChannelDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChangeChannel default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeChannel default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change channel default body based on the context it is used
func (o *ChangeChannelDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeChannelDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeChannel default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeChannel default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeChannelDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeChannelDefaultBodyDetailsItems0 change channel default body details items0
swagger:model ChangeChannelDefaultBodyDetailsItems0
*/
type ChangeChannelDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this change channel default body details items0
func (o *ChangeChannelDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change channel default body details items0 based on context it is used
func (o *ChangeChannelDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeChannelDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeChannelParamsBodyEmailConfig EmailConfig represents email configuration.
swagger:model ChangeChannelParamsBodyEmailConfig
*/
type ChangeChannelParamsBodyEmailConfig struct {
	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// to
	To []string `json:"to"`
}

// Validate validates this change channel params body email config
func (o *ChangeChannelParamsBodyEmailConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change channel params body email config based on context it is used
func (o *ChangeChannelParamsBodyEmailConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelParamsBodyEmailConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelParamsBodyEmailConfig) UnmarshalBinary(b []byte) error {
	var res ChangeChannelParamsBodyEmailConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeChannelParamsBodyPagerdutyConfig PagerDutyConfig represents PagerDuty configuration.
swagger:model ChangeChannelParamsBodyPagerdutyConfig
*/
type ChangeChannelParamsBodyPagerdutyConfig struct {
	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// The PagerDuty key for "Events API v2" integration type. Exactly one key should be set.
	RoutingKey string `json:"routing_key,omitempty"`

	// The PagerDuty key for "Prometheus" integration type. Exactly one key should be set.
	ServiceKey string `json:"service_key,omitempty"`
}

// Validate validates this change channel params body pagerduty config
func (o *ChangeChannelParamsBodyPagerdutyConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change channel params body pagerduty config based on context it is used
func (o *ChangeChannelParamsBodyPagerdutyConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelParamsBodyPagerdutyConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelParamsBodyPagerdutyConfig) UnmarshalBinary(b []byte) error {
	var res ChangeChannelParamsBodyPagerdutyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeChannelParamsBodySlackConfig SlackConfig represents Slack configuration.
swagger:model ChangeChannelParamsBodySlackConfig
*/
type ChangeChannelParamsBodySlackConfig struct {
	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// channel
	Channel string `json:"channel,omitempty"`
}

// Validate validates this change channel params body slack config
func (o *ChangeChannelParamsBodySlackConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change channel params body slack config based on context it is used
func (o *ChangeChannelParamsBodySlackConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelParamsBodySlackConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelParamsBodySlackConfig) UnmarshalBinary(b []byte) error {
	var res ChangeChannelParamsBodySlackConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeChannelParamsBodyWebhookConfig WebhookConfig represents webhook configuration.
swagger:model ChangeChannelParamsBodyWebhookConfig
*/
type ChangeChannelParamsBodyWebhookConfig struct {
	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// max alerts
	MaxAlerts int32 `json:"max_alerts,omitempty"`

	// http config
	HTTPConfig *ChangeChannelParamsBodyWebhookConfigHTTPConfig `json:"http_config,omitempty"`
}

// Validate validates this change channel params body webhook config
func (o *ChangeChannelParamsBodyWebhookConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeChannelParamsBodyWebhookConfig) validateHTTPConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.HTTPConfig) { // not required
		return nil
	}

	if o.HTTPConfig != nil {
		if err := o.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config" + "." + "http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "webhook_config" + "." + "http_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change channel params body webhook config based on the context it is used
func (o *ChangeChannelParamsBodyWebhookConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeChannelParamsBodyWebhookConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {
	if o.HTTPConfig != nil {
		if err := o.HTTPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config" + "." + "http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "webhook_config" + "." + "http_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelParamsBodyWebhookConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelParamsBodyWebhookConfig) UnmarshalBinary(b []byte) error {
	var res ChangeChannelParamsBodyWebhookConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeChannelParamsBodyWebhookConfigHTTPConfig HTTPConfig represents HTTP client configuration.
swagger:model ChangeChannelParamsBodyWebhookConfigHTTPConfig
*/
type ChangeChannelParamsBodyWebhookConfigHTTPConfig struct {
	// bearer token
	BearerToken string `json:"bearer_token,omitempty"`

	// bearer token file
	BearerTokenFile string `json:"bearer_token_file,omitempty"`

	// proxy url
	ProxyURL string `json:"proxy_url,omitempty"`

	// basic auth
	BasicAuth *ChangeChannelParamsBodyWebhookConfigHTTPConfigBasicAuth `json:"basic_auth,omitempty"`

	// tls config
	TLSConfig *ChangeChannelParamsBodyWebhookConfigHTTPConfigTLSConfig `json:"tls_config,omitempty"`
}

// Validate validates this change channel params body webhook config HTTP config
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBasicAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfig) validateBasicAuth(formats strfmt.Registry) error {
	if swag.IsZero(o.BasicAuth) { // not required
		return nil
	}

	if o.BasicAuth != nil {
		if err := o.BasicAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "basic_auth")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfig) validateTLSConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.TLSConfig) { // not required
		return nil
	}

	if o.TLSConfig != nil {
		if err := o.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "tls_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change channel params body webhook config HTTP config based on the context it is used
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBasicAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTLSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfig) contextValidateBasicAuth(ctx context.Context, formats strfmt.Registry) error {
	if o.BasicAuth != nil {
		if err := o.BasicAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "basic_auth")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfig) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {
	if o.TLSConfig != nil {
		if err := o.TLSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "webhook_config" + "." + "http_config" + "." + "tls_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfig) UnmarshalBinary(b []byte) error {
	var res ChangeChannelParamsBodyWebhookConfigHTTPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeChannelParamsBodyWebhookConfigHTTPConfigBasicAuth BasicAuth represents basic HTTP auth configuration.
swagger:model ChangeChannelParamsBodyWebhookConfigHTTPConfigBasicAuth
*/
type ChangeChannelParamsBodyWebhookConfigHTTPConfigBasicAuth struct {
	// username
	Username string `json:"username,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// password file
	PasswordFile string `json:"password_file,omitempty"`
}

// Validate validates this change channel params body webhook config HTTP config basic auth
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfigBasicAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change channel params body webhook config HTTP config basic auth based on context it is used
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfigBasicAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfigBasicAuth) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfigBasicAuth) UnmarshalBinary(b []byte) error {
	var res ChangeChannelParamsBodyWebhookConfigHTTPConfigBasicAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeChannelParamsBodyWebhookConfigHTTPConfigTLSConfig TLSConfig represents TLS configuration for alertmanager
// https://prometheus.io/docs/alerting/latest/configuration/#tls_config
swagger:model ChangeChannelParamsBodyWebhookConfigHTTPConfigTLSConfig
*/
type ChangeChannelParamsBodyWebhookConfigHTTPConfigTLSConfig struct {
	// A path to the CA certificate file to validate the server certificate with.
	// ca_file and ca_file_content should not be set at the same time.
	CaFile string `json:"ca_file,omitempty"`

	// A path to the certificate file for client cert authentication to the server.
	// cert_file and cert_file_content should not be set at the same time.
	CertFile string `json:"cert_file,omitempty"`

	// A path to the key file for client cert authentication to the server.
	// key_file and key_file_content should not be set at the same time.
	KeyFile string `json:"key_file,omitempty"`

	// Name of the server.
	ServerName string `json:"server_name,omitempty"`

	// Disable validation of the server certificate.
	InsecureSkipVerify bool `json:"insecure_skip_verify,omitempty"`

	// CA certificate to validate the server certificate with.
	// ca_file and ca_file_content should not be set at the same time.
	CaFileContent string `json:"ca_file_content,omitempty"`

	// A certificate for client cert authentication to the server.
	// cert_file and cert_file_content should not be set at the same time.
	CertFileContent string `json:"cert_file_content,omitempty"`

	// A key for client cert authentication to the server.
	// key_file and key_file_content should not be set at the same time.
	KeyFileContent string `json:"key_file_content,omitempty"`
}

// Validate validates this change channel params body webhook config HTTP config TLS config
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfigTLSConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change channel params body webhook config HTTP config TLS config based on context it is used
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfigTLSConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfigTLSConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeChannelParamsBodyWebhookConfigHTTPConfigTLSConfig) UnmarshalBinary(b []byte) error {
	var res ChangeChannelParamsBodyWebhookConfigHTTPConfigTLSConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
