package main

import (
	"pizza/controllers"
	"pizza/middlewares"

	"context"
	"log"
	"net/http"
	"pizza/dao/factory"

	"github.com/gorilla/mux"
)

func handleRequest(c *controllers.Controller, m *middlewares.Middleware) {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/pizza", c.AddNewPizza).Methods("POST")          //curl -d "{\"Name\":\"Stodola\", \"Size\":\"45\"}" -X POST localhost:8081/pizza
	myRouter.HandleFunc("/opinion", c.AddNewOpinion).Methods("POST")      //curl -d "{\"mainbase\":\"6100533dcb7ebec302f8f773\", \"score\":\"23\", \"opinions\":\"ale pizzocha\"}" -X POST localhost:8081/opinion
	myRouter.HandleFunc("/pizza", c.ListPizzasWithOpinins).Methods("GET") //curl localhost:8081/pizza
	myRouter.HandleFunc("/pizza", c.DeletePizza).Methods("DELETE")        //curl -d "{\"Name\":\"K2\"}" -X DELETE localhost:8081/pizza
	myRouter.HandleFunc("/pizza", c.UpdatePizza).Methods("PUT")           //curl -d "{\"Name\":\"K2\", \"Size\":\"60\"}" -X PUT localhost:8081/pizza
	myRouter.HandleFunc("/user", c.AddNewUser).Methods("POST")            //curl -d "{\"Username\":\"Mateusz\", \"Password\":\"12345678\", \"Email\":\"syru1988@gmail.com\"}" -X POST localhost:8081/user
	myRouter.HandleFunc("/login", c.Login)                                //with cookie JWT token

	myRouter.Use(m.IsAutorised)

	log.Fatal(http.ListenAndServe(":8085", myRouter))
}

func main() {
	ctx := context.Background()

	factoryDao := factory.FactoryDao{Ctx: ctx}

	dao := factoryDao.FactoryDao("mongodb") //postgresql or mongodb

	defer dao.Disconnect()

	controllers := &controllers.Controller{Dao: dao}
	middlewares := &middlewares.Middleware{Dao: dao}

	handleRequest(controllers, middlewares)
}
