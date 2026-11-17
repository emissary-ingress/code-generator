package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type DesiredReplicas int32
type DisplayLabel string
type RuleKey string
type RuleText string

type Policy struct {
	Mode      string `json:"mode,omitempty"`
	Threshold int32  `json:"threshold,omitempty"`
}

type Settings struct {
	Policies []Policy `json:"policies,omitempty"`
}

type WidgetSpec struct {
	Title    DisplayLabel         `json:"title,omitempty"`
	Replicas DesiredReplicas      `json:"replicas,omitempty"`
	Rules    map[RuleKey]RuleText `json:"rules,omitempty"`
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
