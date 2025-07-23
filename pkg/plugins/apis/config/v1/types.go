/*
Copyright 2024 The Gocrane Authors.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DynamicArgs is the args struction of Dynamic scheduler plugin.
type DynamicArgs struct {
	metav1.TypeMeta `json:",inline"`
	// PolicyConfigPath specified the path of policy config.
	PolicyConfigPath *string `json:"policyConfigPath,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeResourceTopologyMatchArgs holds arguments used to configure the NodeResourceTopologyMatch plugin.
type NodeResourceTopologyMatchArgs struct {
	metav1.TypeMeta `json:",inline"`
	// TopologyAwareResources represents the resource names of topology.
	TopologyAwareResources []string `json:"topologyAwareResources,omitempty"`
}
