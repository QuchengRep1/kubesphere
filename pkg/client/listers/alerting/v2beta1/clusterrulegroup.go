/*
Copyright 2020 The KubeSphere Authors.

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

package v2beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v2beta1 "kubesphere.io/api/alerting/v2beta1"
)

// ClusterRuleGroupLister helps list ClusterRuleGroups.
// All objects returned here must be treated as read-only.
type ClusterRuleGroupLister interface {
	// List lists all ClusterRuleGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta1.ClusterRuleGroup, err error)
	// ClusterRuleGroups returns an object that can list and get ClusterRuleGroups.
	ClusterRuleGroups(namespace string) ClusterRuleGroupNamespaceLister
	ClusterRuleGroupListerExpansion
}

// clusterRuleGroupLister implements the ClusterRuleGroupLister interface.
type clusterRuleGroupLister struct {
	indexer cache.Indexer
}

// NewClusterRuleGroupLister returns a new ClusterRuleGroupLister.
func NewClusterRuleGroupLister(indexer cache.Indexer) ClusterRuleGroupLister {
	return &clusterRuleGroupLister{indexer: indexer}
}

// List lists all ClusterRuleGroups in the indexer.
func (s *clusterRuleGroupLister) List(selector labels.Selector) (ret []*v2beta1.ClusterRuleGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2beta1.ClusterRuleGroup))
	})
	return ret, err
}

// ClusterRuleGroups returns an object that can list and get ClusterRuleGroups.
func (s *clusterRuleGroupLister) ClusterRuleGroups(namespace string) ClusterRuleGroupNamespaceLister {
	return clusterRuleGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterRuleGroupNamespaceLister helps list and get ClusterRuleGroups.
// All objects returned here must be treated as read-only.
type ClusterRuleGroupNamespaceLister interface {
	// List lists all ClusterRuleGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta1.ClusterRuleGroup, err error)
	// Get retrieves the ClusterRuleGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2beta1.ClusterRuleGroup, error)
	ClusterRuleGroupNamespaceListerExpansion
}

// clusterRuleGroupNamespaceLister implements the ClusterRuleGroupNamespaceLister
// interface.
type clusterRuleGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterRuleGroups in the indexer for a given namespace.
func (s clusterRuleGroupNamespaceLister) List(selector labels.Selector) (ret []*v2beta1.ClusterRuleGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2beta1.ClusterRuleGroup))
	})
	return ret, err
}

// Get retrieves the ClusterRuleGroup from the indexer for a given namespace and name.
func (s clusterRuleGroupNamespaceLister) Get(name string) (*v2beta1.ClusterRuleGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2beta1.Resource("clusterrulegroup"), name)
	}
	return obj.(*v2beta1.ClusterRuleGroup), nil
}
