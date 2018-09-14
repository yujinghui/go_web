package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"
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
		db, err := sql.Open("mysql", "localhost:3306/chitchat?charset=utf8")
		stmt, err := db.Prepare("INSERT users SET uuid=?,name=?,password=?,email=?,created_at=?")
		res, err := stmt.Exec("111",
			r.Form["username"][0],
			r.Form["password"][0],
			r.Form["vacation"][0],
			time.Now())
		id, err := res.LastInsertId()
		fmt.Println(id, err)
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", index)
	http.HandleFunc("/say", sayHelloName)
	http.ListenAndServe(":8888", nil)
}
