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

type BackupStrategyObservation struct {
}

type BackupStrategyParameters struct {

	// Specifies the number of days to retain the generated backup files.
	// The value range is from 0 to 732.
	// +kubebuilder:validation:Required
	KeepDays *float64 `json:"keepDays" tf:"keep_days,omitempty"`

	// Specifies the backup time window. Automated backups will be triggered
	// during the backup time window. The value cannot be empty. It must be a valid value in the
	// "hh:mm-HH:MM" format. The current time is in the UTC format.
	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type DatastoreObservation struct {
}

type DatastoreParameters struct {

	// Specifies the storage engine of the DB instance. Only wiredTiger is supported now.
	// +kubebuilder:validation:Optional
	StorageEngine *string `json:"storageEngine,omitempty" tf:"storage_engine,omitempty"`

	// Specifies the DB engine. Only DDS-Community is supported now.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the DB instance version. Only 3.4 and 4.0 are supported now.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type FlavorObservation struct {
}

type FlavorParameters struct {

	// Specifies the node quantity. Valid value:
	// +kubebuilder:validation:Required
	Num *float64 `json:"num" tf:"num,omitempty"`

	// Specifies the disk size. The value must be a multiple of 10. The unit is GB. This parameter
	// is mandatory for nodes except mongos and invalid for mongos.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// Specifies the resource specification code. Valid values:
	// +kubebuilder:validation:Required
	SpecCode *string `json:"specCode" tf:"spec_code,omitempty"`

	// Specifies the disk type. Valid value: ULTRAHIGH which indicates the type SSD.
	// +kubebuilder:validation:Optional
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// Specifies the node type. Valid value: mongos, shard, config, replica.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceObservation struct {

	// Indicates the DB Administator name.
	DBUsername *string `json:"dbUsername,omitempty" tf:"db_username,omitempty"`

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the instance nodes information. Structure is documented below.
	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// Indicates the database port number. The port range is 2100 to 9500.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Indicates the the DB instance status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type InstanceParameters struct {

	// Specifies the ID of the availability zone. Changing
	// this creates a new instance.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// Specifies the advanced backup policy. The structure is
	// described below. Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	BackupStrategy []BackupStrategyParameters `json:"backupStrategy,omitempty" tf:"backup_strategy,omitempty"`

	// Specifies database information. The structure is described
	// below. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	Datastore []DatastoreParameters `json:"datastore" tf:"datastore,omitempty"`

	// Specifies the disk encryption ID of the instance.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	DiskEncryptionIDSecretRef *v1.SecretKeySelector `json:"diskEncryptionIdSecretRef,omitempty" tf:"-"`

	// Specifies the flavors information. The structure is described below.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Required
	Flavor []FlavorParameters `json:"flavor" tf:"flavor,omitempty"`

	// Specifies the mode of the database instance. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// Specifies the DB instance name. The DB instance name of the same
	// type is unique in the same tenant. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the Administrator password of the database instance.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Specifies the region of the DDS instance. Changing this creates
	// a new instance.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies whether to enable or disable SSL. Defaults to true.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	// Specifies the security group ID of the DDS instance.
	// Changing this creates a new instance.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecurityGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Specifies the ID of the VPC Subnet. Changing this creates a new instance.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools.ExtractorParamPathfunc(true, "id")
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The key/value pairs to associate with the DDS instance.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the VPC ID. Changing this creates a new instance.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type NodesObservation struct {

	// Indicates the node ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates the node name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates the private IP address of a node. This parameter is valid only for
	// mongos nodes, replica set instances, and single node instances.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Indicates the EIP that has been bound on a node. This parameter is valid only for
	// mongos nodes of cluster instances, primary nodes and secondary nodes of replica set instances,
	// and single node instances.
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Indicates the node role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Indicates the node status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Indicates the node type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NodesParameters struct {
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. ""page_title: "flexibleengine_dds_instance_v3"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
