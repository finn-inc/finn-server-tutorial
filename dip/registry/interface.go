package registry

import "database/sql"

type Registry interface {
	DBConn() *sql.DB
}
