/*
Copyright 2022 The Kruise Authors.

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
	"context"
	"time"

	v1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	scheme "github.com/openkruise/kruise-api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ImagesPullJobsGetter has a method to return a ImagesPullJobInterface.
// A group's client should implement this interface.
type ImagesPullJobsGetter interface {
	ImagesPullJobs(namespace string) ImagesPullJobInterface
}

// ImagesPullJobInterface has methods to work with ImagesPullJob resources.
type ImagesPullJobInterface interface {
	Create(ctx context.Context, imagesPullJob *v1alpha1.ImagesPullJob, opts v1.CreateOptions) (*v1alpha1.ImagesPullJob, error)
	Update(ctx context.Context, imagesPullJob *v1alpha1.ImagesPullJob, opts v1.UpdateOptions) (*v1alpha1.ImagesPullJob, error)
	UpdateStatus(ctx context.Context, imagesPullJob *v1alpha1.ImagesPullJob, opts v1.UpdateOptions) (*v1alpha1.ImagesPullJob, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ImagesPullJob, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ImagesPullJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImagesPullJob, err error)
	ImagesPullJobExpansion
}

// imagesPullJobs implements ImagesPullJobInterface
type imagesPullJobs struct {
	client rest.Interface
	ns     string
}

// newImagesPullJobs returns a ImagesPullJobs
func newImagesPullJobs(c *AppsV1alpha1Client, namespace string) *imagesPullJobs {
	return &imagesPullJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the imagesPullJob, and returns the corresponding imagesPullJob object, and an error if there is any.
func (c *imagesPullJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImagesPullJob, err error) {
	result = &v1alpha1.ImagesPullJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagespulljobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ImagesPullJobs that match those selectors.
func (c *imagesPullJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImagesPullJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ImagesPullJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagespulljobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested imagesPullJobs.
func (c *imagesPullJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("imagespulljobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a imagesPullJob and creates it.  Returns the server's representation of the imagesPullJob, and an error, if there is any.
func (c *imagesPullJobs) Create(ctx context.Context, imagesPullJob *v1alpha1.ImagesPullJob, opts v1.CreateOptions) (result *v1alpha1.ImagesPullJob, err error) {
	result = &v1alpha1.ImagesPullJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("imagespulljobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(imagesPullJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a imagesPullJob and updates it. Returns the server's representation of the imagesPullJob, and an error, if there is any.
func (c *imagesPullJobs) Update(ctx context.Context, imagesPullJob *v1alpha1.ImagesPullJob, opts v1.UpdateOptions) (result *v1alpha1.ImagesPullJob, err error) {
	result = &v1alpha1.ImagesPullJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagespulljobs").
		Name(imagesPullJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(imagesPullJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *imagesPullJobs) UpdateStatus(ctx context.Context, imagesPullJob *v1alpha1.ImagesPullJob, opts v1.UpdateOptions) (result *v1alpha1.ImagesPullJob, err error) {
	result = &v1alpha1.ImagesPullJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagespulljobs").
		Name(imagesPullJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(imagesPullJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the imagesPullJob and deletes it. Returns an error if one occurs.
func (c *imagesPullJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagespulljobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *imagesPullJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagespulljobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched imagesPullJob.
func (c *imagesPullJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImagesPullJob, err error) {
	result = &v1alpha1.ImagesPullJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("imagespulljobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}