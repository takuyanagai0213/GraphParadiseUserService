package main

import (
	"api"
	"html/template"
	"net/http"
	"os"
)

func main() {
	dir, _ := os.Getwd()

	http.HandleFunc("/users", users)
	http.HandleFunc("/user/new", api.NewUser)
	http.HandleFunc("/user/get", api.GetUsers)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	// port
	http.ListenAndServe(":80", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
