package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func (c *Controller) DeletePizza(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Usuwanie pizzy")

	deleted := ReadBody(r)

	result, err := c.Dao.DeletePizza(string(deleted.Name))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	fmt.Println("ok")

	w.WriteHeader(http.StatusNoContent)
}
