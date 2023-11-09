package main

import (
	"context"
	"fmt"
	"github.com/lidongqi/demo-controller/generated/clientset/versioned"
	nginxinformer "github.com/lidongqi/demo-controller/generated/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	Watch()
}

func List() {
	config, _ := clientcmd.BuildConfigFromFlags("", "/home/lidongqi/.kube/config")
	clientSet, _ := versioned.NewForConfig(config)
	nginxClient := clientSet.LdqV1().Nginxes("default")
	nginxs, err := nginxClient.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, ng := range nginxs.Items {
		fmt.Println("nginx", ng.Name)
	}
}

func Watch() {
	config, _ := clientcmd.BuildConfigFromFlags("", "/home/lidongqi/.kube/config")
	ngclientSet, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	//Create nginx informer factory
	ngInformerFactory := nginxinformer.NewSharedInformerFactoryWithOptions(ngclientSet, 0)

	//Get the nginx informer
	ngInformer := ngInformerFactory.Ldq().V1().Nginxes().Informer()

	ngInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("nginx added", obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("nginx update", oldObj, newObj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("nginx delete", obj)
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	go ngInformer.Run(stopCh)

	select {
	case val := <-stopCh:
		fmt.Println(val)
	}
}
