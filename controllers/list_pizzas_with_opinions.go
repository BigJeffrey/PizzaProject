package controllers

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) ListPizzasWithOpinins(w http.ResponseWriter, r *http.Request) {
	result, err := c.Dao.ListPizzasWithOpinins()
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
