package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("--->WithServerHeader()")
		writer.Header().Set("Server", "HelloServer v0.0.1")
		h(writer, request)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved Reuest %s from %s\n", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}

func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}

func main() {
	//http.Handle("/v1/hello", WithServerHeader(hello))
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServer: ", err)
	//}

}
