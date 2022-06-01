// Code generated by go-swagger; DO NOT EDIT.

package services

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

// ListServicesReader is a Reader for the ListServices structure.
type ListServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListServicesOK creates a ListServicesOK with default headers values
func NewListServicesOK() *ListServicesOK {
	return &ListServicesOK{}
}

/* ListServicesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListServicesOK struct {
	Payload *ListServicesOKBody
}

func (o *ListServicesOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/List][%d] listServicesOk  %+v", 200, o.Payload)
}

func (o *ListServicesOK) GetPayload() *ListServicesOKBody {
	return o.Payload
}

func (o *ListServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesDefault creates a ListServicesDefault with default headers values
func NewListServicesDefault(code int) *ListServicesDefault {
	return &ListServicesDefault{
		_statusCode: code,
	}
}

/* ListServicesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListServicesDefault struct {
	_statusCode int

	Payload *ListServicesDefaultBody
}

// Code gets the status code for the list services default response
func (o *ListServicesDefault) Code() int {
	return o._statusCode
}

func (o *ListServicesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/List][%d] ListServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListServicesDefault) GetPayload() *ListServicesDefaultBody {
	return o.Payload
}

func (o *ListServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListServicesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListServicesBody list services body
swagger:model ListServicesBody
*/
type ListServicesBody struct {
	// Return only Services running on that Node.
	NodeID string `json:"node_id,omitempty"`

	// ServiceType describes supported Service types.
	// Enum: [SERVICE_TYPE_INVALID MYSQL_SERVICE MONGODB_SERVICE POSTGRESQL_SERVICE PROXYSQL_SERVICE HAPROXY_SERVICE EXTERNAL_SERVICE]
	ServiceType *string `json:"service_type,omitempty"`

	// Return only services in this external group.
	ExternalGroup string `json:"external_group,omitempty"`
}

// Validate validates this list services body
func (o *ListServicesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listServicesBodyTypeServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SERVICE_TYPE_INVALID","MYSQL_SERVICE","MONGODB_SERVICE","POSTGRESQL_SERVICE","PROXYSQL_SERVICE","HAPROXY_SERVICE","EXTERNAL_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listServicesBodyTypeServiceTypePropEnum = append(listServicesBodyTypeServiceTypePropEnum, v)
	}
}

const (

	// ListServicesBodyServiceTypeSERVICETYPEINVALID captures enum value "SERVICE_TYPE_INVALID"
	ListServicesBodyServiceTypeSERVICETYPEINVALID string = "SERVICE_TYPE_INVALID"

	// ListServicesBodyServiceTypeMYSQLSERVICE captures enum value "MYSQL_SERVICE"
	ListServicesBodyServiceTypeMYSQLSERVICE string = "MYSQL_SERVICE"

	// ListServicesBodyServiceTypeMONGODBSERVICE captures enum value "MONGODB_SERVICE"
	ListServicesBodyServiceTypeMONGODBSERVICE string = "MONGODB_SERVICE"

	// ListServicesBodyServiceTypePOSTGRESQLSERVICE captures enum value "POSTGRESQL_SERVICE"
	ListServicesBodyServiceTypePOSTGRESQLSERVICE string = "POSTGRESQL_SERVICE"

	// ListServicesBodyServiceTypePROXYSQLSERVICE captures enum value "PROXYSQL_SERVICE"
	ListServicesBodyServiceTypePROXYSQLSERVICE string = "PROXYSQL_SERVICE"

	// ListServicesBodyServiceTypeHAPROXYSERVICE captures enum value "HAPROXY_SERVICE"
	ListServicesBodyServiceTypeHAPROXYSERVICE string = "HAPROXY_SERVICE"

	// ListServicesBodyServiceTypeEXTERNALSERVICE captures enum value "EXTERNAL_SERVICE"
	ListServicesBodyServiceTypeEXTERNALSERVICE string = "EXTERNAL_SERVICE"
)

// prop value enum
func (o *ListServicesBody) validateServiceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listServicesBodyTypeServiceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListServicesBody) validateServiceType(formats strfmt.Registry) error {
	if swag.IsZero(o.ServiceType) { // not required
		return nil
	}

	// value enum
	if err := o.validateServiceTypeEnum("body"+"."+"service_type", "body", *o.ServiceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list services body based on context it is used
func (o *ListServicesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesBody) UnmarshalBinary(b []byte) error {
	var res ListServicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesDefaultBody list services default body
swagger:model ListServicesDefaultBody
*/
type ListServicesDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListServicesDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list services default body
func (o *ListServicesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListServices default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list services default body based on the context it is used
func (o *ListServicesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListServices default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListServices default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListServicesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesDefaultBodyDetailsItems0 list services default body details items0
swagger:model ListServicesDefaultBodyDetailsItems0
*/
type ListServicesDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this list services default body details items0
func (o *ListServicesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services default body details items0 based on context it is used
func (o *ListServicesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesOKBody list services OK body
swagger:model ListServicesOKBody
*/
type ListServicesOKBody struct {
	// mysql
	Mysql []*ListServicesOKBodyMysqlItems0 `json:"mysql"`

	// mongodb
	Mongodb []*ListServicesOKBodyMongodbItems0 `json:"mongodb"`

	// postgresql
	Postgresql []*ListServicesOKBodyPostgresqlItems0 `json:"postgresql"`

	// proxysql
	Proxysql []*ListServicesOKBodyProxysqlItems0 `json:"proxysql"`

	// haproxy
	Haproxy []*ListServicesOKBodyHaproxyItems0 `json:"haproxy"`

	// external
	External []*ListServicesOKBodyExternalItems0 `json:"external"`
}

// Validate validates this list services OK body
func (o *ListServicesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongodb(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePostgresql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHaproxy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateExternal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesOKBody) validateMysql(formats strfmt.Registry) error {
	if swag.IsZero(o.Mysql) { // not required
		return nil
	}

	for i := 0; i < len(o.Mysql); i++ {
		if swag.IsZero(o.Mysql[i]) { // not required
			continue
		}

		if o.Mysql[i] != nil {
			if err := o.Mysql[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListServicesOKBody) validateMongodb(formats strfmt.Registry) error {
	if swag.IsZero(o.Mongodb) { // not required
		return nil
	}

	for i := 0; i < len(o.Mongodb); i++ {
		if swag.IsZero(o.Mongodb[i]) { // not required
			continue
		}

		if o.Mongodb[i] != nil {
			if err := o.Mongodb[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListServicesOKBody) validatePostgresql(formats strfmt.Registry) error {
	if swag.IsZero(o.Postgresql) { // not required
		return nil
	}

	for i := 0; i < len(o.Postgresql); i++ {
		if swag.IsZero(o.Postgresql[i]) { // not required
			continue
		}

		if o.Postgresql[i] != nil {
			if err := o.Postgresql[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "postgresql" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "postgresql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListServicesOKBody) validateProxysql(formats strfmt.Registry) error {
	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	for i := 0; i < len(o.Proxysql); i++ {
		if swag.IsZero(o.Proxysql[i]) { // not required
			continue
		}

		if o.Proxysql[i] != nil {
			if err := o.Proxysql[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "proxysql" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "proxysql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListServicesOKBody) validateHaproxy(formats strfmt.Registry) error {
	if swag.IsZero(o.Haproxy) { // not required
		return nil
	}

	for i := 0; i < len(o.Haproxy); i++ {
		if swag.IsZero(o.Haproxy[i]) { // not required
			continue
		}

		if o.Haproxy[i] != nil {
			if err := o.Haproxy[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "haproxy" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "haproxy" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListServicesOKBody) validateExternal(formats strfmt.Registry) error {
	if swag.IsZero(o.External) { // not required
		return nil
	}

	for i := 0; i < len(o.External); i++ {
		if swag.IsZero(o.External[i]) { // not required
			continue
		}

		if o.External[i] != nil {
			if err := o.External[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "external" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "external" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list services OK body based on the context it is used
func (o *ListServicesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMysql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMongodb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePostgresql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateProxysql(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateHaproxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateExternal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesOKBody) contextValidateMysql(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Mysql); i++ {
		if o.Mysql[i] != nil {
			if err := o.Mysql[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *ListServicesOKBody) contextValidateMongodb(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Mongodb); i++ {
		if o.Mongodb[i] != nil {
			if err := o.Mongodb[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *ListServicesOKBody) contextValidatePostgresql(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Postgresql); i++ {
		if o.Postgresql[i] != nil {
			if err := o.Postgresql[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "postgresql" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "postgresql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *ListServicesOKBody) contextValidateProxysql(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Proxysql); i++ {
		if o.Proxysql[i] != nil {
			if err := o.Proxysql[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "proxysql" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "proxysql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *ListServicesOKBody) contextValidateHaproxy(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Haproxy); i++ {
		if o.Haproxy[i] != nil {
			if err := o.Haproxy[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "haproxy" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "haproxy" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *ListServicesOKBody) contextValidateExternal(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.External); i++ {
		if o.External[i] != nil {
			if err := o.External[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "external" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listServicesOk" + "." + "external" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBody) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesOKBodyExternalItems0 ExternalService represents a generic External service instance.
swagger:model ListServicesOKBodyExternalItems0
*/
type ListServicesOKBodyExternalItems0 struct {
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this service instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Group name of external service.
	Group string `json:"group,omitempty"`
}

// Validate validates this list services OK body external items0
func (o *ListServicesOKBodyExternalItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body external items0 based on context it is used
func (o *ListServicesOKBodyExternalItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyExternalItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyExternalItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyExternalItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesOKBodyHaproxyItems0 HAProxyService represents a generic HAProxy service instance.
swagger:model ListServicesOKBodyHaproxyItems0
*/
type ListServicesOKBodyHaproxyItems0 struct {
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this service instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this list services OK body haproxy items0
func (o *ListServicesOKBodyHaproxyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body haproxy items0 based on context it is used
func (o *ListServicesOKBodyHaproxyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyHaproxyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyHaproxyItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyHaproxyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesOKBodyMongodbItems0 MongoDBService represents a generic MongoDB instance.
swagger:model ListServicesOKBodyMongodbItems0
*/
type ListServicesOKBodyMongodbItems0 struct {
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this list services OK body mongodb items0
func (o *ListServicesOKBodyMongodbItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body mongodb items0 based on context it is used
func (o *ListServicesOKBodyMongodbItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyMongodbItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyMongodbItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyMongodbItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesOKBodyMysqlItems0 MySQLService represents a generic MySQL instance.
swagger:model ListServicesOKBodyMysqlItems0
*/
type ListServicesOKBodyMysqlItems0 struct {
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this list services OK body mysql items0
func (o *ListServicesOKBodyMysqlItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body mysql items0 based on context it is used
func (o *ListServicesOKBodyMysqlItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyMysqlItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyMysqlItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyMysqlItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesOKBodyPostgresqlItems0 PostgreSQLService represents a generic PostgreSQL instance.
swagger:model ListServicesOKBodyPostgresqlItems0
*/
type ListServicesOKBodyPostgresqlItems0 struct {
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Database name.
	DatabaseName string `json:"database_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this list services OK body postgresql items0
func (o *ListServicesOKBodyPostgresqlItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body postgresql items0 based on context it is used
func (o *ListServicesOKBodyPostgresqlItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyPostgresqlItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyPostgresqlItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyPostgresqlItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesOKBodyProxysqlItems0 ProxySQLService represents a generic ProxySQL instance.
swagger:model ListServicesOKBodyProxysqlItems0
*/
type ListServicesOKBodyProxysqlItems0 struct {
	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this list services OK body proxysql items0
func (o *ListServicesOKBodyProxysqlItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list services OK body proxysql items0 based on context it is used
func (o *ListServicesOKBodyProxysqlItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBodyProxysqlItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBodyProxysqlItems0) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBodyProxysqlItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
