package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	// "github.com/jinzhu/gorm"
)

func GetData1(w http.ResponseWriter, r *http.Request) {
	data := createDummyData()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func GetData2(w http.ResponseWriter, r *http.Request) {
	data := createDummyData()

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
