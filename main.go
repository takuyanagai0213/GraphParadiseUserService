package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// controller
	dir, _ := os.Getwd() //追記
	log.Print(http.Dir(dir + "/static/"))
	http.HandleFunc("/", echoHello)
	http.HandleFunc("/getData", getData)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/")))) //追記
	// port
	http.ListenAndServe(":3000", nil)
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	log.Print("api/")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	data := []int{10, 20, 20, 15, 40, 100}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
