package api

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"graph_paradise/database"
	"net/http"
	// "github.com/jinzhu/gorm"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	data := []int{10, 20, 20, 15, 40, 100}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

type User struct {
	gorm.Model
	Name  string
	Email string
}

func New(w http.ResponseWriter, r *http.Request) {
	db := database.DBConnect()
	db.AutoMigrate(&User{})

	name := "takuya"
	email := "aaabbb@exmaplle.com"
	fmt.Println("create user " + name + " with email " + email)
	db.Create(&User{Name: name, Email: email})
	// defer db.Close()
}
func Get(w http.ResponseWriter, r *http.Request) {
	db := database.DBConnect()

	// fmt.Println("create user " + name + " with email " + email)
	var users []User
	db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
