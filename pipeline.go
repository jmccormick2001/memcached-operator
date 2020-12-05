package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	v1 "github.com/example-inc/memcached-operator/api/v1"
	"go.uber.org/zap"
	"k8s.io/api/node/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	_, config, err := GetKubeClient()
	if err != nil {
		panic(err)
	}

	ns := "default"
	crClient, err := NewClient(config, ns)
	if err != nil {
		panic(err)
	}

	cr := v1.Memcached{}
	cr.ObjectMeta.Name = "doo"
	cr.Spec.Foo = "dooby"
	cr.Spec.Goo = "gooby"
	result, err := crClient.Create(&cr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("result cr %+v\n", result)
}

func GetKubeClient() (client *kubernetes.Clientset, config *rest.Config, err error) {

	config, err = ctrl.GetConfig()
	if err != nil {
		return client, config, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return client, config, err
	}

	return clientset, config, err
}

func GetMemcached(logger *zap.SugaredLogger) (v1.Memcached, error) {

	instance := v1.Memcached{}
	ns := os.Getenv("CHURRO_NAMESPACE")
	pipelineName := os.Getenv("CHURRO_PIPELINE")
	if ns == "" {
		return instance, errors.New("CHURRO_PIPELINE not set")
	}
	if pipelineName == "" {
		return instance, errors.New("CHURRO_PIPELINE not set")
	}
	logger.Infof("CHURRO_NAMESPACE %s\n", ns)
	logger.Infof("CHURRO_PIPELINE %s\n", pipelineName)

	_, cfg, err := GetKubeClient()
	if err != nil {
		return instance, err
	}

	err = v1alpha1.AddToScheme(clientgoscheme.Scheme)
	if err != nil {
		return instance, err
	}

	var k8sClient client.Client
	k8sClient, err = client.New(cfg, client.Options{Scheme: clientgoscheme.Scheme})
	if err != nil {
		return instance, err
	}

	namespacedName := types.NamespacedName{
		Namespace: ns,
		Name:      pipelineName,
	}

	err = k8sClient.Get(context.TODO(), namespacedName, &instance)
	if err != nil {
		return instance, err
	}
	return instance, nil
}

var SchemeGroupVersion = schema.GroupVersion{Group: "cache.example.com", Version: "v1"}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&v1.Memcached{},
		&v1.MemcachedList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func NewClient(cfg *rest.Config, namespace string) (*MemcachedClient, error) {
	scheme := runtime.NewScheme()
	SchemeBuilder := runtime.NewSchemeBuilder(addKnownTypes)
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, err
	}
	config := *cfg
	config.GroupVersion = &SchemeGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MemcachedClient{restClient: client, ns: namespace}, nil
}

type MemcachedClient struct {
	restClient rest.Interface
	ns         string
}

func (c *MemcachedClient) Get(name string) (*v1.Memcached, error) {
	result := &v1.Memcached{}
	err := c.restClient.Get().
		Namespace(c.ns).Resource("memcacheds").
		Name(name).Do(context.TODO()).Into(result)
	return result, err
}
func (c *MemcachedClient) Create(obj *v1.Memcached) (*v1.Memcached, error) {
	result := &v1.Memcached{}
	err := c.restClient.Post().
		Namespace(c.ns).Resource("memcacheds").
		Body(obj).Do(context.TODO()).Into(result)
	return result, err
}

func (c *MemcachedClient) Update(obj *v1.Memcached) (*v1.Memcached, error) {
	result := &v1.Memcached{}
	err := c.restClient.Put().
		Namespace(c.ns).Resource("memcacheds").
		Body(obj).Do(context.TODO()).Into(result)
	return result, err
}

func (c *MemcachedClient) Delete(name string, options *metav1.DeleteOptions) error {
	return c.restClient.Delete().
		Namespace(c.ns).Resource("memcacheds").
		Name(name).Body(options).Do(context.TODO()).
		Error()
}
