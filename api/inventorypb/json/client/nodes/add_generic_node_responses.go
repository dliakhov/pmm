// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// AddGenericNodeReader is a Reader for the AddGenericNode structure.
type AddGenericNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddGenericNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddGenericNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddGenericNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddGenericNodeOK creates a AddGenericNodeOK with default headers values
func NewAddGenericNodeOK() *AddGenericNodeOK {
	return &AddGenericNodeOK{}
}

/* AddGenericNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddGenericNodeOK struct {
	Payload *AddGenericNodeOKBody
}

func (o *AddGenericNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/AddGeneric][%d] addGenericNodeOk  %+v", 200, o.Payload)
}

func (o *AddGenericNodeOK) GetPayload() *AddGenericNodeOKBody {
	return o.Payload
}

func (o *AddGenericNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddGenericNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGenericNodeDefault creates a AddGenericNodeDefault with default headers values
func NewAddGenericNodeDefault(code int) *AddGenericNodeDefault {
	return &AddGenericNodeDefault{
		_statusCode: code,
	}
}

/* AddGenericNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddGenericNodeDefault struct {
	_statusCode int

	Payload *AddGenericNodeDefaultBody
}

// Code gets the status code for the add generic node default response
func (o *AddGenericNodeDefault) Code() int {
	return o._statusCode
}

func (o *AddGenericNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/AddGeneric][%d] AddGenericNode default  %+v", o._statusCode, o.Payload)
}

func (o *AddGenericNodeDefault) GetPayload() *AddGenericNodeDefaultBody {
	return o.Payload
}

func (o *AddGenericNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(AddGenericNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddGenericNodeBody add generic node body
swagger:model AddGenericNodeBody
*/
type AddGenericNodeBody struct {
	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add generic node body
func (o *AddGenericNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add generic node body based on context it is used
func (o *AddGenericNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddGenericNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddGenericNodeBody) UnmarshalBinary(b []byte) error {
	var res AddGenericNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddGenericNodeDefaultBody add generic node default body
swagger:model AddGenericNodeDefaultBody
*/
type AddGenericNodeDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddGenericNodeDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add generic node default body
func (o *AddGenericNodeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddGenericNodeDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddGenericNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddGenericNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add generic node default body based on the context it is used
func (o *AddGenericNodeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddGenericNodeDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddGenericNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddGenericNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddGenericNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddGenericNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddGenericNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddGenericNodeDefaultBodyDetailsItems0 add generic node default body details items0
swagger:model AddGenericNodeDefaultBodyDetailsItems0
*/
type AddGenericNodeDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this add generic node default body details items0
func (o *AddGenericNodeDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add generic node default body details items0 based on context it is used
func (o *AddGenericNodeDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddGenericNodeDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddGenericNodeDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddGenericNodeDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddGenericNodeOKBody add generic node OK body
swagger:model AddGenericNodeOKBody
*/
type AddGenericNodeOKBody struct {
	// generic
	Generic *AddGenericNodeOKBodyGeneric `json:"generic,omitempty"`
}

// Validate validates this add generic node OK body
func (o *AddGenericNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddGenericNodeOKBody) validateGeneric(formats strfmt.Registry) error {
	if swag.IsZero(o.Generic) { // not required
		return nil
	}

	if o.Generic != nil {
		if err := o.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addGenericNodeOk" + "." + "generic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addGenericNodeOk" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add generic node OK body based on the context it is used
func (o *AddGenericNodeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGeneric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddGenericNodeOKBody) contextValidateGeneric(ctx context.Context, formats strfmt.Registry) error {
	if o.Generic != nil {
		if err := o.Generic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addGenericNodeOk" + "." + "generic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addGenericNodeOk" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddGenericNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddGenericNodeOKBody) UnmarshalBinary(b []byte) error {
	var res AddGenericNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddGenericNodeOKBodyGeneric GenericNode represents a bare metal server or virtual machine.
swagger:model AddGenericNodeOKBodyGeneric
*/
type AddGenericNodeOKBodyGeneric struct {
	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add generic node OK body generic
func (o *AddGenericNodeOKBodyGeneric) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add generic node OK body generic based on context it is used
func (o *AddGenericNodeOKBodyGeneric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddGenericNodeOKBodyGeneric) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddGenericNodeOKBodyGeneric) UnmarshalBinary(b []byte) error {
	var res AddGenericNodeOKBodyGeneric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
