package pkg

import (
	"reflect"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	coreLister "k8s.io/client-go/listers/core/v1"
	v1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	v12 "k8s.io/api/networking/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	workNum  = 5
	maxRetry = 5
)

type contoller struct {
	client        kubernetes.Interface
	ingressLister v1.IngressClassLister
	serviceLister coreLister.ServiceLister
	queue workqueue.RateLimitingInterface
}


func (c *controller) updateService(oldObj interface{}, newObj interface{}){
	if reflect.DeepEqual(oldObj, newObj){
		return
	}
	c.enqueue(newObj)
}

func deleteIngress(obj interface{}){
	ingress := obj.(*v12.Ingress)
	ownerReference  := v13.GetControllerOf(ingress)
	if ownerReference == nil{
		return
	}
	if ownerReference.Kind !=  "Service" {
		return
	}
	c.queue.Add(igress.namespace +  "/" + ingress.Name)
}

func (c *controller) processNextItem() bool{
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}

	defer c.queue.Done(item)

	key := item(string)
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

func (c *contoller)enqueue(newObj, interface{}){
	key, err :=  cache.MetaNamespaceKeyFunc(newObj)
	if err != nil {
		runtime.HandleError(err)
	}
	c.queue.Add(newObj)
}

func (c *contoller)addService(obj interface{}){
	c.enqueue(obj)
	
}


//构造contreoller
func Newcontroller(client kubernetes.Interface, serviceInformer coreLister.ServiceLister, ingressLister ) concontoller{
	c := concontoller{
		client: client,
		ingressLister: ingressInformer.Lister(),
		serviceLister: serviceInformer.Lister(),
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ingressManager"),
	}
	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: c.addService,
		UpdateFunc: c.updateService,
	})

	ingressInformer.Informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,

	})

	return c

	
}
