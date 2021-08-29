package usecase

import (
	"github.com/takuyanagai0213/GraphParadiseUserService/domain/repository"
)

// UserにおけるUseCaseのインターフェース
type UserUseCase interface {
	GetUserByName(name string) ([]*repository.User, error)
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
func (uu userUseCase) GetUserByName(name string) (user []*repository.User, err error) {
	user, err = uu.userRepository.GetUserByName(name)

	if err != nil {
		return nil, err
	}
	return user, nil
}
