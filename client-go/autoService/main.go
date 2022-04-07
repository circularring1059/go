package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	//"time"

	"k8s.io/client-go/informers"
	// "k8s.io/client-go/tools/cache"
	"github.com/go/client-go/autoservice/controller"

	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// // create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// listPods
	// pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{} )
	// if err != nil {
	// 	fmt.Println("list pods fail")
	// }
	// for _, pod := range pods.Items {
	// 	fmt.Println(pod.Name)
	// }

	// get deplyment
	deploymentIns, err := clientset.AppsV1().Deployments("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	if err != nil {
		return
	}
	fmt.Printf("%d\n",len(deploymentIns.Spec.Template.Spec.Containers[0].Ports))
	fmt.Println(deploymentIns.Spec.Selector.MatchLabels)
	
	// //informer
	// factory := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithNamespace("default"))
	// informer := factory.Core().V1().Pods().Informer()

	// //事件处理
	// informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	// 	AddFunc: func(obj interface{}) {
	// 		fmt.Println("add event")
	// 	},

	// 	UpdateFunc: func(oldObj, newObj interface{}) {
	// 		fmt.Println("Update Event")
	// 	},

	// 	DeleteFunc: func(obj interface{}) {
	// 		fmt.Println("delete event")
	// 	},
	// })

	//create informers
	factory := informers.NewSharedInformerFactory(clientset, 0)
	deploymentInformer := factory.Apps().V1().Deployments()
	servicesInformer := factory.Core().V1().Services()

	// fmt.Println(deploymentInformer,  servicesInformer)

	// run informer
	// stopCache := make(chan struct{})
	// factory.Start(stopCache)
	// factory.WaitForCacheSync(stopCache)
	// x := <-stopCache
	// fmt.Println(x)

	fmt.Println("strat job")

	controller := controller.Newcontroller(clientset, deploymentInformer, servicesInformer)
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)

	controller.Run(stopCh)

}

// func  getContainerPort(client, kubernetes.Interface)(int error){
	
// }
