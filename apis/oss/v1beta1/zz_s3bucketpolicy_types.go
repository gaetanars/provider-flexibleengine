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

type S3BucketPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type S3BucketPolicyParameters struct {

	// The name of the bucket to which to apply the policy.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/oss/v1beta1.S3Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a S3Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a S3Bucket in oss to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The text of the policy.
	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`
}

// S3BucketPolicySpec defines the desired state of S3BucketPolicy
type S3BucketPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     S3BucketPolicyParameters `json:"forProvider"`
}

// S3BucketPolicyStatus defines the observed state of S3BucketPolicy.
type S3BucketPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        S3BucketPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketPolicy is the Schema for the S3BucketPolicys API. ""page_title: "flexibleengine_s3_bucket_policy"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type S3BucketPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketPolicySpec   `json:"spec"`
	Status            S3BucketPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketPolicyList contains a list of S3BucketPolicys
type S3BucketPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketPolicy `json:"items"`
}

// Repository type metadata.
var (
	S3BucketPolicy_Kind             = "S3BucketPolicy"
	S3BucketPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: S3BucketPolicy_Kind}.String()
	S3BucketPolicy_KindAPIVersion   = S3BucketPolicy_Kind + "." + CRDGroupVersion.String()
	S3BucketPolicy_GroupVersionKind = CRDGroupVersion.WithKind(S3BucketPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&S3BucketPolicy{}, &S3BucketPolicyList{})
}
