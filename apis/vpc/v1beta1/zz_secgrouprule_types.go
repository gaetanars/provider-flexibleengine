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

type SecGroupRuleObservation struct {

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecGroupRuleParameters struct {

	// Specifies the supplementary information about the security group rule.
	// This parameter can contain a maximum of 255 characters and cannot contain angle brackets (< or >).
	// Changing this creates a new security group rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction of the rule, valid values are ingress
	// or egress. Changing this creates a new security group rule.
	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// The layer 3 protocol type, valid values are IPv4
	// or IPv6. Changing this creates a new security group rule.
	// +kubebuilder:validation:Required
	Ethertype *string `json:"ethertype" tf:"ethertype,omitempty"`

	// The higher part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	// +kubebuilder:validation:Optional
	PortRangeMax *float64 `json:"portRangeMax,omitempty" tf:"port_range_max,omitempty"`

	// The lower part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	// +kubebuilder:validation:Optional
	PortRangeMin *float64 `json:"portRangeMin,omitempty" tf:"port_range_min,omitempty"`

	// The layer 4 protocol type, valid values are following.
	// Changing this creates a new security group rule. This is required if you want to specify a port range.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// security group rule.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The remote group id, the value needs to be an
	// FlexibleEngine ID of a security group in the same tenant. Changing this creates
	// a new security group rule.
	// +crossplane:generate:reference:type=SecGroup
	// +kubebuilder:validation:Optional
	RemoteGroupID *string `json:"remoteGroupId,omitempty" tf:"remote_group_id,omitempty"`

	// Reference to a SecGroup to populate remoteGroupId.
	// +kubebuilder:validation:Optional
	RemoteGroupIDRef *v1.Reference `json:"remoteGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecGroup to populate remoteGroupId.
	// +kubebuilder:validation:Optional
	RemoteGroupIDSelector *v1.Selector `json:"remoteGroupIdSelector,omitempty" tf:"-"`

	// The remote CIDR, the value needs to be a valid
	// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
	// +kubebuilder:validation:Optional
	RemoteIPPrefix *string `json:"remoteIpPrefix,omitempty" tf:"remote_ip_prefix,omitempty"`

	// The security group ID the rule should belong
	// to. Changing this creates a new security group rule.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.SecGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// The resource ID in UUID format.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/iam/v1beta1.Project
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Reference to a Project in iam to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDRef *v1.Reference `json:"tenantIdRef,omitempty" tf:"-"`

	// Selector for a Project in iam to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDSelector *v1.Selector `json:"tenantIdSelector,omitempty" tf:"-"`
}

// SecGroupRuleSpec defines the desired state of SecGroupRule
type SecGroupRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecGroupRuleParameters `json:"forProvider"`
}

// SecGroupRuleStatus defines the observed state of SecGroupRule.
type SecGroupRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecGroupRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecGroupRule is the Schema for the SecGroupRules API. ""page_title: "flexibleengine_networking_secgroup_rule_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type SecGroupRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecGroupRuleSpec   `json:"spec"`
	Status            SecGroupRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecGroupRuleList contains a list of SecGroupRules
type SecGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecGroupRule `json:"items"`
}

// Repository type metadata.
var (
	SecGroupRule_Kind             = "SecGroupRule"
	SecGroupRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecGroupRule_Kind}.String()
	SecGroupRule_KindAPIVersion   = SecGroupRule_Kind + "." + CRDGroupVersion.String()
	SecGroupRule_GroupVersionKind = CRDGroupVersion.WithKind(SecGroupRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SecGroupRule{}, &SecGroupRuleList{})
}
