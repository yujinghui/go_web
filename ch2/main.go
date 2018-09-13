package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHelloName(writer http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("value", v)
	}
	fmt.Fprintf(writer, "Hello yujinghui")
}

func index(writer http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, val := template.ParseFiles("template/index.gtpl")
		fmt.Println(val)
		t.Execute(writer, nil)
	} else {
		r.ParseForm()
		fmt.Println(r.Form["username"])
		fmt.Println(r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/say", sayHelloName)
	http.ListenAndServe(":8888", nil)
}
