package api

type Users []*User
type User struct {
	Id string
	Admin bool
	UserName string
	Password string
	Name string
}

type UserStore interface {
	Get(userid string) (User, error)
}