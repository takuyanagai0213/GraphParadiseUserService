package api

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"graph_paradise/database"
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

type User struct {
	gorm.Model
	Name     string
	Password string
}

func New(w http.ResponseWriter, r *http.Request) {
	var name string = r.FormValue("name")
	var password string = r.FormValue("password")

	if name == "" || password == "" {
		fmt.Println("Empty user or Empty password")
		return
	}
	db := database.DBConnect()
	db.AutoMigrate(&User{})

	// パスワードのハッシュ化
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	fmt.Println("create user " + name + " with password " + string(hashed_password))
	db.Create(&User{Name: name, Password: string(hashed_password)})
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
