/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type KeyObservation struct {

	// Creation time (time stamp) of a key.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// Identification of a Master Key. The value 1 indicates a Default
	// Master Key, and the value 0 indicates a key.
	DefaultKeyFlag *string `json:"defaultKeyFlag,omitempty" tf:"default_key_flag,omitempty"`

	// ID of a user domain for the key.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The globally unique identifier for the key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Origin of a key. The default value is kms.
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

	// The total number of key rotations.
	RotationNumber *float64 `json:"rotationNumber,omitempty" tf:"rotation_number,omitempty"`
}

type KeyParameters struct {

	// Specifies whether the key is enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// Specifies the name of a KMS key.
	// +kubebuilder:validation:Required
	KeyAlias *string `json:"keyAlias" tf:"key_alias,omitempty"`

	// Specifies the description of a KMS key.
	// +kubebuilder:validation:Optional
	KeyDescription *string `json:"keyDescription,omitempty" tf:"key_description,omitempty"`

	// Specifies the duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 1096 days. Defaults to 7.
	// It only be used when delete a key.
	// +kubebuilder:validation:Optional
	PendingDays *string `json:"pendingDays,omitempty" tf:"pending_days,omitempty"`

	// Region where a key resides. Changing this creates a new key.
	// +kubebuilder:validation:Optional
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`

	// Specifies whether the key rotation is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	RotationEnabled *bool `json:"rotationEnabled,omitempty" tf:"rotation_enabled,omitempty"`

	// Specifies the key rotation interval. The valid value is range from 30 to 365,
	// defaults to 365.
	// +kubebuilder:validation:Optional
	RotationInterval *float64 `json:"rotationInterval,omitempty" tf:"rotation_interval,omitempty"`
}

// KeySpec defines the desired state of Key
type KeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyParameters `json:"forProvider"`
}

// KeyStatus defines the observed state of Key.
type KeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Key is the Schema for the Keys API. ""page_title: "flexibleengine_kms_key_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Key struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeySpec   `json:"spec"`
	Status            KeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyList contains a list of Keys
type KeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Key `json:"items"`
}

// Repository type metadata.
var (
	Key_Kind             = "Key"
	Key_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Key_Kind}.String()
	Key_KindAPIVersion   = Key_Kind + "." + CRDGroupVersion.String()
	Key_GroupVersionKind = CRDGroupVersion.WithKind(Key_Kind)
)

func init() {
	SchemeBuilder.Register(&Key{}, &KeyList{})
}
