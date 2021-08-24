package model

type User struct {
	// gorm.Model
	Name     string
	Password string
	Area     string
}
