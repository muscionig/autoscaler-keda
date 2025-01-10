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

	v1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	scheme "knative.dev/autoscaler-keda/pkg/client/clientset/versioned/scheme"
)

// ScaledObjectsGetter has a method to return a ScaledObjectInterface.
// A group's client should implement this interface.
type ScaledObjectsGetter interface {
	ScaledObjects(namespace string) ScaledObjectInterface
}

// ScaledObjectInterface has methods to work with ScaledObject resources.
type ScaledObjectInterface interface {
	Create(ctx context.Context, scaledObject *v1alpha1.ScaledObject, opts v1.CreateOptions) (*v1alpha1.ScaledObject, error)
	Update(ctx context.Context, scaledObject *v1alpha1.ScaledObject, opts v1.UpdateOptions) (*v1alpha1.ScaledObject, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, scaledObject *v1alpha1.ScaledObject, opts v1.UpdateOptions) (*v1alpha1.ScaledObject, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ScaledObject, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ScaledObjectList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScaledObject, err error)
	ScaledObjectExpansion
}

// scaledObjects implements ScaledObjectInterface
type scaledObjects struct {
	*gentype.ClientWithList[*v1alpha1.ScaledObject, *v1alpha1.ScaledObjectList]
}

// newScaledObjects returns a ScaledObjects
func newScaledObjects(c *KedaV1alpha1Client, namespace string) *scaledObjects {
	return &scaledObjects{
		gentype.NewClientWithList[*v1alpha1.ScaledObject, *v1alpha1.ScaledObjectList](
			"scaledobjects",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.ScaledObject { return &v1alpha1.ScaledObject{} },
			func() *v1alpha1.ScaledObjectList { return &v1alpha1.ScaledObjectList{} }),
	}
}
