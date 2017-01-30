package rest

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/auth0/go-jwt-middleware"
)

var mySigningKey = []byte("secret")


var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Set token claims */
	token.Header["admin"] = true
	token.Header["name"] = "Ado Kukic"
	token.Header["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
})
