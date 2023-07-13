package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Tooluloope/gourl/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

func JWTAuth(
	original func(http.ResponseWriter, *http.Request),
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]

		if authHeader == nil {
			utils.WriteJSONError(w, http.StatusUnauthorized, errors.New("User is not authorized"))
			return
		}

		bearerToken := strings.Split(authHeader[0], " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			utils.WriteJSONError(w, http.StatusUnauthorized, errors.New("Malformed token"))
			return
		}

		token := bearerToken[1]

		if validateToken(token) {
			original(w, r)
		} else {
			utils.WriteJSONError(w, http.StatusUnauthorized, errors.New("User is not authorized"))
			return
		}
	}
}

func validateToken(token string) bool {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET"))
	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok != true {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return accessToken.Valid
}
