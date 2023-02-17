/*
Copyright 2023 The Kruise Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ImagesPullJobSpec defines the desired state of ImagesPullJob
type ImagesPullJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Images is the image list to be pulled by the job
	Images []string `json:"images"`

	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling the image.
	// If specified, these secrets will be passed to individual puller implementations for them to use.  For example,
	// in the case of docker, only DockerConfig type secrets are honored.
	// +optional
	PullSecrets []string `json:"pullSecrets,omitempty"`
}

// ImagesPullJobStatus defines the observed state of ImagesPullJob
type ImagesPullJobStatus struct {
	// Represents time when the job was acknowledged by the job controller.
	// It is not guaranteed to be set in happens-before order across separate operations.
	// It is represented in RFC3339 form and is in UTC.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// LastUpdatedTime is the last updated timestamp for status.
	// It is represented in RFC3339 form and is in UTC.
	// +optional
	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`

	// Phase indicates the progressing phase of ImagesPullJob.
	Phase ImagesPullJobPhase `json:"phase"`

	// Message during pulling or caching images.
	// Defaults to "".
	Message string `json:"message,omitempty"`

	// Progress is a percentage which describe the overall job progress.
	// Within ["0%", "100%"].
	Progress string `json:"progress,omitempty"`
}

type ImagesPullJobPhase string

// +genclient
// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ImagesPullJob is the Schema for the imagespulljobs API
type ImagesPullJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImagesPullJobSpec   `json:"spec,omitempty"`
	Status ImagesPullJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImagesPullJobList contains a list of ImagesPullJob
type ImagesPullJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagesPullJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImagesPullJob{}, &ImagesPullJobList{})
}
