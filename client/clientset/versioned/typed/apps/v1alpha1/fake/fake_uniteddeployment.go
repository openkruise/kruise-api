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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUnitedDeployments implements UnitedDeploymentInterface
type FakeUnitedDeployments struct {
	Fake *FakeAppsV1alpha1
	ns   string
}

var uniteddeploymentsResource = schema.GroupVersionResource{Group: "apps.kruise.io", Version: "v1alpha1", Resource: "uniteddeployments"}

var uniteddeploymentsKind = schema.GroupVersionKind{Group: "apps.kruise.io", Version: "v1alpha1", Kind: "UnitedDeployment"}

// Get takes name of the unitedDeployment, and returns the corresponding unitedDeployment object, and an error if there is any.
func (c *FakeUnitedDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UnitedDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(uniteddeploymentsResource, c.ns, name), &v1alpha1.UnitedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UnitedDeployment), err
}

// List takes label and field selectors, and returns the list of UnitedDeployments that match those selectors.
func (c *FakeUnitedDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UnitedDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(uniteddeploymentsResource, uniteddeploymentsKind, c.ns, opts), &v1alpha1.UnitedDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UnitedDeploymentList{ListMeta: obj.(*v1alpha1.UnitedDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.UnitedDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested unitedDeployments.
func (c *FakeUnitedDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(uniteddeploymentsResource, c.ns, opts))

}

// Create takes the representation of a unitedDeployment and creates it.  Returns the server's representation of the unitedDeployment, and an error, if there is any.
func (c *FakeUnitedDeployments) Create(ctx context.Context, unitedDeployment *v1alpha1.UnitedDeployment, opts v1.CreateOptions) (result *v1alpha1.UnitedDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(uniteddeploymentsResource, c.ns, unitedDeployment), &v1alpha1.UnitedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UnitedDeployment), err
}

// Update takes the representation of a unitedDeployment and updates it. Returns the server's representation of the unitedDeployment, and an error, if there is any.
func (c *FakeUnitedDeployments) Update(ctx context.Context, unitedDeployment *v1alpha1.UnitedDeployment, opts v1.UpdateOptions) (result *v1alpha1.UnitedDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(uniteddeploymentsResource, c.ns, unitedDeployment), &v1alpha1.UnitedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UnitedDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUnitedDeployments) UpdateStatus(ctx context.Context, unitedDeployment *v1alpha1.UnitedDeployment, opts v1.UpdateOptions) (*v1alpha1.UnitedDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(uniteddeploymentsResource, "status", c.ns, unitedDeployment), &v1alpha1.UnitedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UnitedDeployment), err
}

// Delete takes name of the unitedDeployment and deletes it. Returns an error if one occurs.
func (c *FakeUnitedDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(uniteddeploymentsResource, c.ns, name, opts), &v1alpha1.UnitedDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUnitedDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(uniteddeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UnitedDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched unitedDeployment.
func (c *FakeUnitedDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UnitedDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(uniteddeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.UnitedDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UnitedDeployment), err
}

// GetScale takes name of the unitedDeployment, and returns the corresponding scale object, and an error if there is any.
func (c *FakeUnitedDeployments) GetScale(ctx context.Context, unitedDeploymentName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(uniteddeploymentsResource, c.ns, "scale", unitedDeploymentName), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeUnitedDeployments) UpdateScale(ctx context.Context, unitedDeploymentName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(uniteddeploymentsResource, "scale", c.ns, scale), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}
