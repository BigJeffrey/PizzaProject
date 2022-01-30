package controllers

import (
	"net/http"
)

func (c *Controller) UpdatePizza(w http.ResponseWriter, r *http.Request) {
	updated, err := ReadBody(r, w)
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	_, err = c.Dao.UpdatePizza(updated)
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	ReturnMessage("The pizza was succesfully updated", nil, w, http.StatusNoContent)
}
