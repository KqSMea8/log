package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET" {
		t,_ := template.ParseFiles("login.html")
		beego.Info(t.Execute(w, nil))
	}else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Printf("url:%+v\n", r.URL.RawQuery)
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err!= nil {
		beego.Info("err:", err)
	}
	http.SetCookie()
}