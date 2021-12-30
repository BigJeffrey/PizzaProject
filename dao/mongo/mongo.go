package mongodao

import (
	"context"
	"log"
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

	m.client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://admin_user:admin_password@77.55.216.72:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = m.client.Connect(m.ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Mongo) Disconnect() {
	m.client.Disconnect(m.ctx)
}

func (m *Mongo) GetListaPizzy() *mongo.Collection {
	pizzaDatabase := m.client.Database("pizzowaBaza")
	return pizzaDatabase.Collection("pizze")
}

func (m *Mongo) GetOpinie() *mongo.Collection {
	opinie := m.client.Database("pizzowaBaza")
	return opinie.Collection("opinie")
}

func (m *Mongo) GetUsers() *mongo.Collection {
	users := m.client.Database("pizzowaBaza")
	return users.Collection("users")
}
