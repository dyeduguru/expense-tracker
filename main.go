package main

import (
	"fmt"
	"database/sql"
	"github.com/dyeduguru/expense-tracker/expense"
	_ "github.com/lib/pq"
	"github.com/dyeduguru/expense-tracker/api"
	"time"
)

type Config struct {
	User string
	Database string
}

const (
	PostgresDriver = "postgres"
)

func main() {
	config := &Config{
		User: "dinesh",
		Database: "expense-tracker",
	}

	connString := fmt.Sprintf("user=%s dbname=%s sslmode=disable", config.User, config.Database)
	db, err := sql.Open(PostgresDriver, connString)
	if err != nil {
		panic(err)
	}

	var expenseStore api.ExpenseStore
	expenseStore = expense.NewStore(db)
	if err = expenseStore.Update(&api.Expense{
		Id: "1",
		UserId: "1",
		Amount: 11.02,
		Description: "This is a test description",
		Timestamp: time.Now(),
	}); err != nil {
		panic(err)
	}
}

