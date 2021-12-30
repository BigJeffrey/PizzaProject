package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserMongo struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitemty"`
	Password string             `bson:"password,omitempty"`
	Email    string             `bson:"email,omitempty"`
}

type PizzaMongo struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
	Size string             `bson:"size,omitempty"`
}

type OpinionMongo struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Score    string             `bson:"score,omitempty"`
	Opinions string             `bson:"opinion,omitempty"`
	PizzaId  primitive.ObjectID `bson:"pizzaid,omitempty"`
}
