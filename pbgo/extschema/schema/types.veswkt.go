// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package schema

import (
	"fmt"
	"strings"
	"time"

	google_protobuf "github.com/gogo/protobuf/types"
	"github.com/google/uuid"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/db/sro"
	"gopkg.volterra.us/stdlib/util/secret"
)

// DeepCopy2 allows ObjectRefType to satisfy DirectRef
func (m *ObjectRefType) DeepCopy2() db.DirectRef {
	return m.DeepCopy()
}

// ToStatusObjectConditions takes in []*ConditionType returns a []sro.StatusObjectCondition form
func ToStatusObjectConditions(conditions []*ConditionType) []sro.StatusObjectCondition {
	soc := []sro.StatusObjectCondition{}
	for _, c := range conditions {
		soc = append(soc, c)
	}
	return soc
}

// FromStatusObjectConditions takes in []sro.StatusObjectCondition form and returns []*ConditionType
func FromStatusObjectConditions(socSet []sro.StatusObjectCondition) []*ConditionType {
	conditions := []*ConditionType{}
	for _, soc := range socSet {
		conditions = append(conditions, &ConditionType{
			Type:           soc.GetType(),
			Status:         soc.GetStatus(),
			Reason:         soc.GetReason(),
			LastUpdateTime: soc.GetLastUpdateTime(),
		})
	}
	return conditions
}

// ShouldPublish returns true if m.Publish is set to STATUS_PUBLISH
func (m *StatusMetaType) ShouldPublish() bool {
	if m == nil {
		return false
	}
	return m.Publish == STATUS_PUBLISH
}

// DoPublish sets m.Publish to STATUS_PUBLISH
func (m *StatusMetaType) DoPublish() {
	if m == nil {
		return
	}
	m.Publish = STATUS_PUBLISH
}

// DoNotPublish sets m.Publish to STATUS_DO_NOT_PUBLISH
func (m *StatusMetaType) DoNotPublish() {
	if m == nil {
		return
	}
	m.Publish = STATUS_DO_NOT_PUBLISH
}

// NewStatusMetaType returns a new status metadata with uid and creation timestamp already filled
func NewStatusMetaType() *StatusMetaType {
	timestamp, _ := google_protobuf.TimestampProto(time.Now())
	return &StatusMetaType{
		Uid:               uuid.New().String(),
		CreationTimestamp: timestamp,
		Publish:           STATUS_DO_NOT_PUBLISH,
	}
}

// GetOwnerViewRef returns the owner view if it is present in the object
func (m *SystemObjectMetaType) GetOwnerViewRef() string {
	ref := m.GetOwnerView()
	if ref == nil {
		return ""
	}
	return fmt.Sprintf("%s:%s:%s:%s", ref.Kind, ref.Uid, ref.Namespace, ref.Name)
}

// GetOwnerViewPublic returns public representation of the owner view if it is present in the object
func (m *SystemObjectMetaType) GetOwnerViewPublic() string {
	ref := m.GetOwnerView()
	if ref == nil {
		return ""
	}
	return fmt.Sprintf("%s:%s/%s", ref.Kind, ref.Namespace, ref.Name)
}

// SetOwnerView creates a reference to owner view.
// The argument owner should be colon separated string containing kind, uid, namespace and name
func (m *SystemObjectMetaType) SetOwnerView(owner string) error {
	parts := strings.Split(owner, ":")
	if len(parts) != 4 {
		return fmt.Errorf("Expected 4 parts in owner, but got %s", owner)
	}
	m.OwnerView = &ViewRefType{Kind: parts[0], Uid: parts[1], Namespace: parts[2], Name: parts[3]}
	return nil
}

// GetNamespaceRef returns values of tenant and namespace an object is present in
func (m *SystemObjectMetaType) GetNamespaceRef() (string, string) {
	if m == nil || m.Namespace == nil {
		return "", ""
	}
	return m.Namespace[0].Tenant, m.Namespace[0].Name
}

// SetNamespaceRef creates a reference to namespace object
func (m *SystemObjectMetaType) SetNamespaceRef(tenant, nsName string) error {
	m.Namespace = []*ObjectRefType{
		{Kind: "namespace", Tenant: tenant, Name: nsName},
	}
	return nil
}

// GetType() implements stdlib/util/secret/Secret interface
func (m *SecretType) GetType() secret.Type {
	if m == nil {
		return secret.TypeInvalid
	}
	switch m.SecretInfoOneof.(type) {
	case *SecretType_BlindfoldSecretInfo:
		return secret.TypeBlindfold
	case *SecretType_VaultSecretInfo:
		return secret.TypeVault
	case *SecretType_ClearSecretInfo:
		return secret.TypeClear
	case *SecretType_WingmanSecretInfo:
		return secret.TypeWingman
	default:
		return secret.TypeInvalid
	}
}

// GetBlindfoldType() implements stdlib/util/secret/Secret interface
func (m *SecretType) GetBlindfoldType() *secret.BlindfoldType {
	bfInfo := m.GetBlindfoldSecretInfo()
	if bfInfo == nil {
		return nil
	}
	return &secret.BlindfoldType{
		DecryptionProvider: bfInfo.DecryptionProvider,
		StoreProvider:      bfInfo.StoreProvider,
		Location:           bfInfo.Location,
	}
}

// GetVaultType() implements stdlib/util/secret/Secret interface
func (m *SecretType) GetVaultType() *secret.VaultType {
	vInfo := m.GetVaultSecretInfo()
	if vInfo == nil {
		return nil
	}
	return &secret.VaultType{
		Provider:     vInfo.Provider,
		Location:     vInfo.Location,
		Key:          vInfo.Key,
		Version:      vInfo.Version,
		EncodingType: vInfo.SecretEncoding.String(),
	}
}

// GetClearType() implements stdlib/util/secret/Secret interface
func (m *SecretType) GetClearType() *secret.ClearType {
	cInfo := m.GetClearSecretInfo()
	if cInfo == nil {
		return nil
	}
	return &secret.ClearType{
		Provider: cInfo.Provider,
		URL:      cInfo.Url,
	}
}

// GetWingmanType() implements stdlib/util/secret/Secret interface
func (m *SecretType) GetWingmanType() *secret.WingmanType {
	wInfo := m.GetWingmanSecretInfo()
	if wInfo == nil {
		return nil
	}
	return &secret.WingmanType{
		Name: wInfo.Name,
	}
}

// GetEncodingType() implements stdlib/util/secret/Secret interface
func (m *SecretType) GetEncodingType() string {
	return m.GetSecretEncodingType().String()
}

// GetReEncryptedBlindfoldSecret() gets the re-encrypted blindfold secret in the blindfold_secret_info_internal field
func (m *SecretType) GetReEncryptedBlindfoldSecret() *secret.BlindfoldType {
	bfInfoInt := m.GetBlindfoldSecretInfoInternal()
	if bfInfoInt == nil {
		return nil
	}
	return &secret.BlindfoldType{
		DecryptionProvider: bfInfoInt.DecryptionProvider,
		StoreProvider:      bfInfoInt.StoreProvider,
		Location:           bfInfoInt.Location,
	}
}

// SetReEncryptedBlindfoldSecret() sets the re-encrypted blindfold secret with different
// policy-id into blindfold_secret_info_internal field.
func (m *SecretType) SetReEncryptedBlindfoldSecret(reEnc *secret.BlindfoldType) error {
	if reEnc == nil {
		m.BlindfoldSecretInfoInternal = nil
	} else {
		m.BlindfoldSecretInfoInternal = &BlindfoldSecretInfoType{
			DecryptionProvider: reEnc.DecryptionProvider,
			StoreProvider:      reEnc.StoreProvider,
			Location:           reEnc.Location,
		}
	}
	return nil
}
