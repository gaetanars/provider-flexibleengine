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

type CustomImageObservation struct {
}

type CustomImageParameters struct {

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type FuncMountsObservation struct {
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type FuncMountsParameters struct {

	// Specifies the function access path.
	// +kubebuilder:validation:Required
	LocalMountPath *string `json:"localMountPath" tf:"local_mount_path,omitempty"`

	// Specifies the ID of the mounted resource (corresponding cloud service).
	// +kubebuilder:validation:Required
	MountResource *string `json:"mountResource" tf:"mount_resource,omitempty"`

	// Specifies the remote mount path. Example: 192.168.0.12:/data.
	// +kubebuilder:validation:Required
	MountSharePath *string `json:"mountSharePath" tf:"mount_share_path,omitempty"`

	// Specifies the mount type. Options: sfs, sfsTurbo, and ecs.
	// +kubebuilder:validation:Required
	MountType *string `json:"mountType" tf:"mount_type,omitempty"`
}

type FunctionObservation struct {

	// Specifies the file system list. The func_mounts object structure is documented
	// below.
	// +kubebuilder:validation:Optional
	FuncMounts []FuncMountsObservation `json:"funcMounts,omitempty" tf:"func_mounts,omitempty"`

	// Specifies a resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Uniform Resource Name
	Urn *string `json:"urn,omitempty" tf:"urn,omitempty"`

	// The version of the function
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FunctionParameters struct {

	// Specifies the agency. This parameter is mandatory if the function needs to access other
	// cloud services.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/iam/v1beta1.Agency
	// +kubebuilder:validation:Optional
	Agency *string `json:"agency,omitempty" tf:"agency,omitempty"`

	// Reference to a Agency in iam to populate agency.
	// +kubebuilder:validation:Optional
	AgencyRef *v1.Reference `json:"agencyRef,omitempty" tf:"-"`

	// Selector for a Agency in iam to populate agency.
	// +kubebuilder:validation:Optional
	AgencySelector *v1.Selector `json:"agencySelector,omitempty" tf:"-"`

	// Specifies the group to which the function belongs.
	// +kubebuilder:validation:Optional
	App *string `json:"app,omitempty" tf:"app,omitempty"`

	// Specifies An execution agency enables you to obtain a token or an AK/SK for
	// accessing other cloud services.
	// +kubebuilder:validation:Optional
	AppAgency *string `json:"appAgency,omitempty" tf:"app_agency,omitempty"`

	// Specifies the name of a function file, This field is mandatory only when coe_type
	// is set to jar or zip.
	// +kubebuilder:validation:Optional
	CodeFilename *string `json:"codeFilename,omitempty" tf:"code_filename,omitempty"`

	// Specifies the function code type, which can be inline: inline code, zip: ZIP file,
	// jar: JAR file or java functions, obs: function code stored in an OBS bucket.
	// +kubebuilder:validation:Optional
	CodeType *string `json:"codeType,omitempty" tf:"code_type,omitempty"`

	// Specifies the code url. This parameter is mandatory when code_type is set to obs.
	// +kubebuilder:validation:Optional
	CodeURL *string `json:"codeUrl,omitempty" tf:"code_url,omitempty"`

	// +kubebuilder:validation:Optional
	CustomImage []CustomImageParameters `json:"customImage,omitempty" tf:"custom_image,omitempty"`

	// Specifies the ID list of the dependencies.
	// +kubebuilder:validation:Optional
	DependList []*string `json:"dependList,omitempty" tf:"depend_list,omitempty"`

	// Specifies the description of the function.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the key/value information defined to be encrypted for the
	// function. The format is the same as user_data.
	// +kubebuilder:validation:Optional
	EncryptedUserDataSecretRef *v1.SecretKeySelector `json:"encryptedUserDataSecretRef,omitempty" tf:"-"`

	// Specifies the enterprise project id of the function. Changing
	// this creates a new function.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies the function code. When code_type is set to inline, zip, or jar, this
	// parameter is mandatory, and the code can be encoded using Base64 or just with the text code.
	// +kubebuilder:validation:Optional
	FuncCode *string `json:"funcCode,omitempty" tf:"func_code,omitempty"`

	// Specifies the file system list. The func_mounts object structure is documented
	// below.
	// +kubebuilder:validation:Optional
	FuncMounts []FuncMountsParameters `json:"funcMounts,omitempty" tf:"func_mounts,omitempty"`

	// The version of the function
	// +kubebuilder:validation:Optional
	FunctiongraphVersion *string `json:"functiongraphVersion,omitempty" tf:"functiongraph_version,omitempty"`

	// Specifies the entry point of the function.
	// +kubebuilder:validation:Optional
	Handler *string `json:"handler,omitempty" tf:"handler,omitempty"`

	// Specifies the initializer of the function.
	// +kubebuilder:validation:Optional
	InitializerHandler *string `json:"initializerHandler,omitempty" tf:"initializer_handler,omitempty"`

	// Specifies the maximum duration the function can be initialized. Value range:
	// 1s to 300s.
	// +kubebuilder:validation:Optional
	InitializerTimeout *float64 `json:"initializerTimeout,omitempty" tf:"initializer_timeout,omitempty"`

	// Specifies the memory size(MB) allocated to the function.
	// +kubebuilder:validation:Required
	MemorySize *float64 `json:"memorySize" tf:"memory_size,omitempty"`

	// Specifies the user group ID, a non-0 integer from –1 to 65534. Default to
	// -1.
	// +kubebuilder:validation:Optional
	MountUserGroupID *float64 `json:"mountUserGroupId,omitempty" tf:"mount_user_group_id,omitempty"`

	// Specifies the user ID, a non-0 integer from –1 to 65534. Default to -1.
	// +kubebuilder:validation:Optional
	MountUserID *float64 `json:"mountUserId,omitempty" tf:"mount_user_id,omitempty"`

	// Specifies the name of the function.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the network ID of subnet.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools.ExtractorParamPathfunc(true, "id")
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Package *string `json:"package,omitempty" tf:"package,omitempty"`

	// Specifies the region in which to create the Function resource. If omitted, the
	// provider-level region will be used. Changing this creates a new Function resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the environment for executing the function. Changing this creates a
	// new Function resource.
	// +kubebuilder:validation:Required
	Runtime *string `json:"runtime" tf:"runtime,omitempty"`

	// Specifies the timeout interval of the function, ranges from 3s to 900s.
	// +kubebuilder:validation:Required
	Timeout *float64 `json:"timeout" tf:"timeout,omitempty"`

	// Specifies the Key/Value information defined for the function.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// Specifies the ID of VPC.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Xrole *string `json:"xrole,omitempty" tf:"xrole,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API. ""page_title: "flexibleengine_fgs_function"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
