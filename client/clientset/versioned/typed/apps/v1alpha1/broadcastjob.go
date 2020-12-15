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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	scheme "github.com/openkruise/kruise-api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BroadcastJobsGetter has a method to return a BroadcastJobInterface.
// A group's client should implement this interface.
type BroadcastJobsGetter interface {
	BroadcastJobs(namespace string) BroadcastJobInterface
}

// BroadcastJobInterface has methods to work with BroadcastJob resources.
type BroadcastJobInterface interface {
	Create(*v1alpha1.BroadcastJob) (*v1alpha1.BroadcastJob, error)
	Update(*v1alpha1.BroadcastJob) (*v1alpha1.BroadcastJob, error)
	UpdateStatus(*v1alpha1.BroadcastJob) (*v1alpha1.BroadcastJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BroadcastJob, error)
	List(opts v1.ListOptions) (*v1alpha1.BroadcastJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BroadcastJob, err error)
	BroadcastJobExpansion
}

// broadcastJobs implements BroadcastJobInterface
type broadcastJobs struct {
	client rest.Interface
	ns     string
}

// newBroadcastJobs returns a BroadcastJobs
func newBroadcastJobs(c *AppsV1alpha1Client, namespace string) *broadcastJobs {
	return &broadcastJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the broadcastJob, and returns the corresponding broadcastJob object, and an error if there is any.
func (c *broadcastJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.BroadcastJob, err error) {
	result = &v1alpha1.BroadcastJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("broadcastjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BroadcastJobs that match those selectors.
func (c *broadcastJobs) List(opts v1.ListOptions) (result *v1alpha1.BroadcastJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BroadcastJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("broadcastjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested broadcastJobs.
func (c *broadcastJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("broadcastjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a broadcastJob and creates it.  Returns the server's representation of the broadcastJob, and an error, if there is any.
func (c *broadcastJobs) Create(broadcastJob *v1alpha1.BroadcastJob) (result *v1alpha1.BroadcastJob, err error) {
	result = &v1alpha1.BroadcastJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("broadcastjobs").
		Body(broadcastJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a broadcastJob and updates it. Returns the server's representation of the broadcastJob, and an error, if there is any.
func (c *broadcastJobs) Update(broadcastJob *v1alpha1.BroadcastJob) (result *v1alpha1.BroadcastJob, err error) {
	result = &v1alpha1.BroadcastJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("broadcastjobs").
		Name(broadcastJob.Name).
		Body(broadcastJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *broadcastJobs) UpdateStatus(broadcastJob *v1alpha1.BroadcastJob) (result *v1alpha1.BroadcastJob, err error) {
	result = &v1alpha1.BroadcastJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("broadcastjobs").
		Name(broadcastJob.Name).
		SubResource("status").
		Body(broadcastJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the broadcastJob and deletes it. Returns an error if one occurs.
func (c *broadcastJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("broadcastjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *broadcastJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("broadcastjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched broadcastJob.
func (c *broadcastJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BroadcastJob, err error) {
	result = &v1alpha1.BroadcastJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("broadcastjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
