package postgresqldao

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSql struct {
	client *sql.DB
	ctx    context.Context
}

var once sync.Once
var instance *PostgreSql

func NewPostgreSql(ctx context.Context) *PostgreSql {

	once.Do(func() {
		instance = new(PostgreSql)
		instance.ctx = ctx
		instance.connect()
	})

	return instance
}

func (m *PostgreSql) connect() {
	var err error

	m.client, err = sql.Open("postgres", "host=130.61.54.93 port=49153 user=postgres dbname=pizzas sslmode=disable password=")
	if err != nil {
		fmt.Println("To jest brak połaczenia z bazą")
		log.Fatal(err)
	}
}

func (m *PostgreSql) Disconnect() {
	m.client.Close()
}
