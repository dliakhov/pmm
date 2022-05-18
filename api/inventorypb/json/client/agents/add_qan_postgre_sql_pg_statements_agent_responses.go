// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
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

// AddQANPostgreSQLPgStatementsAgentReader is a Reader for the AddQANPostgreSQLPgStatementsAgent structure.
type AddQANPostgreSQLPgStatementsAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddQANPostgreSQLPgStatementsAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddQANPostgreSQLPgStatementsAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddQANPostgreSQLPgStatementsAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddQANPostgreSQLPgStatementsAgentOK creates a AddQANPostgreSQLPgStatementsAgentOK with default headers values
func NewAddQANPostgreSQLPgStatementsAgentOK() *AddQANPostgreSQLPgStatementsAgentOK {
	return &AddQANPostgreSQLPgStatementsAgentOK{}
}

/* AddQANPostgreSQLPgStatementsAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddQANPostgreSQLPgStatementsAgentOK struct {
	Payload *AddQANPostgreSQLPgStatementsAgentOKBody
}

func (o *AddQANPostgreSQLPgStatementsAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANPostgreSQLPgStatementsAgent][%d] addQanPostgreSqlPgStatementsAgentOk  %+v", 200, o.Payload)
}
func (o *AddQANPostgreSQLPgStatementsAgentOK) GetPayload() *AddQANPostgreSQLPgStatementsAgentOKBody {
	return o.Payload
}

func (o *AddQANPostgreSQLPgStatementsAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANPostgreSQLPgStatementsAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddQANPostgreSQLPgStatementsAgentDefault creates a AddQANPostgreSQLPgStatementsAgentDefault with default headers values
func NewAddQANPostgreSQLPgStatementsAgentDefault(code int) *AddQANPostgreSQLPgStatementsAgentDefault {
	return &AddQANPostgreSQLPgStatementsAgentDefault{
		_statusCode: code,
	}
}

/* AddQANPostgreSQLPgStatementsAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddQANPostgreSQLPgStatementsAgentDefault struct {
	_statusCode int

	Payload *AddQANPostgreSQLPgStatementsAgentDefaultBody
}

// Code gets the status code for the add QAN postgre SQL pg statements agent default response
func (o *AddQANPostgreSQLPgStatementsAgentDefault) Code() int {
	return o._statusCode
}

func (o *AddQANPostgreSQLPgStatementsAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANPostgreSQLPgStatementsAgent][%d] AddQANPostgreSQLPgStatementsAgent default  %+v", o._statusCode, o.Payload)
}
func (o *AddQANPostgreSQLPgStatementsAgentDefault) GetPayload() *AddQANPostgreSQLPgStatementsAgentDefaultBody {
	return o.Payload
}

func (o *AddQANPostgreSQLPgStatementsAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANPostgreSQLPgStatementsAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddQANPostgreSQLPgStatementsAgentBody add QAN postgre SQL pg statements agent body
swagger:model AddQANPostgreSQLPgStatementsAgentBody
*/
type AddQANPostgreSQLPgStatementsAgentBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// PostgreSQL username for getting pg stat statements data.
	Username string `json:"username,omitempty"`

	// PostgreSQL password for getting pg stat statements data.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// TLS CA certificate.
	TLSCa string `json:"tls_ca,omitempty"`

	// TLS Certifcate.
	TLSCert string `json:"tls_cert,omitempty"`

	// TLS Certificate Key.
	TLSKey string `json:"tls_key,omitempty"`
}

// Validate validates this add QAN postgre SQL pg statements agent body
func (o *AddQANPostgreSQLPgStatementsAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add QAN postgre SQL pg statements agent body based on context it is used
func (o *AddQANPostgreSQLPgStatementsAgentBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentBody) UnmarshalBinary(b []byte) error {
	var res AddQANPostgreSQLPgStatementsAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANPostgreSQLPgStatementsAgentDefaultBody add QAN postgre SQL pg statements agent default body
swagger:model AddQANPostgreSQLPgStatementsAgentDefaultBody
*/
type AddQANPostgreSQLPgStatementsAgentDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddQANPostgreSQLPgStatementsAgentDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add QAN postgre SQL pg statements agent default body
func (o *AddQANPostgreSQLPgStatementsAgentDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANPostgreSQLPgStatementsAgentDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddQANPostgreSQLPgStatementsAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddQANPostgreSQLPgStatementsAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add QAN postgre SQL pg statements agent default body based on the context it is used
func (o *AddQANPostgreSQLPgStatementsAgentDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANPostgreSQLPgStatementsAgentDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddQANPostgreSQLPgStatementsAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddQANPostgreSQLPgStatementsAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddQANPostgreSQLPgStatementsAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANPostgreSQLPgStatementsAgentDefaultBodyDetailsItems0 add QAN postgre SQL pg statements agent default body details items0
swagger:model AddQANPostgreSQLPgStatementsAgentDefaultBodyDetailsItems0
*/
type AddQANPostgreSQLPgStatementsAgentDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this add QAN postgre SQL pg statements agent default body details items0
func (o *AddQANPostgreSQLPgStatementsAgentDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add QAN postgre SQL pg statements agent default body details items0 based on context it is used
func (o *AddQANPostgreSQLPgStatementsAgentDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddQANPostgreSQLPgStatementsAgentDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANPostgreSQLPgStatementsAgentOKBody add QAN postgre SQL pg statements agent OK body
swagger:model AddQANPostgreSQLPgStatementsAgentOKBody
*/
type AddQANPostgreSQLPgStatementsAgentOKBody struct {

	// qan postgresql pgstatements agent
	QANPostgresqlPgstatementsAgent *AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent `json:"qan_postgresql_pgstatements_agent,omitempty"`
}

// Validate validates this add QAN postgre SQL pg statements agent OK body
func (o *AddQANPostgreSQLPgStatementsAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANPostgresqlPgstatementsAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANPostgreSQLPgStatementsAgentOKBody) validateQANPostgresqlPgstatementsAgent(formats strfmt.Registry) error {
	if swag.IsZero(o.QANPostgresqlPgstatementsAgent) { // not required
		return nil
	}

	if o.QANPostgresqlPgstatementsAgent != nil {
		if err := o.QANPostgresqlPgstatementsAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addQanPostgreSqlPgStatementsAgentOk" + "." + "qan_postgresql_pgstatements_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addQanPostgreSqlPgStatementsAgentOk" + "." + "qan_postgresql_pgstatements_agent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add QAN postgre SQL pg statements agent OK body based on the context it is used
func (o *AddQANPostgreSQLPgStatementsAgentOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateQANPostgresqlPgstatementsAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANPostgreSQLPgStatementsAgentOKBody) contextValidateQANPostgresqlPgstatementsAgent(ctx context.Context, formats strfmt.Registry) error {

	if o.QANPostgresqlPgstatementsAgent != nil {
		if err := o.QANPostgresqlPgstatementsAgent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addQanPostgreSqlPgStatementsAgentOk" + "." + "qan_postgresql_pgstatements_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addQanPostgreSqlPgStatementsAgentOk" + "." + "qan_postgresql_pgstatements_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentOKBody) UnmarshalBinary(b []byte) error {
	var res AddQANPostgreSQLPgStatementsAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent QANPostgreSQLPgStatementsAgent runs within pmm-agent and sends PostgreSQL Query Analytics data to the PMM Server.
swagger:model AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent
*/
type AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// PostgreSQL username for getting pg stat statements data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

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

	// Path to exec process.
	ProcessExecPath string `json:"process_exec_path,omitempty"`
}

// Validate validates this add QAN postgre SQL pg statements agent OK body QAN postgresql pgstatements agent
func (o *AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addQanPostgreSqlPgStatementsAgentOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addQanPostgreSqlPgStatementsAgentOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum = append(addQanPostgreSqlPgStatementsAgentOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum, v)
	}
}

const (

	// AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusSTARTING captures enum value "STARTING"
	AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusSTARTING string = "STARTING"

	// AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusRUNNING captures enum value "RUNNING"
	AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusRUNNING string = "RUNNING"

	// AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusWAITING captures enum value "WAITING"
	AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusWAITING string = "WAITING"

	// AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusSTOPPING captures enum value "STOPPING"
	AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusSTOPPING string = "STOPPING"

	// AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusDONE captures enum value "DONE"
	AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusDONE string = "DONE"

	// AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusUNKNOWN captures enum value "UNKNOWN"
	AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgentStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addQanPostgreSqlPgStatementsAgentOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addQanPostgreSqlPgStatementsAgentOk"+"."+"qan_postgresql_pgstatements_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add QAN postgre SQL pg statements agent OK body QAN postgresql pgstatements agent based on context it is used
func (o *AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent) UnmarshalBinary(b []byte) error {
	var res AddQANPostgreSQLPgStatementsAgentOKBodyQANPostgresqlPgstatementsAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
