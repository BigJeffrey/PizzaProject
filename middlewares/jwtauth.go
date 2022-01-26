package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySignedKey = []byte("mySecredPhrase")

func (m *Middleware) IsAutorised(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Sprawdzam czy user jest zalogowany - IsAutorised middleware")

		log.Println(r.RequestURI)
		if r.RequestURI == "/login" || r.RequestURI == "/dodajUsera" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("Najpierw się zaloguj")
			}
			log.Println(err)
		}

		tokenStr := cookie.Value

		tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return mySignedKey, nil
		})
		if err != nil {
			log.Fatal(err)
		}

		if tkn.Valid {
			next.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "Brak autoryzacji")
		}
	})
}
