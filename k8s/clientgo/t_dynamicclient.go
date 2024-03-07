package main

import (
	"context"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hpfish/.kube/config")
	if err != nil {
		panic(err)
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return
	}

	gvr := schema.GroupVersionResource{
		Version:  "v1",
		Resource: "pods",
	}
	unstructObj, err := dynamicClient.Resource(gvr).Namespace(apiv1.NamespaceDefault).List(context.Background(), metav1.ListOptions{Limit: 50})
	if err != nil {
		return
	}
	podList := &apiv1.PodList{}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), podList)
	if err != nil {
		return
	}

	for _, item := range podList.Items {
		fmt.Println(item)
	}
}
