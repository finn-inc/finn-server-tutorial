package registry

import (
	"database/sql"
)

type RegistryImpl struct {
	Registry
	dbConn *sql.DB
}

func (r RegistryImpl) DBConn() *sql.DB {
	return r.dbConn
}

type RegistryConfig struct {
	DBConn *sql.DB
}

func NewRegistryImpl(cfg RegistryConfig) RegistryImpl {
	return RegistryImpl{
		dbConn: cfg.DBConn,
	}
}
