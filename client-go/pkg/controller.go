package pkg

import (
	"context"
	"reflect"
	"time"

	v14 "k8s.io/api/core/v1"
	v12 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	coreLister "k8s.io/client-go/listers/core/v1"
	v1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
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

func (c *controller) Run(stopCache struct{}) {
	for i := 0;  i < workNum; i++{
		go wait.Until(c.woker, time.Minute, stopCache)
	}
	<-  stopCache
}

func (c *controller)syncService(key string) error{
	namespaceKey, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return  err
	}

	//delete
	service, err := c.ServiceLister.Service(namespaceKey).Get(name)
	if errors.IsNotFound(err) {
		return nil
	}

	if err != nil{
		return  err
	}

	//add && delete 
	_, ok  := service.GetAnnotaions()["ingress/http"]
	ingress, err := c.ingressLister.Ingress(namespaceKey).Get(name)
	if err !=  nil &&  !erros.IsNotFound(err) {
		return err
	}

	if ok && errors.IsNotFound(err) {
		ig  := c.constructIngress(service)
		_, err := c.client.NetworkingV1().Ingress(namespaceKey).Create(context.TODO(), ig, v13.CreateOptions{})
		if err != nil {
			return err
		}
	}else if !ok && ingress != nil {
		err := c.client.NetworkingV1.Ingress(namespaceKey).delete(context.TODO(), name, v13.DeleteOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func  (c *controller)HandleError(key string, err error){
	if c.queue.NumRequeues(key) <=  maxRetry {
		c.queue.AddRateLimited(key)
		return
	}

	runtime.HandleError(err)
	c.queue.Forget(key)
}

func  (c *controller) constructIngress(service *v14.Service) *v12.Ingress{
	ingress := v12.Ingress{}

	ingress.ObjectMeta.OwnerReferences =  []v13.OwnerReference{
		*v13.NewControllerRef(service, v14.SchemeGroupVersion.WithKind("Service")),
	}

	ingress.Name = service.Name
	ingress.Namespace = service.Namespace
	pathType := v12.PathType
	icn := "nginx"
	
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
