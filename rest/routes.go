package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

func AddRoutes(router *mux.Router,expenseResource *ExpenseResource, tokenResource *TokenResource) {
	//expense
	expense := router.PathPrefix(ExpensesRoot).Subrouter()
	expense.Methods("GET").Handler(jwtMiddleware.Handler(http.HandlerFunc(expenseResource.List)))

	//security
	token := router.PathPrefix(TokenRoot).Subrouter()
	token.Methods("POST").Handler(http.HandlerFunc(tokenResource.GetToken))
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})