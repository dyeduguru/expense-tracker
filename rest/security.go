package rest

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/dyeduguru/expense-tracker/api"
	"github.com/dyeduguru/expense-tracker/stores"
	"github.com/gorilla/context"
	"github.com/palantir/stacktrace"
	"log"
)

type TokenResource struct {
	userStore *stores.UserStore
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewTokenResource(userStore *stores.UserStore) *TokenResource {
	return &TokenResource{userStore: userStore}
}

func (tr *TokenResource) GetToken(w http.ResponseWriter, r *http.Request) {
	var user UserCredentials
	inputBytes, err := ReadBody(r)
	if err != nil {
		WriteJSON(w, fmt.Errorf("Invalid Request: %v, error:%v", r.Body, err), http.StatusBadRequest)
		return
	}
	//decode request into UserCredentials struct
	if err := json.Unmarshal(inputBytes, &user); err != nil {
		WriteJSON(w, fmt.Errorf("Invalid Request: %v, error:%v", string(inputBytes), err), http.StatusBadRequest)
		return
	}

	retrievedUser, err := tr.userStore.Read(user.Username)
	if err != nil {
		errorMsg := fmt.Sprintf("Error while fetching user info:%v", err)
		log.Print(errorMsg)
		WriteJSON(w, errors.New(errorMsg), http.StatusInternalServerError)
		return
	}

	if retrievedUser.Password != user.Password {
		WriteJSON(w, errors.New("Invalid Username/Password"), http.StatusUnauthorized)
		return
	}

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Set token claims */
	token.Header["name"] = user.Username
	token.Header["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
}

func GetUserFromRequest(r *http.Request, userStore api.UserStore) (*api.User, error) {
	parsedToken := context.Get(r, "user").(*jwt.Token)
	if parsedToken == nil {
		return nil, stacktrace.NewError("User not set on request")
	}
	username := parsedToken.Header["name"].(string)
	user, err := userStore.Read(username)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	return user, nil
}
