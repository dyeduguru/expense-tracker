package rest

import (
	"github.com/dyeduguru/expense-tracker/api"
	"net/http"
	"log"
)

type ExpenseResource struct {
	store api.ExpenseStore
}

func NewExpenseResource(expenseStore api.ExpenseStore) *ExpenseResource {
	return &ExpenseResource{store:expenseStore}
}

func (expenseResource *ExpenseResource) List(w http.ResponseWriter, r *http.Request) {
	expenses, err := expenseResource.store.ReadAll()
	if err != nil {
		log.Print(err)
	}
	WriteJSON(w, expenses, http.StatusOK)
}