package usecase

import (
	"context"
	"fmt"

	"github.com/takuyanagai0213/GraphParadiseUserService/database"
	"github.com/takuyanagai0213/GraphParadiseUserService/domain/model"
)

// UserにおけるUseCaseのインターフェース
type UserUseCase struct {
}

// GetUserByName Nameを元にユーザを1件取得する
func (uu *UserUseCase) GetUserByName(name string) (model.User, error) {
	var user model.User
	db := database.DBConnect()
	db.Where("name = ?", name).First(&user)

	return user, nil
}

// List ユーザを全件取得
func (uu *UserUseCase) List() ([]model.User, error) {
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
