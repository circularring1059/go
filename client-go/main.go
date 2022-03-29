package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
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

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//获取default namespace 下的所有pod
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic("list pod faile")
	}

	//fmt.Println(len(pods.Items))

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	//获取单个pod
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "etcd-56f5bf55b6-vqvjc", metav1.GetOptions{})

	if err != nil {
		fmt.Println("not such pod")
	} else {
		fmt.Println("pod:", pod.S)
	}

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

	// //run informer
	// stopCache := make(chan struct{})
	// factory.Start(stopCache)
	// factory.WaitForCacheSync(stopCache)
	// x := <-stopCache
	// fmt.Println(x)

	factory := informers.NewSharedInformerFactory(clientset, 0)
	serviceInformer := factory.Core().V1().Services()
	serviceInformer := factory.Networking().V1().Ingresses()

	controller := pkg.N

}
