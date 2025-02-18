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

type MonitoringSubscriptionInitParameters struct {

	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	MonitoringSubscription []MonitoringSubscriptionMonitoringSubscriptionInitParameters `json:"monitoringSubscription,omitempty" tf:"monitoring_subscription,omitempty"`
}

type MonitoringSubscriptionMonitoringSubscriptionInitParameters struct {

	// A subscription configuration for additional CloudWatch metrics. See below.
	RealtimeMetricsSubscriptionConfig []RealtimeMetricsSubscriptionConfigInitParameters `json:"realtimeMetricsSubscriptionConfig,omitempty" tf:"realtime_metrics_subscription_config,omitempty"`
}

type MonitoringSubscriptionMonitoringSubscriptionObservation struct {

	// A subscription configuration for additional CloudWatch metrics. See below.
	RealtimeMetricsSubscriptionConfig []RealtimeMetricsSubscriptionConfigObservation `json:"realtimeMetricsSubscriptionConfig,omitempty" tf:"realtime_metrics_subscription_config,omitempty"`
}

type MonitoringSubscriptionMonitoringSubscriptionParameters struct {

	// A subscription configuration for additional CloudWatch metrics. See below.
	// +kubebuilder:validation:Optional
	RealtimeMetricsSubscriptionConfig []RealtimeMetricsSubscriptionConfigParameters `json:"realtimeMetricsSubscriptionConfig" tf:"realtime_metrics_subscription_config,omitempty"`
}

type MonitoringSubscriptionObservation struct {

	// The ID of the distribution that you are enabling metrics for.
	DistributionID *string `json:"distributionId,omitempty" tf:"distribution_id,omitempty"`

	// The ID of the CloudFront monitoring subscription, which corresponds to the distribution_id.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	MonitoringSubscription []MonitoringSubscriptionMonitoringSubscriptionObservation `json:"monitoringSubscription,omitempty" tf:"monitoring_subscription,omitempty"`
}

type MonitoringSubscriptionParameters struct {

	// The ID of the distribution that you are enabling metrics for.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudfront/v1beta1.Distribution
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DistributionID *string `json:"distributionId,omitempty" tf:"distribution_id,omitempty"`

	// Reference to a Distribution in cloudfront to populate distributionId.
	// +kubebuilder:validation:Optional
	DistributionIDRef *v1.Reference `json:"distributionIdRef,omitempty" tf:"-"`

	// Selector for a Distribution in cloudfront to populate distributionId.
	// +kubebuilder:validation:Optional
	DistributionIDSelector *v1.Selector `json:"distributionIdSelector,omitempty" tf:"-"`

	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	// +kubebuilder:validation:Optional
	MonitoringSubscription []MonitoringSubscriptionMonitoringSubscriptionParameters `json:"monitoringSubscription,omitempty" tf:"monitoring_subscription,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type RealtimeMetricsSubscriptionConfigInitParameters struct {

	// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution. Valid values are Enabled and Disabled. See below.
	RealtimeMetricsSubscriptionStatus *string `json:"realtimeMetricsSubscriptionStatus,omitempty" tf:"realtime_metrics_subscription_status,omitempty"`
}

type RealtimeMetricsSubscriptionConfigObservation struct {

	// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution. Valid values are Enabled and Disabled. See below.
	RealtimeMetricsSubscriptionStatus *string `json:"realtimeMetricsSubscriptionStatus,omitempty" tf:"realtime_metrics_subscription_status,omitempty"`
}

type RealtimeMetricsSubscriptionConfigParameters struct {

	// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution. Valid values are Enabled and Disabled. See below.
	// +kubebuilder:validation:Optional
	RealtimeMetricsSubscriptionStatus *string `json:"realtimeMetricsSubscriptionStatus" tf:"realtime_metrics_subscription_status,omitempty"`
}

// MonitoringSubscriptionSpec defines the desired state of MonitoringSubscription
type MonitoringSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitoringSubscriptionParameters `json:"forProvider"`
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
	InitProvider MonitoringSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// MonitoringSubscriptionStatus defines the observed state of MonitoringSubscription.
type MonitoringSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitoringSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitoringSubscription is the Schema for the MonitoringSubscriptions API. Provides a CloudFront monitoring subscription resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MonitoringSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.monitoringSubscription) || (has(self.initProvider) && has(self.initProvider.monitoringSubscription))",message="spec.forProvider.monitoringSubscription is a required parameter"
	Spec   MonitoringSubscriptionSpec   `json:"spec"`
	Status MonitoringSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitoringSubscriptionList contains a list of MonitoringSubscriptions
type MonitoringSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringSubscription `json:"items"`
}

// Repository type metadata.
var (
	MonitoringSubscription_Kind             = "MonitoringSubscription"
	MonitoringSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitoringSubscription_Kind}.String()
	MonitoringSubscription_KindAPIVersion   = MonitoringSubscription_Kind + "." + CRDGroupVersion.String()
	MonitoringSubscription_GroupVersionKind = CRDGroupVersion.WithKind(MonitoringSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitoringSubscription{}, &MonitoringSubscriptionList{})
}
