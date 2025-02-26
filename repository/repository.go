package repository

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetUser(id int) (*User, error)
}
