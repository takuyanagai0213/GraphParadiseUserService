package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	// "github.com/jinzhu/gorm"
)

type Room struct {
	room_no   string
	room_name string
}

func GetRooms(w http.ResponseWriter, r *http.Request) {
	var data []interface{}

	var buf map[string]interface{}
	for i := 1; i < 11; i++ {
		buf = map[string]interface{}{}
		buf["room_name"] = "部屋" + strconv.Itoa(i)
		buf["room_no"] = strconv.Itoa(i)
		data = append(data, buf)
	}

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
func GetDataForTable(w http.ResponseWriter, r *http.Request) {
	data := createDummyData()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
