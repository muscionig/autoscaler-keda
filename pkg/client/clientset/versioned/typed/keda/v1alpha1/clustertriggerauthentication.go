/*
Copyright 2024 The Knative Authors

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

	v1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "knative.dev/autoscaler-keda/pkg/client/clientset/versioned/scheme"
)

// ClusterTriggerAuthenticationsGetter has a method to return a ClusterTriggerAuthenticationInterface.
// A group's client should implement this interface.
type ClusterTriggerAuthenticationsGetter interface {
	ClusterTriggerAuthentications() ClusterTriggerAuthenticationInterface
}

// ClusterTriggerAuthenticationInterface has methods to work with ClusterTriggerAuthentication resources.
type ClusterTriggerAuthenticationInterface interface {
	Create(ctx context.Context, clusterTriggerAuthentication *v1alpha1.ClusterTriggerAuthentication, opts v1.CreateOptions) (*v1alpha1.ClusterTriggerAuthentication, error)
	Update(ctx context.Context, clusterTriggerAuthentication *v1alpha1.ClusterTriggerAuthentication, opts v1.UpdateOptions) (*v1alpha1.ClusterTriggerAuthentication, error)
	UpdateStatus(ctx context.Context, clusterTriggerAuthentication *v1alpha1.ClusterTriggerAuthentication, opts v1.UpdateOptions) (*v1alpha1.ClusterTriggerAuthentication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterTriggerAuthentication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterTriggerAuthenticationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterTriggerAuthentication, err error)
	ClusterTriggerAuthenticationExpansion
}

// clusterTriggerAuthentications implements ClusterTriggerAuthenticationInterface
type clusterTriggerAuthentications struct {
	client rest.Interface
}

// newClusterTriggerAuthentications returns a ClusterTriggerAuthentications
func newClusterTriggerAuthentications(c *KedaV1alpha1Client) *clusterTriggerAuthentications {
	return &clusterTriggerAuthentications{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterTriggerAuthentication, and returns the corresponding clusterTriggerAuthentication object, and an error if there is any.
func (c *clusterTriggerAuthentications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterTriggerAuthentication, err error) {
	result = &v1alpha1.ClusterTriggerAuthentication{}
	err = c.client.Get().
		Resource("clustertriggerauthentications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterTriggerAuthentications that match those selectors.
func (c *clusterTriggerAuthentications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterTriggerAuthenticationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterTriggerAuthenticationList{}
	err = c.client.Get().
		Resource("clustertriggerauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterTriggerAuthentications.
func (c *clusterTriggerAuthentications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustertriggerauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterTriggerAuthentication and creates it.  Returns the server's representation of the clusterTriggerAuthentication, and an error, if there is any.
func (c *clusterTriggerAuthentications) Create(ctx context.Context, clusterTriggerAuthentication *v1alpha1.ClusterTriggerAuthentication, opts v1.CreateOptions) (result *v1alpha1.ClusterTriggerAuthentication, err error) {
	result = &v1alpha1.ClusterTriggerAuthentication{}
	err = c.client.Post().
		Resource("clustertriggerauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTriggerAuthentication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterTriggerAuthentication and updates it. Returns the server's representation of the clusterTriggerAuthentication, and an error, if there is any.
func (c *clusterTriggerAuthentications) Update(ctx context.Context, clusterTriggerAuthentication *v1alpha1.ClusterTriggerAuthentication, opts v1.UpdateOptions) (result *v1alpha1.ClusterTriggerAuthentication, err error) {
	result = &v1alpha1.ClusterTriggerAuthentication{}
	err = c.client.Put().
		Resource("clustertriggerauthentications").
		Name(clusterTriggerAuthentication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTriggerAuthentication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterTriggerAuthentications) UpdateStatus(ctx context.Context, clusterTriggerAuthentication *v1alpha1.ClusterTriggerAuthentication, opts v1.UpdateOptions) (result *v1alpha1.ClusterTriggerAuthentication, err error) {
	result = &v1alpha1.ClusterTriggerAuthentication{}
	err = c.client.Put().
		Resource("clustertriggerauthentications").
		Name(clusterTriggerAuthentication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTriggerAuthentication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterTriggerAuthentication and deletes it. Returns an error if one occurs.
func (c *clusterTriggerAuthentications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustertriggerauthentications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterTriggerAuthentications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustertriggerauthentications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterTriggerAuthentication.
func (c *clusterTriggerAuthentications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterTriggerAuthentication, err error) {
	result = &v1alpha1.ClusterTriggerAuthentication{}
	err = c.client.Patch(pt).
		Resource("clustertriggerauthentications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
