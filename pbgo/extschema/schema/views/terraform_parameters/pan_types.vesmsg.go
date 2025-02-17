// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package terraform_parameters

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *PANAWSType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *PANAWSType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *PANAWSType) String() string {
	if m == nil {
		return ""
	}
	copy := m.DeepCopy()
	copy.PanAuthorizationKeyBlindfolded = ""

	return copy.string()
}

func (m *PANAWSType) GoString() string {
	copy := m.DeepCopy()
	copy.PanAuthorizationKeyBlindfolded = ""

	return copy.goString()
}

// Redact squashes sensitive info in m (in-place)
func (m *PANAWSType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.PanAuthorizationKeyBlindfolded = ""

	return nil
}

func (m *PANAWSType) DeepCopy() *PANAWSType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &PANAWSType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *PANAWSType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *PANAWSType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return PANAWSTypeValidator().Validate(ctx, m, opts...)
}

type ValidatePANAWSType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidatePANAWSType) AwsRegionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for aws_region")
	}

	return validatorFn, nil
}

func (v *ValidatePANAWSType) AwsNamePrefixValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for aws_name_prefix")
	}

	return validatorFn, nil
}

func (v *ValidatePANAWSType) VpcIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for vpc_id")
	}

	return validatorFn, nil
}

func (v *ValidatePANAWSType) AdminUsernameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for admin_username")
	}

	return validatorFn, nil
}

func (v *ValidatePANAWSType) DevicesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for devices")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*PaloAltoServicesNodeType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := PaloAltoServicesNodeTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for devices")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*PaloAltoServicesNodeType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*PaloAltoServicesNodeType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated devices")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items devices")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidatePANAWSType) SshKeyValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ssh_key")
	}

	return validatorFn, nil
}

func (v *ValidatePANAWSType) VolterraSubnetIdsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for volterra_subnet_ids")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema_views.AWSSubnetIdsType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema_views.AWSSubnetIdsTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for volterra_subnet_ids")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema_views.AWSSubnetIdsType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema_views.AWSSubnetIdsType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated volterra_subnet_ids")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items volterra_subnet_ids")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidatePANAWSType) SiteNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for site_name")
	}

	return validatorFn, nil
}

func (v *ValidatePANAWSType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*PANAWSType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *PANAWSType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["admin_username"]; exists {

		vOpts := append(opts, db.WithValidateField("admin_username"))
		if err := fv(ctx, m.GetAdminUsername(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["aws_name"]; exists {

		vOpts := append(opts, db.WithValidateField("aws_name"))
		if err := fv(ctx, m.GetAwsName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["aws_name_prefix"]; exists {

		vOpts := append(opts, db.WithValidateField("aws_name_prefix"))
		if err := fv(ctx, m.GetAwsNamePrefix(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["aws_region"]; exists {

		vOpts := append(opts, db.WithValidateField("aws_region"))
		if err := fv(ctx, m.GetAwsRegion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["devices"]; exists {
		vOpts := append(opts, db.WithValidateField("devices"))
		if err := fv(ctx, m.GetDevices(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["pan_authorization_key"]; exists {

		vOpts := append(opts, db.WithValidateField("pan_authorization_key"))
		if err := fv(ctx, m.GetPanAuthorizationKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["pan_authorization_key_blindfolded"]; exists {

		vOpts := append(opts, db.WithValidateField("pan_authorization_key_blindfolded"))
		if err := fv(ctx, m.GetPanAuthorizationKeyBlindfolded(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["pan_authorization_key_clear_b64"]; exists {

		vOpts := append(opts, db.WithValidateField("pan_authorization_key_clear_b64"))
		if err := fv(ctx, m.GetPanAuthorizationKeyClearB64(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["panorama_server"]; exists {

		vOpts := append(opts, db.WithValidateField("panorama_server"))
		if err := fv(ctx, m.GetPanoramaServer(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["panorama_template_name"]; exists {

		vOpts := append(opts, db.WithValidateField("panorama_template_name"))
		if err := fv(ctx, m.GetPanoramaTemplateName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["product_code"]; exists {

		vOpts := append(opts, db.WithValidateField("product_code"))
		if err := fv(ctx, m.GetProductCode(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site_name"]; exists {

		vOpts := append(opts, db.WithValidateField("site_name"))
		if err := fv(ctx, m.GetSiteName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ssh_key"]; exists {

		vOpts := append(opts, db.WithValidateField("ssh_key"))
		if err := fv(ctx, m.GetSshKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tags"]; exists {

		vOpts := append(opts, db.WithValidateField("tags"))
		for key, value := range m.GetTags() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["vmseries_version"]; exists {

		vOpts := append(opts, db.WithValidateField("vmseries_version"))
		if err := fv(ctx, m.GetVmseriesVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["volterra_subnet_ids"]; exists {
		vOpts := append(opts, db.WithValidateField("volterra_subnet_ids"))
		if err := fv(ctx, m.GetVolterraSubnetIds(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["vpc_id"]; exists {

		vOpts := append(opts, db.WithValidateField("vpc_id"))
		if err := fv(ctx, m.GetVpcId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultPANAWSTypeValidator = func() *ValidatePANAWSType {
	v := &ValidatePANAWSType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAwsRegion := v.AwsRegionValidationRuleHandler
	rulesAwsRegion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAwsRegion(rulesAwsRegion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PANAWSType.aws_region: %s", err)
		panic(errMsg)
	}
	v.FldValidators["aws_region"] = vFn

	vrhAwsNamePrefix := v.AwsNamePrefixValidationRuleHandler
	rulesAwsNamePrefix := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAwsNamePrefix(rulesAwsNamePrefix)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PANAWSType.aws_name_prefix: %s", err)
		panic(errMsg)
	}
	v.FldValidators["aws_name_prefix"] = vFn

	vrhVpcId := v.VpcIdValidationRuleHandler
	rulesVpcId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVpcId(rulesVpcId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PANAWSType.vpc_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vpc_id"] = vFn

	vrhAdminUsername := v.AdminUsernameValidationRuleHandler
	rulesAdminUsername := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAdminUsername(rulesAdminUsername)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PANAWSType.admin_username: %s", err)
		panic(errMsg)
	}
	v.FldValidators["admin_username"] = vFn

	vrhDevices := v.DevicesValidationRuleHandler
	rulesDevices := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "2",
		"ves.io.schema.rules.repeated.min_items": "1",
	}
	vFn, err = vrhDevices(rulesDevices)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PANAWSType.devices: %s", err)
		panic(errMsg)
	}
	v.FldValidators["devices"] = vFn

	vrhSshKey := v.SshKeyValidationRuleHandler
	rulesSshKey := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSshKey(rulesSshKey)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PANAWSType.ssh_key: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ssh_key"] = vFn

	vrhVolterraSubnetIds := v.VolterraSubnetIdsValidationRuleHandler
	rulesVolterraSubnetIds := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVolterraSubnetIds(rulesVolterraSubnetIds)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PANAWSType.volterra_subnet_ids: %s", err)
		panic(errMsg)
	}
	v.FldValidators["volterra_subnet_ids"] = vFn

	vrhSiteName := v.SiteNameValidationRuleHandler
	rulesSiteName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSiteName(rulesSiteName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PANAWSType.site_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_name"] = vFn

	return v
}()

func PANAWSTypeValidator() db.Validator {
	return DefaultPANAWSTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *PaloAltoServicesNodeType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *PaloAltoServicesNodeType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *PaloAltoServicesNodeType) DeepCopy() *PaloAltoServicesNodeType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &PaloAltoServicesNodeType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *PaloAltoServicesNodeType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *PaloAltoServicesNodeType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return PaloAltoServicesNodeTypeValidator().Validate(ctx, m, opts...)
}

type ValidatePaloAltoServicesNodeType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidatePaloAltoServicesNodeType) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidatePaloAltoServicesNodeType) InstanceTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for instance_type")
	}

	return validatorFn, nil
}

func (v *ValidatePaloAltoServicesNodeType) SliSubnetIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for sli_subnet_id")
	}

	return validatorFn, nil
}

func (v *ValidatePaloAltoServicesNodeType) MgmtSubnetCidrValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mgmt_subnet_cidr")
	}

	return validatorFn, nil
}

func (v *ValidatePaloAltoServicesNodeType) MgmtSubnetIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for mgmt_subnet_id")
	}

	return validatorFn, nil
}

func (v *ValidatePaloAltoServicesNodeType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*PaloAltoServicesNodeType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *PaloAltoServicesNodeType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["availability_zone"]; exists {

		vOpts := append(opts, db.WithValidateField("availability_zone"))
		if err := fv(ctx, m.GetAvailabilityZone(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["instance_type"]; exists {

		vOpts := append(opts, db.WithValidateField("instance_type"))
		if err := fv(ctx, m.GetInstanceType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mgmt_subnet_cidr"]; exists {

		vOpts := append(opts, db.WithValidateField("mgmt_subnet_cidr"))
		if err := fv(ctx, m.GetMgmtSubnetCidr(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["mgmt_subnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("mgmt_subnet_id"))
		if err := fv(ctx, m.GetMgmtSubnetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["sli_subnet_id"]; exists {

		vOpts := append(opts, db.WithValidateField("sli_subnet_id"))
		if err := fv(ctx, m.GetSliSubnetId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultPaloAltoServicesNodeTypeValidator = func() *ValidatePaloAltoServicesNodeType {
	v := &ValidatePaloAltoServicesNodeType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PaloAltoServicesNodeType.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhInstanceType := v.InstanceTypeValidationRuleHandler
	rulesInstanceType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhInstanceType(rulesInstanceType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PaloAltoServicesNodeType.instance_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["instance_type"] = vFn

	vrhSliSubnetId := v.SliSubnetIdValidationRuleHandler
	rulesSliSubnetId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSliSubnetId(rulesSliSubnetId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PaloAltoServicesNodeType.sli_subnet_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["sli_subnet_id"] = vFn

	vrhMgmtSubnetCidr := v.MgmtSubnetCidrValidationRuleHandler
	rulesMgmtSubnetCidr := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMgmtSubnetCidr(rulesMgmtSubnetCidr)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PaloAltoServicesNodeType.mgmt_subnet_cidr: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mgmt_subnet_cidr"] = vFn

	vrhMgmtSubnetId := v.MgmtSubnetIdValidationRuleHandler
	rulesMgmtSubnetId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhMgmtSubnetId(rulesMgmtSubnetId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for PaloAltoServicesNodeType.mgmt_subnet_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["mgmt_subnet_id"] = vFn

	return v
}()

func PaloAltoServicesNodeTypeValidator() db.Validator {
	return DefaultPaloAltoServicesNodeTypeValidator
}
