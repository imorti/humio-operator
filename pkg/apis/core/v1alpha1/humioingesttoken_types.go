package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HumioIngestTokenSpec defines the desired state of HumioIngestToken
type HumioIngestTokenSpec struct {
	// Which cluster
	ManagedClusterName  string `json:"managedClusterName,omitempty"`
	ExternalClusterName string `json:"externalClusterName,omitempty"`

	// Input
	Name           string `json:"name,omitempty"`
	ParserName     string `json:"parserName,omitempty"`
	RepositoryName string `json:"repositoryName,omitempty"`

	// Output
	TokenSecretName string `json:"tokenSecretName,omitempty"`
}

// HumioIngestTokenStatus defines the observed state of HumioIngestToken
type HumioIngestTokenStatus struct {
	Created bool `json:"created,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HumioIngestToken is the Schema for the humioingesttokens API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=humioingesttokens,scope=Namespaced
type HumioIngestToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HumioIngestTokenSpec   `json:"spec,omitempty"`
	Status HumioIngestTokenStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HumioIngestTokenList contains a list of HumioIngestToken
type HumioIngestTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HumioIngestToken `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HumioIngestToken{}, &HumioIngestTokenList{})
}
