package mongodao

import (
	"fmt"
	"pizza/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *Mongo) AddNewPizza(p models.Pizza) (interface{}, error) {
	pizzaResult, err := m.GetPizzaList().InsertOne(m.ctx, bson.D{
		{Key: "Name", Value: p.Name},
		{Key: "Size", Value: p.Size},
	})

	return pizzaResult, err
}

func (m *Mongo) UpdatePizza(p models.Pizza) (interface{}, error) {
	updatedBson := bson.D{
		{"$set", bson.D{
			{"Size", p.Size},
		}},
	}

	result, err := m.GetPizzaList().UpdateOne(m.ctx, bson.M{"Name": p.Name}, updatedBson)
	return result, err
}

func (m *Mongo) DeletePizza(name string) (interface{}, error) {
	result, err := m.GetPizzaList().DeleteOne(m.ctx, bson.M{"Name": name})
	return result, err
}

func (m *Mongo) ListPizzasWithOpinins() (models.ListPizzaOpinions, error) {
	cursor, err := m.GetPizzaList().Find(m.ctx, bson.M{})
	if err != nil {
		return models.ListPizzaOpinions{}, err
	}

	var main models.Pizza
	var opin models.Opinion

	var tab []models.Together
	var tabOpin []models.Opinion
	var listPO models.ListPizzaOpinions

	for cursor.Next(m.ctx) {
		cursor2, err := m.GetOpinions().Find(m.ctx, bson.M{})
		if err != nil {
			return models.ListPizzaOpinions{}, err
		}

		err = cursor.Decode(&main)
		if err != nil {
			return models.ListPizzaOpinions{}, err
		}

		for cursor2.Next(m.ctx) {
			err := cursor2.Decode(&opin)
			if err != nil {
				return models.ListPizzaOpinions{}, err
			}
			pid := fmt.Sprint("ObjectID(\"", opin.PizzaId, "\")")
			mid := fmt.Sprint(main.ID)

			if pid == mid {
				s := fmt.Sprintf("%s, %s, %s, %s", main.Name, main.Size, opin.Opinions, opin.Score)
				fmt.Println(s)
				tabOpin = append(tabOpin, models.Opinion{
					ID:       opin.ID,
					Score:    opin.Score,
					Opinions: opin.Opinions,
					PizzaId:  opin.PizzaId,
				})
			}
		}
		tab = append(tab, models.Together{
			MName: main.Name,
			MSize: main.Size,
			Ops:   tabOpin,
		})
		listPO = models.ListPizzaOpinions{
			ListPizzaWithOpinions: tab,
		}
		tabOpin = nil
	}
	return listPO, nil
}

func (m *Mongo) AddNewOpinion(o models.Opinion) (interface{}, error) {
	nowaResult, err := m.GetOpinions().InsertOne(m.ctx, bson.D{
		{Key: "Score", Value: o.Score},
		{Key: "Opinion", Value: o.Opinions},
		{Key: "PizzaId", Value: o.PizzaId},
	})

	return nowaResult, err
}

func (m *Mongo) Login(u models.User) (bool, error) {
	result, err := m.GetUsers().Find(m.ctx, bson.M{})
	if err != nil {
		return false, err
	}

	for result.Next(m.ctx) {
		var daneBaza models.User
		err := result.Decode(&daneBaza)
		if err != nil {
			return false, err
		}

		if daneBaza.Username == u.Username && daneBaza.Password == u.Password {
			return true, nil
		}
	}

	return false, nil
}

func (m *Mongo) AddNewUser(u models.User) (interface{}, error) {
	nowaResult, err := m.GetUsers().InsertOne(m.ctx, bson.D{
		{Key: "Username", Value: u.Username},
		{Key: "Password", Value: u.Password},
		{Key: "Email", Value: u.Email},
	})

	return nowaResult, err
}
