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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Key is numa ID
type ResourceInfoMap map[string]ResourceInfo

type ResourceInfo struct {
	Allocatable int `json:"allocatable,omitempty"`
	Capacity    int `json:"capacity,omitempty"`
}

type PolicyName string

const (
	CPUManagerPolicy      PolicyName = "CPUManagerPolicy"
	TopologyManagerPolicy PolicyName = "TopologyManagerPolicy"
)

type ResourceName string

// NumatopoSpec defines the desired state of Numatopo
type NumatopoSpec struct {
	// Specifies the policy of the manager
	// +optional
	Policies map[PolicyName]string `json:"policies,omitempty"`

	// Specifies the numa info for the resource
	// Key is resource name
	// +optional
	NumaResMap map[ResourceName]ResourceInfoMap `json:"numares,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true

// Numatopo is the Schema for the numatopoes API
type Numatopo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the numa information of the worker node
	Spec NumatopoSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// NumatopoList contains a list of Numatopo
type NumatopoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Numatopo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Numatopo{}, &NumatopoList{})
}
