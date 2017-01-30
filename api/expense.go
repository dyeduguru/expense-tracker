package api

import "time"

type Expenses []*Expense
type Expense struct {
	Id          string
	UserId      string
	Timestamp   time.Time
	Amount      float32
	Description string
}

type ExpenseStore interface {
	Create(exp *Expense) error
	Read(userid string) (Expenses, error)
	ReadAll() (Expenses, error)
	Update(exp *Expense) error
	Delete(id string) error
}
