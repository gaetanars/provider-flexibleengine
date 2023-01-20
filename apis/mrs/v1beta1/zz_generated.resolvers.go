/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta11 "github.com/FrangipaneTeam/provider-flexibleengine/apis/ecs/v1beta1"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/eip/v1beta1"
	v1beta12 "github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1"
	tools "github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EIPID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EIPIDRef,
		Selector:     mg.Spec.ForProvider.EIPIDSelector,
		To: reference.To{
			List:    &v1beta1.EIPList{},
			Managed: &v1beta1.EIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EIPID")
	}
	mg.Spec.ForProvider.EIPID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EIPIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeKeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NodeKeyPairRef,
		Selector:     mg.Spec.ForProvider.NodeKeyPairSelector,
		To: reference.To{
			List:    &v1beta11.KeyPairList{},
			Managed: &v1beta11.KeyPair{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NodeKeyPair")
	}
	mg.Spec.ForProvider.NodeKeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NodeKeyPairRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PublicIP),
		Extract:      tools.ExtractorParamPathfunc(true, "address"),
		Reference:    mg.Spec.ForProvider.PublicIPRef,
		Selector:     mg.Spec.ForProvider.PublicIPSelector,
		To: reference.To{
			List:    &v1beta1.EIPList{},
			Managed: &v1beta1.EIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PublicIP")
	}
	mg.Spec.ForProvider.PublicIP = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PublicIPRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta12.SecurityGroupList{},
			Managed: &v1beta12.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      tools.ExtractorParamPathfunc(true, "id"),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta12.VPCSubnetList{},
			Managed: &v1beta12.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1beta12.VPCList{},
			Managed: &v1beta12.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Job.
func (mg *Job) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}
