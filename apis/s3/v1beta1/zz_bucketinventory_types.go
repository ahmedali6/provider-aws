// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type BucketInventoryDestinationInitParameters struct {

	// Name of the source bucket that inventory lists the objects for.
	Bucket []DestinationBucketInitParameters `json:"bucket,omitempty" tf:"bucket,omitempty"`
}

type BucketInventoryDestinationObservation struct {

	// Name of the source bucket that inventory lists the objects for.
	Bucket []DestinationBucketObservation `json:"bucket,omitempty" tf:"bucket,omitempty"`
}

type BucketInventoryDestinationParameters struct {

	// Name of the source bucket that inventory lists the objects for.
	// +kubebuilder:validation:Optional
	Bucket []DestinationBucketParameters `json:"bucket" tf:"bucket,omitempty"`
}

type BucketInventoryFilterInitParameters struct {

	// Prefix that an object must have to be included in the inventory results.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type BucketInventoryFilterObservation struct {

	// Prefix that an object must have to be included in the inventory results.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type BucketInventoryFilterParameters struct {

	// Prefix that an object must have to be included in the inventory results.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type BucketInventoryInitParameters struct {

	// Contains information about where to publish the inventory results (documented below).
	Destination []BucketInventoryDestinationInitParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Specifies whether the inventory is enabled or disabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
	Filter []BucketInventoryFilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Object versions to include in the inventory list. Valid values: All, Current.
	IncludedObjectVersions *string `json:"includedObjectVersions,omitempty" tf:"included_object_versions,omitempty"`

	// Unique identifier of the inventory configuration for the bucket.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of optional fields that are included in the inventory results. Please refer to the S3 documentation for more details.
	OptionalFields []*string `json:"optionalFields,omitempty" tf:"optional_fields,omitempty"`

	// Specifies the schedule for generating inventory results (documented below).
	Schedule []ScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type BucketInventoryObservation struct {

	// Name of the source bucket that inventory lists the objects for.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Contains information about where to publish the inventory results (documented below).
	Destination []BucketInventoryDestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// Specifies whether the inventory is enabled or disabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
	Filter []BucketInventoryFilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Object versions to include in the inventory list. Valid values: All, Current.
	IncludedObjectVersions *string `json:"includedObjectVersions,omitempty" tf:"included_object_versions,omitempty"`

	// Unique identifier of the inventory configuration for the bucket.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of optional fields that are included in the inventory results. Please refer to the S3 documentation for more details.
	OptionalFields []*string `json:"optionalFields,omitempty" tf:"optional_fields,omitempty"`

	// Specifies the schedule for generating inventory results (documented below).
	Schedule []ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type BucketInventoryParameters struct {

	// Name of the source bucket that inventory lists the objects for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Contains information about where to publish the inventory results (documented below).
	// +kubebuilder:validation:Optional
	Destination []BucketInventoryDestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Specifies whether the inventory is enabled or disabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
	// +kubebuilder:validation:Optional
	Filter []BucketInventoryFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Object versions to include in the inventory list. Valid values: All, Current.
	// +kubebuilder:validation:Optional
	IncludedObjectVersions *string `json:"includedObjectVersions,omitempty" tf:"included_object_versions,omitempty"`

	// Unique identifier of the inventory configuration for the bucket.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of optional fields that are included in the inventory results. Please refer to the S3 documentation for more details.
	// +kubebuilder:validation:Optional
	OptionalFields []*string `json:"optionalFields,omitempty" tf:"optional_fields,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the schedule for generating inventory results (documented below).
	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type DestinationBucketInitParameters struct {

	// ID of the account that owns the destination bucket. Recommended to be set to prevent problems if the destination bucket ownership changes.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Contains the type of server-side encryption to use to encrypt the inventory (documented below).
	Encryption []EncryptionInitParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Specifies the output format of the inventory results. Can be CSV, ORC or Parquet.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Prefix that an object must have to be included in the inventory results.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type DestinationBucketObservation struct {

	// ID of the account that owns the destination bucket. Recommended to be set to prevent problems if the destination bucket ownership changes.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Amazon S3 bucket ARN of the destination.
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// Contains the type of server-side encryption to use to encrypt the inventory (documented below).
	Encryption []EncryptionObservation `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Specifies the output format of the inventory results. Can be CSV, ORC or Parquet.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// Prefix that an object must have to be included in the inventory results.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type DestinationBucketParameters struct {

	// ID of the account that owns the destination bucket. Recommended to be set to prevent problems if the destination bucket ownership changes.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Amazon S3 bucket ARN of the destination.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnRef *v1.Reference `json:"bucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnSelector *v1.Selector `json:"bucketArnSelector,omitempty" tf:"-"`

	// Contains the type of server-side encryption to use to encrypt the inventory (documented below).
	// +kubebuilder:validation:Optional
	Encryption []EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// Specifies the output format of the inventory results. Can be CSV, ORC or Parquet.
	// +kubebuilder:validation:Optional
	Format *string `json:"format" tf:"format,omitempty"`

	// Prefix that an object must have to be included in the inventory results.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type EncryptionInitParameters struct {

	// Specifies to use server-side encryption with AWS KMS-managed keys to encrypt the inventory file (documented below).
	SseKMS []SseKMSInitParameters `json:"sseKms,omitempty" tf:"sse_kms,omitempty"`

	// Specifies to use server-side encryption with Amazon S3-managed keys (SSE-S3) to encrypt the inventory file.
	SseS3 []SseS3InitParameters `json:"sseS3,omitempty" tf:"sse_s3,omitempty"`
}

type EncryptionObservation struct {

	// Specifies to use server-side encryption with AWS KMS-managed keys to encrypt the inventory file (documented below).
	SseKMS []SseKMSObservation `json:"sseKms,omitempty" tf:"sse_kms,omitempty"`

	// Specifies to use server-side encryption with Amazon S3-managed keys (SSE-S3) to encrypt the inventory file.
	SseS3 []SseS3Parameters `json:"sseS3,omitempty" tf:"sse_s3,omitempty"`
}

type EncryptionParameters struct {

	// Specifies to use server-side encryption with AWS KMS-managed keys to encrypt the inventory file (documented below).
	// +kubebuilder:validation:Optional
	SseKMS []SseKMSParameters `json:"sseKms,omitempty" tf:"sse_kms,omitempty"`

	// Specifies to use server-side encryption with Amazon S3-managed keys (SSE-S3) to encrypt the inventory file.
	// +kubebuilder:validation:Optional
	SseS3 []SseS3Parameters `json:"sseS3,omitempty" tf:"sse_s3,omitempty"`
}

type ScheduleInitParameters struct {

	// Specifies how frequently inventory results are produced. Valid values: Daily, Weekly.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`
}

type ScheduleObservation struct {

	// Specifies how frequently inventory results are produced. Valid values: Daily, Weekly.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`
}

type ScheduleParameters struct {

	// Specifies how frequently inventory results are produced. Valid values: Daily, Weekly.
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency" tf:"frequency,omitempty"`
}

type SseKMSInitParameters struct {

	// ARN of the KMS customer master key (CMK) used to encrypt the inventory file.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type SseKMSObservation struct {

	// ARN of the KMS customer master key (CMK) used to encrypt the inventory file.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type SseKMSParameters struct {

	// ARN of the KMS customer master key (CMK) used to encrypt the inventory file.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId" tf:"key_id,omitempty"`
}

type SseS3InitParameters struct {
}

type SseS3Observation struct {
}

type SseS3Parameters struct {
}

// BucketInventorySpec defines the desired state of BucketInventory
type BucketInventorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketInventoryParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider BucketInventoryInitParameters `json:"initProvider,omitempty"`
}

// BucketInventoryStatus defines the observed state of BucketInventory.
type BucketInventoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketInventoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketInventory is the Schema for the BucketInventorys API. Provides a S3 bucket inventory configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketInventory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination) || (has(self.initProvider) && has(self.initProvider.destination))",message="spec.forProvider.destination is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.includedObjectVersions) || (has(self.initProvider) && has(self.initProvider.includedObjectVersions))",message="spec.forProvider.includedObjectVersions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schedule) || (has(self.initProvider) && has(self.initProvider.schedule))",message="spec.forProvider.schedule is a required parameter"
	Spec   BucketInventorySpec   `json:"spec"`
	Status BucketInventoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketInventoryList contains a list of BucketInventorys
type BucketInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketInventory `json:"items"`
}

// Repository type metadata.
var (
	BucketInventory_Kind             = "BucketInventory"
	BucketInventory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketInventory_Kind}.String()
	BucketInventory_KindAPIVersion   = BucketInventory_Kind + "." + CRDGroupVersion.String()
	BucketInventory_GroupVersionKind = CRDGroupVersion.WithKind(BucketInventory_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketInventory{}, &BucketInventoryList{})
}
