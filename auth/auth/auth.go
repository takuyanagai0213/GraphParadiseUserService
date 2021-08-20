package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"graph_paradise/database"
	"graph_paradise/sessions"
	"html/template"
	"net/http"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

func Auth(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("auth.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	fmt.Println("username:", r.FormValue("username"))
	fmt.Println("password:", r.FormValue("password"))

	db := database.DBConnect()
	name := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("get user " + name + " with password " + password)

	var user User
	db.Where("name = ?", name).Find(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		// セッションを開始
		manager := sessions.NewManager()
		fmt.Println("---------------")
		fmt.Println(manager)
		fmt.Println("---------------")
		sessions.StartDefaultSession(manager)
		fmt.Println(sessions.DefaultCookieName)
		// セッション変数をセット
		session, _ := manager.Get(r, sessions.DefaultCookieName)
		fmt.Println(session)
		session.Set("account", name)
		// セッションを保存
		session.Save()
		fmt.Println(session.Get("username"))
		http.Redirect(w, r, "/graph", 301)
	}
	fmt.Println(err)
	http.Redirect(w, r, "/auth", 301)
}
