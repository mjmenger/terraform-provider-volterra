// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package waf

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *MetricLabelFilter) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *MetricLabelFilter) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *MetricLabelFilter) DeepCopy() *MetricLabelFilter {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &MetricLabelFilter{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *MetricLabelFilter) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *MetricLabelFilter) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return MetricLabelFilterValidator().Validate(ctx, m, opts...)
}

type ValidateMetricLabelFilter struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateMetricLabelFilter) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*MetricLabelFilter)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *MetricLabelFilter got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["label"]; exists {

		vOpts := append(opts, db.WithValidateField("label"))
		if err := fv(ctx, m.GetLabel(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["op"]; exists {

		vOpts := append(opts, db.WithValidateField("op"))
		if err := fv(ctx, m.GetOp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["value"]; exists {

		vOpts := append(opts, db.WithValidateField("value"))
		if err := fv(ctx, m.GetValue(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultMetricLabelFilterValidator = func() *ValidateMetricLabelFilter {
	v := &ValidateMetricLabelFilter{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func MetricLabelFilterValidator() db.Validator {
	return DefaultMetricLabelFilterValidator
}

// augmented methods on protoc/std generated struct

func (m *RuleHitsCountRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RuleHitsCountRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RuleHitsCountRequest) DeepCopy() *RuleHitsCountRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RuleHitsCountRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RuleHitsCountRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RuleHitsCountRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RuleHitsCountRequestValidator().Validate(ctx, m, opts...)
}

type ValidateRuleHitsCountRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRuleHitsCountRequest) StartTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for start_time")
	}

	return validatorFn, nil
}

func (v *ValidateRuleHitsCountRequest) EndTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for end_time")
	}

	return validatorFn, nil
}

func (v *ValidateRuleHitsCountRequest) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateRuleHitsCountRequest) RangeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for range")
	}

	return validatorFn, nil
}

func (v *ValidateRuleHitsCountRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RuleHitsCountRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RuleHitsCountRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_time"]; exists {

		vOpts := append(opts, db.WithValidateField("end_time"))
		if err := fv(ctx, m.GetEndTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["group_by"]; exists {

		vOpts := append(opts, db.WithValidateField("group_by"))
		for idx, item := range m.GetGroupBy() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["label_filter"]; exists {

		vOpts := append(opts, db.WithValidateField("label_filter"))
		for idx, item := range m.GetLabelFilter() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["range"]; exists {

		vOpts := append(opts, db.WithValidateField("range"))
		if err := fv(ctx, m.GetRange(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["step"]; exists {

		vOpts := append(opts, db.WithValidateField("step"))
		if err := fv(ctx, m.GetStep(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRuleHitsCountRequestValidator = func() *ValidateRuleHitsCountRequest {
	v := &ValidateRuleHitsCountRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhStartTime := v.StartTimeValidationRuleHandler
	rulesStartTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhStartTime(rulesStartTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RuleHitsCountRequest.start_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["start_time"] = vFn

	vrhEndTime := v.EndTimeValidationRuleHandler
	rulesEndTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhEndTime(rulesEndTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RuleHitsCountRequest.end_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["end_time"] = vFn

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.time_interval": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RuleHitsCountRequest.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	vrhRange := v.RangeValidationRuleHandler
	rulesRange := map[string]string{
		"ves.io.schema.rules.string.time_interval": "true",
	}
	vFn, err = vrhRange(rulesRange)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RuleHitsCountRequest.range: %s", err)
		panic(errMsg)
	}
	v.FldValidators["range"] = vFn

	return v
}()

func RuleHitsCountRequestValidator() db.Validator {
	return DefaultRuleHitsCountRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *RuleHitsCountResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RuleHitsCountResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RuleHitsCountResponse) DeepCopy() *RuleHitsCountResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RuleHitsCountResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RuleHitsCountResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RuleHitsCountResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RuleHitsCountResponseValidator().Validate(ctx, m, opts...)
}

type ValidateRuleHitsCountResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRuleHitsCountResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RuleHitsCountResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RuleHitsCountResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["data"]; exists {

		vOpts := append(opts, db.WithValidateField("data"))
		for idx, item := range m.GetData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRuleHitsCountResponseValidator = func() *ValidateRuleHitsCountResponse {
	v := &ValidateRuleHitsCountResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RuleHitsCountResponseValidator() db.Validator {
	return DefaultRuleHitsCountResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *RuleHitsCounter) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RuleHitsCounter) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RuleHitsCounter) DeepCopy() *RuleHitsCounter {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RuleHitsCounter{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RuleHitsCounter) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RuleHitsCounter) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RuleHitsCounterValidator().Validate(ctx, m, opts...)
}

type ValidateRuleHitsCounter struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRuleHitsCounter) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RuleHitsCounter)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RuleHitsCounter got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["metric"]; exists {

		vOpts := append(opts, db.WithValidateField("metric"))
		for idx, item := range m.GetMetric() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRuleHitsCounterValidator = func() *ValidateRuleHitsCounter {
	v := &ValidateRuleHitsCounter{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RuleHitsCounterValidator() db.Validator {
	return DefaultRuleHitsCounterValidator
}

// augmented methods on protoc/std generated struct

func (m *RuleHitsId) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RuleHitsId) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RuleHitsId) DeepCopy() *RuleHitsId {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RuleHitsId{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RuleHitsId) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RuleHitsId) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RuleHitsIdValidator().Validate(ctx, m, opts...)
}

type ValidateRuleHitsId struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRuleHitsId) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RuleHitsId)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RuleHitsId got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["app_type"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type"))
		if err := fv(ctx, m.GetAppType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["bot_name"]; exists {

		vOpts := append(opts, db.WithValidateField("bot_name"))
		if err := fv(ctx, m.GetBotName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["bot_type"]; exists {

		vOpts := append(opts, db.WithValidateField("bot_type"))
		if err := fv(ctx, m.GetBotType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["instance"]; exists {

		vOpts := append(opts, db.WithValidateField("instance"))
		if err := fv(ctx, m.GetInstance(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_id"]; exists {

		vOpts := append(opts, db.WithValidateField("rule_id"))
		if err := fv(ctx, m.GetRuleId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_severity"]; exists {

		vOpts := append(opts, db.WithValidateField("rule_severity"))
		if err := fv(ctx, m.GetRuleSeverity(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rule_tag"]; exists {

		vOpts := append(opts, db.WithValidateField("rule_tag"))
		if err := fv(ctx, m.GetRuleTag(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["service"]; exists {

		vOpts := append(opts, db.WithValidateField("service"))
		if err := fv(ctx, m.GetService(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site"]; exists {

		vOpts := append(opts, db.WithValidateField("site"))
		if err := fv(ctx, m.GetSite(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_host"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_host"))
		if err := fv(ctx, m.GetVirtualHost(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["waf_instance_id"]; exists {

		vOpts := append(opts, db.WithValidateField("waf_instance_id"))
		if err := fv(ctx, m.GetWafInstanceId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRuleHitsIdValidator = func() *ValidateRuleHitsId {
	v := &ValidateRuleHitsId{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RuleHitsIdValidator() db.Validator {
	return DefaultRuleHitsIdValidator
}

// augmented methods on protoc/std generated struct

func (m *SecurityEventsCountRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SecurityEventsCountRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SecurityEventsCountRequest) DeepCopy() *SecurityEventsCountRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SecurityEventsCountRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SecurityEventsCountRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SecurityEventsCountRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SecurityEventsCountRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSecurityEventsCountRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSecurityEventsCountRequest) StartTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for start_time")
	}

	return validatorFn, nil
}

func (v *ValidateSecurityEventsCountRequest) EndTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for end_time")
	}

	return validatorFn, nil
}

func (v *ValidateSecurityEventsCountRequest) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateSecurityEventsCountRequest) RangeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for range")
	}

	return validatorFn, nil
}

func (v *ValidateSecurityEventsCountRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SecurityEventsCountRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SecurityEventsCountRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_time"]; exists {

		vOpts := append(opts, db.WithValidateField("end_time"))
		if err := fv(ctx, m.GetEndTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["group_by"]; exists {

		vOpts := append(opts, db.WithValidateField("group_by"))
		for idx, item := range m.GetGroupBy() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["label_filter"]; exists {

		vOpts := append(opts, db.WithValidateField("label_filter"))
		for idx, item := range m.GetLabelFilter() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["range"]; exists {

		vOpts := append(opts, db.WithValidateField("range"))
		if err := fv(ctx, m.GetRange(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["step"]; exists {

		vOpts := append(opts, db.WithValidateField("step"))
		if err := fv(ctx, m.GetStep(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSecurityEventsCountRequestValidator = func() *ValidateSecurityEventsCountRequest {
	v := &ValidateSecurityEventsCountRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhStartTime := v.StartTimeValidationRuleHandler
	rulesStartTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhStartTime(rulesStartTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SecurityEventsCountRequest.start_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["start_time"] = vFn

	vrhEndTime := v.EndTimeValidationRuleHandler
	rulesEndTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhEndTime(rulesEndTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SecurityEventsCountRequest.end_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["end_time"] = vFn

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.time_interval": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SecurityEventsCountRequest.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	vrhRange := v.RangeValidationRuleHandler
	rulesRange := map[string]string{
		"ves.io.schema.rules.string.time_interval": "true",
	}
	vFn, err = vrhRange(rulesRange)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SecurityEventsCountRequest.range: %s", err)
		panic(errMsg)
	}
	v.FldValidators["range"] = vFn

	return v
}()

func SecurityEventsCountRequestValidator() db.Validator {
	return DefaultSecurityEventsCountRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SecurityEventsCountResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SecurityEventsCountResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SecurityEventsCountResponse) DeepCopy() *SecurityEventsCountResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SecurityEventsCountResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SecurityEventsCountResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SecurityEventsCountResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SecurityEventsCountResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSecurityEventsCountResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSecurityEventsCountResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SecurityEventsCountResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SecurityEventsCountResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["data"]; exists {

		vOpts := append(opts, db.WithValidateField("data"))
		for idx, item := range m.GetData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSecurityEventsCountResponseValidator = func() *ValidateSecurityEventsCountResponse {
	v := &ValidateSecurityEventsCountResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SecurityEventsCountResponseValidator() db.Validator {
	return DefaultSecurityEventsCountResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SecurityEventsCounter) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SecurityEventsCounter) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SecurityEventsCounter) DeepCopy() *SecurityEventsCounter {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SecurityEventsCounter{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SecurityEventsCounter) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SecurityEventsCounter) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SecurityEventsCounterValidator().Validate(ctx, m, opts...)
}

type ValidateSecurityEventsCounter struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSecurityEventsCounter) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SecurityEventsCounter)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SecurityEventsCounter got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["metric"]; exists {

		vOpts := append(opts, db.WithValidateField("metric"))
		for idx, item := range m.GetMetric() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSecurityEventsCounterValidator = func() *ValidateSecurityEventsCounter {
	v := &ValidateSecurityEventsCounter{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SecurityEventsCounterValidator() db.Validator {
	return DefaultSecurityEventsCounterValidator
}

// augmented methods on protoc/std generated struct

func (m *SecurityEventsId) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SecurityEventsId) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SecurityEventsId) DeepCopy() *SecurityEventsId {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SecurityEventsId{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SecurityEventsId) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SecurityEventsId) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SecurityEventsIdValidator().Validate(ctx, m, opts...)
}

type ValidateSecurityEventsId struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSecurityEventsId) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SecurityEventsId)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SecurityEventsId got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["app_type"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type"))
		if err := fv(ctx, m.GetAppType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["bot_name"]; exists {

		vOpts := append(opts, db.WithValidateField("bot_name"))
		if err := fv(ctx, m.GetBotName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["bot_type"]; exists {

		vOpts := append(opts, db.WithValidateField("bot_type"))
		if err := fv(ctx, m.GetBotType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["instance"]; exists {

		vOpts := append(opts, db.WithValidateField("instance"))
		if err := fv(ctx, m.GetInstance(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["service"]; exists {

		vOpts := append(opts, db.WithValidateField("service"))
		if err := fv(ctx, m.GetService(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site"]; exists {

		vOpts := append(opts, db.WithValidateField("site"))
		if err := fv(ctx, m.GetSite(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["virtual_host"]; exists {

		vOpts := append(opts, db.WithValidateField("virtual_host"))
		if err := fv(ctx, m.GetVirtualHost(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["waf_instance_id"]; exists {

		vOpts := append(opts, db.WithValidateField("waf_instance_id"))
		if err := fv(ctx, m.GetWafInstanceId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["waf_mode"]; exists {

		vOpts := append(opts, db.WithValidateField("waf_mode"))
		if err := fv(ctx, m.GetWafMode(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSecurityEventsIdValidator = func() *ValidateSecurityEventsId {
	v := &ValidateSecurityEventsId{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SecurityEventsIdValidator() db.Validator {
	return DefaultSecurityEventsIdValidator
}
