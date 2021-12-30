package interfaces

import (
	"quickstart/models"
)

type AppDao interface {
	AddNewPizza(w models.Pizza) (interface{}, error)
	UpdatePizza(p models.Pizza) (interface{}, error)
	DeletePizza(name string) (interface{}, error)
	ListPizzasWithOpinins()
	AddNewOpinion(o models.Opinion) (interface{}, error)
	AddNewUser(u models.User) (interface{}, error)
	Login(u models.User) bool
	Disconnect()
}
