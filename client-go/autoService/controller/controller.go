package controller

import (
	"context"
	"time"
	"fmt"

	ApiAppsV1 "k8s.io/api/apps/v1"
	ApiCoreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	informerAppV1 "k8s.io/client-go/informers/apps/v1"
	informerCoreV1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	listerAppV1 "k8s.io/client-go/listers/apps/v1"
	listerCoreV1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/apimachinery/pkg/util/runtime"

)

const (
	workNum = 2
	maxRetry = 3
)

type controller struct {
	client kubernetes.Interface
	deploymentLister listerAppV1.DeploymentLister
	serviceLister listerCoreV1.ServiceLister
	queue  workqueue.RateLimitingInterface
}

func (c *controller)enqueue(obj interface{}){
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
	}
	c.queue.Add(key)
	
}
func (c *controller)addDeployment(obj interface{}){
	c.enqueue(obj)
}

func (c *controller) processNextItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	
	defer c.queue.Done(item)

	key := item.(string)
	err := c.syncDeployment(key)
	if err != nil {
		c.HandleError(key, err)
	}
	return true
}

func  (c *controller) HandleError (key string, err error){
	if c.queue.NumRequeues(key) <= maxRetry {
		c.queue.AddRateLimited(key)
		return
	}
	runtime.HandleError(err)
	c.queue.Forget(key)
}

func (c *controller) work () {
	for c.processNextItem() {
	}

}

func (c *controller) syncDeployment(key string) error{
	namespacekey, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	deployment, err := c.deploymentLister.Deployments(namespacekey).Get(name)
	if errors.IsNotFound(err) {
		return nil
	}

	if err != nil {
		return err
	}
	_, ok := deployment.GetAnnotations()["createService"]
	service, err := c.serviceLister.Services(namespacekey).Get(name)
	if err != nil && !errors.IsNotFound(err){
		return err
	}
	
	if ok && errors.IsNotFound(err) {
		svc := c.CreateService(deployment)
		_, err := c.client.CoreV1().Services(namespacekey).Create(context.TODO(), svc, metaV1.CreateOptions{})
		if err != nil {
			return err
		}
	}else if !ok  &&  service != nil {
		err := c.client.CoreV1().Services(namespacekey).Delete(context.TODO(), key,  metaV1.DeleteOptions{})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

}

func (c *controller) CreateService(deployment *ApiAppsV1.Deployment) *ApiCoreV1.Service{
	service := &ApiCoreV1.Service{
		ObjectMeta: metaV1.Object{
			Name:  deployment.ObjectMeta.Name,
			Namespace:  deployment.objectMeta.Namespace,
			OwnerReference:  []metaV1.OwnerReference{
				{*metaV1.NewControllerRef(deployment, ApiAppsV1.SchemeGroupVersion.WithKind("Deployment"))},
			},
		},
		Spec: ApiCoreV1.ServiceSpec{
			Ports: []ApiCoreV1.ServicePort{
				{Name: deployment.ObjectMeta.Name, Port: 80, TargetPort: 80,},
			},
			// selector: *ApiAppsV1.Deployment.Spec.Selector.MatchLabels,
			selector: map[string]string{"one": "two"},
		},
	}
}

func (c *controller) updateDeployment(oldObj interface{}, newObj interface{}){
	oldKey, _:= oldObj.(*ApiAppsV1.Deployment).GetAnnotations()["createService"]
	newKey, ok := newObj.(*ApiAppsV1.Deployment).GetAnnotations()["createService"]
	if !ok {
		// add queue
		c.enqueue(newObj)
	}else {
		if ok := CompareInsensitive(oldKey, newKey); !ok{
			//annotation change add  workqueue
			c.enqueue(newObj)
		}
	}

}

func (c *controller) deleteService(obj interface{}){
	service := obj.(*ApiCoreV1.Service)
	ownerReference := metaV1.GetControllerOf(service)
	if ownerReference == nil {
		return
	}
	if ownerReference.Kind !=  "Deployment" {
		return
	}
	
	//add key workqueue 重建services
	c.enqueue(service.ObjectMeta.Namespace + "/" + service.ObjectMeta.Name)
}

func Newcontroller(client kubernetes.Interface, deploymentInformer informerAppV1.DeploymentInformer, serviceInformer informerCoreV1.ServiceInformer ) controller{
	c := controller{
		client: client,
		deploymentLister: deploymentInformer.Lister(),
		serviceLister: serviceInformer.Lister(),
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "queue"),
	}

	//注册事件
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: c.addDeployment,
		UpdateFunc:  c.updateDeployment,
	})

	serviceInformer.informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteService,
	})
}


func CompareInsensitive(str1, str2 string) bool{
	if len(str1) != len(str2) {
		return false
	}

	for i := 0;  i < len(str1); i++{
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func (c *controller) run(stopCache chan struct{}) {
	for i := 0; i < workNum; i++ {
		go wait.Until(c.work, time.Minute, stopCache)
	}
	<- stopCache
}