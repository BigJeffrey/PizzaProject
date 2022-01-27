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
		log.Fatal(err)
	}
	fmt.Println(result)

	w.Header().Set("Content-Type", "application/json")
	for _, val := range result {
		json.NewEncoder(w).Encode(val)
	}
}
