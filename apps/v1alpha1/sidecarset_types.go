/*
Copyright 2020 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// SidecarSetSpec defines the desired state of SidecarSet
type SidecarSetSpec struct {
	// selector is a label query over pods that should be injected
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// Containers is the list of sidecar containers to be injected into the selected pod
	Containers []SidecarContainer `json:"containers,omitempty"`

	// List of volumes that can be mounted by sidecar containers
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// Paused indicates that the sidecarset is paused and will not be processed by the sidecarset controller.
	Paused bool `json:"paused,omitempty"`

	// The sidecarset strategy to use to replace existing pods with new ones.
	Strategy SidecarSetUpdateStrategy `json:"strategy,omitempty"`
}

// SidecarContainer defines the container of Sidecar
type SidecarContainer struct {
	corev1.Container `json:",inline"`
}

// SidecarSetUpdateStrategy indicates the strategy that the SidecarSet
// controller will use to perform updates. It includes any additional parameters
// necessary to perform the update for the indicated strategy.
type SidecarSetUpdateStrategy struct {
	RollingUpdate *RollingUpdateSidecarSet `json:"rollingUpdate,omitempty"`
}

// RollingUpdateSidecarSet is used to communicate parameter
type RollingUpdateSidecarSet struct {
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
}

// SidecarSetStatus defines the observed state of SidecarSet
type SidecarSetStatus struct {
	// observedGeneration is the most recent generation observed for this SidecarSet. It corresponds to the
	// SidecarSet's generation, which is updated on mutation by the API Server.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// matchedPods is the number of Pods whose labels are matched with this SidecarSet's selector and are created after sidecarset creates
	MatchedPods int32 `json:"matchedPods"`

	// updatedPods is the number of matched Pods that are injected with the latest SidecarSet's containers
	UpdatedPods int32 `json:"updatedPods"`

	// readyPods is the number of matched Pods that have a ready condition
	ReadyPods int32 `json:"readyPods"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="MATCHED",type="integer",JSONPath=".status.matchedPods",description="The number of pods matched."
// +kubebuilder:printcolumn:name="UPDATED",type="integer",JSONPath=".status.updatedPods",description="The number of pods matched and updated."
// +kubebuilder:printcolumn:name="READY",type="integer",JSONPath=".status.readyPods",description="The number of pods matched and ready."
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC."

// SidecarSet is the Schema for the sidecarsets API
type SidecarSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SidecarSetSpec   `json:"spec,omitempty"`
	Status SidecarSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SidecarSetList contains a list of SidecarSet
type SidecarSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SidecarSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SidecarSet{}, &SidecarSetList{})
}
