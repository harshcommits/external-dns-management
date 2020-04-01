/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1alpha1 "github.com/gardener/external-dns-management/pkg/apis/dns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSResourceWatches implements DNSResourceWatchInterface
type FakeDNSResourceWatches struct {
	Fake *FakeDnsV1alpha1
	ns   string
}

var dnsresourcewatchesResource = schema.GroupVersionResource{Group: "dns.gardener.cloud", Version: "v1alpha1", Resource: "dnsresourcewatches"}

var dnsresourcewatchesKind = schema.GroupVersionKind{Group: "dns.gardener.cloud", Version: "v1alpha1", Kind: "DNSAnnotation"}

// Get takes name of the dNSResourceWatch, and returns the corresponding dNSResourceWatch object, and an error if there is any.
func (c *FakeDNSResourceWatches) Get(name string, options v1.GetOptions) (result *v1alpha1.DNSAnnotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsresourcewatchesResource, c.ns, name), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}

// List takes label and field selectors, and returns the list of DNSResourceWatches that match those selectors.
func (c *FakeDNSResourceWatches) List(opts v1.ListOptions) (result *v1alpha1.DNSAnnottaionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsresourcewatchesResource, dnsresourcewatchesKind, c.ns, opts), &v1alpha1.DNSAnnottaionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DNSAnnottaionList{ListMeta: obj.(*v1alpha1.DNSAnnottaionList).ListMeta}
	for _, item := range obj.(*v1alpha1.DNSAnnottaionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSResourceWatches.
func (c *FakeDNSResourceWatches) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsresourcewatchesResource, c.ns, opts))

}

// Create takes the representation of a dNSResourceWatch and creates it.  Returns the server's representation of the dNSResourceWatch, and an error, if there is any.
func (c *FakeDNSResourceWatches) Create(dNSResourceWatch *v1alpha1.DNSAnnotation) (result *v1alpha1.DNSAnnotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsresourcewatchesResource, c.ns, dNSResourceWatch), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}

// Update takes the representation of a dNSResourceWatch and updates it. Returns the server's representation of the dNSResourceWatch, and an error, if there is any.
func (c *FakeDNSResourceWatches) Update(dNSResourceWatch *v1alpha1.DNSAnnotation) (result *v1alpha1.DNSAnnotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsresourcewatchesResource, c.ns, dNSResourceWatch), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSResourceWatches) UpdateStatus(dNSResourceWatch *v1alpha1.DNSAnnotation) (*v1alpha1.DNSAnnotation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsresourcewatchesResource, "status", c.ns, dNSResourceWatch), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}

// Delete takes name of the dNSResourceWatch and deletes it. Returns an error if one occurs.
func (c *FakeDNSResourceWatches) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnsresourcewatchesResource, c.ns, name), &v1alpha1.DNSAnnotation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSResourceWatches) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsresourcewatchesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DNSAnnottaionList{})
	return err
}

// Patch applies the patch and returns the patched dNSResourceWatch.
func (c *FakeDNSResourceWatches) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DNSAnnotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsresourcewatchesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DNSAnnotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DNSAnnotation), err
}
