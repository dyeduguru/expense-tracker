package main

import (
	"database/sql"
	"fmt"
	"github.com/dyeduguru/expense-tracker/rest"
	"github.com/dyeduguru/expense-tracker/stores"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

type Config struct {
	User     string
	Database string
	Port     int
}

const (
	PostgresDriver = "postgres"
)

var (
	StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is up and running"))
	})
)

func main() {
	config := &Config{
		User:     "dinesh",
		Database: "expense-tracker",
	}
	db := initDB(config)

	//initialize resources
	expenseStore := stores.NewExpenseStore(db)
	userStore := stores.NewUserStore(db)
	expenseResource := rest.NewExpenseResource(expenseStore, userStore)
	userResource := rest.NewUserResource(userStore)
	tokenResource := rest.NewTokenResource(userStore)

	r := mux.NewRouter()
	//static
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	//status
	r.Handle("/status", StatusHandler).Methods("GET")

	rest.AddRoutes(r, expenseResource, userResource, tokenResource)

	certPath := "keys/server.pem"
	keyPath := "keys/server.key"
	http.ListenAndServeTLS(":3000", certPath, keyPath, handlers.LoggingHandler(os.Stdout, r))

}

func initDB(config *Config) *sql.DB {
	connString := fmt.Sprintf("user=%s dbname=%s sslmode=disable", config.User, config.Database)
	db, err := sql.Open(PostgresDriver, connString)
	if err != nil {
		panic(err)
	}
	return db
}
