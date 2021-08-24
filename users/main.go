package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/takuyanagai0213/GraphParadiseUserService/api"
	"github.com/takuyanagai0213/GraphParadiseUserService/database"
	"github.com/takuyanagai0213/GraphParadiseUserService/infrastructure/persistence"
	handler "github.com/takuyanagai0213/GraphParadiseUserService/interfaces/handler"
	"github.com/takuyanagai0213/GraphParadiseUserService/usecase"
)

func main() {
	dir, _ := os.Getwd()
	// 依存関係を定義
	userPersistence := persistence.NewUserPersistence(database.Connect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)
	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/users", userHandler.Index)

	http.HandleFunc("/users", users)
	http.HandleFunc("/users/all", user_list)
	http.HandleFunc("/user/new", api.CreateUser)
	http.HandleFunc("/user/get", api.GetUsers)
	http.HandleFunc("/user/update", api.UpdateUsers)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	// port
	http.ListenAndServe(":80", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("users.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
func user_list(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("user_list.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
