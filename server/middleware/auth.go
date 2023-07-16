package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Tooluloope/gourl/server/utils"
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

		if isValid, userID := validateToken(token); isValid {
			ctx := context.WithValue(r.Context(), utils.ContextKeyUser, userID)

			original(w, r.WithContext(ctx))
		} else {
			utils.WriteJSONError(w, http.StatusUnauthorized, errors.New("User is not authorized"))
			return
		}
	}
}

func validateToken(token string) (bool, string) {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET"))
	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok != true {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false, ""
	}

	if _, ok := accessToken.Claims.(jwt.MapClaims); ok && accessToken.Valid {
		claims := accessToken.Claims.(jwt.MapClaims)
		return true, claims["jti"].(string)
	}

	return accessToken.Valid, ""
}
