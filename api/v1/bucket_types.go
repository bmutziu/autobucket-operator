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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	// Cloud platform
	// +kubebuilder:validation:Enum=gcp
	// +kubebuilder:validation:Required
	Cloud BucketCloud `json:"cloud"`

	// FullName is the cloud storage bucket full name
	// +kubebuilder:validation:Required
	FullName string `json:"fullName"`

	// OnDeletePolicy defines the behavior when the Deployment/Bucket objects are deleted
	// +kubebuilder:validation:Enum=destroy;ignore
	// +kubebuilder:validation:Required
	OnDeletePolicy BucketOnDeletePolicy `json:"onDeletePolicy"`
}

type BucketCloud string

const (
	// BucketCloudGCP gcp cloud
	BucketCloudGCP BucketCloud = "gcp"
)

// BucketStatus defines the observed state of Bucket
type BucketStatus struct {
	// CreatedAt is the cloud storage bucket creation time
	CreatedAt string `json:"createdAt,omitempty"`
}

type BucketOnDeletePolicy string

const (
	// BucketOnDeletePolicyIgnore ignore object deleted
	BucketOnDeletePolicyIgnore BucketOnDeletePolicy = "ignore"
	// BucketOnDeletePolicyDestroy destroy storage bucket on object delete
	BucketOnDeletePolicyDestroy BucketOnDeletePolicy = "destroy"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Cloud",type=string,JSONPath=`.spec.cloud`
// +kubebuilder:printcolumn:name="FullName",type=string,JSONPath=`.spec.fullName`
// +kubebuilder:printcolumn:name="CreatedAt",type=string,JSONPath=`.status.createdAt`

// Bucket is the Schema for the buckets API
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec   `json:"spec,omitempty"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Bucket
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
