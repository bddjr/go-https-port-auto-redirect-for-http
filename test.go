package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start test server")
	var addr = "127.0.0.1:5678"
	fmt.Println("http://" + addr)
	server := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(httpResponseHandle),
	}
	server.ListenAndServeTLS("localhost.crt", "localhost.key")
}

func httpResponseHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("hello"))
		return
	}
	w.WriteHeader(404)
	w.Write([]byte("404 Not Found"))
}
