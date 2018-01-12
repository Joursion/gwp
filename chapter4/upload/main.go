package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
)

type Post struct{
	User string
	Threads []string
}

/**
Having read the earlier part of this chapter, you might wonder why the ServeHTTP func- tion takes 
two parameters—the ResponseWriter interface and a pointer to a Request struct. The reason why 
it’s a pointer to Request is simple: changes to Request by the handler need to be visible to 
the server, so we’re only passing it by reference in- stead of by value. But why are we passing 
in a ResponseWriter by value? The server needs to know the changes to ResponseWriter too, doesn’t it?
If you dig into the net/http library code, you’ll find that ResponseWriter is an interface to a 
nonexported struct response, and we’re passing the struct by reference (we’re passing in a pointer 
to response) and not by value.In other words, both the parameters are passed in by reference; it’s 
just that the meth- od signature takes a ResponseWriter that’s an interface to a pointer to a struct, 
so it looks as if it’s passed in by value.
**/

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	//  parses a request body as multipart/form-data
	fileHeader := r.MultipartForm.File["uploaded"][0] // why
	file, err := fileHeader.Open()
	if err == nil {
		data , err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data)) 
		}
	}
}

func main () {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func processs(w http.ResponseWriter, r *http.Request) {
	file, _ , err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "A",
		Threads: []string{"1", "2"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}