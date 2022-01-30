package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (c *Controller) ListPizzasWithOpinins(w http.ResponseWriter, r *http.Request) {
	result, err := c.Dao.ListPizzasWithOpinins()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
