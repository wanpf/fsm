/*
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

	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/policyattachment/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSessionStickyPolicies implements SessionStickyPolicyInterface
type FakeSessionStickyPolicies struct {
	Fake *FakeGatewayV1alpha1
	ns   string
}

var sessionstickypoliciesResource = schema.GroupVersionResource{Group: "gateway.flomesh.io", Version: "v1alpha1", Resource: "sessionstickypolicies"}

var sessionstickypoliciesKind = schema.GroupVersionKind{Group: "gateway.flomesh.io", Version: "v1alpha1", Kind: "SessionStickyPolicy"}

// Get takes name of the sessionStickyPolicy, and returns the corresponding sessionStickyPolicy object, and an error if there is any.
func (c *FakeSessionStickyPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SessionStickyPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sessionstickypoliciesResource, c.ns, name), &v1alpha1.SessionStickyPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionStickyPolicy), err
}

// List takes label and field selectors, and returns the list of SessionStickyPolicies that match those selectors.
func (c *FakeSessionStickyPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SessionStickyPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sessionstickypoliciesResource, sessionstickypoliciesKind, c.ns, opts), &v1alpha1.SessionStickyPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SessionStickyPolicyList{ListMeta: obj.(*v1alpha1.SessionStickyPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.SessionStickyPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sessionStickyPolicies.
func (c *FakeSessionStickyPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sessionstickypoliciesResource, c.ns, opts))

}

// Create takes the representation of a sessionStickyPolicy and creates it.  Returns the server's representation of the sessionStickyPolicy, and an error, if there is any.
func (c *FakeSessionStickyPolicies) Create(ctx context.Context, sessionStickyPolicy *v1alpha1.SessionStickyPolicy, opts v1.CreateOptions) (result *v1alpha1.SessionStickyPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sessionstickypoliciesResource, c.ns, sessionStickyPolicy), &v1alpha1.SessionStickyPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionStickyPolicy), err
}

// Update takes the representation of a sessionStickyPolicy and updates it. Returns the server's representation of the sessionStickyPolicy, and an error, if there is any.
func (c *FakeSessionStickyPolicies) Update(ctx context.Context, sessionStickyPolicy *v1alpha1.SessionStickyPolicy, opts v1.UpdateOptions) (result *v1alpha1.SessionStickyPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sessionstickypoliciesResource, c.ns, sessionStickyPolicy), &v1alpha1.SessionStickyPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionStickyPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSessionStickyPolicies) UpdateStatus(ctx context.Context, sessionStickyPolicy *v1alpha1.SessionStickyPolicy, opts v1.UpdateOptions) (*v1alpha1.SessionStickyPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sessionstickypoliciesResource, "status", c.ns, sessionStickyPolicy), &v1alpha1.SessionStickyPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionStickyPolicy), err
}

// Delete takes name of the sessionStickyPolicy and deletes it. Returns an error if one occurs.
func (c *FakeSessionStickyPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(sessionstickypoliciesResource, c.ns, name, opts), &v1alpha1.SessionStickyPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSessionStickyPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sessionstickypoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SessionStickyPolicyList{})
	return err
}

// Patch applies the patch and returns the patched sessionStickyPolicy.
func (c *FakeSessionStickyPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SessionStickyPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sessionstickypoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SessionStickyPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SessionStickyPolicy), err
}
