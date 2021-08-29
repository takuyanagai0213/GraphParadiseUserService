package repository

type User struct {
	// gorm.Model
	Name     string
	Password string
	Area     string
}

type UserRepository interface {
	Search(name string) ([]*User, error)
	GetUserByName(name string) ([]*User, error)
}
