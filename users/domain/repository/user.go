package repository

type UserRepository interface {
	Search(name string) ([]*model.User, error)
}
