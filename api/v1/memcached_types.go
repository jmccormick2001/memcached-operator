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

type WatchSocket struct {
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Scheme    string   `json:"scheme"`
	Stocks    []string `json:"stocks"`
	Tablename string   `json:"tablename"`
}

type Endpoint struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Scheme string `json:"scheme"`
}

type Source struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	Path      string `json:"path"`
	Port      int    `json:"port"`
	Scheme    string `json:"scheme"`
	Username  string `json:"username"`
	Database  string `json:"database"`
	Tablename string `json:"tablename"`
}
type DBCreds struct {
	CAKey         string `json:"cakey"`
	CACrt         string `json:"cacrt"`
	NodeKey       string `json:"nodekey"`
	NodeCrt       string `json:"nodecrt"`
	ClientRootCrt string `json:"clientrootcrt"`
	ClientRootKey string `json:"clientrootkey"`
	PipelineCrt   string `json:"pipelinecrt"`
	PipelineKey   string `json:"pipelinekey"`
}
type ServiceCreds struct {
	ServiceCrt string `json:"servicecrt"`
	ServiceKey string `json:"servicekey"`
}

type WatchConfigStruct struct {
	Location Endpoint `json:"location"`
}

type LoaderConfigStruct struct {
	Location    Endpoint `json:"location"`
	QueueSize   int      `json:"queueSize"`
	PctHeadRoom int      `json:"pctHeadRoom"`
	DataSource  Source   `json:"dataSource"`
}

// MemcachedSpec defines the desired state of Memcached
type MemcachedSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Memcached. Edit Memcached_types.go to remove/update
	Id   string `json:"id"`
	Port int    `json:"port"`

	AdminDataSource Source             `json:"adminDataSource,omitempty"`
	DataSource      Source             `json:"dataSource,omitempty"`
	WatchSockets    []WatchSocket      `json:"watchSockets"`
	WatchConfig     WatchConfigStruct  `json:"watchConfig"`
	LoaderConfig    LoaderConfigStruct `json:"loaderConfig"`
	/**
	WatchConfig     struct {
		Location Endpoint `json:"location"`
	} `json:"watchConfig"`
	LoaderConfig struct {
		Location    Endpoint `json:"location"`
		QueueSize   int      `json:"queueSize"`
		PctHeadRoom int      `json:"pctHeadRoom"`
		DataSource  Source   `json:"dataSource"`
	} `json:"loaderConfig"`
	*/
	DatabaseCredentials DBCreds      `json:"dbcreds,omitempty"`
	ServiceCredentials  ServiceCreds `json:"servicecreds,omitempty"`

	Foo string `json:"foo,omitempty"`
	Goo string `json:"goo,omitempty"`
}

// MemcachedStatus defines the observed state of Memcached
type MemcachedStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Active string `json:"active"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Memcached is the Schema for the memcacheds API
type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec   `json:"spec,omitempty"`
	Status MemcachedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemcachedList contains a list of Memcached
type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcached `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Memcached{}, &MemcachedList{})
}
