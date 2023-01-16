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

type AnalysisCoreNodesObservation struct {

	// The host list of this nodes group in the cluster.
	HostIps []*string `json:"hostIps,omitempty" tf:"host_ips,omitempty"`
}

type AnalysisCoreNodesParameters struct {

	// Specifies the roles deployed in a node group. This argument is mandatory
	// when the cluster type is CUSTOM. Each character string represents a role expression.
	// +kubebuilder:validation:Optional
	AssignedRoles []*string `json:"assignedRoles,omitempty" tf:"assigned_roles,omitempty"`

	// Specifies the data disk number of the nodes. The number configuration
	// of each node are as follows:
	// +kubebuilder:validation:Required
	DataVolumeCount *float64 `json:"dataVolumeCount" tf:"data_volume_count,omitempty"`

	// Specifies the data disk size of the nodes,in GB. The value range is 10
	// to 32768. Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	DataVolumeSize *float64 `json:"dataVolumeSize,omitempty" tf:"data_volume_size,omitempty"`

	// Specifies the data disk flavor of the nodes.
	// Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// The following disk types are supported:
	// +kubebuilder:validation:Optional
	DataVolumeType *string `json:"dataVolumeType,omitempty" tf:"data_volume_type,omitempty"`

	// Specifies the instance specifications for each nodes in node group.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Specifies the number of nodes for the node group.
	// Only the core group and task group updations are allowed. The number of nodes after scaling cannot be
	// less than the number of nodes originally created.
	// +kubebuilder:validation:Required
	NodeNumber *float64 `json:"nodeNumber" tf:"node_number,omitempty"`

	// Specifies the system disk size of the nodes. Changing this will create
	// a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeSize *float64 `json:"rootVolumeSize" tf:"root_volume_size,omitempty"`

	// Specifies the system disk flavor of the nodes. Changing this will
	// create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeType *string `json:"rootVolumeType" tf:"root_volume_type,omitempty"`
}

type AnalysisTaskNodesObservation struct {

	// The host list of this nodes group in the cluster.
	HostIps []*string `json:"hostIps,omitempty" tf:"host_ips,omitempty"`
}

type AnalysisTaskNodesParameters struct {

	// Specifies the roles deployed in a node group. This argument is mandatory
	// when the cluster type is CUSTOM. Each character string represents a role expression.
	// +kubebuilder:validation:Optional
	AssignedRoles []*string `json:"assignedRoles,omitempty" tf:"assigned_roles,omitempty"`

	// Specifies the data disk number of the nodes. The number configuration
	// of each node are as follows:
	// +kubebuilder:validation:Required
	DataVolumeCount *float64 `json:"dataVolumeCount" tf:"data_volume_count,omitempty"`

	// Specifies the data disk size of the nodes,in GB. The value range is 10
	// to 32768. Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	DataVolumeSize *float64 `json:"dataVolumeSize,omitempty" tf:"data_volume_size,omitempty"`

	// Specifies the data disk flavor of the nodes.
	// Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// The following disk types are supported:
	// +kubebuilder:validation:Optional
	DataVolumeType *string `json:"dataVolumeType,omitempty" tf:"data_volume_type,omitempty"`

	// Specifies the instance specifications for each nodes in node group.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Specifies the number of nodes for the node group.
	// Only the core group and task group updations are allowed. The number of nodes after scaling cannot be
	// less than the number of nodes originally created.
	// +kubebuilder:validation:Required
	NodeNumber *float64 `json:"nodeNumber" tf:"node_number,omitempty"`

	// Specifies the system disk size of the nodes. Changing this will create
	// a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeSize *float64 `json:"rootVolumeSize" tf:"root_volume_size,omitempty"`

	// Specifies the system disk flavor of the nodes. Changing this will
	// create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeType *string `json:"rootVolumeType" tf:"root_volume_type,omitempty"`
}

type ClusterObservation struct {

	// Specifies a list of the informations about the analysis core nodes in the
	// MRS cluster.
	// The nodes object structure of the analysis_core_nodes is documented below.
	// +kubebuilder:validation:Optional
	AnalysisCoreNodes []AnalysisCoreNodesObservation `json:"analysisCoreNodes,omitempty" tf:"analysis_core_nodes,omitempty"`

	// Specifies a list of the informations about the analysis task nodes in the
	// MRS cluster.
	// The nodes object structure of the analysis_task_nodes is documented below.
	// +kubebuilder:validation:Optional
	AnalysisTaskNodes []AnalysisTaskNodesObservation `json:"analysisTaskNodes,omitempty" tf:"analysis_task_nodes,omitempty"`

	// The charging start time which is the start time of billing, in RFC-3339 format.
	ChargingStartTime *string `json:"chargingStartTime,omitempty" tf:"charging_start_time,omitempty"`

	Components []ComponentsObservation `json:"components,omitempty" tf:"components,omitempty"`

	// The cluster creation time, in RFC-3339 format.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Specifies a list of the informations about the custom nodes in the MRS cluster.
	// The nodes object structure of the custom_nodes is documented below.
	// Unlike other nodes, it needs to specify group_name.
	// +kubebuilder:validation:Optional
	CustomNodes []CustomNodesObservation `json:"customNodes,omitempty" tf:"custom_nodes,omitempty"`

	// The cluster ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address of the master node.
	MasterNodeIP *string `json:"masterNodeIp,omitempty" tf:"master_node_ip,omitempty"`

	// Specifies a list of the informations about the master nodes in the
	// MRS cluster.
	// The nodes object structure of the master_nodes is documented below.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	MasterNodes []MasterNodesObservation `json:"masterNodes,omitempty" tf:"master_nodes,omitempty"`

	// The preferred private IP address of the master node.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// The cluster state, which include: running, frozen, abnormal and failed.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specifies a list of the informations about the streaming core nodes in the
	// MRS cluster.
	// The nodes object structure of the streaming_core_nodes is documented below.
	// +kubebuilder:validation:Optional
	StreamingCoreNodes []StreamingCoreNodesObservation `json:"streamingCoreNodes,omitempty" tf:"streaming_core_nodes,omitempty"`

	// Specifies a list of the informations about the streaming task nodes in the
	// MRS cluster.
	// The nodes object structure of the streaming_task_nodes is documented below.
	// +kubebuilder:validation:Optional
	StreamingTaskNodes []StreamingTaskNodesObservation `json:"streamingTaskNodes,omitempty" tf:"streaming_task_nodes,omitempty"`

	// The total number of nodes deployed in the cluster.
	TotalNodeNumber *float64 `json:"totalNodeNumber,omitempty" tf:"total_node_number,omitempty"`

	// The cluster update time, in RFC-3339 format.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ClusterParameters struct {

	// Specifies a list of the informations about the analysis core nodes in the
	// MRS cluster.
	// The nodes object structure of the analysis_core_nodes is documented below.
	// +kubebuilder:validation:Optional
	AnalysisCoreNodes []AnalysisCoreNodesParameters `json:"analysisCoreNodes,omitempty" tf:"analysis_core_nodes,omitempty"`

	// Specifies a list of the informations about the analysis task nodes in the
	// MRS cluster.
	// The nodes object structure of the analysis_task_nodes is documented below.
	// +kubebuilder:validation:Optional
	AnalysisTaskNodes []AnalysisTaskNodesParameters `json:"analysisTaskNodes,omitempty" tf:"analysis_task_nodes,omitempty"`

	// Specifies the availability zone in which to create the cluster.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// Specifies the list of component names.
	// Changing this will create a new MRS cluster resource. The supported components are as follows:
	// +kubebuilder:validation:Required
	ComponentList []*string `json:"componentList" tf:"component_list,omitempty"`

	// Specifies a list of the informations about the custom nodes in the MRS cluster.
	// The nodes object structure of the custom_nodes is documented below.
	// Unlike other nodes, it needs to specify group_name.
	// +kubebuilder:validation:Optional
	CustomNodes []CustomNodesParameters `json:"customNodes,omitempty" tf:"custom_nodes,omitempty"`

	// Specifies the EIP ID which bound to the MRS cluster.
	// The EIP must have been created and must be in the same region as the cluster.
	// Changing this will create a new MRS cluster resource.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/eip/v1beta1.EIP
	// +kubebuilder:validation:Optional
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	// Reference to a EIP in eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDRef *v1.Reference `json:"eipIdRef,omitempty" tf:"-"`

	// Selector for a EIP in eip to populate eipId.
	// +kubebuilder:validation:Optional
	EIPIDSelector *v1.Selector `json:"eipIdSelector,omitempty" tf:"-"`

	// Specifies whether logs are collected when cluster installation fails.
	// Default to true. If log_collection set true, the OBS buckets will be created and only used to collect logs that
	// record MRS cluster creation failures. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	LogCollection *bool `json:"logCollection,omitempty" tf:"log_collection,omitempty"`

	// Specifies the administrator password, which is used to login to
	// the cluster management page. The password can contain 8 to 26 charactors and cannot be the username or the username
	// spelled backwards. The password must contain lowercase letters, uppercase letters, digits, spaces and the special
	// characters: !?,.:-_{}[]@$^+=/. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	ManagerAdminPwdSecretRef v1.SecretKeySelector `json:"managerAdminPwdSecretRef" tf:"-"`

	// Specifies a list of the informations about the master nodes in the
	// MRS cluster.
	// The nodes object structure of the master_nodes is documented below.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	MasterNodes []MasterNodesParameters `json:"masterNodes" tf:"master_nodes,omitempty"`

	// Specifies the name of the MRS cluster. The name can contain 2 to 64
	// characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
	// MRS cluster resource.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the name of a key pair, which is used to login to the each
	// nodes(ECSs). Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	NodeKeyPairSecretRef v1.SecretKeySelector `json:"nodeKeyPairSecretRef" tf:"-"`

	// Specifies the EIP address which bound to the MRS cluster.
	// The EIP must have been created and must be in the same region as the cluster.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// The region in which to create the MRS cluster resource. If omitted, the
	// provider-level region will be used. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies whether the running mode of the MRS cluster is secure,
	// default to true.
	// +kubebuilder:validation:Optional
	SafeMode *bool `json:"safeMode,omitempty" tf:"safe_mode,omitempty"`

	// Specifies an array of one or more security group ID to attach to the
	// MRS cluster. If using the specified security group, the group need to open the specified port (9022) rules.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Specifies a list of the informations about the streaming core nodes in the
	// MRS cluster.
	// The nodes object structure of the streaming_core_nodes is documented below.
	// +kubebuilder:validation:Optional
	StreamingCoreNodes []StreamingCoreNodesParameters `json:"streamingCoreNodes,omitempty" tf:"streaming_core_nodes,omitempty"`

	// Specifies a list of the informations about the streaming task nodes in the
	// MRS cluster.
	// The nodes object structure of the streaming_task_nodes is documented below.
	// +kubebuilder:validation:Optional
	StreamingTaskNodes []StreamingTaskNodesParameters `json:"streamingTaskNodes,omitempty" tf:"streaming_task_nodes,omitempty"`

	// Specifies the ID of the VPC Subnet which bound to the MRS cluster.
	// Changing this will create a new MRS cluster resource.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FrangipaneTeam/provider-flexibleengine/config/common.IDExtractor()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Specifies the key/value pairs to associate with the cluster.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the template used for node deployment when the cluster type is
	// CUSTOM.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// Specifies the type of the MRS cluster. The valid values are ANALYSIS,
	// STREAMING, MIXED and CUSTOM (supported in MRS 3.x only), default to ANALYSIS.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the ID of the VPC which bound to the MRS cluster.
	// Changing this will create a new MRS cluster resource.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Specifies the MRS cluster version. Currently, MRS 1.8.9,
	// MRS 2.0.1, and MRS 3.1.0-LTS.1 are supported. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type ComponentsObservation struct {

	// Component description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The cluster ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the MRS cluster. The name can contain 2 to 64
	// characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
	// MRS cluster resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the MRS cluster version. Currently, MRS 1.8.9,
	// MRS 2.0.1, and MRS 3.1.0-LTS.1 are supported. Changing this will create a new MRS cluster resource.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ComponentsParameters struct {
}

type CustomNodesObservation struct {

	// The host list of this nodes group in the cluster.
	HostIps []*string `json:"hostIps,omitempty" tf:"host_ips,omitempty"`
}

type CustomNodesParameters struct {

	// Specifies the roles deployed in a node group. This argument is mandatory
	// when the cluster type is CUSTOM. Each character string represents a role expression.
	// +kubebuilder:validation:Optional
	AssignedRoles []*string `json:"assignedRoles,omitempty" tf:"assigned_roles,omitempty"`

	// Specifies the data disk number of the nodes. The number configuration
	// of each node are as follows:
	// +kubebuilder:validation:Required
	DataVolumeCount *float64 `json:"dataVolumeCount" tf:"data_volume_count,omitempty"`

	// Specifies the data disk size of the nodes,in GB. The value range is 10
	// to 32768. Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	DataVolumeSize *float64 `json:"dataVolumeSize,omitempty" tf:"data_volume_size,omitempty"`

	// Specifies the data disk flavor of the nodes.
	// Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// The following disk types are supported:
	// +kubebuilder:validation:Optional
	DataVolumeType *string `json:"dataVolumeType,omitempty" tf:"data_volume_type,omitempty"`

	// Specifies the instance specifications for each nodes in node group.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Specifies the name of nodes for the node group.
	// This argument is mandatory when the cluster type is CUSTOM.
	// +kubebuilder:validation:Required
	GroupName *string `json:"groupName" tf:"group_name,omitempty"`

	// Specifies the number of nodes for the node group.
	// Only the core group and task group updations are allowed. The number of nodes after scaling cannot be
	// less than the number of nodes originally created.
	// +kubebuilder:validation:Required
	NodeNumber *float64 `json:"nodeNumber" tf:"node_number,omitempty"`

	// Specifies the system disk size of the nodes. Changing this will create
	// a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeSize *float64 `json:"rootVolumeSize" tf:"root_volume_size,omitempty"`

	// Specifies the system disk flavor of the nodes. Changing this will
	// create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeType *string `json:"rootVolumeType" tf:"root_volume_type,omitempty"`
}

type MasterNodesObservation struct {

	// The host list of this nodes group in the cluster.
	HostIps []*string `json:"hostIps,omitempty" tf:"host_ips,omitempty"`
}

type MasterNodesParameters struct {

	// Specifies the roles deployed in a node group. This argument is mandatory
	// when the cluster type is CUSTOM. Each character string represents a role expression.
	// +kubebuilder:validation:Optional
	AssignedRoles []*string `json:"assignedRoles,omitempty" tf:"assigned_roles,omitempty"`

	// Specifies the data disk number of the nodes. The number configuration
	// of each node are as follows:
	// +kubebuilder:validation:Required
	DataVolumeCount *float64 `json:"dataVolumeCount" tf:"data_volume_count,omitempty"`

	// Specifies the data disk size of the nodes,in GB. The value range is 10
	// to 32768. Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	DataVolumeSize *float64 `json:"dataVolumeSize,omitempty" tf:"data_volume_size,omitempty"`

	// Specifies the data disk flavor of the nodes.
	// Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// The following disk types are supported:
	// +kubebuilder:validation:Optional
	DataVolumeType *string `json:"dataVolumeType,omitempty" tf:"data_volume_type,omitempty"`

	// Specifies the instance specifications for each nodes in node group.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Specifies the number of nodes for the node group.
	// Only the core group and task group updations are allowed. The number of nodes after scaling cannot be
	// less than the number of nodes originally created.
	// +kubebuilder:validation:Required
	NodeNumber *float64 `json:"nodeNumber" tf:"node_number,omitempty"`

	// Specifies the system disk size of the nodes. Changing this will create
	// a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeSize *float64 `json:"rootVolumeSize" tf:"root_volume_size,omitempty"`

	// Specifies the system disk flavor of the nodes. Changing this will
	// create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeType *string `json:"rootVolumeType" tf:"root_volume_type,omitempty"`
}

type StreamingCoreNodesObservation struct {

	// The host list of this nodes group in the cluster.
	HostIps []*string `json:"hostIps,omitempty" tf:"host_ips,omitempty"`
}

type StreamingCoreNodesParameters struct {

	// Specifies the roles deployed in a node group. This argument is mandatory
	// when the cluster type is CUSTOM. Each character string represents a role expression.
	// +kubebuilder:validation:Optional
	AssignedRoles []*string `json:"assignedRoles,omitempty" tf:"assigned_roles,omitempty"`

	// Specifies the data disk number of the nodes. The number configuration
	// of each node are as follows:
	// +kubebuilder:validation:Required
	DataVolumeCount *float64 `json:"dataVolumeCount" tf:"data_volume_count,omitempty"`

	// Specifies the data disk size of the nodes,in GB. The value range is 10
	// to 32768. Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	DataVolumeSize *float64 `json:"dataVolumeSize,omitempty" tf:"data_volume_size,omitempty"`

	// Specifies the data disk flavor of the nodes.
	// Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// The following disk types are supported:
	// +kubebuilder:validation:Optional
	DataVolumeType *string `json:"dataVolumeType,omitempty" tf:"data_volume_type,omitempty"`

	// Specifies the instance specifications for each nodes in node group.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Specifies the number of nodes for the node group.
	// Only the core group and task group updations are allowed. The number of nodes after scaling cannot be
	// less than the number of nodes originally created.
	// +kubebuilder:validation:Required
	NodeNumber *float64 `json:"nodeNumber" tf:"node_number,omitempty"`

	// Specifies the system disk size of the nodes. Changing this will create
	// a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeSize *float64 `json:"rootVolumeSize" tf:"root_volume_size,omitempty"`

	// Specifies the system disk flavor of the nodes. Changing this will
	// create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeType *string `json:"rootVolumeType" tf:"root_volume_type,omitempty"`
}

type StreamingTaskNodesObservation struct {

	// The host list of this nodes group in the cluster.
	HostIps []*string `json:"hostIps,omitempty" tf:"host_ips,omitempty"`
}

type StreamingTaskNodesParameters struct {

	// Specifies the roles deployed in a node group. This argument is mandatory
	// when the cluster type is CUSTOM. Each character string represents a role expression.
	// +kubebuilder:validation:Optional
	AssignedRoles []*string `json:"assignedRoles,omitempty" tf:"assigned_roles,omitempty"`

	// Specifies the data disk number of the nodes. The number configuration
	// of each node are as follows:
	// +kubebuilder:validation:Required
	DataVolumeCount *float64 `json:"dataVolumeCount" tf:"data_volume_count,omitempty"`

	// Specifies the data disk size of the nodes,in GB. The value range is 10
	// to 32768. Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Optional
	DataVolumeSize *float64 `json:"dataVolumeSize,omitempty" tf:"data_volume_size,omitempty"`

	// Specifies the data disk flavor of the nodes.
	// Required if data_volume_count is greater than zero. Changing this will create a new MRS cluster resource.
	// The following disk types are supported:
	// +kubebuilder:validation:Optional
	DataVolumeType *string `json:"dataVolumeType,omitempty" tf:"data_volume_type,omitempty"`

	// Specifies the instance specifications for each nodes in node group.
	// Changing this will create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Specifies the number of nodes for the node group.
	// Only the core group and task group updations are allowed. The number of nodes after scaling cannot be
	// less than the number of nodes originally created.
	// +kubebuilder:validation:Required
	NodeNumber *float64 `json:"nodeNumber" tf:"node_number,omitempty"`

	// Specifies the system disk size of the nodes. Changing this will create
	// a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeSize *float64 `json:"rootVolumeSize" tf:"root_volume_size,omitempty"`

	// Specifies the system disk flavor of the nodes. Changing this will
	// create a new MRS cluster resource.
	// +kubebuilder:validation:Required
	RootVolumeType *string `json:"rootVolumeType" tf:"root_volume_type,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. ""page_title: "flexibleengine_mrs_cluster_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
