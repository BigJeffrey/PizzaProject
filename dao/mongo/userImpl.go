package mongodao

import (
	"fmt"
	"log"
	"pizza/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *Mongo) AddNewPizza(p models.Pizza) (interface{}, error) {
	fmt.Println("Dodawanie pizzy")

	pizzaResult, err := m.GetListaPizzy().InsertOne(m.ctx, bson.D{
		{Key: "Name", Value: p.Name},
		{Key: "Size", Value: p.Size},
	})

	return pizzaResult, err
}

func (m *Mongo) UpdatePizza(p models.Pizza) (interface{}, error) {
	fmt.Println("Aktualizacja pizzy")

	updatedBson := bson.D{
		{"$set", bson.D{
			{"Size", p.Size},
		}},
	}

	result, err := m.GetListaPizzy().UpdateOne(m.ctx, bson.M{"Name": p.Name}, updatedBson)

	return result, err
}

func (m *Mongo) DeletePizza(name string) (interface{}, error) {
	fmt.Println("Usuwanie pizzy")

	result, err := m.GetListaPizzy().DeleteOne(m.ctx, bson.M{"Name": name})

	return result, err
}

func (m *Mongo) ListPizzasWithOpinins() {
	fmt.Println("Lista pizzy z opiniami")

	cursor, err := m.GetListaPizzy().Find(m.ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var main models.PizzaMongo
	var opin models.OpinionMongo

	for cursor.Next(m.ctx) {
		cursor2, err := m.GetOpinie().Find(m.ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}

		err2 := cursor.Decode(&main)
		if err2 != nil {
			log.Fatal(err2)
		}

		for cursor2.Next(m.ctx) {

			err := cursor2.Decode(&opin)
			if err != nil {
				log.Fatal(err)
			}

			if opin.PizzaId == main.ID {
				s := fmt.Sprintf("%s, %s, %s, %s", main.Name, main.Size, opin.Opinions, opin.Score)
				fmt.Println(s)
			}
		}
	}

	fmt.Println("ok")
}

func (m *Mongo) AddNewOpinion(o models.Opinion) (interface{}, error) {
	fmt.Println("Dodawanie opinii")

	nowaResult, err := m.GetOpinie().InsertOne(m.ctx, bson.D{
		{Key: "MainBase", Value: o.PizzaId},
		{Key: "Score", Value: o.Opinions},
		{Key: "Opinions", Value: o.Score},
	})

	return nowaResult, err
}

func (m *Mongo) Login(u models.User) bool {
	result, err := m.GetUsers().Find(m.ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for result.Next(m.ctx) {
		var daneBaza models.User
		err := result.Decode(&daneBaza)
		if err != nil {
			log.Fatal(err)
		}

		if daneBaza.Username == u.Username && daneBaza.Password == u.Password {
			return true
		}
	}

	return false
}

func (m *Mongo) AddNewUser(u models.User) (interface{}, error) {
	nowaResult, err := m.GetUsers().InsertOne(m.ctx, bson.D{
		{Key: "Username", Value: u.Username},
		{Key: "Password", Value: u.Password},
		{Key: "Email", Value: u.Email},
	})

	return nowaResult, err
}
