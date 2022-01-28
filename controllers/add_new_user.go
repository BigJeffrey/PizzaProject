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
	var newUser models.User

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		log.Fatal(err)
	}

	insertedID, err := c.Dao.AddNewUser(newUser)

	fmt.Println(newUser.Username)
	fmt.Println(newUser.Password)
	fmt.Println(newUser.Email)

	if err != nil {
		log.Fatal(err)
	}

	rabbit.SendRabbitMessage(newUser.Email, "new_user")

	fmt.Println(insertedID)
	fmt.Println("ok")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
