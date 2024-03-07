package main

import (
	"context"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hpfish/.kube/config")
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}

	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	list, err := podClient.List(context.Background(), metav1.ListOptions{Limit: 500})
	if err != nil {
		return
	}
	fmt.Println(list)

	pod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				apiv1.Container{
					Name:  "test",
					Image: "nginx:latest",
				},
			},
		},
		Status: apiv1.PodStatus{},
	}
	_, err = podClient.Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		return
	}
}
