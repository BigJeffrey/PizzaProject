package postgresqldao

import (
	"fmt"
	"log"
	"pizza/models"
)

func (m *PostgreSql) AddNewPizza(p models.Pizza) (interface{}, error) {
	sqlStatement := `insert into pizza (name, size) values ($1, $2) returning id`
	var id int64
	err := m.client.QueryRow(sqlStatement, p.Name, p.Size).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	return id, err
}

func (m *PostgreSql) AddNewOpinion(o models.Opinion) (interface{}, error) {
	sqlStatement := `insert into opinions (score, opinion, pizzaid) values ($1, $2, $3) returning opinionid`
	var id int64
	err := m.client.QueryRow(sqlStatement, o.Score, o.Opinions, o.PizzaId).Scan(&id)
	fmt.Println(o)

	if err != nil {
		fmt.Println(err)
	}
	return id, err
}

func (m *PostgreSql) AddNewUser(u models.User) (interface{}, error) {
	sqlStatement := `insert into users (username, password, email) values ($1, $2, $3) returning userid`
	var id int64
	err := m.client.QueryRow(sqlStatement, u.Username, u.Password, u.Email).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	return id, err
}

func (m *PostgreSql) UpdatePizza(p models.Pizza) (interface{}, error) {
	sqlStatement := `update pizza set name=$1, size=$2 WHERE name=$1`
	res, err := m.client.Exec(sqlStatement, p.Name, p.Size)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}

func (m *PostgreSql) DeletePizza(name string) (interface{}, error) {
	sqlStatement := `delete from pizza where name=$1`
	res, err := m.client.Exec(sqlStatement, name)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}

type ListPizzaOpinions struct {
	Name     string
	Size     string
	Opinions string
	Score    string
}

func (m *PostgreSql) ListPizzasWithOpinins() (models.ListPizzaOpinions, error) {
	sqlStatement := `select p.name, p.size, o.score, o.opinion from pizza p, opinions o where o.pizzaid=p.id`
	var listPO models.ListPizzaOpinions
	var tab []models.Together
	var opinTab []models.Opinion

	rows, err := m.client.Query(sqlStatement)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var pizza models.Pizza
		var opinion models.Opinion

		err = rows.Scan(&pizza.Name, &pizza.Size, &opinion.Score, &opinion.Opinions)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		opinTab = append(opinTab, models.Opinion{
			Score:    opinion.Score,
			Opinions: opinion.Opinions,
		})
		tab = append(tab, models.Together{
			MName: pizza.Name,
			MSize: pizza.Size,
			Ops:   opinTab,
		})
	}
	listPO = models.ListPizzaOpinions{
		ListPizzaWithOpinions: tab,
	}
	return listPO, nil
}

func (m *PostgreSql) Login(u models.User) bool {
	var users models.User
	sqlStatement := `select username, password from users`
	rows, err := m.client.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&users.Username, &users.Password)
		if err != nil {
			log.Fatal(err)
		}
		if users.Username == u.Username && users.Password == u.Password {
			fmt.Println("Witaj ", u.Username)
			return true
		}
	}

	return false
}
