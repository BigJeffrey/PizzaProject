package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Controller) UpdatePizza(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Aktualizacja wpisu/pizzy")

	updated := ReadBody(r)

	result, err := c.Dao.UpdatePizza(updated)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println("ok")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}
