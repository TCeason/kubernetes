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

package v1beta1

import (
	context "context"

	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsrbacv1beta1 "k8s.io/client-go/applyconfigurations/rbac/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// ClusterRoleBindingsGetter has a method to return a ClusterRoleBindingInterface.
// A group's client should implement this interface.
type ClusterRoleBindingsGetter interface {
	ClusterRoleBindings() ClusterRoleBindingInterface
}

// ClusterRoleBindingInterface has methods to work with ClusterRoleBinding resources.
type ClusterRoleBindingInterface interface {
	Create(ctx context.Context, clusterRoleBinding *rbacv1beta1.ClusterRoleBinding, opts v1.CreateOptions) (*rbacv1beta1.ClusterRoleBinding, error)
	Update(ctx context.Context, clusterRoleBinding *rbacv1beta1.ClusterRoleBinding, opts v1.UpdateOptions) (*rbacv1beta1.ClusterRoleBinding, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*rbacv1beta1.ClusterRoleBinding, error)
	List(ctx context.Context, opts v1.ListOptions) (*rbacv1beta1.ClusterRoleBindingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *rbacv1beta1.ClusterRoleBinding, err error)
	Apply(ctx context.Context, clusterRoleBinding *applyconfigurationsrbacv1beta1.ClusterRoleBindingApplyConfiguration, opts v1.ApplyOptions) (result *rbacv1beta1.ClusterRoleBinding, err error)
	ClusterRoleBindingExpansion
}

// clusterRoleBindings implements ClusterRoleBindingInterface
type clusterRoleBindings struct {
	*gentype.ClientWithListAndApply[*rbacv1beta1.ClusterRoleBinding, *rbacv1beta1.ClusterRoleBindingList, *applyconfigurationsrbacv1beta1.ClusterRoleBindingApplyConfiguration]
}

// newClusterRoleBindings returns a ClusterRoleBindings
func newClusterRoleBindings(c *RbacV1beta1Client) *clusterRoleBindings {
	return &clusterRoleBindings{
		gentype.NewClientWithListAndApply[*rbacv1beta1.ClusterRoleBinding, *rbacv1beta1.ClusterRoleBindingList, *applyconfigurationsrbacv1beta1.ClusterRoleBindingApplyConfiguration](
			"clusterrolebindings",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *rbacv1beta1.ClusterRoleBinding { return &rbacv1beta1.ClusterRoleBinding{} },
			func() *rbacv1beta1.ClusterRoleBindingList { return &rbacv1beta1.ClusterRoleBindingList{} },
			gentype.PrefersProtobuf[*rbacv1beta1.ClusterRoleBinding](),
		),
	}
}
