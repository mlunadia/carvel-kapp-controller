// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	versions "github.com/vmware-tanzu/carvel-vendir/pkg/vendir/versions/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=ipkg
// +kubebuilder:printcolumn:name=Package name,JSONPath=.spec.packageName,description=Package public name,type=string
// +kubebuilder:printcolumn:name=Package version,JSONPath=.status.version,description=Package version,type=string
// +kubebuilder:printcolumn:name=Description,JSONPath=.status.friendlyDescription,description=Friendly description,type=string
// +kubebuilder:printcolumn:name=Age,JSONPath=.metadata.creationTimestamp,description=Time since creation,type=date
type InstalledPackage struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec InstalledPackageSpec `json:"spec"`
	// +optional
	Status InstalledPackageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type InstalledPackageList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []InstalledPackage `json:"items"`
}

type InstalledPackageSpec struct {
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
	// +optional
	Cluster *v1alpha1.AppCluster `json:"cluster,omitempty"`
	// +optional
	PackageVersionRef *PackageVersionRef `json:"packageVersionRef,omitempty"`
	// +optional
	Values []InstalledPackageValues `json:"values,omitempty"`
	// Paused when set to true will ignore all pending changes,
	// once it set back to false, pending changes will be applied
	// +optional
	Paused bool `json:"paused,omitempty"`
	// Canceled when set to true will stop all active changes
	// +optional
	Canceled bool `json:"canceled,omitempty"`
	// Controls frequency of App reconciliation in time + unit
	// format. Always >= 30s. If value below 30s is specified,
	// 30s will be used.
	// +optional
	SyncPeriod *metav1.Duration `json:"syncPeriod,omitempty"`
	// When NoopDelete set to true, InstalledPackage deletion
	// should delete InstalledPackage/App CR but preserve App's
	// associated resources.
	// +optional
	NoopDelete bool `json:"noopDelete,omitempty"`
}

type PackageVersionRef struct {
	// +optional
	PackageName string `json:"packageName,omitempty"`
	// +optional
	VersionSelection *versions.VersionSelectionSemver `json:"versionSelection,omitempty"`
}

type InstalledPackageValues struct {
	// +optional
	SecretRef *InstalledPackageValuesSecretRef `json:"secretRef,omitempty"`
}

type InstalledPackageValuesSecretRef struct {
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Key string `json:"key,omitempty"`
}

type InstalledPackageStatus struct {
	// +optional
	v1alpha1.GenericStatus `json:",inline"`
	// TODO this is desired resolved version (not actually deployed)
	// +optional
	Version string `json:"version,omitempty"`
}