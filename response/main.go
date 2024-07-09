package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct{
	User 		string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	// レスポンスボディにHTMLを書き込む
	html := `<html>
<head><title>Example</title></head>
<body><h1>Hello, World!</h1></body>
</html>`

	w.Write([]byte(html))
}

func WriteHeaderExample(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません。")
}

func HeaderExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func JsonExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "TestUser",
		Threads: []string{"1", "2", "3"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", WriteHeaderExample)
	http.HandleFunc("/redirect", HeaderExample)
	http.HandleFunc("/json", JsonExample)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
