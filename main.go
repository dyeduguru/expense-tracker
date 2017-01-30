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
	"io/ioutil"
	"log"
)

type Config struct {
	User string
	Database string
	Port int
}

const (
	PostgresDriver = "postgres"
	privKeyPath = "keys/app.rsa"
	pubKeyPath = "keys/app.rsa.pub"
)

var VerifyKey, SignKey []byte


func initKeys(){
	var err error

	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

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
	http.ListenAndServeTLS(":3000", certPath, keyPath,handlers.LoggingHandler(os.Stdout, r))

}

func initDB(config *Config) *sql.DB {
	connString := fmt.Sprintf("user=%s dbname=%s sslmode=disable", config.User, config.Database)
	db, err := sql.Open(PostgresDriver, connString)
	if err != nil {
		panic(err)
	}
	return db
}