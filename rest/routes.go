package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

func AddRoutes(router *mux.Router,expenseResource *ExpenseResource, tokenResource *TokenResource) {
	//expense
	expense := router.PathPrefix(ExpensesRoot).Subrouter()
	expense.Methods("GET").Handler(jwtMiddleware.Handler(http.HandlerFunc(expenseResource.Read)))
	expense.Methods("POST").Handler(jwtMiddleware.Handler(http.HandlerFunc(expenseResource.Create)))
	expense.Methods("PATCH").Handler(jwtMiddleware.Handler(http.HandlerFunc(expenseResource.Update)))
	expense.Methods("PUT").Handler(jwtMiddleware.Handler(http.HandlerFunc(expenseResource.Delete)))

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