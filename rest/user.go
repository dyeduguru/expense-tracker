package rest

import (
	"encoding/json"
	"github.com/dyeduguru/expense-tracker/api"
	"net/http"
)

type UserResource struct {
	userStore api.UserStore
}

func NewUserResource(userStore api.UserStore) *UserResource {
	return &UserResource{
		userStore: userStore,
	}
}

func (ur *UserResource) Create(w http.ResponseWriter, r *http.Request) {
	var user api.User
	inputBytes, err := ReadBody(r)
	if err != nil {
		WriteJSON(w, err, http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(inputBytes, &user); err != nil {
		WriteJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := ur.userStore.Create(&user); err != nil {
		WriteJSON(w, err, http.StatusBadRequest)
		return
	}
	WriteJSON(w, user, http.StatusOK)
}
