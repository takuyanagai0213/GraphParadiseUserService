package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/takuyanagai0213/GraphParadiseUserService/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Area     string
}

func CreateUser(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	var name string r.FormValue("name")
	var password string = r.FormValue("password")

	if name == "" || password == "" {
		var message string = "ユーザを作成できませんでした"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		fmt.Println("Empty user or Empty password")
		return
	}
	db := database.DBConnect()
	db.AutoMigrate(&User{})

	// パスワードのハッシュ化
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	fmt.Println("create user " + name + " with password " + string(hashed_password))
	db.Create(&User{Name: name, Password: string(hashed_password)})

	var message string = "ユーザを作成しました"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := database.DBConnect()

	var users []User
	db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// func UpdateUsers(w http.ResponseWriter, r *http.Request) {
// 	db := database.DBConnect()

// 	var user []User
// 	db.First(&user)
// 	fmt.Println(&user)
// 	user.Name = "jinzhu 2"
// 	db.Save(&user)
// 	var message string = "更新しました"
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)
// }
