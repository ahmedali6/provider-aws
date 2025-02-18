/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Connector); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Connector[i3].VoiceConnectorID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.Connector[i3].VoiceConnectorIDRef,
			Selector:     mg.Spec.ForProvider.Connector[i3].VoiceConnectorIDSelector,
			To: reference.To{
				List:    &VoiceConnectorList{},
				Managed: &VoiceConnector{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Connector[i3].VoiceConnectorID")
		}
		mg.Spec.ForProvider.Connector[i3].VoiceConnectorID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Connector[i3].VoiceConnectorIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VoiceConnectorID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VoiceConnectorIDRef,
		Selector:     mg.Spec.ForProvider.VoiceConnectorIDSelector,
		To: reference.To{
			List:    &VoiceConnectorList{},
			Managed: &VoiceConnector{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VoiceConnectorID")
	}
	mg.Spec.ForProvider.VoiceConnectorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VoiceConnectorIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VoiceConnectorID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VoiceConnectorIDRef,
		Selector:     mg.Spec.ForProvider.VoiceConnectorIDSelector,
		To: reference.To{
			List:    &VoiceConnectorList{},
			Managed: &VoiceConnector{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VoiceConnectorID")
	}
	mg.Spec.ForProvider.VoiceConnectorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VoiceConnectorIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VoiceConnectorID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VoiceConnectorIDRef,
		Selector:     mg.Spec.ForProvider.VoiceConnectorIDSelector,
		To: reference.To{
			List:    &VoiceConnectorList{},
			Managed: &VoiceConnector{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VoiceConnectorID")
	}
	mg.Spec.ForProvider.VoiceConnectorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VoiceConnectorIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VoiceConnectorID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VoiceConnectorIDRef,
		Selector:     mg.Spec.ForProvider.VoiceConnectorIDSelector,
		To: reference.To{
			List:    &VoiceConnectorList{},
			Managed: &VoiceConnector{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VoiceConnectorID")
	}
	mg.Spec.ForProvider.VoiceConnectorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VoiceConnectorIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VoiceConnectorID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VoiceConnectorIDRef,
		Selector:     mg.Spec.ForProvider.VoiceConnectorIDSelector,
		To: reference.To{
			List:    &VoiceConnectorList{},
			Managed: &VoiceConnector{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VoiceConnectorID")
	}
	mg.Spec.ForProvider.VoiceConnectorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VoiceConnectorIDRef = rsp.ResolvedReference

	return nil
}
