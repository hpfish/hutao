package main

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"time"
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

	stopCh := make(chan struct{})
	defer close(stopCh)

	sharedInformerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)
	informer := sharedInformerFactory.Core().V1().Pods().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			object := obj.(v1.Object)
			fmt.Printf("New pod added to store: %s \n", object.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(v1.Object)
			nObj := newObj.(v1.Object)
			fmt.Printf("%s pod Updated to %s \n", oObj.GetName(), nObj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			object := obj.(v1.Object)
			fmt.Printf("Pod deleted from store: %s \n", object.GetName())
		},
	})
	informer.Run(stopCh)
}
