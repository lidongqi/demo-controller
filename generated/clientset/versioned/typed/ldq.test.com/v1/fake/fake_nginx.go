/*
Copyright The Kubernetes Authors.

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

	v1 "github.com/lidongqi/demo-controller/apis/ldq.test.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNginxes implements NginxInterface
type FakeNginxes struct {
	Fake *FakeLdqV1
	ns   string
}

var nginxesResource = v1.SchemeGroupVersion.WithResource("nginxes")

var nginxesKind = v1.SchemeGroupVersion.WithKind("Nginx")

// Get takes name of the nginx, and returns the corresponding nginx object, and an error if there is any.
func (c *FakeNginxes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Nginx, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nginxesResource, c.ns, name), &v1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Nginx), err
}

// List takes label and field selectors, and returns the list of Nginxes that match those selectors.
func (c *FakeNginxes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NginxList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nginxesResource, nginxesKind, c.ns, opts), &v1.NginxList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.NginxList{ListMeta: obj.(*v1.NginxList).ListMeta}
	for _, item := range obj.(*v1.NginxList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nginxes.
func (c *FakeNginxes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nginxesResource, c.ns, opts))

}

// Create takes the representation of a nginx and creates it.  Returns the server's representation of the nginx, and an error, if there is any.
func (c *FakeNginxes) Create(ctx context.Context, nginx *v1.Nginx, opts metav1.CreateOptions) (result *v1.Nginx, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nginxesResource, c.ns, nginx), &v1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Nginx), err
}

// Update takes the representation of a nginx and updates it. Returns the server's representation of the nginx, and an error, if there is any.
func (c *FakeNginxes) Update(ctx context.Context, nginx *v1.Nginx, opts metav1.UpdateOptions) (result *v1.Nginx, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nginxesResource, c.ns, nginx), &v1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Nginx), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNginxes) UpdateStatus(ctx context.Context, nginx *v1.Nginx, opts metav1.UpdateOptions) (*v1.Nginx, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nginxesResource, "status", c.ns, nginx), &v1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Nginx), err
}

// Delete takes name of the nginx and deletes it. Returns an error if one occurs.
func (c *FakeNginxes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(nginxesResource, c.ns, name, opts), &v1.Nginx{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNginxes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nginxesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.NginxList{})
	return err
}

// Patch applies the patch and returns the patched nginx.
func (c *FakeNginxes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Nginx, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nginxesResource, c.ns, name, pt, data, subresources...), &v1.Nginx{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Nginx), err
}
