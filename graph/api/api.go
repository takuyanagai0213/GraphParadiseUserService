package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	// "github.com/jinzhu/gorm"
)

func GetRooms(w http.ResponseWriter, r *http.Request) {
	var data = make([]string, 10)
	data[0] = "部屋1"
	data[1] = "部屋2"
	data[2] = "部屋3"
	data[3] = "部屋4"
	data[4] = "部屋5"
	data[5] = "部屋6"
	data[6] = "部屋7"
	data[7] = "部屋8"
	data[8] = "部屋9"
	data[9] = "部屋10"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func GetDates(w http.ResponseWriter, r *http.Request) {
	var data = make([]int, 30)
	for i := 0; i < 30; i++ {
		data[i] = rand.Intn(100)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func GetData1(w http.ResponseWriter, r *http.Request) {
	data := createDummyData()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func GetDataForDaily(w http.ResponseWriter, r *http.Request) {
	data := createDummyDataForMonth()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func GetData3(w http.ResponseWriter, r *http.Request) {
	data := createDummyData()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func createDummyData() []int {
	var data = make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = rand.Intn(100)
	}

	return data
}
func createDummyDataForMonth() []int {
	var data = make([]int, 30)
	for i := 0; i < 30; i++ {
		data[i] = rand.Intn(100)
	}

	return data
}
