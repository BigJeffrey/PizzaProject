package controllers

import (
	"net/http"
)

func (c *Controller) DeletePizza(w http.ResponseWriter, r *http.Request) {
	deleted, err := ReadBody(r, w)
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}
	_, err = c.Dao.DeletePizza(string(deleted.Name))
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	ReturnMessage("Pizza was successfully deleted", nil, w, http.StatusNoContent)
}
