package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pizza/models"
	"pizza/rabbit"
)

func (c *Controller) AddNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dodawnie nowego usera")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	var nowyUser models.User

	err = json.Unmarshal(body, &nowyUser)
	if err != nil {
		log.Fatal(err)
	}

	insertedID, err := c.Dao.AddNewUser(nowyUser)

	fmt.Println(nowyUser.Username)
	fmt.Println(nowyUser.Password)
	fmt.Println(nowyUser.Email)

	if err != nil {
		log.Fatal(err)
	}

	rabbit.SendRabbitMessage(nowyUser.Email, "new_user")

	fmt.Println(insertedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nowyUser)
}
