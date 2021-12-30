package controllers

import (
	"fmt"
	"net/http"
)

func (c *Controller) ListPizzasWithOpinins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Wyświetlanie wszystkich pizz z opiniami")

	c.Dao.ListPizzasWithOpinins()

	fmt.Println("ok")

	w.Header().Set("Content-Type", "application/json")
}
