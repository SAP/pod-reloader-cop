/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and pod-reloader-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	"github.com/sap/component-operator-runtime/pkg/component"
	componentoperatorruntimetypes "github.com/sap/component-operator-runtime/pkg/types"
)

// PodReloaderSpec defines the desired state of PodReloader.
type PodReloaderSpec struct {
	component.Spec `json:",inline"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1
	ReplicaCount int `json:"replicaCount,omitempty"`
	// +optional
	Image                          component.ImageSpec `json:"image"`
	component.KubernetesProperties `json:",inline"`
	ObjectSelector                 *metav1.LabelSelector `json:"objectSelector,omitempty"`
	NamespaceSelector              *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
}

// PodReloaderStatus defines the observed state of PodReloader.
type PodReloaderStatus struct {
	component.Status `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +genclient

// PodReloader is the Schema for the podreloaders API.
type PodReloader struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PodReloaderSpec `json:"spec,omitempty"`
	// +kubebuilder:default={"observedGeneration":-1}
	Status PodReloaderStatus `json:"status,omitempty"`
}

var _ component.Component = &PodReloader{}

// +kubebuilder:object:root=true

// PodReloaderList contains a list of PodReloader.
type PodReloaderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodReloader `json:"items"`
}

func (s *PodReloaderSpec) ToUnstructured() map[string]any {
	result, err := runtime.DefaultUnstructuredConverter.ToUnstructured(s)
	if err != nil {
		panic(err)
	}
	return result
}

func (c *PodReloader) GetDeploymentNamespace() string {
	if c.Spec.Namespace != "" {
		return c.Spec.Namespace
	}
	return c.Namespace
}

func (c *PodReloader) GetDeploymentName() string {
	if c.Spec.Name != "" {
		return c.Spec.Name
	}
	return c.Name
}

func (c *PodReloader) GetSpec() componentoperatorruntimetypes.Unstructurable {
	return &c.Spec
}

func (c *PodReloader) GetStatus() *component.Status {
	return &c.Status.Status
}

func init() {
	SchemeBuilder.Register(&PodReloader{}, &PodReloaderList{})
}
