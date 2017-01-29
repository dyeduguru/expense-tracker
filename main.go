package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/dyeduguru/expense-tracker/api"
	"github.com/dyeduguru/expense-tracker/stores"
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

	var userStore api.UserStore
	userStore = stores.NewUserStore(db)
	if err := userStore.Create(&api.User{
		Id: "1",
		Admin:true,
		UserName: "yvdinesh",
		Password: "2121",
		Name: "Dinesh Yeduguru",
	}); err != nil {
		panic(err)
	}
}

