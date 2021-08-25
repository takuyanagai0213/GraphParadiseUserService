package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/takuyanagai0213/GraphParadiseUserService/database"
	"github.com/takuyanagai0213/GraphParadiseUserService/infrastructure/persistence"
	handler "github.com/takuyanagai0213/GraphParadiseUserService/interfaces/handler"
	"github.com/takuyanagai0213/GraphParadiseUserService/usecase"
)

func main() {
	// dir, _ := os.Getwd()
	// 依存関係を定義
	userPersistence := persistence.NewUserPersistence(database.DBConnect())
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)
	// ルーティングの設定
	router := httprouter.New()
	router.GET("/users", users)
	router.GET("/users/all", user_list)
	router.GET("/api/users", userHandler.Index)
	// Serve static files from the ./public directory
	// router.ServeFiles("/static/", http.Dir(dir+"/static/"))

	// router.NotFound = http.FileServer(http.Dir(dir + "/static/"))

	// サーバ起動
	// fmt.Println("========================")
	// fmt.Println("Server Start >> http://localhost:3000")
	// fmt.Println("========================")
	// log.Fatal(http.ListenAndServe(":80", router))

	// http.HandleFunc("/users", users)
	// http.HandleFunc("/users/all", user_list)
	// http.HandleFunc("/user/new", api.CreateUser)
	// http.Handle("/user/get", new(userHandler.Index))
	// http.HandleFunc("/user/update", api.UpdateUsers)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))
	// port
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:80")
	fmt.Println("========================")
	http.ListenAndServe(":80", router)
}

func users(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("users.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
func user_list(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("user_list.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
