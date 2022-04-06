package controller

import (
	"runtime/trace"

	"golang.org/x/vuln/client"
	"k8s.io/client-go/1.5/kubernetes"
	AppsV1 "k8s.io/client-go/apps/v1"
	informerAppV1 "k8s.io/client-go/iformeres/apps/v1"
	informerCoreV1 "k8s.io/client-go/iformeres/core/v1"
	"k8s.io/client-go/kubernetes"
	listerAppV1 "k8s.io/client-go/listers/apps/v1"
	listerCoreV1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	CoreV1 "k8s.ip/client-go/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	workNum = 2
	maxRetry = 3
)

type controller struct {
	client kubernetes.Interface
	deploymentLister listerAppV1.deploymentLister
	serviceLister listerCoreV1.serviceLister
	queue  workqueue.DelayingInterface

}

func (c *controller)enqueue(obj interface{}){
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandldError(err)
	}
	c.queue.Add(key)
	
}
func (c *controller)addDeployment(obj interface{}){
	c.enqueue(obj)
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
	service := obj.(*CoreV1.service)
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
