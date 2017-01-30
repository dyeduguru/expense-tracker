package api

type Users []*User
type User struct {
	Id       string
	Admin    bool
	UserName string
	Password string
	Name     string
}

type UserStore interface {
	Create(user *User) error
	Read(username string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}
