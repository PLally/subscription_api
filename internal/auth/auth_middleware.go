package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func AuthMiddleware(secretString string) mux.MiddlewareFunc {
	secret := []byte(secretString)
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			unparsedToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)

			token, err := jwt.Parse(unparsedToken, func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized bad token"))
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid && claims["super"].(bool) {
				handler.ServeHTTP(w, r)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		})
	}
}
