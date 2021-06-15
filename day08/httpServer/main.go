package main

import (
	"log"
	"net/http"
)

//type myHandler func(w http.ResponseWriter, r *http.Request)

func myHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	page := data.Get("page")
	log.Println("url请求参数:" + data.Encode())
	str := page
	r.ParseForm()
	log.Println("post参数" + r.PostForm.Encode())
	name := r.PostForm.Get("name")
	str += "post参数" + name

	w.Write([]byte(str))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", myHandle)
	s := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}
	s.ListenAndServe()
}
