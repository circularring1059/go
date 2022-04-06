package controller

import (
	"context"
	"runtime/trace"

	"golang.org/x/vuln/client"
	ApiAppsV1 "k8s.io/api/apps/v1"
	ApiCoreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/pkg/util/runtime"
	AppsV1 "k8s.io/client-go/apps/v1"
	informerAppV1 "k8s.io/client-go/iformeres/apps/v1"
	informerCoreV1 "k8s.io/client-go/iformeres/core/v1"
	"k8s.io/client-go/kubernetes"
	listerAppV1 "k8s.io/client-go/listers/apps/v1"
	listerCoreV1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	workNum = 2
	maxRetry = 3
)

type controller struct {
	client kubernetes.Interface
	deploymentLister listerAppV1.DeploymentLister
	serviceLister listerCoreV1.ServiceLister
	queue  workqueue.DelayingInterface

}

func (c *controller)enqueue(obj interface{}){
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandledError(err)
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

func  (c *controller) HandleError (){
	if c.queue.NumRequeues(key) <= maxRety {
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

func syncDeployment(key string) error{
	namespacekey, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	deployment, err := c.deploymentLister.Deployment(namespacekey).Get(name)
	if errors.IsNotFound(err) {
		return nil
	}

	if err != nil {
		return err
	}
	_, ok := deployment.GetAnnotations()["createService"]
	service, err := c.serviceLister(namespacekey).Get(name)
	if err != nil && !erros.IsNotFound(err){
		return err
	}
	
	if ok && errors.IsNotFound(err) {
		sev := c.createService
		_, err := c.client.CoreV1(namespacekey).Services().Create(context.t, svc, metaV1.CreateOptions{})
		if err != nil {
			return err
		}
	}else if !ok  &&  service != nil {
		err := c.client.CoreV1(namespacekey).Services().Delete(context.T, svc,  metaV1.DeleteOptions{})
	}

}
func (c *controller) CreateService(deployment *ApiAppsV1.Deployment) *ApiCoreV1Service{
	service := &ApiCoreV1.service{
		objectMeat: metaV1.Object{
			Name:  deployment.ojbectMeta.Name,
			Namespace:  deployment.objectMeta.Namespace,
			OwnerReference:  []metaV1.OwnerReference{*metaV1.NewControllerRef(deployment, ApiAppsV1.SchemeGroupVersion.WithKind("Deployment"))}
		}
		spec: ApiCoreV1.ServiceSpec{
			Ports: []ApiCoreV1.ServicePort{
				{Name: deploy.objectMeta.Name, Port: 80, targetPort: 80,}
			},
			selector: *ApiAppsV1.Deployment.Spec.Selector.MatchLabels,
		}

	}
}

func (c *controller) updateDeployment(oldObj interface{}, newObj interface{}){
	oldKey, _:= oldObj.(*AppsV1.Deployment).GetAnnotations()["createService"]
	newKey, ok := newObj.(*AppsV1.Deployment).GetAnnotations()["createService"]
	if !ok {
		// add queue
		c.enqueue(newObj)
	}else {
		if ok := CompareInsensitive(oldkey, newKey); !ok{
			//annotation change add  workqueue
			c.enqueue(newObj)
		}
	}

}

func (c *controller) deleteService(obj interface{}){
	service := obj.(*ApiCoreV1.service)
	ownerReference := metaV1.GetControllerOf(service)
	if ownerReference == nil {
		return
	}
	if ownerReference.Kind !=  "Deployment" {
		return
	}
	
}

func Newcontroller(client kubernetes.Interface, deploymentInformer informerAppV1.deploymentInformer, serviceInformer informerCoreV1.serviceInformer ) controller{
	c := controller{
		client: client,
		deploymentLister: deploymentInformer.Lister(),
		serviceLister: serviceInformer.Listter(),
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
