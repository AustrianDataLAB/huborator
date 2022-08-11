/*
Copyright 2022 Thomas Weber.

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

package v1alpha1

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type KubespawnerOverride struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image                    string              `json:"image,omitempty"`
	CPU_Guarantee            resource.Quantity   `json:"cpu_guarantee,omitempty"`
	CPU_Limit                resource.Quantity   `json:"cpu_limit,omitempty"`
	MEM_Limit                resource.Quantity   `json:"mem_limit,omitempty"`
	MEM_Guarantee            resource.Quantity   `json:"mem_guarantee,omitempty"`
	Environment              *apiextensions.JSON `json:"environment,omitempty"`
	ImagePullPolicy          *apiextensions.JSON `json:"image_pull_policy,omitempty"`
	ImagePullSecrets         *apiextensions.JSON `json:"image_pull_secrets,omitempty"`
	InitContainers           *apiextensions.JSON `json:"init_containers,omitempty"`
	UID                      *apiextensions.JSON `json:"uid,omitempty"`
	GID                      *apiextensions.JSON `json:"gid,omitempty"`
	FS_GID                   *apiextensions.JSON `json:"fs_gid,omitempty"`
	ExtraAnnotations         *apiextensions.JSON `json:"extra_annotations,omitempty"`
	ExtraContainerConfig     *apiextensions.JSON `json:"extra_container_config,omitempty"`
	ExtraContainers          *apiextensions.JSON `json:"extra_containers,omitempty"`
	ExtraLabels              *apiextensions.JSON `json:"extra_labels,omitempty"`
	ExtraPodConfig           *apiextensions.JSON `json:"extra_pod_config,omitempty"`
	ExtraResourceGuarantees  *apiextensions.JSON `json:"extra_resource_guarantees,omitempty"`
	ExtraResourceLimits      *apiextensions.JSON `json:"extra_resource_limits,omitempty"`
	NodeAffinityPreferred    *apiextensions.JSON `json:"node_affinity_preferred,omitempty"`
	NodeAffinityRequired     *apiextensions.JSON `json:"node_affinity_required,omitempty"`
	NodeSelector             *apiextensions.JSON `json:"node_selector,omitempty"`
	PodAffinityPreferred     *apiextensions.JSON `json:"pod_affinity_preferred,omitempty"`
	PodAffinityRequired      *apiextensions.JSON `json:"pod_affinity_required,omitempty"`
	PodAntiAffinityPreferred *apiextensions.JSON `json:"pod_anti_affinity_preferred,omitempty"`
	PodAntiAffinityRequired  *apiextensions.JSON `json:"pod_anti_affinity_required,omitempty"`
	PodSecurityContext       *apiextensions.JSON `json:"pod_security_context,omitempty"`
	LifecycleHooks           *apiextensions.JSON `json:"lifecycle_hooks,omitempty"`
}

type Choice struct {
	DisplayName         string              `json:"display_name"`
	KubespawnerOverride KubespawnerOverride `json:"kubespawner_override"`
}

type ProfileOption struct {
	DisplayName string            `json:"display_name"`
	Choices     map[string]Choice `json:"choices"`
}

// ProfileSpec defines the desired state of Profile
type ProfileSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Slug                string                   `json:"slug"`
	DisplayName         string                   `json:"display_name"`
	AllowedGroups       []string                 `json:"allowed_groups,omitempty"`
	Description         string                   `json:"description,omitempty"`
	KubespawnerOverride KubespawnerOverride      `json:"kubespawner_override,omitempty"`
	ProfileOptions      map[string]ProfileOption `json:"profile_options,omitempty"`
}

// ProfileStatus defines the observed state of Profile
type ProfileStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Profile is the Schema for the profiles API
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProfileSpec   `json:"spec,omitempty"`
	Status ProfileStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProfileList contains a list of Profile
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}
