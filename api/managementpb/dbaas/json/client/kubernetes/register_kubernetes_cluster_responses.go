// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// RegisterKubernetesClusterReader is a Reader for the RegisterKubernetesCluster structure.
type RegisterKubernetesClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterKubernetesClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterKubernetesClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRegisterKubernetesClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRegisterKubernetesClusterOK creates a RegisterKubernetesClusterOK with default headers values
func NewRegisterKubernetesClusterOK() *RegisterKubernetesClusterOK {
	return &RegisterKubernetesClusterOK{}
}

/* RegisterKubernetesClusterOK describes a response with status code 200, with default header values.

A successful response.
*/
type RegisterKubernetesClusterOK struct {
	Payload interface{}
}

func (o *RegisterKubernetesClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/Register][%d] registerKubernetesClusterOk  %+v", 200, o.Payload)
}

func (o *RegisterKubernetesClusterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RegisterKubernetesClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterKubernetesClusterDefault creates a RegisterKubernetesClusterDefault with default headers values
func NewRegisterKubernetesClusterDefault(code int) *RegisterKubernetesClusterDefault {
	return &RegisterKubernetesClusterDefault{
		_statusCode: code,
	}
}

/* RegisterKubernetesClusterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RegisterKubernetesClusterDefault struct {
	_statusCode int

	Payload *RegisterKubernetesClusterDefaultBody
}

// Code gets the status code for the register kubernetes cluster default response
func (o *RegisterKubernetesClusterDefault) Code() int {
	return o._statusCode
}

func (o *RegisterKubernetesClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/Register][%d] RegisterKubernetesCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RegisterKubernetesClusterDefault) GetPayload() *RegisterKubernetesClusterDefaultBody {
	return o.Payload
}

func (o *RegisterKubernetesClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RegisterKubernetesClusterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RegisterKubernetesClusterBody register kubernetes cluster body
swagger:model RegisterKubernetesClusterBody
*/
type RegisterKubernetesClusterBody struct {
	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`

	// AWS access key id, only needed when registering EKS cluster and kubeconfig does not contain it.
	AWSAccessKeyID string `json:"aws_access_key_id,omitempty"`

	// AWS secret access key, only needed when registering EKS cluster and kubeconfig does not contain it.
	AWSSecretAccessKey string `json:"aws_secret_access_key,omitempty"`

	// kube auth
	KubeAuth *RegisterKubernetesClusterParamsBodyKubeAuth `json:"kube_auth,omitempty"`
}

// Validate validates this register kubernetes cluster body
func (o *RegisterKubernetesClusterBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKubeAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterKubernetesClusterBody) validateKubeAuth(formats strfmt.Registry) error {
	if swag.IsZero(o.KubeAuth) { // not required
		return nil
	}

	if o.KubeAuth != nil {
		if err := o.KubeAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "kube_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "kube_auth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this register kubernetes cluster body based on the context it is used
func (o *RegisterKubernetesClusterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateKubeAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterKubernetesClusterBody) contextValidateKubeAuth(ctx context.Context, formats strfmt.Registry) error {
	if o.KubeAuth != nil {
		if err := o.KubeAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "kube_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "kube_auth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RegisterKubernetesClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterKubernetesClusterBody) UnmarshalBinary(b []byte) error {
	var res RegisterKubernetesClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RegisterKubernetesClusterDefaultBody register kubernetes cluster default body
swagger:model RegisterKubernetesClusterDefaultBody
*/
type RegisterKubernetesClusterDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RegisterKubernetesClusterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this register kubernetes cluster default body
func (o *RegisterKubernetesClusterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterKubernetesClusterDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("RegisterKubernetesCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RegisterKubernetesCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this register kubernetes cluster default body based on the context it is used
func (o *RegisterKubernetesClusterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterKubernetesClusterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RegisterKubernetesCluster default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RegisterKubernetesCluster default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RegisterKubernetesClusterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterKubernetesClusterDefaultBody) UnmarshalBinary(b []byte) error {
	var res RegisterKubernetesClusterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RegisterKubernetesClusterDefaultBodyDetailsItems0 register kubernetes cluster default body details items0
swagger:model RegisterKubernetesClusterDefaultBodyDetailsItems0
*/
type RegisterKubernetesClusterDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this register kubernetes cluster default body details items0
func (o *RegisterKubernetesClusterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this register kubernetes cluster default body details items0 based on context it is used
func (o *RegisterKubernetesClusterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RegisterKubernetesClusterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterKubernetesClusterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RegisterKubernetesClusterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RegisterKubernetesClusterParamsBodyKubeAuth KubeAuth represents Kubernetes / kubectl authentication and authorization information.
swagger:model RegisterKubernetesClusterParamsBodyKubeAuth
*/
type RegisterKubernetesClusterParamsBodyKubeAuth struct {
	// Kubeconfig file content.
	Kubeconfig string `json:"kubeconfig,omitempty"`
}

// Validate validates this register kubernetes cluster params body kube auth
func (o *RegisterKubernetesClusterParamsBodyKubeAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this register kubernetes cluster params body kube auth based on context it is used
func (o *RegisterKubernetesClusterParamsBodyKubeAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RegisterKubernetesClusterParamsBodyKubeAuth) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterKubernetesClusterParamsBodyKubeAuth) UnmarshalBinary(b []byte) error {
	var res RegisterKubernetesClusterParamsBodyKubeAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
