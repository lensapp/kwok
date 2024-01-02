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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/kwok/pkg/apis/v1alpha1"
	scheme "sigs.k8s.io/kwok/pkg/client/clientset/versioned/scheme"
)

// ClusterAttachesGetter has a method to return a ClusterAttachInterface.
// A group's client should implement this interface.
type ClusterAttachesGetter interface {
	ClusterAttaches() ClusterAttachInterface
}

// ClusterAttachInterface has methods to work with ClusterAttach resources.
type ClusterAttachInterface interface {
	Create(ctx context.Context, clusterAttach *v1alpha1.ClusterAttach, opts v1.CreateOptions) (*v1alpha1.ClusterAttach, error)
	Update(ctx context.Context, clusterAttach *v1alpha1.ClusterAttach, opts v1.UpdateOptions) (*v1alpha1.ClusterAttach, error)
	UpdateStatus(ctx context.Context, clusterAttach *v1alpha1.ClusterAttach, opts v1.UpdateOptions) (*v1alpha1.ClusterAttach, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterAttach, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterAttachList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterAttach, err error)
	ClusterAttachExpansion
}

// clusterAttaches implements ClusterAttachInterface
type clusterAttaches struct {
	client rest.Interface
}

// newClusterAttaches returns a ClusterAttaches
func newClusterAttaches(c *KwokV1alpha1Client) *clusterAttaches {
	return &clusterAttaches{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterAttach, and returns the corresponding clusterAttach object, and an error if there is any.
func (c *clusterAttaches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterAttach, err error) {
	result = &v1alpha1.ClusterAttach{}
	err = c.client.Get().
		Resource("clusterattaches").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterAttaches that match those selectors.
func (c *clusterAttaches) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterAttachList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterAttachList{}
	err = c.client.Get().
		Resource("clusterattaches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterAttaches.
func (c *clusterAttaches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterattaches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterAttach and creates it.  Returns the server's representation of the clusterAttach, and an error, if there is any.
func (c *clusterAttaches) Create(ctx context.Context, clusterAttach *v1alpha1.ClusterAttach, opts v1.CreateOptions) (result *v1alpha1.ClusterAttach, err error) {
	result = &v1alpha1.ClusterAttach{}
	err = c.client.Post().
		Resource("clusterattaches").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAttach).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterAttach and updates it. Returns the server's representation of the clusterAttach, and an error, if there is any.
func (c *clusterAttaches) Update(ctx context.Context, clusterAttach *v1alpha1.ClusterAttach, opts v1.UpdateOptions) (result *v1alpha1.ClusterAttach, err error) {
	result = &v1alpha1.ClusterAttach{}
	err = c.client.Put().
		Resource("clusterattaches").
		Name(clusterAttach.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAttach).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterAttaches) UpdateStatus(ctx context.Context, clusterAttach *v1alpha1.ClusterAttach, opts v1.UpdateOptions) (result *v1alpha1.ClusterAttach, err error) {
	result = &v1alpha1.ClusterAttach{}
	err = c.client.Put().
		Resource("clusterattaches").
		Name(clusterAttach.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterAttach).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterAttach and deletes it. Returns an error if one occurs.
func (c *clusterAttaches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterattaches").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterAttaches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterattaches").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterAttach.
func (c *clusterAttaches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterAttach, err error) {
	result = &v1alpha1.ClusterAttach{}
	err = c.client.Patch(pt).
		Resource("clusterattaches").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
