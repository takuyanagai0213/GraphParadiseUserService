package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/takuyanagai0213/GraphParadiseUserService/api"
)

func main() {
	dir, _ := os.Getwd()

	http.HandleFunc("/users", users)
	http.HandleFunc("/user/new", api.CreateUser)
	http.HandleFunc("/user/get", api.GetUsers)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	// port
	http.ListenAndServe(":80", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("users.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
