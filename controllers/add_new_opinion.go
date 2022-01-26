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

	var newOpinion models.Opinion
	err = json.Unmarshal(body, &newOpinion)
	if err != nil {
		log.Fatal(err)
	}

	newResult, err := c.Dao.AddNewOpinion(newOpinion)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newResult)
	fmt.Println("ok")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOpinion)
}
