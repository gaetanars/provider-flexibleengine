/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta11 "github.com/FrangipaneTeam/provider-flexibleengine/apis/iam/v1beta1"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1"
	tools "github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ACL.
func (mg *ACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.InboundRules),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.InboundRuleRefs,
		Selector:      mg.Spec.ForProvider.InboundRuleSelector,
		To: reference.To{
			List:    &ACLRuleList{},
			Managed: &ACLRule{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InboundRules")
	}
	mg.Spec.ForProvider.InboundRules = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.InboundRuleRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.OutboundRules),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.OutboundRuleRefs,
		Selector:      mg.Spec.ForProvider.OutboundRuleSelector,
		To: reference.To{
			List:    &ACLRuleList{},
			Managed: &ACLRule{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OutboundRules")
	}
	mg.Spec.ForProvider.OutboundRules = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.OutboundRuleRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Subnets),
		Extract:       tools.ExtractorParamPathfunc(true, "id"),
		References:    mg.Spec.ForProvider.SubnetRefs,
		Selector:      mg.Spec.ForProvider.SubnetSelector,
		To: reference.To{
			List:    &v1beta1.VPCSubnetList{},
			Managed: &v1beta1.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this FirewallGroup.
func (mg *FirewallGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta11.ProjectList{},
			Managed: &v1beta11.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta11.ProjectList{},
			Managed: &v1beta11.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Rule.
func (mg *Rule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TenantID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TenantIDRef,
		Selector:     mg.Spec.ForProvider.TenantIDSelector,
		To: reference.To{
			List:    &v1beta11.ProjectList{},
			Managed: &v1beta11.Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TenantID")
	}
	mg.Spec.ForProvider.TenantID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TenantIDRef = rsp.ResolvedReference

	return nil
}
