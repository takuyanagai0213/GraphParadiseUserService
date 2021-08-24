package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/sample/sample-api/domain/model"
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

// 検索
func (up *userPersistence) Search(name string) ([]*model.User, error) {
	var user []model.User

	// DB接続確認
	if err := up.Conn.Take(&user).Error; err != nil {
		return nil, err
	}

	db := up.Conn.Find(&user)

	// 名前検索
	if name != "" {
		db = db.Where("name = ?", name).Find(&user)
	}

	return user, nil
}
