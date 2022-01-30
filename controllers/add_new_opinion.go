package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pizza/models"
)

func (c *Controller) AddNewOpinion(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	var newOpinion models.Opinion
	err = json.Unmarshal(body, &newOpinion)
	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	_, err = c.Dao.AddNewOpinion(newOpinion)

	if err != nil {
		ReturnMessage("Something went wrong", err, w, http.StatusBadRequest)
		return
	}

	ReturnMessage("This opinion was succesfully added", nil, w, http.StatusCreated)
}
