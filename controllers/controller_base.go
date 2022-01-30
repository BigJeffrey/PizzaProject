package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"pizza/dao/interfaces"
	"pizza/models"
)

type Controller struct {
	Dao interfaces.AppDao
}

func ReadBody(r *http.Request, w http.ResponseWriter) (models.Pizza, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return models.Pizza{}, err
	}

	var converted models.Pizza
	err = json.Unmarshal(body, &converted)
	if err != nil {
		return models.Pizza{}, err
	}

	return converted, nil
}

func ReturnMessage(message string, err error, w http.ResponseWriter, status int) {
	var messageFmt string
	if err != nil {
		messageFmt = fmt.Sprintf("%s: %s", message, err)
	} else {
		messageFmt = message
	}

	jstr := models.ErrorJson{
		Message: messageFmt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(jstr)
}
