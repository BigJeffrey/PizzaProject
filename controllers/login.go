package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"quickstart/models"
	"quickstart/rabbit"

	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySignedKey = []byte("mySecredPhrase")

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logowanie")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var daneLogowania models.User

	err = json.Unmarshal(body, &daneLogowania)
	if err != nil {
		log.Fatal(err)
	}

	result := c.Dao.Login(daneLogowania)

	if err != nil {
		log.Fatal(err)
	}

	if result {
		validToken, err := generateJWT()
		if err != nil {
			log.Fatal(err)
		}

		http.SetCookie(w,
			&http.Cookie{
				Name:    "token",
				Value:   validToken,
				Expires: time.Now().Add(time.Minute * 3),
			})

		fmt.Println("Zalogowany poprawnie")
		rabbit.SendRabbitMessage(daneLogowania.Email, "login")
	} else {
		fmt.Println("Nie udana próba logowania")
	}
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["autorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

	tokenString, err := token.SignedString(mySignedKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
