package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type YoyoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Yoyo `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Yoyo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              YoyoSpec   `json:"spec"`
	Status            YoyoStatus `json:"status,omitempty"`
}

type YoyoSpec struct {
	YoyoDemo bool `json:",yoyoDemo"`
}
type YoyoStatus struct {
	Status string `json:",status"`
}
