package models

type User struct {
	ID       int    `bson:"_id,omitempty"`
	Username string `bson:"username,omitemty"`
	Password string `bson:"password,omitempty"`
	Email    string `bson:"email,omitempty"`
}

type Pizza struct {
	ID   int    `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Size string `bson:"size,omitempty"`
}

type Opinion struct {
	ID       int    `bson:"_id,omitempty"`
	Score    string `bson:"score,omitempty"`
	Opinions string `bson:"opinion,omitempty"`
	PizzaId  int    `bson:"pizzaid,omitempty"`
}
