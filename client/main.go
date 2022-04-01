package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	//list pod
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	//get pod
	pod, err := clientset.CoreV1().Pods("default").Get(context.Background(), "ring", v1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pod.Namespace)

	//delete pod
	// err = clientset.CoreV1().Pods("default").Delete(context.TODO(), "ring", v1.DeleteOptions{})
	// if err != nil {
	// 	fmt.Println("delete pod fail")
	// }

	//create pod
	//构造pod struct
	// podStruct := &corev1.Pod{
	// 	ObjectMeta: v1.ObjectMeta{
	// 		Name:      "ring",
	// 		Namespace: "default",
	// 	},
	// 	Spec: corev1.PodSpec{
	// 		Containers: []corev1.Container{
	// 			{
	// 				Name:  "ring",
	// 				Image: "nginx:1.18",
	// 			},
	// 		},
	// 	},
	// }

	// podIns, err := clientset.CoreV1().Pods("default").Create(context.TODO(), podStruct, v1.CreateOptions{})
	// if err != nil {
	// 	fmt.Println("create pod fail", err) //pods "ring" already exists
	// 	return
	// }

	// fmt.Println(podIns)

	//list deployment
	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), v1.ListOptions{Limit: 1})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("deployments number:", len(deployments.Items))
	for _, deployment := range deployments.Items {
		fmt.Println(deployment.Name)
	}

	//get  deployment
	deployment, err := clientset.AppsV1().Deployments("default").Get(context.TODO(), "etcd", v1.GetOptions{})
	if err != nil {
		fmt.Println("get deployment fail")
	}
	fmt.Println(deployment.Name)

	//create deploy
	// 构造 deployment struct
	dp := &appsv1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name: "ring",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{"ring": "yuan"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Name:   "ring",
					Labels: map[string]string{"ring": "yuan"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:1.18",
						},
					},
				},
			},
		},
	}

	deploymentIns, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), dp, v1.CreateOptions{})
	if err != nil {
		fmt.Println("create deployment fail")
	}
	fmt.Println(deploymentIns)
}

func int32Ptr(i int32) *int32 { return &i }
