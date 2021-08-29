package usecase

import (
	"context"
	"fmt"

	"github.com/takuyanagai0213/GraphParadiseUserService/database"
	"github.com/takuyanagai0213/GraphParadiseUserService/domain/model"
	"github.com/takuyanagai0213/GraphParadiseUserService/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

// UserにおけるUseCaseのインターフェース
type UserUseCase interface {
	List() ([]model.User, error)
	CreateUser(string, string) string
	GetUserByName(string) (model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// Userデータに関するUseCaseを生成
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

// GetUserByName Nameを元にユーザを1件取得する
func (uu userUseCase) CreateUser(name string, password string) string {
	var message string
	if name == "" || password == "" {
		message = "ユーザを作成できませんでした"
		fmt.Println("Empty user or Empty password")
		return message
	}
	db := database.DBConnect()
	// db.AutoMigrate(&user)

	// パスワードのハッシュ化
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	user := model.User{Name: name, Password: string(hashed_password)}
	fmt.Println("create user " + name + " with password " + string(hashed_password))
	db.Create(&user)

	message = "ユーザを作成しました"

	return message
}

// GetUserByName Nameを元にユーザを1件取得する
func (uu userUseCase) GetUserByName(name string) (model.User, error) {
	var user model.User
	db := database.DBConnect()
	db.Where("name = ?", name).First(&user)

	return user, nil
}

// List ユーザを全件取得
func (uu userUseCase) List() ([]model.User, error) {
	var userList []model.User
	rows, err := listAll(context.Background())
	if err != nil {
		fmt.Println("Error happened")
		return []model.User{}, err
	}
	for _, row := range rows {
		userList = append(userList, row)
	}

	return userList, nil
}

// listAll 全件取得
func listAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	db := database.DBConnect()

	rows, err := db.Find(&users).Rows()
	if err != nil {
		return nil, err
	}

	fmt.Println(rows)
	for rows.Next() {
		// db.ScanRows(rows, &user)
		// users = append(users, user)
	}
	return users, nil
}
