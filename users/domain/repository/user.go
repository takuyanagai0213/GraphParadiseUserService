package repository

import (
	"github.com/takuyanagai0213/GraphParadiseUserService/domain/model"
)

type UserRepository interface {
	Search(name string) ([]*model.User, error)
}
