package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pizza/models"
	"pizza/rabbit"
)

func (c *Controller) AddNewUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	var newUser models.User

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	_, err = c.Dao.AddNewUser(newUser)

	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	err = rabbit.SendRabbitMessage(newUser.Email, "new_user")
	if err != nil {
		ReturnMessage("Unable to send email but user was succesfully added", err, w, http.StatusInternalServerError)
		return
	}

	ReturnMessage("New user was succesfully added and welcome email was sent", nil, w, http.StatusCreated)
}
