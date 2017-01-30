package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AddRoutes(router *mux.Router,expenseResource *ExpenseResource) {
	//expense
	expense := router.PathPrefix(ExpensesRoot).Subrouter()
	expense.Methods("GET").Handler(http.HandlerFunc(expenseResource.List))
}