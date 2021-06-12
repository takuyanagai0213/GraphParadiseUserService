package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// controller
	http.HandleFunc("/", echoHello)
	http.HandleFunc("/getData", getData)
	// port
	http.ListenAndServe(":8000", nil)
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func getData(w http.ResponseWriter, r *http.Request) {
	data := []int{10, 20, 20, 15, 40, 100}

	// fmt.Fprintf(w, "<h1>data</h1>")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
