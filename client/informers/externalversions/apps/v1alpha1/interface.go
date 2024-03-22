/*
Copyright 2024 The Kruise Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openkruise/kruise-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AdvancedCronJobs returns a AdvancedCronJobInformer.
	AdvancedCronJobs() AdvancedCronJobInformer
	// BroadcastJobs returns a BroadcastJobInformer.
	BroadcastJobs() BroadcastJobInformer
	// CloneSets returns a CloneSetInformer.
	CloneSets() CloneSetInformer
	// ContainerRecreateRequests returns a ContainerRecreateRequestInformer.
	ContainerRecreateRequests() ContainerRecreateRequestInformer
	// DaemonSets returns a DaemonSetInformer.
	DaemonSets() DaemonSetInformer
	// EphemeralJobs returns a EphemeralJobInformer.
	EphemeralJobs() EphemeralJobInformer
	// ImageListPullJobs returns a ImageListPullJobInformer.
	ImageListPullJobs() ImageListPullJobInformer
	// ImagePullJobs returns a ImagePullJobInformer.
	ImagePullJobs() ImagePullJobInformer
	// NodeImages returns a NodeImageInformer.
	NodeImages() NodeImageInformer
	// NodePodProbes returns a NodePodProbeInformer.
	NodePodProbes() NodePodProbeInformer
	// PersistentPodStates returns a PersistentPodStateInformer.
	PersistentPodStates() PersistentPodStateInformer
	// PodProbeMarkers returns a PodProbeMarkerInformer.
	PodProbeMarkers() PodProbeMarkerInformer
	// ResourceDistributions returns a ResourceDistributionInformer.
	ResourceDistributions() ResourceDistributionInformer
	// SidecarSets returns a SidecarSetInformer.
	SidecarSets() SidecarSetInformer
	// StatefulSets returns a StatefulSetInformer.
	StatefulSets() StatefulSetInformer
	// UnitedDeployments returns a UnitedDeploymentInformer.
	UnitedDeployments() UnitedDeploymentInformer
	// WorkloadSpreads returns a WorkloadSpreadInformer.
	WorkloadSpreads() WorkloadSpreadInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AdvancedCronJobs returns a AdvancedCronJobInformer.
func (v *version) AdvancedCronJobs() AdvancedCronJobInformer {
	return &advancedCronJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BroadcastJobs returns a BroadcastJobInformer.
func (v *version) BroadcastJobs() BroadcastJobInformer {
	return &broadcastJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloneSets returns a CloneSetInformer.
func (v *version) CloneSets() CloneSetInformer {
	return &cloneSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ContainerRecreateRequests returns a ContainerRecreateRequestInformer.
func (v *version) ContainerRecreateRequests() ContainerRecreateRequestInformer {
	return &containerRecreateRequestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DaemonSets returns a DaemonSetInformer.
func (v *version) DaemonSets() DaemonSetInformer {
	return &daemonSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EphemeralJobs returns a EphemeralJobInformer.
func (v *version) EphemeralJobs() EphemeralJobInformer {
	return &ephemeralJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ImageListPullJobs returns a ImageListPullJobInformer.
func (v *version) ImageListPullJobs() ImageListPullJobInformer {
	return &imageListPullJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ImagePullJobs returns a ImagePullJobInformer.
func (v *version) ImagePullJobs() ImagePullJobInformer {
	return &imagePullJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NodeImages returns a NodeImageInformer.
func (v *version) NodeImages() NodeImageInformer {
	return &nodeImageInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// NodePodProbes returns a NodePodProbeInformer.
func (v *version) NodePodProbes() NodePodProbeInformer {
	return &nodePodProbeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PersistentPodStates returns a PersistentPodStateInformer.
func (v *version) PersistentPodStates() PersistentPodStateInformer {
	return &persistentPodStateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PodProbeMarkers returns a PodProbeMarkerInformer.
func (v *version) PodProbeMarkers() PodProbeMarkerInformer {
	return &podProbeMarkerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ResourceDistributions returns a ResourceDistributionInformer.
func (v *version) ResourceDistributions() ResourceDistributionInformer {
	return &resourceDistributionInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// SidecarSets returns a SidecarSetInformer.
func (v *version) SidecarSets() SidecarSetInformer {
	return &sidecarSetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// StatefulSets returns a StatefulSetInformer.
func (v *version) StatefulSets() StatefulSetInformer {
	return &statefulSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UnitedDeployments returns a UnitedDeploymentInformer.
func (v *version) UnitedDeployments() UnitedDeploymentInformer {
	return &unitedDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WorkloadSpreads returns a WorkloadSpreadInformer.
func (v *version) WorkloadSpreads() WorkloadSpreadInformer {
	return &workloadSpreadInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
