// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/dbaas/psmdb_clusters.proto

package dbaasv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *PSMDBClusterParams) Validate() error {
	if this.Replicaset != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Replicaset); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Replicaset", err)
		}
	}
	return nil
}

func (this *PSMDBClusterParams_ReplicaSet) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}

func (this *GetPSMDBClusterCredentialsRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}

func (this *GetPSMDBClusterCredentialsResponse) Validate() error {
	if this.ConnectionCredentials != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionCredentials); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionCredentials", err)
		}
	}
	return nil
}

func (this *GetPSMDBClusterCredentialsResponse_PSMDBCredentials) Validate() error {
	return nil
}
func (this *CreatePSMDBClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}

func (this *CreatePSMDBClusterResponse) Validate() error {
	return nil
}

func (this *UpdatePSMDBClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}

func (this *UpdatePSMDBClusterRequest_UpdatePSMDBClusterParams) Validate() error {
	if this.Replicaset != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Replicaset); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Replicaset", err)
		}
	}
	return nil
}

func (this *UpdatePSMDBClusterRequest_UpdatePSMDBClusterParams_ReplicaSet) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}

func (this *UpdatePSMDBClusterResponse) Validate() error {
	return nil
}

func (this *GetPSMDBClusterResourcesRequest) Validate() error {
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}

func (this *GetPSMDBClusterResourcesResponse) Validate() error {
	if this.Expected != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Expected); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Expected", err)
		}
	}
	return nil
}
