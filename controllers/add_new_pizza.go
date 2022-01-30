package controllers

import (
	"net/http"
)

func (c *Controller) AddNewPizza(w http.ResponseWriter, r *http.Request) {
	newPizza, err := ReadBody(r, w)
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}
	_, err = c.Dao.AddNewPizza(newPizza)

	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	ReturnMessage("Pizza was succesfully added", nil, w, http.StatusCreated)
}
