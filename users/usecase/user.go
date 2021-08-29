package usecase

import (
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
