package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Miner describes a Miner.
type Miner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MinerSpec `json:"spec"`
}

// MinerSpec is the spec for a Foo resource
type MinerSpec struct {
	Replicas int `json:"replicas"`
	Kind    string `json:"kind"`
	Gpu 	bool `json:"gpu"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MinerList is a list of Miner resources
type MinerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items []Miner `json:"items"`
}