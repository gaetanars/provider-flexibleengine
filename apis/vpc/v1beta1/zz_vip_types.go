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

type VipObservation struct {

	// The device owner of the vip.
	DeviceOwner *string `json:"deviceOwner,omitempty" tf:"device_owner,omitempty"`

	// The ID of the vip.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The status of vip.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The tenant ID of the vip.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type VipParameters struct {

	// IP address desired in the subnet for this vip.
	// If you don't specify ip_address, an available IP address from
	// the specified subnet will be allocated to this vip.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// A unique name for the vip.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the network to attach the vip to.
	// Changing this creates a new vip.
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Subnet in which to allocate IP address for this vip.
	// Changing this creates a new vip.
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.NetworkingSubnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a NetworkingSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a NetworkingSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// VipSpec defines the desired state of Vip
type VipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VipParameters `json:"forProvider"`
}

// VipStatus defines the observed state of Vip.
type VipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vip is the Schema for the Vips API. ""page_title: "flexibleengine_networking_vip_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Vip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VipSpec   `json:"spec"`
	Status            VipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VipList contains a list of Vips
type VipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vip `json:"items"`
}

// Repository type metadata.
var (
	Vip_Kind             = "Vip"
	Vip_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vip_Kind}.String()
	Vip_KindAPIVersion   = Vip_Kind + "." + CRDGroupVersion.String()
	Vip_GroupVersionKind = CRDGroupVersion.WithKind(Vip_Kind)
)

func init() {
	SchemeBuilder.Register(&Vip{}, &VipList{})
}
