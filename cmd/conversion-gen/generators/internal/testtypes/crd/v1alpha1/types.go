package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ReplicaCount int32
type Label string
type RuleID string
type RuleValue string

type Policy struct {
	Mode      string `json:"mode,omitempty"`
	Threshold int32  `json:"threshold,omitempty"`
}

type Settings struct {
	Policies []Policy `json:"policies,omitempty"`
}

type WidgetSpec struct {
	// +k8s:conversion-gen:rename=Title
	Name Label `json:"name,omitempty"`

	Replicas ReplicaCount         `json:"replicas,omitempty"`
	Rules    map[RuleID]RuleValue `json:"rules,omitempty"`
	Settings *Settings            `json:"settings,omitempty"`
}

type WidgetStatus struct {
	Ready bool `json:"ready,omitempty"`
}

// +kubebuilder:object:root=true
type Widget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WidgetSpec   `json:"spec,omitempty"`
	Status            WidgetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type WidgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Widget `json:"items"`
}
