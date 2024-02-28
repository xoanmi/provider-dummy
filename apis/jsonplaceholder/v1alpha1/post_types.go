/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// PostParameters are the configurable fields of a Post.
type PostParameters struct {
	// Title is the title of the Post
	Title string `json:"title,omitempty"`
	// Body is the body of the Post
	Body string `json:"body,omitempty"`
	// UserId is the user id of the Post
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default=1
	UserID int64 `json:"userId,omitempty"`
}

// PostObservation are the observable fields of a Post.
type PostObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A PostSpec defines the desired state of a Post.
type PostSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PostParameters `json:"forProvider"`
}

// A PostStatus represents the observed state of a Post.
type PostStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Post is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dummy}
type Post struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostSpec   `json:"spec"`
	Status PostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostList contains a list of Post
type PostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Post `json:"items"`
}

// Post type metadata.
var (
	PostKind             = reflect.TypeOf(Post{}).Name()
	PostGroupKind        = schema.GroupKind{Group: Group, Kind: PostKind}.String()
	PostKindAPIVersion   = PostKind + "." + SchemeGroupVersion.String()
	PostGroupVersionKind = SchemeGroupVersion.WithKind(PostKind)
)

func init() {
	SchemeBuilder.Register(&Post{}, &PostList{})
}
