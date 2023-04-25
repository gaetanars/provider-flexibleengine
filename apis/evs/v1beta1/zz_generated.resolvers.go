/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta11 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/csbs/v1beta1"
	v1beta12 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/ecs/v1beta1"
	v1beta1 "github.com/FlexibleEngineCloud/provider-flexibleengine/apis/ims/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BlockStorageVolume.
func (mg *BlockStorageVolume) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ImageID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ImageIDRef,
		Selector:     mg.Spec.ForProvider.ImageIDSelector,
		To: reference.To{
			List:    &v1beta1.ImageList{},
			Managed: &v1beta1.Image{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ImageID")
	}
	mg.Spec.ForProvider.ImageID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ImageIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotIDRef,
		Selector:     mg.Spec.ForProvider.SnapshotIDSelector,
		To: reference.To{
			List:    &v1beta11.BackupList{},
			Managed: &v1beta11.Backup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotID")
	}
	mg.Spec.ForProvider.SnapshotID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ComputeVolumeAttach.
func (mg *ComputeVolumeAttach) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceIDRef,
		Selector:     mg.Spec.ForProvider.InstanceIDSelector,
		To: reference.To{
			List:    &v1beta12.InstanceList{},
			Managed: &v1beta12.Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VolumeID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VolumeIDRef,
		Selector:     mg.Spec.ForProvider.VolumeIDSelector,
		To: reference.To{
			List:    &BlockStorageVolumeList{},
			Managed: &BlockStorageVolume{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VolumeID")
	}
	mg.Spec.ForProvider.VolumeID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VolumeIDRef = rsp.ResolvedReference

	return nil
}
