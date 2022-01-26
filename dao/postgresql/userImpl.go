package postgresqldao

import (
	"encoding/json"
	"fmt"
	"log"
	"pizza/models"
)

func (m *PostgreSql) AddNewPizza(p models.Pizza) (interface{}, error) {
	fmt.Println("postgresql")
	sqlStatement := `insert into pizza (name, size) values ($1, $2) returning id`
	var id int64
	err := m.client.QueryRow(sqlStatement, p.Name, p.Size).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	return id, err
}

func (m *PostgreSql) AddNewOpinion(o models.Opinion) (interface{}, error) {
	fmt.Println("postgresql")
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
	fmt.Println("postgresql")
	sqlStatement := `insert into users (username, password, email) values ($1, $2, $3) returning userid`
	var id int64
	err := m.client.QueryRow(sqlStatement, u.Username, u.Password, u.Email).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

	return id, err
}

func (m *PostgreSql) UpdatePizza(p models.Pizza) (interface{}, error) {
	fmt.Println("postgresql")
	sqlStatement := `update pizza set name=$1, size=$2 WHERE name=$1`
	res, err := m.client.Exec(sqlStatement, p.Name, p.Size)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}

func (m *PostgreSql) DeletePizza(name string) (interface{}, error) {
	fmt.Println("postgresql")
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

func (m *PostgreSql) ListPizzasWithOpinins() {
	fmt.Println("postgresql")

	sqlStatement := `select p.name, p.size, o.score, o.opinion from pizza p, opinions o where o.pizzaid=p.id`

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

		jsoned, err := json.Marshal(ListPizzaOpinions{Name: pizza.Name, Size: pizza.Size, Opinions: opinion.Opinions, Score: opinion.Score})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsoned))
	}
}

func (m *PostgreSql) Login(u models.User) bool {
	fmt.Println("postgresql")
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
