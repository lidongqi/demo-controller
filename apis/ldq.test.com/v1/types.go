package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type NginxSpec struct {
	// +optional
	Image    string `json:"image,omitempty"`
	Replicas int32  `json:"replicas"`
}

type NginxStatus struct {
	// +optional
	Message string `json:"message,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Nginx type
type Nginx struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec NginxSpec `json:"spec,omitempty"`

	// +optional
	Status NginxStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NginxList is a list of Nginx resources
type NginxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Nginx `json:"items"`
}
