package main

import (
	"time"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
		fmt.Println(r.Form["username"][0])
		fmt.Println(len(r.Form["username"][0]))
		fmt.Println(r.Form["password"])
		fmt.Println(r.Form["vacation"][0])
		fmt.Println(r.Form["interest"][1])
		fmt.Println(time.Now())
		intage, err := strconv.Atoi(r.Form.Get("age"))
		fmt.Println(intage, err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/say", sayHelloName)
	http.ListenAndServe(":8888", nil)
}
