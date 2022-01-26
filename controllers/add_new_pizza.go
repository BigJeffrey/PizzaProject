package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Controller) AddNewPizza(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dodawnie nowej pizzy")

	newPizza := ReadBody(r)

	insertedID, err := c.Dao.AddNewPizza(newPizza)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertedID)

	fmt.Println("ok")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPizza)
}
