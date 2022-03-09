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
		return 0, nil
	}
	return id, nil
}

func (m *PostgreSql) AddNewOpinion(o models.Opinion) (interface{}, error) {
	sqlStatement := `insert into opinions (score, opinion, pizzaid) values ($1, $2, $3) returning opinionid`
	var id int64
	err := m.client.QueryRow(sqlStatement, o.Score, o.Opinions, o.PizzaId).Scan(&id)
	fmt.Println(o)

	if err != nil {
		return 0, err
	}
	return id, nil
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

func (m *PostgreSql) ListPizzasWithOpinins() (models.ListPizzaOpinions, error) {
	//sqlStatement := `select p.name, p.size, o.score, o.opinion from pizza p, opinions o where o.pizzaid=p.id`
	sqlStatementAllPizza := `select * from pizza`
	sqlStatementOpinions := `select * from opinions`
	pizzas, err := m.client.Query(sqlStatementAllPizza)
	if err != nil {
		log.Println(err)
	}
	opinions, err := m.client.Query(sqlStatementOpinions)
	if err != nil {
		log.Println(err)
	}

	var pizza models.Pizza
	var opinion models.Opinion
	var tabPizza []models.Pizza
	var tabOpinion []models.Opinion
	for pizzas.Next() {
		err = pizzas.Scan(&pizza.ID, &pizza.Name, &pizza.Size)
		if err != nil {
			log.Println(err)
		}
		tabPizza = append(tabPizza, models.Pizza{
			ID:   pizza.ID,
			Name: pizza.Name,
			Size: pizza.Size,
		})
	}
	for opinions.Next() {
		err = opinions.Scan(&opinion.ID, &opinion.Opinions, &opinion.Score, &opinion.PizzaId)
		if err != nil {
			log.Println(err)
		}
		tabOpinion = append(tabOpinion, models.Opinion{
			ID:       opinion.ID,
			Score:    opinion.Score,
			Opinions: opinion.Opinions,
			PizzaId:  opinion.PizzaId,
		})
	}

	var opinTab []models.Opinion
	var tab []models.Together
	var listPO models.ListPizzaOpinions

	for _, val := range tabPizza {
		for _, val2 := range tabOpinion {
			if val2.PizzaId == val.ID {
				opinTab = append(opinTab, models.Opinion{
					ID:       val2.ID,
					Score:    val2.Score,
					Opinions: val2.Opinions,
					PizzaId:  val2.PizzaId,
				})
			}
		}
		tab = append(tab, models.Together{
			MName: val.Name,
			MSize: val.Size,
			Ops:   opinTab,
		})
		opinTab = nil
	}

	listPO = models.ListPizzaOpinions{
		ListPizzaWithOpinions: tab,
	}

	return listPO, nil
}

func (m *PostgreSql) Login(u models.User) (bool, error) {
	var users models.User
	sqlStatement := `select username, password from users`
	rows, err := m.client.Query(sqlStatement)

	if err != nil {
		return false, err
	}

	for rows.Next() {
		err := rows.Scan(&users.Username, &users.Password)
		if err != nil {
			return false, err
		}
		if users.Username == u.Username && users.Password == u.Password {
			fmt.Println("Witaj ", u.Username)
			return true, nil
		}
	}

	return false, nil
}
