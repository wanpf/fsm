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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/plugin/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PluginChainLister helps list PluginChains.
// All objects returned here must be treated as read-only.
type PluginChainLister interface {
	// List lists all PluginChains in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PluginChain, err error)
	// PluginChains returns an object that can list and get PluginChains.
	PluginChains(namespace string) PluginChainNamespaceLister
	PluginChainListerExpansion
}

// pluginChainLister implements the PluginChainLister interface.
type pluginChainLister struct {
	indexer cache.Indexer
}

// NewPluginChainLister returns a new PluginChainLister.
func NewPluginChainLister(indexer cache.Indexer) PluginChainLister {
	return &pluginChainLister{indexer: indexer}
}

// List lists all PluginChains in the indexer.
func (s *pluginChainLister) List(selector labels.Selector) (ret []*v1alpha1.PluginChain, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PluginChain))
	})
	return ret, err
}

// PluginChains returns an object that can list and get PluginChains.
func (s *pluginChainLister) PluginChains(namespace string) PluginChainNamespaceLister {
	return pluginChainNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PluginChainNamespaceLister helps list and get PluginChains.
// All objects returned here must be treated as read-only.
type PluginChainNamespaceLister interface {
	// List lists all PluginChains in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PluginChain, err error)
	// Get retrieves the PluginChain from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PluginChain, error)
	PluginChainNamespaceListerExpansion
}

// pluginChainNamespaceLister implements the PluginChainNamespaceLister
// interface.
type pluginChainNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PluginChains in the indexer for a given namespace.
func (s pluginChainNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PluginChain, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PluginChain))
	})
	return ret, err
}

// Get retrieves the PluginChain from the indexer for a given namespace and name.
func (s pluginChainNamespaceLister) Get(name string) (*v1alpha1.PluginChain, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pluginchain"), name)
	}
	return obj.(*v1alpha1.PluginChain), nil
}
