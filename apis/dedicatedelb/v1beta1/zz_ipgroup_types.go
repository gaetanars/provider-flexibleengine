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

type IPGroupObservation struct {

	// The uuid of the ip group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IPGroupParameters struct {

	// Specifies the description of the ip group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The uuid of the ip group.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies an array of one or more ip addresses. The ip_list object structure is
	// documented below.
	// +kubebuilder:validation:Required
	IPList []IPListParameters `json:"ipList" tf:"ip_list,omitempty"`

	// Specifies the name of the ip group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The region in which to create the ip group resource. If omitted, the
	// provider-level region will be used. Changing this creates a new ip group.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type IPListObservation struct {
}

type IPListParameters struct {

	// Human-readable description for the ip.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IP address or CIDR block.
	// +kubebuilder:validation:Required
	IP *string `json:"ip" tf:"ip,omitempty"`
}

// IPGroupSpec defines the desired state of IPGroup
type IPGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IPGroupParameters `json:"forProvider"`
}

// IPGroupStatus defines the observed state of IPGroup.
type IPGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IPGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IPGroup is the Schema for the IPGroups API. ""page_title: "flexibleengine_elb_ipgroup"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type IPGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPGroupSpec   `json:"spec"`
	Status            IPGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPGroupList contains a list of IPGroups
type IPGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPGroup `json:"items"`
}

// Repository type metadata.
var (
	IPGroup_Kind             = "IPGroup"
	IPGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IPGroup_Kind}.String()
	IPGroup_KindAPIVersion   = IPGroup_Kind + "." + CRDGroupVersion.String()
	IPGroup_GroupVersionKind = CRDGroupVersion.WithKind(IPGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&IPGroup{}, &IPGroupList{})
}
