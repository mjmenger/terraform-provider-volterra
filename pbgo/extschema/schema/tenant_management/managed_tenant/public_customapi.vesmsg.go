// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package managed_tenant

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

func (m *GetManagedTenantListReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetManagedTenantListReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetManagedTenantListReq) DeepCopy() *GetManagedTenantListReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetManagedTenantListReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetManagedTenantListReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetManagedTenantListReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetManagedTenantListReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetManagedTenantListReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetManagedTenantListReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetManagedTenantListReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetManagedTenantListReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["page_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("page_limit"))
		if err := fv(ctx, m.GetPageLimit(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["page_start"]; exists {

		vOpts := append(opts, db.WithValidateField("page_start"))
		if err := fv(ctx, m.GetPageStart(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["search_keyword"]; exists {

		vOpts := append(opts, db.WithValidateField("search_keyword"))
		if err := fv(ctx, m.GetSearchKeyword(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetManagedTenantListReqValidator = func() *ValidateGetManagedTenantListReq {
	v := &ValidateGetManagedTenantListReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetManagedTenantListReqValidator() db.Validator {
	return DefaultGetManagedTenantListReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetManagedTenantListResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetManagedTenantListResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetManagedTenantListResp) DeepCopy() *GetManagedTenantListResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetManagedTenantListResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetManagedTenantListResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetManagedTenantListResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetManagedTenantListRespValidator().Validate(ctx, m, opts...)
}

func (m *GetManagedTenantListResp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetAccessConfigDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetManagedTenantListResp) GetAccessConfigDRefInfo() ([]db.DRefInfo, error) {
	if m.GetAccessConfig() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetAccessConfig() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetAccessConfig() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("access_config[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateGetManagedTenantListResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetManagedTenantListResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetManagedTenantListResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetManagedTenantListResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["access_config"]; exists {

		vOpts := append(opts, db.WithValidateField("access_config"))
		for idx, item := range m.GetAccessConfig() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["next_page"]; exists {

		vOpts := append(opts, db.WithValidateField("next_page"))
		if err := fv(ctx, m.GetNextPage(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["total_access_config_count"]; exists {

		vOpts := append(opts, db.WithValidateField("total_access_config_count"))
		if err := fv(ctx, m.GetTotalAccessConfigCount(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["total_count"]; exists {

		vOpts := append(opts, db.WithValidateField("total_count"))
		if err := fv(ctx, m.GetTotalCount(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["total_filtered_count"]; exists {

		vOpts := append(opts, db.WithValidateField("total_filtered_count"))
		if err := fv(ctx, m.GetTotalFilteredCount(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetManagedTenantListRespValidator = func() *ValidateGetManagedTenantListResp {
	v := &ValidateGetManagedTenantListResp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["access_config"] = AccessInfoValidator().Validate

	return v
}()

func GetManagedTenantListRespValidator() db.Validator {
	return DefaultGetManagedTenantListRespValidator
}
