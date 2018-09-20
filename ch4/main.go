package main

import (
	"database/sql"
	"fmt"
	session "goWeb/ch4/session"
	"html/template"
	"net/http"
	"time"

	memory "goWeb/ch4/memory"
	_ "github.com/go-sql-driver/mysql"
)

var globalSessionManager *session.SessionManager
var err error

func init() {
	memory.Init()
	globalSessionManager, err = session.NewManager("memory","gosessionid",3600)
	if err != nil {
		panic(err)
	}
}

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
		db, err := sql.Open("mysql", "root:root@/chitchat?charset=utf8")
		fmt.Println(err)
		if err != nil {
			fmt.Println("stmt : %s", err)
			panic(err.Error())
		}
		stmt, err := db.Prepare("INSERT into users(uuid,name,password,email,created_at) values(?,?,?,?,?)")
		if err != nil {
			fmt.Printf("stmt : %s", err)
			panic(err.Error())
		}
		res, err := stmt.Exec("222222",
			r.Form["username"][0],
			r.Form["password"][0],
			r.Form["vacation"][0],
			time.Now())
		if err != nil {
			fmt.Println("stmt : %s", err)
			panic(err.Error())
		}
		_, err1 := res.LastInsertId()
		if err1 != nil {

			panic(err.Error())
		}
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", index)
	http.HandleFunc("/say", sayHelloName)
	http.ListenAndServe(":8888", nil)
}
