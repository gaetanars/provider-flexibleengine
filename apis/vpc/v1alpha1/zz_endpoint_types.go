/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EndpointObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PacketID *float64 `json:"packetId,omitempty" tf:"packet_id,omitempty"`

	PrivateDomainName *string `json:"privateDomainName,omitempty" tf:"private_domain_name,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type EndpointParameters struct {

	// +kubebuilder:validation:Optional
	EnableDNS *bool `json:"enableDns,omitempty" tf:"enable_dns,omitempty"`

	// +kubebuilder:validation:Optional
	EnableWhitelist *bool `json:"enableWhitelist,omitempty" tf:"enable_whitelist,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +crossplane:generate:reference:type=EndpointService
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a EndpointService to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a EndpointService to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=Vpc
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a Vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a Vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Whitelist []*string `json:"whitelist,omitempty" tf:"whitelist,omitempty"`
}

// EndpointSpec defines the desired state of Endpoint
type EndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointParameters `json:"forProvider"`
}

// EndpointStatus defines the observed state of Endpoint.
type EndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Endpoint is the Schema for the Endpoints API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec"`
	Status            EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointList contains a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

// Repository type metadata.
var (
	Endpoint_Kind             = "Endpoint"
	Endpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Endpoint_Kind}.String()
	Endpoint_KindAPIVersion   = Endpoint_Kind + "." + CRDGroupVersion.String()
	Endpoint_GroupVersionKind = CRDGroupVersion.WithKind(Endpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
