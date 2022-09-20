/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AdminAdditionConfigSpec defines the desired state of AdminAdditionConfig
type AdminAdditionConfigSpec struct {
	addonv1alpha1.CommonSpec `json:",inline"`
	addonv1alpha1.PatchSpec  `json:",inline"`

	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// AdminAdditionConfigStatus defines the observed state of AdminAdditionConfig
type AdminAdditionConfigStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`

	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AdminAdditionConfig is the Schema for the adminadditionconfigs API
type AdminAdditionConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AdminAdditionConfigSpec   `json:"spec,omitempty"`
	Status AdminAdditionConfigStatus `json:"status,omitempty"`
}

var _ addonv1alpha1.CommonObject = &AdminAdditionConfig{}

func (o *AdminAdditionConfig) ComponentName() string {
	return "adminadditionconfig"
}

func (o *AdminAdditionConfig) CommonSpec() addonv1alpha1.CommonSpec {
	return o.Spec.CommonSpec
}

func (o *AdminAdditionConfig) PatchSpec() addonv1alpha1.PatchSpec {
	return o.Spec.PatchSpec
}

func (o *AdminAdditionConfig) GetCommonStatus() addonv1alpha1.CommonStatus {
	return o.Status.CommonStatus
}

func (o *AdminAdditionConfig) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	o.Status.CommonStatus = s
}

//+kubebuilder:object:root=true

// AdminAdditionConfigList contains a list of AdminAdditionConfig
type AdminAdditionConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AdminAdditionConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdminAdditionConfig{}, &AdminAdditionConfigList{})
}
