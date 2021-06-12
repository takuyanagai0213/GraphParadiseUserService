package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// controller
	http.HandleFunc("/", echoHello)
	http.HandleFunc("/getData", getData)
	// port
	http.ListenAndServe(":3000", nil)
}

func echoHello(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "<h1>data</h1>")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
