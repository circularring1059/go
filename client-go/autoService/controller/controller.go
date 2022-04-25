package controller

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/intstr"

	ApiAppsV1 "k8s.io/api/apps/v1"
	ApiCoreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	informerAppV1 "k8s.io/client-go/informers/apps/v1"
	informerCoreV1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	listerAppV1 "k8s.io/client-go/listers/apps/v1"
	listerCoreV1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	workNum  = 2
	maxRetry = 1
)

type controller struct {
	client           kubernetes.Interface
	deploymentLister listerAppV1.DeploymentLister
	serviceLister    listerCoreV1.ServiceLister
	queue            workqueue.RateLimitingInterface
}

func (c *controller) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	// fmt.Println(key)  //namespace + name
	if err != nil {
		runtime.HandleError(err)
	}
	c.queue.Add(key)

}
func (c *controller) addDeployment(obj interface{}) {
	// fmt.Println("listen new deploy")
	c.enqueue(obj)
}

func (c *controller) updateService(oldObj interface{}, newObj interface{}) {
	if ok := reflect.DeepEqual(oldObj, newObj); !ok {
		fmt.Println(newObj)
		// 通过controller 自动创建的services 会带有 ownerReference, 不会重建其他services
		service := newObj.(*ApiCoreV1.Service)
		ownerReference := metaV1.GetControllerOf(service)
		if ownerReference == nil {
			return
		}
		if ownerReference.Kind != "Deployment" {
			return
		}

		//add key workqueue 重建services
		name := strings.Split(service.ObjectMeta.Name, "-") //deployyemt 和services 名字不相同了，这里需要做下处理
		// c.queue.Add("update-" + service.ObjectMeta.Namespace + "/" + name[0])
		c.queue.Add(service.ObjectMeta.Namespace + "/" + name[0])
	}
}

func (c *controller) updateDeployment(oldObj interface{}, newObj interface{}) {
	// newKey, _ := newObj.(*ApiAppsV1.Deployment).GetAnnotations()["createService"]
	// oldKey, _ := oldObj.(*ApiAppsV1.Deployment).GetAnnotations()["createService"]
	isThesameAnnotation := strings.Compare(oldObj.(*ApiAppsV1.Deployment).GetAnnotations()["createService"], newObj.(*ApiAppsV1.Deployment).GetAnnotations()["createService"])
	isTheSamePorts := reflect.DeepEqual(oldObj.(*ApiAppsV1.Deployment).Spec.Template.Spec.Containers[0].Ports, oldObj.(*ApiAppsV1.Deployment).Spec.Template.Spec.Containers[0].Ports)
	if isThesameAnnotation == -1 || isTheSamePorts {
		fmt.Println("update deploy")
		c.enqueue(newObj)
	}
	// if !ok {
	// 	// add queue
	// 	fmt.Println("deployment without annotations")
	// 	c.enqueue(newObj)
	// } else {
	// 	oldKey, _ := oldObj.(*ApiAppsV1.Deployment).GetAnnotations()["createService"]
	// 	isTheSamePorts := reflect.DeepEqual(oldObj.(*ApiAppsV1.Deployment).Spec.Template.Spec.Containers[0].Ports, oldObj.(*ApiAppsV1.Deployment).Spec.Template.Spec.Containers[0].Ports)
	// 	if ok := CompareInsensitive(oldKey, newKey); !ok || !isTheSamePorts {
	// 		//annotation change add  workqueue
	// 		fmt.Println("deployment annotation has changed")
	// 		c.enqueue(newObj)
	// 	} else {
	// 		fmt.Println("deployment has chnaged but annotation not change")
	// 	}
	// }
}

func (c *controller) deleteService(obj interface{}) {
	service := obj.(*ApiCoreV1.Service)
	// 通过controller 自动创建的services 会带有 ownerReference, 不会重建其他services
	ownerReference := metaV1.GetControllerOf(service)
	if ownerReference == nil {
		return
	}
	if ownerReference.Kind != "Deployment" {
		return
	}

	//add key workqueue 重建services
	name := strings.Split(service.ObjectMeta.Name, "-") //deployyemt 和services 名字不相同了，这里需要做下处理
	c.queue.Add(service.ObjectMeta.Namespace + "/" + name[0])
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

func (c *controller) HandleError(key string, err error) {
	if c.queue.NumRequeues(key) <= maxRetry {
		c.queue.AddRateLimited(key)
		return
	}
	runtime.HandleError(err)
	c.queue.Forget(key)
}

func (c *controller) work() {
	for c.processNextItem() {
	}
}

func (c *controller) syncDeployment(key string) error {
	// update := false
	namespacekey, name, err := cache.SplitMetaNamespaceKey(key)
	// if ok := strings.HasPrefix(namespacekey, "update"); ok {
	// 	namespacekey = strings.Split(namespacekey, "-")[1]
	// 	update = true

	// }

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
	created, _ := deployment.GetAnnotations()["createService"]
	service, err := c.serviceLister.Services(namespacekey).Get(name + "-" + "auto-svc")
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	//手动修改services 后该services 将会被删除，调谐至起始状态
	// if update && service != nil {
	// 	err := c.client.CoreV1().Services(namespacekey).Delete(context.TODO(), name+"-"+"auto-svc", metaV1.DeleteOptions{})
	// 	fmt.Println("delete services")
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// if created == "true" && errors.IsNotFound(err) {
	// 	if containerPortLength := len(deployment.Spec.Template.Spec.Containers[0].Ports); containerPortLength == 0 {
	// 		fmt.Println("deployment:", deployment.GetObjectMeta().GetName(), "may not set ports it will not create auto svc")
	// 		return nil
	// 	}
	// 	svc := c.CreateService(deployment)
	// 	_, err := c.client.CoreV1().Services(namespacekey).Create(context.TODO(), svc, metaV1.CreateOptions{})
	// 	// _, err := c.client.CoreV1().Services(namespacekey).A
	// 	if err != nil {
	// 		return err
	// 	}
	// } else if created != "true" && service != nil {
	// 	err := c.client.CoreV1().Services(namespacekey).Delete(context.TODO(), name+"-"+"auto-svc", metaV1.DeleteOptions{})
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// if containerPortLength := len(deployment.Spec.Template.Spec.Containers[0].Ports); containerPortLength == 0 && service != nil{
	// 	err := c.client.CoreV1().Services(namespacekey).Delete(context.TODO(), name+"-"+"auto-svc", metaV1.DeleteOptions{})
	// 	// if err != nil {
	// 	// 	return err
	// 	// }
	// 	return err  //这里直接return
	// }

	// if created == "true" && errors.IsNotFound(err)  {
	// 	if containerPortLength := len(deployment.Spec.Template.Spec.Containers[0].Ports); containerPortLength == 0 {
	// 		fmt.Println("deployment:", deployment.GetObjectMeta().GetName(), "may not set ports it will not create auto svc ***")
	// 		return nil
	// 	}
	// 	svc := c.CreateService(deployment)
	// 	_, err := c.client.CoreV1().Services(namespacekey).Create(context.TODO(), svc, metaV1.CreateOptions{})
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// if created == "true"  && service != nil {
	// 	if containerPortLength := len(deployment.Spec.Template.Spec.Containers[0].Ports); containerPortLength == 0 {
	// 		fmt.Println("deployment:", deployment.GetObjectMeta().GetName(), "may not set ports it will not create auto svc")
	// 		return nil
	// 	}
	// 	//这里将services  删除后会构建新的services 完成更新
	// 	err := c.client.CoreV1().Services(namespacekey).Delete(context.TODO(), name+"-"+"auto-svc", metaV1.DeleteOptions{})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// svc := c.CreateService(deployment)
	// 	// _, err = c.client.CoreV1().Services(namespacekey).Create(context.TODO(), svc, metaV1.CreateOptions{})
	// 	// if err != nil {
	// 	// 	return err
	// 	// }
	// }
	// if created != "true" && service != nil {
	// 	err := c.client.CoreV1().Services(namespacekey).Delete(context.TODO(), name+"-"+"auto-svc", metaV1.DeleteOptions{})
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	containerPortLength := len(deployment.Spec.Template.Spec.Containers[0].Ports)
	// ret := strings.Compare(created, "true")
    if service != nil {
		if containerPortLength <= 0 || strings.Compare(created, "true") == 0 {
			 err := c.client.CoreV1().Services(namespacekey).Delete(context.TODO(), name + "-" + "auto-svc", metaV1.DeleteOptions{})
			 if err != nil {
				 return err
			 }
		}
	}
	
	if containerPortLength > 0 && strings.Compare(created, "true") == 0 && errors.IsNotFound(err) {
		// 创建 service
		svc := c.CreateService(deployment)
		_, err := c.client.CoreV1().Services(namespacekey).Create(context.TODO(), svc, metaV1.CreateOptions{})
		if err != nil {
			return nil
		}
	}

	if containerPortLength >0 && strings.Compare(created, "true") == 0 && service != nil {
		// 删除servie 产生 delete event  
		err :=  c.client.CoreV1().Services(namespacekey).Delete(context.TODO(), name + "-" + "auto-svc", metaV1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	return nil

}

func (c *controller) CreateService(deployment *ApiAppsV1.Deployment) *ApiCoreV1.Service {
	containerPortLength := len(deployment.Spec.Template.Spec.Containers[0].Ports)
	//获取第一个container 下的ports 作为services 的port
	ports := make([]ApiCoreV1.ServicePort, containerPortLength, containerPortLength)
	for i := 0; i < containerPortLength; i++ {
		// map1 = append(map1, ApiCoreV1.ServicePort{
		// 	Name: deployment.ObjectMeta.Name + strconv.Itoa(i), Port: deployment.Spec.Template.Spec.Containers[0].Ports[i].ContainerPort, TargetPort: intstr.IntOrString{
		// 		Type:   intstr.Int,
		// 		IntVal: deployment.Spec.Template.Spec.Containers[0].Ports[i].ContainerPort,},
		// },)
		ports[i] = ApiCoreV1.ServicePort{
			Name: deployment.ObjectMeta.Name + "-" + strconv.Itoa(i), Port: deployment.Spec.Template.Spec.Containers[0].Ports[i].ContainerPort, TargetPort: intstr.IntOrString{
				Type:   intstr.Int,
				IntVal: deployment.Spec.Template.Spec.Containers[0].Ports[i].ContainerPort},
		}
	}

	service := &ApiCoreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      deployment.ObjectMeta.Name + "-" + "auto-svc",
			Namespace: deployment.ObjectMeta.Namespace,
			OwnerReferences: []metaV1.OwnerReference{
				*metaV1.NewControllerRef(deployment, ApiAppsV1.SchemeGroupVersion.WithKind("Deployment")),
			},
		},
		Spec: ApiCoreV1.ServiceSpec{
			// Ports: []ApiCoreV1.ServicePort{
			// 	{Name: deployment.ObjectMeta.Name, Port: deployment.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort, TargetPort: intstr.IntOrString{
			// 		Type:   intstr.Int,
			// 		IntVal: deployment.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort,
			// 	}},
			// },
			Ports:    ports,
			Selector: deployment.Spec.Selector.MatchLabels,
		},
	}

	// service.Spec = ApiCoreV1.ServiceSpec{
	// 	Ports: ports,
	// }
	return service
}

func Newcontroller(client kubernetes.Interface, deploymentInformer informerAppV1.DeploymentInformer, serviceInformer informerCoreV1.ServiceInformer) controller {
	c := controller{
		client:           client,
		deploymentLister: deploymentInformer.Lister(),
		serviceLister:    serviceInformer.Lister(),
		queue:            workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "queue"),
	}

	//注册事件
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addDeployment,
		UpdateFunc: c.updateDeployment,
	})

	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteService,
		UpdateFunc: c.updateService,
	})

	return c
}

// func CompareInsensitive(str1, str2 string) bool {
// 	if len(str1) != len(str2) {
// 		return false
// 	}

// 	for i := 0; i < len(str1); i++ {
// 		if str1[i] != str2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func (c *controller) Run(stopCache chan struct{}) {
	for i := 0; i < workNum; i++ {
		go wait.Until(c.work, time.Minute, stopCache)
	}
	<-stopCache
}
