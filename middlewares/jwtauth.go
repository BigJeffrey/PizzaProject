package middlewares

import (
	"net/http"
	"pizza/controllers"

	"github.com/dgrijalva/jwt-go"
)

var mySignedKey = []byte("mySecredPhrase")

func (m *Middleware) IsAutorised(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			controllers.ReturnMessage("Uzytkownik nie jest zalogowany lub sesja wygasła", err, w, http.StatusUnauthorized)
		}

		tokenStr := cookie.Value

		tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return mySignedKey, nil
		})
		if err != nil {
			controllers.ReturnMessage("Uzytkownik nie jest zalogowany lub sesja wygasła", err, w, http.StatusUnauthorized)
		}

		if tkn.Valid {
			next.ServeHTTP(w, r)
		} else {
			controllers.ReturnMessage("Uzytkownik nie jest zalogowany lub sesja wygasła", err, w, http.StatusUnauthorized)
		}
	})
}
