package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "hello %s", req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
