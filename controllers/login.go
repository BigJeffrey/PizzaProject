package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pizza/models"
	"pizza/rabbit"

	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySignedKey = []byte("mySecredPhrase")

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	var loginData models.User

	err = json.Unmarshal(body, &loginData)
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	result, err := c.Dao.Login(loginData)

	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	if result {
		validToken, err := generateJWT()
		if err != nil {
			ReturnMessage("Something went wrong", err, w, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w,
			&http.Cookie{
				Name:    "token",
				Value:   validToken,
				Expires: time.Now().Add(time.Minute * 10),
			})
		err = nil
		err = rabbit.SendRabbitMessage(loginData.Email, "login")
		ReturnMessage("User signed in", err, w, http.StatusOK)
	} else {
		ReturnMessage("Incorrect login", err, w, http.StatusUnauthorized)
	}
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["autorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(mySignedKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
