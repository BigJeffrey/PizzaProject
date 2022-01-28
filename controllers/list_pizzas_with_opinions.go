package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Controller) ListPizzasWithOpinins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Wyświetlanie wszystkich pizz z opiniami")

	result, err := c.Dao.ListPizzasWithOpinins()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
