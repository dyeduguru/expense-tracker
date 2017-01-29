package expense

import "time"

type Expenses []*Expense
type Expense struct {
	UserId string
	Timestamp time.Time
	Amount float32
	Description string
}

