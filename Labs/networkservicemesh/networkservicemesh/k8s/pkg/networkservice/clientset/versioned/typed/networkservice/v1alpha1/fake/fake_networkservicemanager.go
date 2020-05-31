// Copyright (c) 2019 Cisco and/or its affiliates.
// Copyright (c) 2019 Red Hat Inc. and/or its affiliates.
// Copyright (c) 2019 VMware, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/networkservicemesh/networkservicemesh/k8s/pkg/apis/networkservice/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServiceManagers implements NetworkServiceManagerInterface
type FakeNetworkServiceManagers struct {
	Fake *FakeNetworkservicemeshV1alpha1
	ns   string
}

var networkservicemanagersResource = schema.GroupVersionResource{Group: "networkservicemesh.io", Version: "v1alpha1", Resource: "networkservicemanagers"}

var networkservicemanagersKind = schema.GroupVersionKind{Group: "networkservicemesh.io", Version: "v1alpha1", Kind: "NetworkServiceManager"}

// Get takes name of the networkServiceManager, and returns the corresponding networkServiceManager object, and an error if there is any.
func (c *FakeNetworkServiceManagers) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkServiceManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkservicemanagersResource, c.ns, name), &v1alpha1.NetworkServiceManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServiceManager), err
}

// List takes label and field selectors, and returns the list of NetworkServiceManagers that match those selectors.
func (c *FakeNetworkServiceManagers) List(opts v1.ListOptions) (result *v1alpha1.NetworkServiceManagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkservicemanagersResource, networkservicemanagersKind, c.ns, opts), &v1alpha1.NetworkServiceManagerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkServiceManagerList{ListMeta: obj.(*v1alpha1.NetworkServiceManagerList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkServiceManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServiceManagers.
func (c *FakeNetworkServiceManagers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkservicemanagersResource, c.ns, opts))

}

// Create takes the representation of a networkServiceManager and creates it.  Returns the server's representation of the networkServiceManager, and an error, if there is any.
func (c *FakeNetworkServiceManagers) Create(networkServiceManager *v1alpha1.NetworkServiceManager) (result *v1alpha1.NetworkServiceManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkservicemanagersResource, c.ns, networkServiceManager), &v1alpha1.NetworkServiceManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServiceManager), err
}

// Update takes the representation of a networkServiceManager and updates it. Returns the server's representation of the networkServiceManager, and an error, if there is any.
func (c *FakeNetworkServiceManagers) Update(networkServiceManager *v1alpha1.NetworkServiceManager) (result *v1alpha1.NetworkServiceManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkservicemanagersResource, c.ns, networkServiceManager), &v1alpha1.NetworkServiceManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServiceManager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkServiceManagers) UpdateStatus(networkServiceManager *v1alpha1.NetworkServiceManager) (*v1alpha1.NetworkServiceManager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkservicemanagersResource, "status", c.ns, networkServiceManager), &v1alpha1.NetworkServiceManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServiceManager), err
}

// Delete takes name of the networkServiceManager and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServiceManagers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkservicemanagersResource, c.ns, name), &v1alpha1.NetworkServiceManager{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServiceManagers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkservicemanagersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkServiceManagerList{})
	return err
}

// Patch applies the patch and returns the patched networkServiceManager.
func (c *FakeNetworkServiceManagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkServiceManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkservicemanagersResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkServiceManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServiceManager), err
}
