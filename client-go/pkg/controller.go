package pkg

import (
	"k8s.io/client-go/kubernetes"
	coreLister "k8s.io/client-go/listers/core/v1"
	v1 "k8s.io/client-go/listers/networking/v1"
)

const (
	workNum  = 5
	maxRetry = 5
)

type contoller struct {
	client        kubernetes.Interface
	ingressLister v1.IngressClassLister
	serviceLister coreLister.ServiceLister
}
