package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func protect(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Protect call")
		h(w, r)
	}
}

func main() {
	//hello := HelloHandler{}
	//world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//http.Handle("/hello", &hello)
	//http.Handle("/world", &world)

	//http.HandleFunc("/hello", hello)
	http.HandleFunc("/hello", protect(log(hello)))
	http.HandleFunc("/world", world)

	//server.ListenAndServeTLS("cert.pem", "key.pem")
	server.ListenAndServe()
}
