package main

import (
	"graph_paradise/api"
	"html/template"
	"net/http"
	"os"
	// "github.com/jinzhu/gorm"
)

func main() {
	dir, _ := os.Getwd()

	http.HandleFunc("/graph", graph)
	http.HandleFunc("/getRooms", api.GetRooms)
	http.HandleFunc("/getAreas", api.GetAreas)
	http.HandleFunc("/getDates", api.GetDates)
	http.HandleFunc("/getData1", api.GetData1)
	http.HandleFunc("/GetDataForTable", api.GetDataForTable)
	http.HandleFunc("/getDataForDaily", api.GetDataForDaily)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	// port
	http.ListenAndServe(":80", nil)
}

func graph(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
