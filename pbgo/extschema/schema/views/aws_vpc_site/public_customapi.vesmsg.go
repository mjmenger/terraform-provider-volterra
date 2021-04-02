//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package aws_vpc_site

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

func (m *SetVPCK8SHostnamesRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPCK8SHostnamesRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPCK8SHostnamesRequest) DeepCopy() *SetVPCK8SHostnamesRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPCK8SHostnamesRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPCK8SHostnamesRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPCK8SHostnamesRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPCK8SHostnamesRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPCK8SHostnamesRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPCK8SHostnamesRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPCK8SHostnamesRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPCK8SHostnamesRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["node_names"]; exists {

		vOpts := append(opts, db.WithValidateField("node_names"))
		for idx, item := range m.GetNodeNames() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPCK8SHostnamesRequestValidator = func() *ValidateSetVPCK8SHostnamesRequest {
	v := &ValidateSetVPCK8SHostnamesRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetVPCK8SHostnamesRequestValidator() db.Validator {
	return DefaultSetVPCK8SHostnamesRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVPCK8SHostnamesResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPCK8SHostnamesResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPCK8SHostnamesResponse) DeepCopy() *SetVPCK8SHostnamesResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPCK8SHostnamesResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPCK8SHostnamesResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPCK8SHostnamesResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPCK8SHostnamesResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPCK8SHostnamesResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPCK8SHostnamesResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPCK8SHostnamesResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPCK8SHostnamesResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPCK8SHostnamesResponseValidator = func() *ValidateSetVPCK8SHostnamesResponse {
	v := &ValidateSetVPCK8SHostnamesResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetVPCK8SHostnamesResponseValidator() db.Validator {
	return DefaultSetVPCK8SHostnamesResponseValidator
}