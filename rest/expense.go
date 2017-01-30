package rest

import (
	"github.com/dyeduguru/expense-tracker/api"
	"net/http"
)

type ExpenseResource struct {
	expenseStore api.ExpenseStore
	userStore api.UserStore
}

func NewExpenseResource(expenseStore api.ExpenseStore, userStore api.UserStore) *ExpenseResource {
	return &ExpenseResource{
		expenseStore:expenseStore,
		userStore:userStore,
	}
}

func (expenseResource *ExpenseResource) List(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromRequest(r, expenseResource.userStore)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}
	var expenses api.Expenses
	if user.Admin {
		expenses, err = expenseResource.expenseStore.ReadAll()
	} else {
		expenses, err = expenseResource.expenseStore.Read(user.Id)
	}
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, expenses, http.StatusOK)
}