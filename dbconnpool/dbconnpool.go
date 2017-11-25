package dbconnpool

import (
	"database/sql"
)

var (
	Connections = make(map[string]*sql.DB)
)
