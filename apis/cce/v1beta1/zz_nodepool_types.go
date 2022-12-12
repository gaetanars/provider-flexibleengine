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

type DataVolumesObservation struct {
}

type DataVolumesParameters struct {

	// Specifies the disk expansion parameters in key/value pair format.
	// +kubebuilder:validation:Optional
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// Specifies the KMS key ID. This is used to encrypt the volume.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Specifies the disk size in GB.
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Specifies the disk type.
	// +kubebuilder:validation:Required
	Volumetype *string `json:"volumetype" tf:"volumetype,omitempty"`
}

type NodePoolObservation struct {

	// Billing mode of a node.
	BillingMode *float64 `json:"billingMode,omitempty" tf:"billing_mode,omitempty"`

	// Specifies a resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Node status information.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NodePoolParameters struct {

	// specify the name of the available partition (AZ).
	// Default value is random to create nodes in a random AZ in the node pool.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// ID of the cluster. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// Represents the data disk to be created.
	// The object structure is documented below. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Required
	DataVolumes []DataVolumesParameters `json:"dataVolumes" tf:"data_volumes,omitempty"`

	// +kubebuilder:validation:Optional
	ExtendParam map[string]*string `json:"extendParam,omitempty" tf:"extend_param,omitempty"`

	// Specifies the flavor id. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Required
	FlavorID *string `json:"flavorId" tf:"flavor_id,omitempty"`

	// Initial number of expected nodes in the node pool.
	// +kubebuilder:validation:Required
	InitialNodeCount *float64 `json:"initialNodeCount" tf:"initial_node_count,omitempty"`

	// Key pair name when logging in to select the key pair mode.
	// This parameter and password are alternative. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair,omitempty"`

	// Tags of a Kubernetes node, key/value pair format.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maximum number of nodes allowed if auto scaling is enabled.
	// +kubebuilder:validation:Optional
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// The maximum number of instances a node is allowed to create.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	MaxPods *float64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`

	// Minimum number of nodes allowed if auto scaling is enabled.
	// +kubebuilder:validation:Optional
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`

	// Node Pool Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Operating System of the node. The value can be EulerOS 2.5 and CentOS 7.6.
	// Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Os *string `json:"os,omitempty" tf:"os,omitempty"`

	// root password when logging in to select the password mode.
	// This parameter must be salted and alternative to key_pair. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Script required after the installation. The input value can be
	// a Base64 encoded string or not. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Postinstall *string `json:"postinstall,omitempty" tf:"postinstall,omitempty"`

	// Script required before installation. The input value can be
	// a Base64 encoded string or not. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Preinstall *string `json:"preinstall,omitempty" tf:"preinstall,omitempty"`

	// Weight of a node pool. A node pool with a higher weight has a higher priority during scaling.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// It corresponds to the system disk related configuration.
	// The object structure is documented below. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Required
	RootVolume []RootVolumeParameters `json:"rootVolume" tf:"root_volume,omitempty"`

	// Interval between two scaling operations, in minutes.
	// +kubebuilder:validation:Optional
	ScaleDownCooldownTime *float64 `json:"scaleDownCooldownTime,omitempty" tf:"scale_down_cooldown_time,omitempty"`

	// Whether to enable auto scaling. If Autoscaler is enabled, install the autoscaler
	// add-on to use the auto scaling feature.
	// +kubebuilder:validation:Optional
	ScaleEnable *bool `json:"scaleEnable,omitempty" tf:"scale_enable,omitempty"`

	// +kubebuilder:validation:Optional
	ScallEnable *bool `json:"scallEnable,omitempty" tf:"scall_enable,omitempty"`

	// The ID of the subnet to which the NIC belongs.
	// Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Tags of a VM node, key/value pair format.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// You can add taints to created nodes to configure anti-affinity.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	Taints []TaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// Node Pool type. Possible values are: "vm" and "ElasticBMS".
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RootVolumeObservation struct {
}

type RootVolumeParameters struct {

	// Specifies the disk expansion parameters in key/value pair format.
	// +kubebuilder:validation:Optional
	ExtendParams map[string]*string `json:"extendParams,omitempty" tf:"extend_params,omitempty"`

	// Specifies the disk size in GB.
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Specifies the disk type.
	// +kubebuilder:validation:Required
	Volumetype *string `json:"volumetype" tf:"volumetype,omitempty"`
}

type TaintsObservation struct {
}

type TaintsParameters struct {

	// Available options are NoSchedule, PreferNoSchedule and NoExecute.
	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// A key must contain 1 to 63 characters starting with a letter or digit.
	// Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed.
	// A DNS subdomain name can be used as the prefix of a key.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// A value must start with a letter or digit and can contain a maximum of 63 characters,
	// including letters, digits, hyphens (-), underscores (_), and periods (.).
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// NodePoolSpec defines the desired state of NodePool
type NodePoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodePoolParameters `json:"forProvider"`
}

// NodePoolStatus defines the observed state of NodePool.
type NodePoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodePoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodePool is the Schema for the NodePools API. ""page_title: "flexibleengine_cce_node_pool_v3"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type NodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodePoolSpec   `json:"spec"`
	Status            NodePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodePoolList contains a list of NodePools
type NodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePool `json:"items"`
}

// Repository type metadata.
var (
	NodePool_Kind             = "NodePool"
	NodePool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodePool_Kind}.String()
	NodePool_KindAPIVersion   = NodePool_Kind + "." + CRDGroupVersion.String()
	NodePool_GroupVersionKind = CRDGroupVersion.WithKind(NodePool_Kind)
)

func init() {
	SchemeBuilder.Register(&NodePool{}, &NodePoolList{})
}
