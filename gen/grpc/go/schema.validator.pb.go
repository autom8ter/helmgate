// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.proto

package hpaaspb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_Dependency_Chart = regexp.MustCompile(`^.{1,225}$`)
var _regex_Dependency_Version = regexp.MustCompile(`^.{1,225}$`)
var _regex_Dependency_Repository = regexp.MustCompile(`^.{1,225}$`)

func (this *Dependency) Validate() error {
	if !_regex_Dependency_Chart.MatchString(this.Chart) {
		return github_com_mwitkow_go_proto_validators.FieldError("Chart", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Chart))
	}
	if !_regex_Dependency_Version.MatchString(this.Version) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Version))
	}
	if !_regex_Dependency_Repository.MatchString(this.Repository) {
		return github_com_mwitkow_go_proto_validators.FieldError("Repository", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Repository))
	}
	return nil
}

var _regex_Maintainer_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_Maintainer_Email = regexp.MustCompile(`^.{1,225}$`)

func (this *Maintainer) Validate() error {
	if !_regex_Maintainer_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_Maintainer_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Email))
	}
	return nil
}

var _regex_ChartFilter_Term = regexp.MustCompile(`^.{1,225}$`)

func (this *ChartFilter) Validate() error {
	if !_regex_ChartFilter_Term.MatchString(this.Term) {
		return github_com_mwitkow_go_proto_validators.FieldError("Term", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Term))
	}
	return nil
}

var _regex_Chart_Name = regexp.MustCompile(`^.{1,225}$`)

func (this *Chart) Validate() error {
	if !_regex_Chart_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	for _, item := range this.Dependencies {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Dependencies", err)
			}
		}
	}
	for _, item := range this.Maintainers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Maintainers", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Charts) Validate() error {
	for _, item := range this.Charts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Charts", err)
			}
		}
	}
	return nil
}

var _regex_App_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_App_Namespace = regexp.MustCompile(`^.{1,225}$`)

func (this *App) Validate() error {
	if !_regex_App_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_App_Namespace.MatchString(this.Namespace) {
		return github_com_mwitkow_go_proto_validators.FieldError("Namespace", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Namespace))
	}
	if this.Release != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Release); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Release", err)
		}
	}
	if this.Chart != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Chart); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Chart", err)
		}
	}
	return nil
}
func (this *Apps) Validate() error {
	for _, item := range this.Apps {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Apps", err)
			}
		}
	}
	return nil
}
func (this *AppFilter) Validate() error {
	return nil
}
func (this *Release) Validate() error {
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	if this.Timestamps != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamps); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamps", err)
		}
	}
	return nil
}
func (this *Timestamps) Validate() error {
	if this.Created != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Created); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Created", err)
		}
	}
	if this.Updated != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Updated); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Updated", err)
		}
	}
	if this.Deleted != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deleted); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deleted", err)
		}
	}
	return nil
}

var _regex_AppRef_Namespace = regexp.MustCompile(`^.{1,225}$`)
var _regex_AppRef_Name = regexp.MustCompile(`^.{1,225}$`)

func (this *AppRef) Validate() error {
	if !_regex_AppRef_Namespace.MatchString(this.Namespace) {
		return github_com_mwitkow_go_proto_validators.FieldError("Namespace", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Namespace))
	}
	if !_regex_AppRef_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	return nil
}

var _regex_AppInput_Namespace = regexp.MustCompile(`^.{1,225}$`)
var _regex_AppInput_Chart = regexp.MustCompile(`^.{1,225}$`)
var _regex_AppInput_AppName = regexp.MustCompile(`^.{1,225}$`)

func (this *AppInput) Validate() error {
	if !_regex_AppInput_Namespace.MatchString(this.Namespace) {
		return github_com_mwitkow_go_proto_validators.FieldError("Namespace", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Namespace))
	}
	if !_regex_AppInput_Chart.MatchString(this.Chart) {
		return github_com_mwitkow_go_proto_validators.FieldError("Chart", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Chart))
	}
	if !_regex_AppInput_AppName.MatchString(this.AppName) {
		return github_com_mwitkow_go_proto_validators.FieldError("AppName", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.AppName))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

var _regex_NamespaceRef_Name = regexp.MustCompile(`^.{1,225}$`)

func (this *NamespaceRef) Validate() error {
	if !_regex_NamespaceRef_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	return nil
}
func (this *NamespaceRefs) Validate() error {
	for _, item := range this.Namespaces {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Namespaces", err)
			}
		}
	}
	return nil
}
func (this *HistoryFilter) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
		}
	}
	return nil
}
