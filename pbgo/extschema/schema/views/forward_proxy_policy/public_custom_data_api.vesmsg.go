// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package forward_proxy_policy

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

func (m *ForwardProxyPolicyHits) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ForwardProxyPolicyHits) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ForwardProxyPolicyHits) DeepCopy() *ForwardProxyPolicyHits {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ForwardProxyPolicyHits{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ForwardProxyPolicyHits) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ForwardProxyPolicyHits) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ForwardProxyPolicyHitsValidator().Validate(ctx, m, opts...)
}

type ValidateForwardProxyPolicyHits struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateForwardProxyPolicyHits) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ForwardProxyPolicyHits)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ForwardProxyPolicyHits got type %s", t)
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
var DefaultForwardProxyPolicyHitsValidator = func() *ValidateForwardProxyPolicyHits {
	v := &ValidateForwardProxyPolicyHits{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ForwardProxyPolicyHitsValidator() db.Validator {
	return DefaultForwardProxyPolicyHitsValidator
}

// augmented methods on protoc/std generated struct

func (m *ForwardProxyPolicyHitsId) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ForwardProxyPolicyHitsId) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ForwardProxyPolicyHitsId) DeepCopy() *ForwardProxyPolicyHitsId {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ForwardProxyPolicyHitsId{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ForwardProxyPolicyHitsId) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ForwardProxyPolicyHitsId) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ForwardProxyPolicyHitsIdValidator().Validate(ctx, m, opts...)
}

type ValidateForwardProxyPolicyHitsId struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateForwardProxyPolicyHitsId) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ForwardProxyPolicyHitsId)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ForwardProxyPolicyHitsId got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["action"]; exists {

		vOpts := append(opts, db.WithValidateField("action"))
		if err := fv(ctx, m.GetAction(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["policy"]; exists {

		vOpts := append(opts, db.WithValidateField("policy"))
		if err := fv(ctx, m.GetPolicy(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["policy_rule"]; exists {

		vOpts := append(opts, db.WithValidateField("policy_rule"))
		if err := fv(ctx, m.GetPolicyRule(), vOpts...); err != nil {
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultForwardProxyPolicyHitsIdValidator = func() *ValidateForwardProxyPolicyHitsId {
	v := &ValidateForwardProxyPolicyHitsId{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ForwardProxyPolicyHitsIdValidator() db.Validator {
	return DefaultForwardProxyPolicyHitsIdValidator
}

// augmented methods on protoc/std generated struct

func (m *ForwardProxyPolicyHitsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ForwardProxyPolicyHitsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ForwardProxyPolicyHitsRequest) DeepCopy() *ForwardProxyPolicyHitsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ForwardProxyPolicyHitsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ForwardProxyPolicyHitsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ForwardProxyPolicyHitsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ForwardProxyPolicyHitsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateForwardProxyPolicyHitsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateForwardProxyPolicyHitsRequest) StartTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for start_time")
	}

	return validatorFn, nil
}

func (v *ValidateForwardProxyPolicyHitsRequest) EndTimeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for end_time")
	}

	return validatorFn, nil
}

func (v *ValidateForwardProxyPolicyHitsRequest) StepValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for step")
	}

	return validatorFn, nil
}

func (v *ValidateForwardProxyPolicyHitsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ForwardProxyPolicyHitsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ForwardProxyPolicyHitsRequest got type %s", t)
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
var DefaultForwardProxyPolicyHitsRequestValidator = func() *ValidateForwardProxyPolicyHitsRequest {
	v := &ValidateForwardProxyPolicyHitsRequest{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for ForwardProxyPolicyHitsRequest.start_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["start_time"] = vFn

	vrhEndTime := v.EndTimeValidationRuleHandler
	rulesEndTime := map[string]string{
		"ves.io.schema.rules.string.query_time": "true",
	}
	vFn, err = vrhEndTime(rulesEndTime)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ForwardProxyPolicyHitsRequest.end_time: %s", err)
		panic(errMsg)
	}
	v.FldValidators["end_time"] = vFn

	vrhStep := v.StepValidationRuleHandler
	rulesStep := map[string]string{
		"ves.io.schema.rules.string.time_interval": "true",
	}
	vFn, err = vrhStep(rulesStep)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ForwardProxyPolicyHitsRequest.step: %s", err)
		panic(errMsg)
	}
	v.FldValidators["step"] = vFn

	return v
}()

func ForwardProxyPolicyHitsRequestValidator() db.Validator {
	return DefaultForwardProxyPolicyHitsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ForwardProxyPolicyHitsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ForwardProxyPolicyHitsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ForwardProxyPolicyHitsResponse) DeepCopy() *ForwardProxyPolicyHitsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ForwardProxyPolicyHitsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ForwardProxyPolicyHitsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ForwardProxyPolicyHitsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ForwardProxyPolicyHitsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateForwardProxyPolicyHitsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateForwardProxyPolicyHitsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ForwardProxyPolicyHitsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ForwardProxyPolicyHitsResponse got type %s", t)
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
var DefaultForwardProxyPolicyHitsResponseValidator = func() *ValidateForwardProxyPolicyHitsResponse {
	v := &ValidateForwardProxyPolicyHitsResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ForwardProxyPolicyHitsResponseValidator() db.Validator {
	return DefaultForwardProxyPolicyHitsResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ForwardProxyPolicyMetricLabelFilter) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ForwardProxyPolicyMetricLabelFilter) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ForwardProxyPolicyMetricLabelFilter) DeepCopy() *ForwardProxyPolicyMetricLabelFilter {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ForwardProxyPolicyMetricLabelFilter{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ForwardProxyPolicyMetricLabelFilter) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ForwardProxyPolicyMetricLabelFilter) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ForwardProxyPolicyMetricLabelFilterValidator().Validate(ctx, m, opts...)
}

type ValidateForwardProxyPolicyMetricLabelFilter struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateForwardProxyPolicyMetricLabelFilter) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ForwardProxyPolicyMetricLabelFilter)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ForwardProxyPolicyMetricLabelFilter got type %s", t)
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
var DefaultForwardProxyPolicyMetricLabelFilterValidator = func() *ValidateForwardProxyPolicyMetricLabelFilter {
	v := &ValidateForwardProxyPolicyMetricLabelFilter{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ForwardProxyPolicyMetricLabelFilterValidator() db.Validator {
	return DefaultForwardProxyPolicyMetricLabelFilterValidator
}
