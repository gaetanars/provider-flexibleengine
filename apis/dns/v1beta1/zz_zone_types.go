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

type RouterObservation struct {
}

type RouterParameters struct {

	// The VPC UUID.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// Reference to a VPC in vpc to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDRef *v1.Reference `json:"routerIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDSelector *v1.Selector `json:"routerIdSelector,omitempty" tf:"-"`

	// The region of the VPC. Defaults to the region.
	// +kubebuilder:validation:Optional
	RouterRegion *string `json:"routerRegion,omitempty" tf:"router_region,omitempty"`
}

type ZoneObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of master DNS servers.
	Masters []*string `json:"masters,omitempty" tf:"masters,omitempty"`
}

type ZoneParameters struct {

	// A description of the zone.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The email contact for the zone record.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The name of the zone. Note the . at the end of the name.
	// Changing this creates a new DNS zone.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The region in which to create the DNS zone.
	// If omitted, the region argument of the provider is used.
	// Changing this creates a new DNS zone.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Router configuration block which is required if zone_type is private.
	// The router structure is documented below.
	// +kubebuilder:validation:Optional
	Router []RouterParameters `json:"router,omitempty" tf:"router,omitempty"`

	// The time to live (TTL) of the zone.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The key/value pairs to associate with the zone.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of additional options. Changing this creates a
	// new DNS zone.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// The type of zone. Can either be public or private.
	// Changing this creates a new DNS zone.
	// +kubebuilder:validation:Optional
	ZoneType *string `json:"zoneType,omitempty" tf:"zone_type,omitempty"`
}

// ZoneSpec defines the desired state of Zone
type ZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneParameters `json:"forProvider"`
}

// ZoneStatus defines the observed state of Zone.
type ZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Zone is the Schema for the Zones API. ""page_title: "flexibleengine_dns_zone_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZoneSpec   `json:"spec"`
	Status            ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

// Repository type metadata.
var (
	Zone_Kind             = "Zone"
	Zone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Zone_Kind}.String()
	Zone_KindAPIVersion   = Zone_Kind + "." + CRDGroupVersion.String()
	Zone_GroupVersionKind = CRDGroupVersion.WithKind(Zone_Kind)
)

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}
