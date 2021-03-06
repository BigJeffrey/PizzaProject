package mongodao

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
	ctx    context.Context
}

var once sync.Once
var instance *Mongo

func NewMongo(ctx context.Context) *Mongo {

	once.Do(func() {
		instance = new(Mongo)
		instance.ctx = ctx
		instance.connect()
	})

	return instance
}

func (m *Mongo) connect() {
	var err error
	mgPass := os.Getenv("PGP")
	conStr := "mongodb://pizza:" + mgPass + "@130.61.54.93:27017"
	m.client, err = mongo.NewClient(options.Client().ApplyURI(conStr))
	if err != nil {
		log.Fatal(err)
	}

	err = m.client.Connect(m.ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Mongo) Disconnect() {
	err := m.client.Disconnect(m.ctx)
	if err != nil {
		log.Println(err)
	}
}

func (m *Mongo) GetPizzaList() *mongo.Collection {
	pizzaDatabase := m.client.Database("pizzowaBaza")
	return pizzaDatabase.Collection("pizze")
}

func (m *Mongo) GetOpinions() *mongo.Collection {
	opinions := m.client.Database("pizzowaBaza")
	return opinions.Collection("opinie")
}

func (m *Mongo) GetUsers() *mongo.Collection {
	users := m.client.Database("pizzowaBaza")
	return users.Collection("users")
}
