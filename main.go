package main

import (
	"context"
	"fmt"
	"github.com/lidongqi/demo-controller/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, _ := clientcmd.BuildConfigFromFlags("", "/home/lidongqi/.kube/config")
	clientSet, _ := versioned.NewForConfig(config)
	nginxClient := clientSet.LdqV1().Nginxes("default")
	nginxs, err := nginxClient.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(nginxs)
}
