//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 Nokia.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/sdcio/config-server/apis/config"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Condition)(nil), (*config.Condition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Condition_To_config_Condition(a.(*Condition), b.(*config.Condition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Condition)(nil), (*Condition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Condition_To_v1alpha1_Condition(a.(*config.Condition), b.(*Condition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConditionedStatus)(nil), (*config.ConditionedStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConditionedStatus_To_config_ConditionedStatus(a.(*ConditionedStatus), b.(*config.ConditionedStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConditionedStatus)(nil), (*ConditionedStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConditionedStatus_To_v1alpha1_ConditionedStatus(a.(*config.ConditionedStatus), b.(*ConditionedStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Config)(nil), (*config.Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Config_To_config_Config(a.(*Config), b.(*config.Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Config)(nil), (*Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Config_To_v1alpha1_Config(a.(*config.Config), b.(*Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigBlob)(nil), (*config.ConfigBlob)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigBlob_To_config_ConfigBlob(a.(*ConfigBlob), b.(*config.ConfigBlob), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigBlob)(nil), (*ConfigBlob)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigBlob_To_v1alpha1_ConfigBlob(a.(*config.ConfigBlob), b.(*ConfigBlob), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigList)(nil), (*config.ConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigList_To_config_ConfigList(a.(*ConfigList), b.(*config.ConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigList)(nil), (*ConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigList_To_v1alpha1_ConfigList(a.(*config.ConfigList), b.(*ConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigSet)(nil), (*config.ConfigSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigSet_To_config_ConfigSet(a.(*ConfigSet), b.(*config.ConfigSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigSet)(nil), (*ConfigSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigSet_To_v1alpha1_ConfigSet(a.(*config.ConfigSet), b.(*ConfigSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigSetList)(nil), (*config.ConfigSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigSetList_To_config_ConfigSetList(a.(*ConfigSetList), b.(*config.ConfigSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigSetList)(nil), (*ConfigSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigSetList_To_v1alpha1_ConfigSetList(a.(*config.ConfigSetList), b.(*ConfigSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigSetSpec)(nil), (*config.ConfigSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigSetSpec_To_config_ConfigSetSpec(a.(*ConfigSetSpec), b.(*config.ConfigSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigSetSpec)(nil), (*ConfigSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigSetSpec_To_v1alpha1_ConfigSetSpec(a.(*config.ConfigSetSpec), b.(*ConfigSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigSetStatus)(nil), (*config.ConfigSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigSetStatus_To_config_ConfigSetStatus(a.(*ConfigSetStatus), b.(*config.ConfigSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigSetStatus)(nil), (*ConfigSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigSetStatus_To_v1alpha1_ConfigSetStatus(a.(*config.ConfigSetStatus), b.(*ConfigSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigSpec)(nil), (*config.ConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigSpec_To_config_ConfigSpec(a.(*ConfigSpec), b.(*config.ConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigSpec)(nil), (*ConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigSpec_To_v1alpha1_ConfigSpec(a.(*config.ConfigSpec), b.(*ConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigStatus)(nil), (*config.ConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigStatus_To_config_ConfigStatus(a.(*ConfigStatus), b.(*config.ConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigStatus)(nil), (*ConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigStatus_To_v1alpha1_ConfigStatus(a.(*config.ConfigStatus), b.(*ConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigStatusLastKnownGoodSchema)(nil), (*config.ConfigStatusLastKnownGoodSchema)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigStatusLastKnownGoodSchema_To_config_ConfigStatusLastKnownGoodSchema(a.(*ConfigStatusLastKnownGoodSchema), b.(*config.ConfigStatusLastKnownGoodSchema), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigStatusLastKnownGoodSchema)(nil), (*ConfigStatusLastKnownGoodSchema)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigStatusLastKnownGoodSchema_To_v1alpha1_ConfigStatusLastKnownGoodSchema(a.(*config.ConfigStatusLastKnownGoodSchema), b.(*ConfigStatusLastKnownGoodSchema), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Lifecycle)(nil), (*config.Lifecycle)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Lifecycle_To_config_Lifecycle(a.(*Lifecycle), b.(*config.Lifecycle), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Lifecycle)(nil), (*Lifecycle)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Lifecycle_To_v1alpha1_Lifecycle(a.(*config.Lifecycle), b.(*Lifecycle), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RunningConfig)(nil), (*config.RunningConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RunningConfig_To_config_RunningConfig(a.(*RunningConfig), b.(*config.RunningConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RunningConfig)(nil), (*RunningConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RunningConfig_To_v1alpha1_RunningConfig(a.(*config.RunningConfig), b.(*RunningConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RunningConfigList)(nil), (*config.RunningConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RunningConfigList_To_config_RunningConfigList(a.(*RunningConfigList), b.(*config.RunningConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RunningConfigList)(nil), (*RunningConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RunningConfigList_To_v1alpha1_RunningConfigList(a.(*config.RunningConfigList), b.(*RunningConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RunningConfigSpec)(nil), (*config.RunningConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RunningConfigSpec_To_config_RunningConfigSpec(a.(*RunningConfigSpec), b.(*config.RunningConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RunningConfigSpec)(nil), (*RunningConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RunningConfigSpec_To_v1alpha1_RunningConfigSpec(a.(*config.RunningConfigSpec), b.(*RunningConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RunningConfigStatus)(nil), (*config.RunningConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RunningConfigStatus_To_config_RunningConfigStatus(a.(*RunningConfigStatus), b.(*config.RunningConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RunningConfigStatus)(nil), (*RunningConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RunningConfigStatus_To_v1alpha1_RunningConfigStatus(a.(*config.RunningConfigStatus), b.(*RunningConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Target)(nil), (*config.Target)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Target_To_config_Target(a.(*Target), b.(*config.Target), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Target)(nil), (*Target)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Target_To_v1alpha1_Target(a.(*config.Target), b.(*Target), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TargetStatus)(nil), (*config.TargetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_TargetStatus_To_config_TargetStatus(a.(*TargetStatus), b.(*config.TargetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TargetStatus)(nil), (*TargetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TargetStatus_To_v1alpha1_TargetStatus(a.(*config.TargetStatus), b.(*TargetStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Condition_To_config_Condition(in *Condition, out *config.Condition, s conversion.Scope) error {
	out.Condition = in.Condition
	return nil
}

// Convert_v1alpha1_Condition_To_config_Condition is an autogenerated conversion function.
func Convert_v1alpha1_Condition_To_config_Condition(in *Condition, out *config.Condition, s conversion.Scope) error {
	return autoConvert_v1alpha1_Condition_To_config_Condition(in, out, s)
}

func autoConvert_config_Condition_To_v1alpha1_Condition(in *config.Condition, out *Condition, s conversion.Scope) error {
	out.Condition = in.Condition
	return nil
}

// Convert_config_Condition_To_v1alpha1_Condition is an autogenerated conversion function.
func Convert_config_Condition_To_v1alpha1_Condition(in *config.Condition, out *Condition, s conversion.Scope) error {
	return autoConvert_config_Condition_To_v1alpha1_Condition(in, out, s)
}

func autoConvert_v1alpha1_ConditionedStatus_To_config_ConditionedStatus(in *ConditionedStatus, out *config.ConditionedStatus, s conversion.Scope) error {
	out.Conditions = *(*[]config.Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v1alpha1_ConditionedStatus_To_config_ConditionedStatus is an autogenerated conversion function.
func Convert_v1alpha1_ConditionedStatus_To_config_ConditionedStatus(in *ConditionedStatus, out *config.ConditionedStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConditionedStatus_To_config_ConditionedStatus(in, out, s)
}

func autoConvert_config_ConditionedStatus_To_v1alpha1_ConditionedStatus(in *config.ConditionedStatus, out *ConditionedStatus, s conversion.Scope) error {
	out.Conditions = *(*[]Condition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_config_ConditionedStatus_To_v1alpha1_ConditionedStatus is an autogenerated conversion function.
func Convert_config_ConditionedStatus_To_v1alpha1_ConditionedStatus(in *config.ConditionedStatus, out *ConditionedStatus, s conversion.Scope) error {
	return autoConvert_config_ConditionedStatus_To_v1alpha1_ConditionedStatus(in, out, s)
}

func autoConvert_v1alpha1_Config_To_config_Config(in *Config, out *config.Config, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ConfigSpec_To_config_ConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ConfigStatus_To_config_ConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Config_To_config_Config is an autogenerated conversion function.
func Convert_v1alpha1_Config_To_config_Config(in *Config, out *config.Config, s conversion.Scope) error {
	return autoConvert_v1alpha1_Config_To_config_Config(in, out, s)
}

func autoConvert_config_Config_To_v1alpha1_Config(in *config.Config, out *Config, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_config_ConfigSpec_To_v1alpha1_ConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_config_ConfigStatus_To_v1alpha1_ConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_Config_To_v1alpha1_Config is an autogenerated conversion function.
func Convert_config_Config_To_v1alpha1_Config(in *config.Config, out *Config, s conversion.Scope) error {
	return autoConvert_config_Config_To_v1alpha1_Config(in, out, s)
}

func autoConvert_v1alpha1_ConfigBlob_To_config_ConfigBlob(in *ConfigBlob, out *config.ConfigBlob, s conversion.Scope) error {
	out.Path = in.Path
	out.Value = in.Value
	return nil
}

// Convert_v1alpha1_ConfigBlob_To_config_ConfigBlob is an autogenerated conversion function.
func Convert_v1alpha1_ConfigBlob_To_config_ConfigBlob(in *ConfigBlob, out *config.ConfigBlob, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigBlob_To_config_ConfigBlob(in, out, s)
}

func autoConvert_config_ConfigBlob_To_v1alpha1_ConfigBlob(in *config.ConfigBlob, out *ConfigBlob, s conversion.Scope) error {
	out.Path = in.Path
	out.Value = in.Value
	return nil
}

// Convert_config_ConfigBlob_To_v1alpha1_ConfigBlob is an autogenerated conversion function.
func Convert_config_ConfigBlob_To_v1alpha1_ConfigBlob(in *config.ConfigBlob, out *ConfigBlob, s conversion.Scope) error {
	return autoConvert_config_ConfigBlob_To_v1alpha1_ConfigBlob(in, out, s)
}

func autoConvert_v1alpha1_ConfigList_To_config_ConfigList(in *ConfigList, out *config.ConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]config.Config)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ConfigList_To_config_ConfigList is an autogenerated conversion function.
func Convert_v1alpha1_ConfigList_To_config_ConfigList(in *ConfigList, out *config.ConfigList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigList_To_config_ConfigList(in, out, s)
}

func autoConvert_config_ConfigList_To_v1alpha1_ConfigList(in *config.ConfigList, out *ConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Config)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_config_ConfigList_To_v1alpha1_ConfigList is an autogenerated conversion function.
func Convert_config_ConfigList_To_v1alpha1_ConfigList(in *config.ConfigList, out *ConfigList, s conversion.Scope) error {
	return autoConvert_config_ConfigList_To_v1alpha1_ConfigList(in, out, s)
}

func autoConvert_v1alpha1_ConfigSet_To_config_ConfigSet(in *ConfigSet, out *config.ConfigSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ConfigSetSpec_To_config_ConfigSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ConfigSetStatus_To_config_ConfigSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ConfigSet_To_config_ConfigSet is an autogenerated conversion function.
func Convert_v1alpha1_ConfigSet_To_config_ConfigSet(in *ConfigSet, out *config.ConfigSet, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigSet_To_config_ConfigSet(in, out, s)
}

func autoConvert_config_ConfigSet_To_v1alpha1_ConfigSet(in *config.ConfigSet, out *ConfigSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_config_ConfigSetSpec_To_v1alpha1_ConfigSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_config_ConfigSetStatus_To_v1alpha1_ConfigSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ConfigSet_To_v1alpha1_ConfigSet is an autogenerated conversion function.
func Convert_config_ConfigSet_To_v1alpha1_ConfigSet(in *config.ConfigSet, out *ConfigSet, s conversion.Scope) error {
	return autoConvert_config_ConfigSet_To_v1alpha1_ConfigSet(in, out, s)
}

func autoConvert_v1alpha1_ConfigSetList_To_config_ConfigSetList(in *ConfigSetList, out *config.ConfigSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]config.ConfigSet)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ConfigSetList_To_config_ConfigSetList is an autogenerated conversion function.
func Convert_v1alpha1_ConfigSetList_To_config_ConfigSetList(in *ConfigSetList, out *config.ConfigSetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigSetList_To_config_ConfigSetList(in, out, s)
}

func autoConvert_config_ConfigSetList_To_v1alpha1_ConfigSetList(in *config.ConfigSetList, out *ConfigSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ConfigSet)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_config_ConfigSetList_To_v1alpha1_ConfigSetList is an autogenerated conversion function.
func Convert_config_ConfigSetList_To_v1alpha1_ConfigSetList(in *config.ConfigSetList, out *ConfigSetList, s conversion.Scope) error {
	return autoConvert_config_ConfigSetList_To_v1alpha1_ConfigSetList(in, out, s)
}

func autoConvert_v1alpha1_ConfigSetSpec_To_config_ConfigSetSpec(in *ConfigSetSpec, out *config.ConfigSetSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_Target_To_config_Target(&in.Target, &out.Target, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_Lifecycle_To_config_Lifecycle(&in.Lifecycle, &out.Lifecycle, s); err != nil {
		return err
	}
	out.Priority = in.Priority
	out.Config = *(*[]config.ConfigBlob)(unsafe.Pointer(&in.Config))
	return nil
}

// Convert_v1alpha1_ConfigSetSpec_To_config_ConfigSetSpec is an autogenerated conversion function.
func Convert_v1alpha1_ConfigSetSpec_To_config_ConfigSetSpec(in *ConfigSetSpec, out *config.ConfigSetSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigSetSpec_To_config_ConfigSetSpec(in, out, s)
}

func autoConvert_config_ConfigSetSpec_To_v1alpha1_ConfigSetSpec(in *config.ConfigSetSpec, out *ConfigSetSpec, s conversion.Scope) error {
	if err := Convert_config_Target_To_v1alpha1_Target(&in.Target, &out.Target, s); err != nil {
		return err
	}
	if err := Convert_config_Lifecycle_To_v1alpha1_Lifecycle(&in.Lifecycle, &out.Lifecycle, s); err != nil {
		return err
	}
	out.Priority = in.Priority
	out.Config = *(*[]ConfigBlob)(unsafe.Pointer(&in.Config))
	return nil
}

// Convert_config_ConfigSetSpec_To_v1alpha1_ConfigSetSpec is an autogenerated conversion function.
func Convert_config_ConfigSetSpec_To_v1alpha1_ConfigSetSpec(in *config.ConfigSetSpec, out *ConfigSetSpec, s conversion.Scope) error {
	return autoConvert_config_ConfigSetSpec_To_v1alpha1_ConfigSetSpec(in, out, s)
}

func autoConvert_v1alpha1_ConfigSetStatus_To_config_ConfigSetStatus(in *ConfigSetStatus, out *config.ConfigSetStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_ConditionedStatus_To_config_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	out.Targets = *(*[]config.TargetStatus)(unsafe.Pointer(&in.Targets))
	return nil
}

// Convert_v1alpha1_ConfigSetStatus_To_config_ConfigSetStatus is an autogenerated conversion function.
func Convert_v1alpha1_ConfigSetStatus_To_config_ConfigSetStatus(in *ConfigSetStatus, out *config.ConfigSetStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigSetStatus_To_config_ConfigSetStatus(in, out, s)
}

func autoConvert_config_ConfigSetStatus_To_v1alpha1_ConfigSetStatus(in *config.ConfigSetStatus, out *ConfigSetStatus, s conversion.Scope) error {
	if err := Convert_config_ConditionedStatus_To_v1alpha1_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	out.Targets = *(*[]TargetStatus)(unsafe.Pointer(&in.Targets))
	return nil
}

// Convert_config_ConfigSetStatus_To_v1alpha1_ConfigSetStatus is an autogenerated conversion function.
func Convert_config_ConfigSetStatus_To_v1alpha1_ConfigSetStatus(in *config.ConfigSetStatus, out *ConfigSetStatus, s conversion.Scope) error {
	return autoConvert_config_ConfigSetStatus_To_v1alpha1_ConfigSetStatus(in, out, s)
}

func autoConvert_v1alpha1_ConfigSpec_To_config_ConfigSpec(in *ConfigSpec, out *config.ConfigSpec, s conversion.Scope) error {
	if err := Convert_v1alpha1_Lifecycle_To_config_Lifecycle(&in.Lifecycle, &out.Lifecycle, s); err != nil {
		return err
	}
	out.Priority = in.Priority
	out.Config = *(*[]config.ConfigBlob)(unsafe.Pointer(&in.Config))
	return nil
}

// Convert_v1alpha1_ConfigSpec_To_config_ConfigSpec is an autogenerated conversion function.
func Convert_v1alpha1_ConfigSpec_To_config_ConfigSpec(in *ConfigSpec, out *config.ConfigSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigSpec_To_config_ConfigSpec(in, out, s)
}

func autoConvert_config_ConfigSpec_To_v1alpha1_ConfigSpec(in *config.ConfigSpec, out *ConfigSpec, s conversion.Scope) error {
	if err := Convert_config_Lifecycle_To_v1alpha1_Lifecycle(&in.Lifecycle, &out.Lifecycle, s); err != nil {
		return err
	}
	out.Priority = in.Priority
	out.Config = *(*[]ConfigBlob)(unsafe.Pointer(&in.Config))
	return nil
}

// Convert_config_ConfigSpec_To_v1alpha1_ConfigSpec is an autogenerated conversion function.
func Convert_config_ConfigSpec_To_v1alpha1_ConfigSpec(in *config.ConfigSpec, out *ConfigSpec, s conversion.Scope) error {
	return autoConvert_config_ConfigSpec_To_v1alpha1_ConfigSpec(in, out, s)
}

func autoConvert_v1alpha1_ConfigStatus_To_config_ConfigStatus(in *ConfigStatus, out *config.ConfigStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_ConditionedStatus_To_config_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	out.LastKnownGoodSchema = (*config.ConfigStatusLastKnownGoodSchema)(unsafe.Pointer(in.LastKnownGoodSchema))
	out.AppliedConfig = (*config.ConfigSpec)(unsafe.Pointer(in.AppliedConfig))
	out.Deviations = *(*[]config.Deviation)(unsafe.Pointer(&in.Deviations))
	return nil
}

// Convert_v1alpha1_ConfigStatus_To_config_ConfigStatus is an autogenerated conversion function.
func Convert_v1alpha1_ConfigStatus_To_config_ConfigStatus(in *ConfigStatus, out *config.ConfigStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigStatus_To_config_ConfigStatus(in, out, s)
}

func autoConvert_config_ConfigStatus_To_v1alpha1_ConfigStatus(in *config.ConfigStatus, out *ConfigStatus, s conversion.Scope) error {
	if err := Convert_config_ConditionedStatus_To_v1alpha1_ConditionedStatus(&in.ConditionedStatus, &out.ConditionedStatus, s); err != nil {
		return err
	}
	out.LastKnownGoodSchema = (*ConfigStatusLastKnownGoodSchema)(unsafe.Pointer(in.LastKnownGoodSchema))
	out.AppliedConfig = (*ConfigSpec)(unsafe.Pointer(in.AppliedConfig))
	out.Deviations = *(*[]Deviation)(unsafe.Pointer(&in.Deviations))
	return nil
}

// Convert_config_ConfigStatus_To_v1alpha1_ConfigStatus is an autogenerated conversion function.
func Convert_config_ConfigStatus_To_v1alpha1_ConfigStatus(in *config.ConfigStatus, out *ConfigStatus, s conversion.Scope) error {
	return autoConvert_config_ConfigStatus_To_v1alpha1_ConfigStatus(in, out, s)
}

func autoConvert_v1alpha1_ConfigStatusLastKnownGoodSchema_To_config_ConfigStatusLastKnownGoodSchema(in *ConfigStatusLastKnownGoodSchema, out *config.ConfigStatusLastKnownGoodSchema, s conversion.Scope) error {
	out.Type = in.Type
	out.Vendor = in.Vendor
	out.Version = in.Version
	return nil
}

// Convert_v1alpha1_ConfigStatusLastKnownGoodSchema_To_config_ConfigStatusLastKnownGoodSchema is an autogenerated conversion function.
func Convert_v1alpha1_ConfigStatusLastKnownGoodSchema_To_config_ConfigStatusLastKnownGoodSchema(in *ConfigStatusLastKnownGoodSchema, out *config.ConfigStatusLastKnownGoodSchema, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigStatusLastKnownGoodSchema_To_config_ConfigStatusLastKnownGoodSchema(in, out, s)
}

func autoConvert_config_ConfigStatusLastKnownGoodSchema_To_v1alpha1_ConfigStatusLastKnownGoodSchema(in *config.ConfigStatusLastKnownGoodSchema, out *ConfigStatusLastKnownGoodSchema, s conversion.Scope) error {
	out.Type = in.Type
	out.Vendor = in.Vendor
	out.Version = in.Version
	return nil
}

// Convert_config_ConfigStatusLastKnownGoodSchema_To_v1alpha1_ConfigStatusLastKnownGoodSchema is an autogenerated conversion function.
func Convert_config_ConfigStatusLastKnownGoodSchema_To_v1alpha1_ConfigStatusLastKnownGoodSchema(in *config.ConfigStatusLastKnownGoodSchema, out *ConfigStatusLastKnownGoodSchema, s conversion.Scope) error {
	return autoConvert_config_ConfigStatusLastKnownGoodSchema_To_v1alpha1_ConfigStatusLastKnownGoodSchema(in, out, s)
}

func autoConvert_v1alpha1_Deviation_To_config_Deviation(in *Deviation, out *config.Deviation, s conversion.Scope) error {
	out.Path = in.Path
	out.DesiredValue = in.DesiredValue
	// WARNING: in.CurrentValue requires manual conversion: does not exist in peer-type
	out.Reason = in.Reason
	return nil
}

func autoConvert_config_Deviation_To_v1alpha1_Deviation(in *config.Deviation, out *Deviation, s conversion.Scope) error {
	out.Path = in.Path
	out.DesiredValue = in.DesiredValue
	// WARNING: in.ActualValue requires manual conversion: does not exist in peer-type
	out.Reason = in.Reason
	return nil
}

func autoConvert_v1alpha1_Lifecycle_To_config_Lifecycle(in *Lifecycle, out *config.Lifecycle, s conversion.Scope) error {
	out.DeletionPolicy = config.DeletionPolicy(in.DeletionPolicy)
	return nil
}

// Convert_v1alpha1_Lifecycle_To_config_Lifecycle is an autogenerated conversion function.
func Convert_v1alpha1_Lifecycle_To_config_Lifecycle(in *Lifecycle, out *config.Lifecycle, s conversion.Scope) error {
	return autoConvert_v1alpha1_Lifecycle_To_config_Lifecycle(in, out, s)
}

func autoConvert_config_Lifecycle_To_v1alpha1_Lifecycle(in *config.Lifecycle, out *Lifecycle, s conversion.Scope) error {
	out.DeletionPolicy = DeletionPolicy(in.DeletionPolicy)
	return nil
}

// Convert_config_Lifecycle_To_v1alpha1_Lifecycle is an autogenerated conversion function.
func Convert_config_Lifecycle_To_v1alpha1_Lifecycle(in *config.Lifecycle, out *Lifecycle, s conversion.Scope) error {
	return autoConvert_config_Lifecycle_To_v1alpha1_Lifecycle(in, out, s)
}

func autoConvert_v1alpha1_RunningConfig_To_config_RunningConfig(in *RunningConfig, out *config.RunningConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_RunningConfigSpec_To_config_RunningConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_RunningConfigStatus_To_config_RunningConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_RunningConfig_To_config_RunningConfig is an autogenerated conversion function.
func Convert_v1alpha1_RunningConfig_To_config_RunningConfig(in *RunningConfig, out *config.RunningConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_RunningConfig_To_config_RunningConfig(in, out, s)
}

func autoConvert_config_RunningConfig_To_v1alpha1_RunningConfig(in *config.RunningConfig, out *RunningConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_config_RunningConfigSpec_To_v1alpha1_RunningConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_config_RunningConfigStatus_To_v1alpha1_RunningConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_RunningConfig_To_v1alpha1_RunningConfig is an autogenerated conversion function.
func Convert_config_RunningConfig_To_v1alpha1_RunningConfig(in *config.RunningConfig, out *RunningConfig, s conversion.Scope) error {
	return autoConvert_config_RunningConfig_To_v1alpha1_RunningConfig(in, out, s)
}

func autoConvert_v1alpha1_RunningConfigList_To_config_RunningConfigList(in *RunningConfigList, out *config.RunningConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]config.RunningConfig)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_RunningConfigList_To_config_RunningConfigList is an autogenerated conversion function.
func Convert_v1alpha1_RunningConfigList_To_config_RunningConfigList(in *RunningConfigList, out *config.RunningConfigList, s conversion.Scope) error {
	return autoConvert_v1alpha1_RunningConfigList_To_config_RunningConfigList(in, out, s)
}

func autoConvert_config_RunningConfigList_To_v1alpha1_RunningConfigList(in *config.RunningConfigList, out *RunningConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]RunningConfig)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_config_RunningConfigList_To_v1alpha1_RunningConfigList is an autogenerated conversion function.
func Convert_config_RunningConfigList_To_v1alpha1_RunningConfigList(in *config.RunningConfigList, out *RunningConfigList, s conversion.Scope) error {
	return autoConvert_config_RunningConfigList_To_v1alpha1_RunningConfigList(in, out, s)
}

func autoConvert_v1alpha1_RunningConfigSpec_To_config_RunningConfigSpec(in *RunningConfigSpec, out *config.RunningConfigSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_RunningConfigSpec_To_config_RunningConfigSpec is an autogenerated conversion function.
func Convert_v1alpha1_RunningConfigSpec_To_config_RunningConfigSpec(in *RunningConfigSpec, out *config.RunningConfigSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_RunningConfigSpec_To_config_RunningConfigSpec(in, out, s)
}

func autoConvert_config_RunningConfigSpec_To_v1alpha1_RunningConfigSpec(in *config.RunningConfigSpec, out *RunningConfigSpec, s conversion.Scope) error {
	return nil
}

// Convert_config_RunningConfigSpec_To_v1alpha1_RunningConfigSpec is an autogenerated conversion function.
func Convert_config_RunningConfigSpec_To_v1alpha1_RunningConfigSpec(in *config.RunningConfigSpec, out *RunningConfigSpec, s conversion.Scope) error {
	return autoConvert_config_RunningConfigSpec_To_v1alpha1_RunningConfigSpec(in, out, s)
}

func autoConvert_v1alpha1_RunningConfigStatus_To_config_RunningConfigStatus(in *RunningConfigStatus, out *config.RunningConfigStatus, s conversion.Scope) error {
	out.Value = in.Value
	return nil
}

// Convert_v1alpha1_RunningConfigStatus_To_config_RunningConfigStatus is an autogenerated conversion function.
func Convert_v1alpha1_RunningConfigStatus_To_config_RunningConfigStatus(in *RunningConfigStatus, out *config.RunningConfigStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_RunningConfigStatus_To_config_RunningConfigStatus(in, out, s)
}

func autoConvert_config_RunningConfigStatus_To_v1alpha1_RunningConfigStatus(in *config.RunningConfigStatus, out *RunningConfigStatus, s conversion.Scope) error {
	out.Value = in.Value
	return nil
}

// Convert_config_RunningConfigStatus_To_v1alpha1_RunningConfigStatus is an autogenerated conversion function.
func Convert_config_RunningConfigStatus_To_v1alpha1_RunningConfigStatus(in *config.RunningConfigStatus, out *RunningConfigStatus, s conversion.Scope) error {
	return autoConvert_config_RunningConfigStatus_To_v1alpha1_RunningConfigStatus(in, out, s)
}

func autoConvert_v1alpha1_Target_To_config_Target(in *Target, out *config.Target, s conversion.Scope) error {
	out.TargetSelector = (*v1.LabelSelector)(unsafe.Pointer(in.TargetSelector))
	return nil
}

// Convert_v1alpha1_Target_To_config_Target is an autogenerated conversion function.
func Convert_v1alpha1_Target_To_config_Target(in *Target, out *config.Target, s conversion.Scope) error {
	return autoConvert_v1alpha1_Target_To_config_Target(in, out, s)
}

func autoConvert_config_Target_To_v1alpha1_Target(in *config.Target, out *Target, s conversion.Scope) error {
	out.TargetSelector = (*v1.LabelSelector)(unsafe.Pointer(in.TargetSelector))
	return nil
}

// Convert_config_Target_To_v1alpha1_Target is an autogenerated conversion function.
func Convert_config_Target_To_v1alpha1_Target(in *config.Target, out *Target, s conversion.Scope) error {
	return autoConvert_config_Target_To_v1alpha1_Target(in, out, s)
}

func autoConvert_v1alpha1_TargetStatus_To_config_TargetStatus(in *TargetStatus, out *config.TargetStatus, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_v1alpha1_Condition_To_config_Condition(&in.Condition, &out.Condition, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_TargetStatus_To_config_TargetStatus is an autogenerated conversion function.
func Convert_v1alpha1_TargetStatus_To_config_TargetStatus(in *TargetStatus, out *config.TargetStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_TargetStatus_To_config_TargetStatus(in, out, s)
}

func autoConvert_config_TargetStatus_To_v1alpha1_TargetStatus(in *config.TargetStatus, out *TargetStatus, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_config_Condition_To_v1alpha1_Condition(&in.Condition, &out.Condition, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_TargetStatus_To_v1alpha1_TargetStatus is an autogenerated conversion function.
func Convert_config_TargetStatus_To_v1alpha1_TargetStatus(in *config.TargetStatus, out *TargetStatus, s conversion.Scope) error {
	return autoConvert_config_TargetStatus_To_v1alpha1_TargetStatus(in, out, s)
}
