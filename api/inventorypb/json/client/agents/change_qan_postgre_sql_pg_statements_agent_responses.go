// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeQANPostgreSQLPgStatementsAgentReader is a Reader for the ChangeQANPostgreSQLPgStatementsAgent structure.
type ChangeQANPostgreSQLPgStatementsAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeQANPostgreSQLPgStatementsAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeQANPostgreSQLPgStatementsAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeQANPostgreSQLPgStatementsAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeQANPostgreSQLPgStatementsAgentOK creates a ChangeQANPostgreSQLPgStatementsAgentOK with default headers values
func NewChangeQANPostgreSQLPgStatementsAgentOK() *ChangeQANPostgreSQLPgStatementsAgentOK {
	return &ChangeQANPostgreSQLPgStatementsAgentOK{}
}

/*ChangeQANPostgreSQLPgStatementsAgentOK handles this case with default header values.

A successful response.
*/
type ChangeQANPostgreSQLPgStatementsAgentOK struct {
	Payload *ChangeQANPostgreSQLPgStatementsAgentOKBody
}

func (o *ChangeQANPostgreSQLPgStatementsAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANPostgreSQLPgStatementsAgent][%d] changeQanPostgreSqlPgStatementsAgentOk  %+v", 200, o.Payload)
}

func (o *ChangeQANPostgreSQLPgStatementsAgentOK) GetPayload() *ChangeQANPostgreSQLPgStatementsAgentOKBody {
	return o.Payload
}

func (o *ChangeQANPostgreSQLPgStatementsAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeQANPostgreSQLPgStatementsAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeQANPostgreSQLPgStatementsAgentDefault creates a ChangeQANPostgreSQLPgStatementsAgentDefault with default headers values
func NewChangeQANPostgreSQLPgStatementsAgentDefault(code int) *ChangeQANPostgreSQLPgStatementsAgentDefault {
	return &ChangeQANPostgreSQLPgStatementsAgentDefault{
		_statusCode: code,
	}
}

/*ChangeQANPostgreSQLPgStatementsAgentDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeQANPostgreSQLPgStatementsAgentDefault struct {
	_statusCode int

	Payload *ChangeQANPostgreSQLPgStatementsAgentDefaultBody
}

// Code gets the status code for the change QAN postgre SQL pg statements agent default response
func (o *ChangeQANPostgreSQLPgStatementsAgentDefault) Code() int {
	return o._statusCode
}

func (o *ChangeQANPostgreSQLPgStatementsAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/ChangeQANPostgreSQLPgStatementsAgent][%d] ChangeQANPostgreSQLPgStatementsAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeQANPostgreSQLPgStatementsAgentDefault) GetPayload() *ChangeQANPostgreSQLPgStatementsAgentDefaultBody {
	return o.Payload
}

func (o *ChangeQANPostgreSQLPgStatementsAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeQANPostgreSQLPgStatementsAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeQANPostgreSQLPgStatementsAgentBody change QAN postgre SQL pg statements agent body
swagger:model ChangeQANPostgreSQLPgStatementsAgentBody
*/
type ChangeQANPostgreSQLPgStatementsAgentBody struct {

	// agent Id
	AgentID string `json:"agentId,omitempty"`

	// common
	Common *ChangeQANPostgreSQLPgStatementsAgentParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this change QAN postgre SQL pg statements agent body
func (o *ChangeQANPostgreSQLPgStatementsAgentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatementsAgentBody) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatementsAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANPostgreSQLPgStatementsAgentDefaultBody change QAN postgre SQL pg statements agent default body
swagger:model ChangeQANPostgreSQLPgStatementsAgentDefaultBody
*/
type ChangeQANPostgreSQLPgStatementsAgentDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change QAN postgre SQL pg statements agent default body
func (o *ChangeQANPostgreSQLPgStatementsAgentDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatementsAgentDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ChangeQANPostgreSQLPgStatementsAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatementsAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANPostgreSQLPgStatementsAgentOKBody change QAN postgre SQL pg statements agent OK body
swagger:model ChangeQANPostgreSQLPgStatementsAgentOKBody
*/
type ChangeQANPostgreSQLPgStatementsAgentOKBody struct {

	// qan postgresql pgstatements agent
	QANPostgresqlPgstatementsAgent *ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent `json:"qanPostgresqlPgstatementsAgent,omitempty"`
}

// Validate validates this change QAN postgre SQL pg statements agent OK body
func (o *ChangeQANPostgreSQLPgStatementsAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANPostgresqlPgstatementsAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatementsAgentOKBody) validateQANPostgresqlPgstatementsAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANPostgresqlPgstatementsAgent) { // not required
		return nil
	}

	if o.QANPostgresqlPgstatementsAgent != nil {
		if err := o.QANPostgresqlPgstatementsAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeQanPostgreSqlPgStatementsAgentOk" + "." + "qanPostgresqlPgstatementsAgent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatementsAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent QANPostgreSQLPgStatementsAgent runs within pmm-agent and sends PostgreSQL Query Analytics data to the PMM Server.
swagger:model ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent
*/
type ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agentId,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmmAgentId,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"serviceId,omitempty"`

	// PostgreSQL username for getting pg stat statements data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tlsSkipVerify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"customLabels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`
}

// Validate validates this change QAN postgre SQL pg statements agent OK body QAN postgresql pgstatements agent
func (o *ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeQanPostgreSqlPgStatementsAgentOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeQanPostgreSqlPgStatementsAgentOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum = append(changeQanPostgreSqlPgStatementsAgentOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum, v)
	}
}

const (

	// ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusSTARTING captures enum value "STARTING"
	ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusSTARTING string = "STARTING"

	// ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusRUNNING captures enum value "RUNNING"
	ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusRUNNING string = "RUNNING"

	// ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusWAITING captures enum value "WAITING"
	ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusWAITING string = "WAITING"

	// ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusSTOPPING captures enum value "STOPPING"
	ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusSTOPPING string = "STOPPING"

	// ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusDONE captures enum value "DONE"
	ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusDONE string = "DONE"

	// ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusUNKNOWN captures enum value "UNKNOWN"
	ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeQanPostgreSqlPgStatementsAgentOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeQanPostgreSqlPgStatementsAgentOk"+"."+"qanPostgresqlPgstatementsAgent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeQANPostgreSQLPgStatementsAgentParamsBodyCommon ChangeCommonAgentParams contains parameters that can be changed for all Agents.
swagger:model ChangeQANPostgreSQLPgStatementsAgentParamsBodyCommon
*/
type ChangeQANPostgreSQLPgStatementsAgentParamsBodyCommon struct {

	// Enable this Agent. Can't be used with disabled.
	Enable bool `json:"enable,omitempty"`

	// Disable this Agent. Can't be used with enabled.
	Disable bool `json:"disable,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"customLabels,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"removeCustomLabels,omitempty"`

	// Enables push metrics with vmagent, can't be used with disable_push_metrics.
	// Can't be used with agent version lower then 2.12 and unsupported agents.
	EnablePushMetrics bool `json:"enablePushMetrics,omitempty"`

	// Disables push metrics, pmm-server starts to pull it, can't be used with enable_push_metrics.
	DisablePushMetrics bool `json:"disablePushMetrics,omitempty"`
}

// Validate validates this change QAN postgre SQL pg statements agent params body common
func (o *ChangeQANPostgreSQLPgStatementsAgentParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeQANPostgreSQLPgStatementsAgentParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res ChangeQANPostgreSQLPgStatementsAgentParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
