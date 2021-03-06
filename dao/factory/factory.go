package factory

import (
	"context"
	"log"

	"pizza/dao/interfaces"
	mongodao "pizza/dao/mongo"
	postgresqldao "pizza/dao/postgresql"
)

type FactoryDao struct {
	Ctx context.Context
}

// FactoryDao returns a DAO object
func (f *FactoryDao) FactoryDao(e string) interfaces.AppDao {
	var i interfaces.AppDao
	switch e {
	case "mongodb":
		i = mongodao.NewMongo(f.Ctx)
	case "postgresql":
		i = postgresqldao.NewPostgreSql(f.Ctx)
	default:
		log.Fatalf("Unsupported %s database", e)
		return nil
	}

	return i
}
