// +kubebuilder:validation:Required

package v1alpha

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UpgradeStatus defines the observed state of Upgrade.
type UpgradeStatus struct {
	// +kubebuilder:validation:Enum=NodeUpgrade;FeatureUpgrade;Completed;Failed
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// Phase indicates the current phase of the upgrade process.
	Phase string `json:"phase,omitempty"`
	// UpgardedNodes is a list of nodes that have been successfully upgraded.
	// +optional
	UpgradedNodes []string `json:"upgradedNodes,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Upgraded Nodes",type="string",JSONPath=".status.upgradedNodes"

// Upgrade is the Schema for the upgrades API.
type Upgrade struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status UpgradeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UpgradeList contains a list of Upgrade.
type UpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Upgrade `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Upgrade{}, &UpgradeList{})
}
