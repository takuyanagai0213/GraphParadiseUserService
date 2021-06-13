package main

import (
	"graph_paradise/api"
	"html/template"
	"log"
	"net/http"
	"os"
	// "github.com/jinzhu/gorm"
)

func main() {
	dir, _ := os.Getwd()

	http.HandleFunc("/graph", graph)
	http.HandleFunc("/new", api.New)
	http.HandleFunc("/get", api.Get)
	http.HandleFunc("/getData", api.GetData)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	// port
	http.ListenAndServe(":3000", nil)
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
