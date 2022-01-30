package middlewares

import (
	"log"
	"net/http"
	"pizza/controllers"

	"github.com/dgrijalva/jwt-go"
)

var mySignedKey = []byte("mySecredPhrase")

func (m *Middleware) IsAutorised(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		if r.RequestURI == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				controllers.ReturnMessage("Pleas log in first", nil, w, http.StatusUnauthorized)
				return
			}
			controllers.ReturnMessage("Pleas log in first", nil, w, http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value

		tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return mySignedKey, nil
		})
		if err != nil {
			controllers.ReturnMessage("Pleas log in first", nil, w, http.StatusUnauthorized)
			return
		}

		if tkn.Valid {
			next.ServeHTTP(w, r)
		} else {
			controllers.ReturnMessage("Pleas log in first", nil, w, http.StatusUnauthorized)
			return
		}
	})
}
