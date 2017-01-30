package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/dyeduguru/expense-tracker/rest"
	"github.com/dyeduguru/expense-tracker/stores"
	"github.com/gorilla/handlers"
	"os"
)

type Config struct {
	User string
	Database string
	Port int
}

const (
	PostgresDriver = "postgres"
)

var(
	StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("API is up and running"))
	})
	NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Not Implemented"))
	})
)



func main() {
	config := &Config{
		User: "dinesh",
		Database: "expense-tracker",
	}
	db := initDB(config)

	//initialize resources
	expenseStore := stores.NewExpenseStore(db)
	expenseResource := rest.NewExpenseResource(expenseStore)

	r := mux.NewRouter()
	//static
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	//status
	r.Handle("/status", StatusHandler).Methods("GET")

	rest.AddRoutes(r, expenseResource)

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}

func initDB(config *Config) *sql.DB {
	connString := fmt.Sprintf("user=%s dbname=%s sslmode=disable", config.User, config.Database)
	db, err := sql.Open(PostgresDriver, connString)
	if err != nil {
		panic(err)
	}
	return db
}