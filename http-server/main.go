package main

import (
	"fmt"
	"net/http"
	// "golang.org/x/net/http2"
)

// type HelloHandler struct{}
// type WorldHandler struct{}

// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w, "Hello!")
// }

// func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w, "World!")
// }

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "World!")
}

func main(){
	// hello := HelloHandler{}
	// world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)


	// http.Handle("/hello", &hello)
	// http.Handle("/world", &world)

	// http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServe()
}
