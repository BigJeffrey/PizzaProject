package middlewares

import (
	"pizza/dao/interfaces"
)

type Middleware struct {
	Dao interfaces.AppDao
}
