package pkg

import (
	"context"
	"reflect"
	"time"
	"fmt"

	informer "k8s.io/client-go/informers/core/v1"

	apiCoreV1 "k8s.io/api/core/v1"
	apiNetV1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	netInformer "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	listenCoreV1 "k8s.io/client-go/listers/core/v1"
	listenNetV1 "k8s.io/client-go/listers/networking/v1"
	cache "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	workNum  = 1
	maxRetry = 5
)

type controller struct {
	client        kubernetes.Interface
	ingressLister listenNetV1.IngressLister
	serviceLister listenCoreV1.ServiceLister
	queue         workqueue.RateLimitingInterface
}

func (c *controller) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)  // key = "namespace"  + "/"" + "name"
	if err != nil {
		runtime.HandleError(err)
	}
	c.queue.Add(key)
}

func (c *controller) addService(obj interface{}) {
	c.enqueue(obj)
}

func (c *controller) updateService(oldObj interface{}, newObj interface{}) {
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	fmt.Println("services update workqueue add")
	c.enqueue(newObj)
}

func (c *controller) deleteIngress(obj interface{}) {
	ingress := obj.(*apiNetV1.Ingress)
	ownerReference := metaV1.GetControllerOf(ingress)  //func GetControllerOf(controllee Object) *OwnerReference
	if ownerReference == nil {
		return
	}
	if ownerReference.Kind != "Service" {
		return
	}
	c.queue.Add(ingress.ObjectMeta.Namespace + "/" + ingress.ObjectMeta.Name)
}

func (c *controller) processNextItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}

	defer c.queue.Done(item)

	key := item.(string)
	err := c.syncService(key)
	if err != nil {
		c.HandleError(key, err)
	}
	return true
}

func (c *controller) worker() {
	for c.processNextItem() {
	}
}

func (c *controller) Run(stopCache chan struct{}) {
	for i := 0; i < workNum; i++ {
		go wait.Until(c.worker, time.Minute, stopCache)
	}
	<-stopCache
}

func (c *controller) syncService(key string) error {
	namespaceKey, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	//delete  //判断该services  是否存在
	service, err := c.serviceLister.Services(namespaceKey).Get(name) //return services resource struct
	if errors.IsNotFound(err) {
		return nil
	}
	
	if err != nil {
		return err
	}

	//add && delete
	_, ok := service.GetAnnotations()["ingress/http"]  //func (meta *ObjectMeta) GetAnnotations() map[string]string            { return meta.Annotations }
	ingress, err := c.ingressLister.Ingresses(namespaceKey).Get(name)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	if ok && errors.IsNotFound(err) {
		ig := c.constructIngress(service)
		_, err := c.client.NetworkingV1().Ingresses(namespaceKey).Create(context.TODO(), ig, metaV1.CreateOptions{})
		if err != nil {
			fmt.Println(err) 
			return err
		}
	} else if !ok && ingress != nil {
		err := c.client.NetworkingV1().Ingresses(namespaceKey).Delete(context.TODO(), name, metaV1.DeleteOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *controller) HandleError(key string, err error) {
	if c.queue.NumRequeues(key) <= maxRetry {
		c.queue.AddRateLimited(key)
		return
	}

	runtime.HandleError(err)
	c.queue.Forget(key)
}

func (c *controller) constructIngress(service *apiCoreV1.Service) *apiNetV1.Ingress {
	ingress := apiNetV1.Ingress{}

	ingress.ObjectMeta.OwnerReferences = []metaV1.OwnerReference{
		*metaV1.NewControllerRef(service, apiCoreV1.SchemeGroupVersion.WithKind("Service")),
	}

	ingress.ObjectMeta.Name = service.Name
	ingress.ObjectMeta.Namespace = service.Namespace
	pathType := apiNetV1.PathTypePrefix
	icn := "nginx"
	ingress.Spec = apiNetV1.IngressSpec{
		IngressClassName: &icn,
		Rules: []apiNetV1.IngressRule{
			{
				Host: "test-ingress.com",
				IngressRuleValue: apiNetV1.IngressRuleValue{
					HTTP: &apiNetV1.HTTPIngressRuleValue{
						Paths: []apiNetV1.HTTPIngressPath{
							{
								Path:     "/",
								PathType: &pathType,
								Backend: apiNetV1.IngressBackend{
									Service: &apiNetV1.IngressServiceBackend{
										Name: service.Name,
										Port: apiNetV1.ServiceBackendPort{
											// Name: "ring",
											Number: 80,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return &ingress

}

//构造contreoller
func Newcontroller(client kubernetes.Interface, serviceInformer informer.ServiceInformer, ingressInformer netInformer.IngressInformer) controller {
	c := controller{
		client:        client,
		ingressLister: ingressInformer.Lister(),
		serviceLister: serviceInformer.Lister(),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ingressManager"),
	}
	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})

	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})

	return c

}
