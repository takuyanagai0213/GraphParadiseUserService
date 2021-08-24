package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/sample/sample-api/domain/repository"
)

// UserにおけるPersistenceのインターフェース
type userPersistence struct {
	Conn *gorm.DB
}

// Userデータに関するPersistenceを生成
func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}
