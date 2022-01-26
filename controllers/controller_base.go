package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"pizza/dao/interfaces"
	"pizza/models"
)

type Controller struct {
	Dao interfaces.AppDao
}

func ReadBody(r *http.Request) models.Pizza {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var converted models.Pizza
	err = json.Unmarshal(body, &converted)
	if err != nil {
		log.Fatal(err)
	}

	return converted
}
