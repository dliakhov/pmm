// Code generated by go-swagger; DO NOT EDIT.

package my_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// AddReader is a Reader for the Add structure.
type AddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddOK creates a AddOK with default headers values
func NewAddOK() *AddOK {
	return &AddOK{}
}

/*AddOK handles this case with default header values.

A successful response.
*/
type AddOK struct {
	Payload *AddOKBody
}

func (o *AddOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/MySQL/Add][%d] addOk  %+v", 200, o.Payload)
}

func (o *AddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDefault creates a AddDefault with default headers values
func NewAddDefault(code int) *AddDefault {
	return &AddDefault{
		_statusCode: code,
	}
}

/*AddDefault handles this case with default header values.

An error response.
*/
type AddDefault struct {
	_statusCode int

	Payload *AddDefaultBody
}

// Code gets the status code for the add default response
func (o *AddDefault) Code() int {
	return o._statusCode
}

func (o *AddDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/MySQL/Add][%d] Add default  %+v", o._statusCode, o.Payload)
}

func (o *AddDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddBody add body
swagger:model AddBody
*/
type AddBody struct {

	// Service Access address (DNS name or IP). Required.
	Address string `json:"address,omitempty"`

	// If true, adds mysqld_exporter for provided service.
	MysqldExporter bool `json:"mysqld_exporter,omitempty"`

	// Node identifier on which a service is been running. Required.
	NodeID string `json:"node_id,omitempty"`

	// MySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The "pmm-agent" identifier which should run agents. Required.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service Access port. Required.
	Port int64 `json:"port,omitempty"`

	// If true, adds qan-mysql-perfschema-agent for provided service.
	QANMysqlPerfschema bool `json:"qan_mysql_perfschema,omitempty"`

	// FIXME remove
	QANPassword string `json:"qan_password,omitempty"`

	// FIXME remove
	QANUsername string `json:"qan_username,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this add body
func (o *AddBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddBody) UnmarshalBinary(b []byte) error {
	var res AddBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddDefaultBody
*/
type AddDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add default body
func (o *AddDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddOKBody add OK body
swagger:model AddOKBody
*/
type AddOKBody struct {

	// mysqld exporter
	MysqldExporter *AddOKBodyMysqldExporter `json:"mysqld_exporter,omitempty"`

	// qan mysql perfschema
	QANMysqlPerfschema *AddOKBodyQANMysqlPerfschema `json:"qan_mysql_perfschema,omitempty"`

	// service
	Service *AddOKBodyService `json:"service,omitempty"`
}

// Validate validates this add OK body
func (o *AddOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANMysqlPerfschema(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	if o.MysqldExporter != nil {
		if err := o.MysqldExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addOk" + "." + "mysqld_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddOKBody) validateQANMysqlPerfschema(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMysqlPerfschema) { // not required
		return nil
	}

	if o.QANMysqlPerfschema != nil {
		if err := o.QANMysqlPerfschema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addOk" + "." + "qan_mysql_perfschema")
			}
			return err
		}
	}

	return nil
}

func (o *AddOKBody) validateService(formats strfmt.Registry) error {

	if swag.IsZero(o.Service) { // not required
		return nil
	}

	if o.Service != nil {
		if err := o.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addOk" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOKBody) UnmarshalBinary(b []byte) error {
	var res AddOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddOKBodyMysqldExporter MySQLdExporter runs on Generic or Container Node and exposes MySQL and AmazonRDSMySQL Service metrics.
swagger:model AddOKBodyMysqldExporter
*/
type AddOKBodyMysqldExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// MySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this add OK body mysqld exporter
func (o *AddOKBodyMysqldExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addOkBodyMysqldExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addOkBodyMysqldExporterTypeStatusPropEnum = append(addOkBodyMysqldExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddOKBodyMysqldExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddOKBodyMysqldExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddOKBodyMysqldExporterStatusSTARTING captures enum value "STARTING"
	AddOKBodyMysqldExporterStatusSTARTING string = "STARTING"

	// AddOKBodyMysqldExporterStatusRUNNING captures enum value "RUNNING"
	AddOKBodyMysqldExporterStatusRUNNING string = "RUNNING"

	// AddOKBodyMysqldExporterStatusWAITING captures enum value "WAITING"
	AddOKBodyMysqldExporterStatusWAITING string = "WAITING"

	// AddOKBodyMysqldExporterStatusSTOPPING captures enum value "STOPPING"
	AddOKBodyMysqldExporterStatusSTOPPING string = "STOPPING"

	// AddOKBodyMysqldExporterStatusDONE captures enum value "DONE"
	AddOKBodyMysqldExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddOKBodyMysqldExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addOkBodyMysqldExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddOKBodyMysqldExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addOk"+"."+"mysqld_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddOKBodyMysqldExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOKBodyMysqldExporter) UnmarshalBinary(b []byte) error {
	var res AddOKBodyMysqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddOKBodyQANMysqlPerfschema QANMySQLPerfSchemaAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model AddOKBodyQANMysqlPerfschema
*/
type AddOKBodyQANMysqlPerfschema struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// MySQL password for getting performance data.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MySQL username for getting performance data.
	Username string `json:"username,omitempty"`
}

// Validate validates this add OK body QAN mysql perfschema
func (o *AddOKBodyQANMysqlPerfschema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addOkBodyQanMysqlPerfschemaTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addOkBodyQanMysqlPerfschemaTypeStatusPropEnum = append(addOkBodyQanMysqlPerfschemaTypeStatusPropEnum, v)
	}
}

const (

	// AddOKBodyQANMysqlPerfschemaStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddOKBodyQANMysqlPerfschemaStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddOKBodyQANMysqlPerfschemaStatusSTARTING captures enum value "STARTING"
	AddOKBodyQANMysqlPerfschemaStatusSTARTING string = "STARTING"

	// AddOKBodyQANMysqlPerfschemaStatusRUNNING captures enum value "RUNNING"
	AddOKBodyQANMysqlPerfschemaStatusRUNNING string = "RUNNING"

	// AddOKBodyQANMysqlPerfschemaStatusWAITING captures enum value "WAITING"
	AddOKBodyQANMysqlPerfschemaStatusWAITING string = "WAITING"

	// AddOKBodyQANMysqlPerfschemaStatusSTOPPING captures enum value "STOPPING"
	AddOKBodyQANMysqlPerfschemaStatusSTOPPING string = "STOPPING"

	// AddOKBodyQANMysqlPerfschemaStatusDONE captures enum value "DONE"
	AddOKBodyQANMysqlPerfschemaStatusDONE string = "DONE"
)

// prop value enum
func (o *AddOKBodyQANMysqlPerfschema) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addOkBodyQanMysqlPerfschemaTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddOKBodyQANMysqlPerfschema) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addOk"+"."+"qan_mysql_perfschema"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddOKBodyQANMysqlPerfschema) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOKBodyQANMysqlPerfschema) UnmarshalBinary(b []byte) error {
	var res AddOKBodyQANMysqlPerfschema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddOKBodyService MySQLService represents a generic MySQL instance.
swagger:model AddOKBodyService
*/
type AddOKBodyService struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this add OK body service
func (o *AddOKBodyService) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddOKBodyService) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOKBodyService) UnmarshalBinary(b []byte) error {
	var res AddOKBodyService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
