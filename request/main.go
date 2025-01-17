package main

import (
	"fmt"
	"io"
	"net/http"
)

// func handler(w http.ResponseWriter, r *http.Request){
// 	r.ParseMultipartForm(1024)
// 	fileHeader := r.MultipartForm.File["uploaded"][0]
// 	file, err := fileHeader.Open()
// 	if err == nil {
// 		data, err := io.ReadAll(file)
// 		if err == nil {
// 			fmt.Fprint(w, string(data))
// 		}
// 	}
// }

func handler(w http.ResponseWriter, r *http.Request){
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := io.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", handler)
	server.ListenAndServe()
}
//
