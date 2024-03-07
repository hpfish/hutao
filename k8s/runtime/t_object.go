package main

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
)

func main() {
	pod := &corev1.Pod{
		TypeMeta: v1.TypeMeta{
			Kind: "Pod",
		},
		ObjectMeta: v1.ObjectMeta{
			Labels: map[string]string{"name": "foo"},
		},
	}

	obj := runtime.Object(pod)

	pod2, ok := obj.(*corev1.Pod)
	if !ok {
		panic("unexpected")
	}

	if !reflect.DeepEqual(pod, pod2) {
		panic("unexpected")
	}
	fmt.Println("completed")
}
