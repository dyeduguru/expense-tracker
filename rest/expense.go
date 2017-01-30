package rest

import (
	"encoding/json"
	"github.com/dyeduguru/expense-tracker/api"
	"net/http"
)

type ExpenseResource struct {
	expenseStore api.ExpenseStore
	userStore    api.UserStore
}

func NewExpenseResource(expenseStore api.ExpenseStore, userStore api.UserStore) *ExpenseResource {
	return &ExpenseResource{
		expenseStore: expenseStore,
		userStore:    userStore,
	}
}

func (expenseResource *ExpenseResource) Create(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromRequest(r, expenseResource.userStore)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}
	var expense api.Expense
	inputBytes, err := ReadBody(r)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(inputBytes, &expense); err != nil {
		WriteJSON(w, err, http.StatusBadRequest)
		return
	}
	expense.UserId = user.Id
	if err := expenseResource.expenseStore.Create(&expense); err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, expense, http.StatusOK)
}

func (expenseResource *ExpenseResource) Read(w http.ResponseWriter, r *http.Request) {
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

func (expenseResource *ExpenseResource) Update(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromRequest(r, expenseResource.userStore)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}
	var expense api.Expense
	inputBytes, err := ReadBody(r)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(inputBytes, &expense); err != nil {
		WriteJSON(w, err, http.StatusBadRequest)
		return
	}
	expense.UserId = user.Id
	if err := expenseResource.expenseStore.Update(&expense); err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, expense, http.StatusOK)
}

func (expenseResource *ExpenseResource) Delete(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromRequest(r, expenseResource.userStore)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}
	var toDelete api.Expense
	inputBytes, err := ReadBody(r)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(inputBytes, &toDelete); err != nil {
		WriteJSON(w, err, http.StatusBadRequest)
		return
	}
	expenses, err := expenseResource.expenseStore.Read(user.Id)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}

	for _, e := range expenses {
		if e.Id == toDelete.Id {
			if err := expenseResource.expenseStore.Delete(toDelete.Id); err != nil {
				WriteJSON(w, err, http.StatusInternalServerError)
				return
			}
			break
		}
	}

	WriteJSON(w, toDelete, http.StatusOK)
}
