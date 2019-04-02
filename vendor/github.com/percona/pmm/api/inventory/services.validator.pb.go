// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inventory/services.proto

package inventorypb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *MySQLService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AmazonRDSMySQLService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *MongoDBService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *PostgreSQLService) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ListServicesRequest) Validate() error {
	return nil
}
func (this *ListServicesResponse) Validate() error {
	for _, item := range this.Mysql {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mysql", err)
			}
		}
	}
	for _, item := range this.AmazonRdsMysql {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AmazonRdsMysql", err)
			}
		}
	}
	for _, item := range this.Mongodb {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mongodb", err)
			}
		}
	}
	for _, item := range this.Postgresql {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Postgresql", err)
			}
		}
	}
	return nil
}
func (this *GetServiceRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	return nil
}
func (this *GetServiceResponse) Validate() error {
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_Mysql); ok {
		if oneOfNester.Mysql != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Mysql); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mysql", err)
			}
		}
	}
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_AmazonRdsMysql); ok {
		if oneOfNester.AmazonRdsMysql != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.AmazonRdsMysql); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AmazonRdsMysql", err)
			}
		}
	}
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_Mongodb); ok {
		if oneOfNester.Mongodb != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Mongodb); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Mongodb", err)
			}
		}
	}
	if oneOfNester, ok := this.GetService().(*GetServiceResponse_Postgresql); ok {
		if oneOfNester.Postgresql != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Postgresql); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Postgresql", err)
			}
		}
	}
	return nil
}
func (this *AddMySQLServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	if !(this.Port > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '0'`, this.Port))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddMySQLServiceResponse) Validate() error {
	if this.Mysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Mysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Mysql", err)
		}
	}
	return nil
}
func (this *AddAmazonRDSMySQLServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	if !(this.Port > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '0'`, this.Port))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddAmazonRDSMySQLServiceResponse) Validate() error {
	if this.AmazonRdsMysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AmazonRdsMysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AmazonRdsMysql", err)
		}
	}
	return nil
}
func (this *AddMongoDBServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	if !(this.Port > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '0'`, this.Port))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddMongoDBServiceResponse) Validate() error {
	if this.Mongodb != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Mongodb); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Mongodb", err)
		}
	}
	return nil
}
func (this *AddPostgreSQLServiceRequest) Validate() error {
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.NodeId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf(`value '%v' must not be an empty string`, this.NodeId))
	}
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	if !(this.Port > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '0'`, this.Port))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddPostgreSQLServiceResponse) Validate() error {
	if this.Postgresql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Postgresql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Postgresql", err)
		}
	}
	return nil
}
func (this *RemoveServiceRequest) Validate() error {
	if this.ServiceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceId", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceId))
	}
	return nil
}
func (this *RemoveServiceResponse) Validate() error {
	return nil
}
