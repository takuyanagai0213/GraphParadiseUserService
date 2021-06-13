package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

func main() {
	// controller
	DBConnect()
	dir, _ := os.Getwd() //追記
	log.Print(http.Dir(dir + "/static/"))
	http.HandleFunc("/", echoHello)
	http.HandleFunc("/new", new)
	http.HandleFunc("/get", get)
	http.HandleFunc("/getData", getData)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/")))) //追記
	// port
	http.ListenAndServe(":3000", nil)
}

type User struct {
	gorm.Model
	Name  string
	Email string
}

func new(w http.ResponseWriter, r *http.Request) {
	db := DBConnect()
	db.AutoMigrate(&User{})

	name := "takuya"
	email := "aaabbb@exmaplle.com"
	fmt.Println("create user " + name + " with email " + email)
	db.Create(&User{Name: name, Email: email})
	// defer db.Close()
}
func get(w http.ResponseWriter, r *http.Request) {
	db := DBConnect()

	// fmt.Println("create user " + name + " with email " + email)
	var users []User
	db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
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

func DBConnect() (database *gorm.DB) {
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "sample"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
