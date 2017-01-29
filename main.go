package main

import (
	"fmt"
	"database/sql"
	"log"
	"github.com/dyeduguru/expense-tracker/expense"
	_ "github.com/lib/pq"
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

	expenseStore := expense.NewStore(db)
	allexpenses, err := expenseStore.GetAll()
	if err != nil {
		panic(err)
	}
	for _, exp := range allexpenses {
		log.Printf("UserId: %v, Amount:%v, Description:%v", exp.UserId, exp.Amount, exp.Description)
	}
}

