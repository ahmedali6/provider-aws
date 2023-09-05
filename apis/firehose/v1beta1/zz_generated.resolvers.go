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
	v1beta1 "github.com/upbound/provider-aws/apis/elasticsearch/v1beta1"
	v1beta13 "github.com/upbound/provider-aws/apis/glue/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta14 "github.com/upbound/provider-aws/apis/opensearch/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/s3/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DeliveryStream.
func (mg *DeliveryStream) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.ElasticsearchConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ElasticsearchConfiguration[i3].DomainArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ElasticsearchConfiguration[i3].DomainArnRef,
			Selector:     mg.Spec.ForProvider.ElasticsearchConfiguration[i3].DomainArnSelector,
			To: reference.To{
				List:    &v1beta1.DomainList{},
				Managed: &v1beta1.Domain{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ElasticsearchConfiguration[i3].DomainArn")
		}
		mg.Spec.ForProvider.ElasticsearchConfiguration[i3].DomainArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ElasticsearchConfiguration[i3].DomainArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ElasticsearchConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ElasticsearchConfiguration[i3].RoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ElasticsearchConfiguration[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.ElasticsearchConfiguration[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ElasticsearchConfiguration[i3].RoleArn")
		}
		mg.Spec.ForProvider.ElasticsearchConfiguration[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ElasticsearchConfiguration[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ElasticsearchConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].BucketArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].BucketArnRef,
				Selector:     mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].BucketArnSelector,
				To: reference.To{
					List:    &v1beta12.BucketList{},
					Managed: &v1beta12.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].BucketArn")
			}
			mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].BucketArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ElasticsearchConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].RoleArn")
			}
			mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ElasticsearchConfiguration[i3].S3Configuration[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ElasticsearchConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ElasticsearchConfiguration[i3].VPCConfig); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ElasticsearchConfiguration[i3].VPCConfig[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.ElasticsearchConfiguration[i3].VPCConfig[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.ElasticsearchConfiguration[i3].VPCConfig[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.ElasticsearchConfiguration[i3].VPCConfig[i4].RoleArn")
			}
			mg.Spec.ForProvider.ElasticsearchConfiguration[i3].VPCConfig[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.ElasticsearchConfiguration[i3].VPCConfig[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ExtendedS3Configuration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExtendedS3Configuration[i3].BucketArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ExtendedS3Configuration[i3].BucketArnRef,
			Selector:     mg.Spec.ForProvider.ExtendedS3Configuration[i3].BucketArnSelector,
			To: reference.To{
				List:    &v1beta12.BucketList{},
				Managed: &v1beta12.Bucket{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ExtendedS3Configuration[i3].BucketArn")
		}
		mg.Spec.ForProvider.ExtendedS3Configuration[i3].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ExtendedS3Configuration[i3].BucketArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ExtendedS3Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].RoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].RoleArnRef,
					Selector:     mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].RoleArnSelector,
					To: reference.To{
						List:    &v1beta11.RoleList{},
						Managed: &v1beta11.Role{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].RoleArn")
				}
				mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].RoleArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ExtendedS3Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].TableName),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].TableNameRef,
					Selector:     mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].TableNameSelector,
					To: reference.To{
						List:    &v1beta13.CatalogTableList{},
						Managed: &v1beta13.CatalogTable{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].TableName")
				}
				mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].TableName = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.ExtendedS3Configuration[i3].DataFormatConversionConfiguration[i4].SchemaConfiguration[i5].TableNameRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ExtendedS3Configuration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExtendedS3Configuration[i3].RoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ExtendedS3Configuration[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.ExtendedS3Configuration[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ExtendedS3Configuration[i3].RoleArn")
		}
		mg.Spec.ForProvider.ExtendedS3Configuration[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ExtendedS3Configuration[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.HTTPEndpointConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].RoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].RoleArn")
		}
		mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.HTTPEndpointConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].BucketArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].BucketArnRef,
				Selector:     mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].BucketArnSelector,
				To: reference.To{
					List:    &v1beta12.BucketList{},
					Managed: &v1beta12.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].BucketArn")
			}
			mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].BucketArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.HTTPEndpointConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].RoleArn")
			}
			mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.HTTPEndpointConfiguration[i3].S3Configuration[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.OpensearchConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OpensearchConfiguration[i3].DomainArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.OpensearchConfiguration[i3].DomainArnRef,
			Selector:     mg.Spec.ForProvider.OpensearchConfiguration[i3].DomainArnSelector,
			To: reference.To{
				List:    &v1beta14.DomainList{},
				Managed: &v1beta14.Domain{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.OpensearchConfiguration[i3].DomainArn")
		}
		mg.Spec.ForProvider.OpensearchConfiguration[i3].DomainArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.OpensearchConfiguration[i3].DomainArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.OpensearchConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OpensearchConfiguration[i3].RoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.OpensearchConfiguration[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.OpensearchConfiguration[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.OpensearchConfiguration[i3].RoleArn")
		}
		mg.Spec.ForProvider.OpensearchConfiguration[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.OpensearchConfiguration[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.OpensearchConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].BucketArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].BucketArnRef,
				Selector:     mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].BucketArnSelector,
				To: reference.To{
					List:    &v1beta12.BucketList{},
					Managed: &v1beta12.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].BucketArn")
			}
			mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].BucketArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.OpensearchConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].RoleArn")
			}
			mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.OpensearchConfiguration[i3].S3Configuration[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.OpensearchConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.OpensearchConfiguration[i3].VPCConfig); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OpensearchConfiguration[i3].VPCConfig[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.OpensearchConfiguration[i3].VPCConfig[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.OpensearchConfiguration[i3].VPCConfig[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.OpensearchConfiguration[i3].VPCConfig[i4].RoleArn")
			}
			mg.Spec.ForProvider.OpensearchConfiguration[i3].VPCConfig[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.OpensearchConfiguration[i3].VPCConfig[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RedshiftConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RedshiftConfiguration[i3].RoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.RedshiftConfiguration[i3].RoleArnRef,
			Selector:     mg.Spec.ForProvider.RedshiftConfiguration[i3].RoleArnSelector,
			To: reference.To{
				List:    &v1beta11.RoleList{},
				Managed: &v1beta11.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RedshiftConfiguration[i3].RoleArn")
		}
		mg.Spec.ForProvider.RedshiftConfiguration[i3].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RedshiftConfiguration[i3].RoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RedshiftConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].BucketArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].BucketArnRef,
				Selector:     mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].BucketArnSelector,
				To: reference.To{
					List:    &v1beta12.BucketList{},
					Managed: &v1beta12.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].BucketArn")
			}
			mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].BucketArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RedshiftConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].RoleArn")
			}
			mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.RedshiftConfiguration[i3].S3BackupConfiguration[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RedshiftConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].BucketArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].BucketArnRef,
				Selector:     mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].BucketArnSelector,
				To: reference.To{
					List:    &v1beta12.BucketList{},
					Managed: &v1beta12.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].BucketArn")
			}
			mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].BucketArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RedshiftConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].RoleArn")
			}
			mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.RedshiftConfiguration[i3].S3Configuration[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SplunkConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].BucketArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].BucketArnRef,
				Selector:     mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].BucketArnSelector,
				To: reference.To{
					List:    &v1beta12.BucketList{},
					Managed: &v1beta12.Bucket{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].BucketArn")
			}
			mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].BucketArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].BucketArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SplunkConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].RoleArn")
			}
			mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SplunkConfiguration[i3].S3Configuration[i4].RoleArnRef = rsp.ResolvedReference

		}
	}

	return nil
}
