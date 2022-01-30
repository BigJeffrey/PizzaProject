package models

type User struct {
	ID       interface{} `bson:"_id,omitempty"`
	Username string      `bson:"username,omitemty"`
	Password string      `bson:"password,omitempty"`
	Email    string      `bson:"email,omitempty"`
}

type Pizza struct {
	ID   interface{} `bson:"_id,omitempty"`
	Name string      `bson:"name,omitempty"`
	Size string      `bson:"size,omitempty"`
}

type Opinion struct {
	ID       interface{} `bson:"_id,omitempty"`
	Score    string      `bson:"score,omitempty"`
	Opinions string      `bson:"opinion,omitempty"`
	PizzaId  interface{} `bson:"pizzaid,omitempty"`
}

type Together struct {
	MName string
	MSize string
	Ops   []Opinion
}

type ListPizzaOpinions struct {
	ListPizzaWithOpinions []Together
}

type ErrorJson struct {
	Message string
}
