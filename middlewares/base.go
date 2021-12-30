package middlewares

import (
	"quickstart/dao/interfaces"
)

type Middleware struct {
	Dao interfaces.AppDao
}
