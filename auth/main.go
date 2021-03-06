package main

import (
	"graph_paradise/api"
	"graph_paradise/auth"
	"html/template"
	"log"
	"net/http"
	"os"
	// "github.com/jinzhu/gorm"
)

func main() {
	dir, _ := os.Getwd()

	http.HandleFunc("/auth", auth.Auth)
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/graph", graph)
	http.HandleFunc("/new", api.New)
	http.HandleFunc("/get", api.Get)
	http.HandleFunc("/getData1", api.GetData1)
	http.HandleFunc("/getData2", api.GetData2)
	http.HandleFunc("/getData3", api.GetData3)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	// port
	http.ListenAndServe(":80", nil)
}

func graph(w http.ResponseWriter, r *http.Request) {
	log.Print("api/")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
