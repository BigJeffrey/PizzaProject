package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pizza/models"
)

func (c *Controller) AddNewOpinion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dodawnie nowej opinii")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var nowaOpinia models.Opinion
	err = json.Unmarshal(body, &nowaOpinia)
	if err != nil {
		log.Fatal(err)
	}

	nowaResult, err := c.Dao.AddNewOpinion(nowaOpinia)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nowaResult)
	fmt.Println("ok")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nowaOpinia)
}
