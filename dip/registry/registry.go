package registry

import (
	"database/sql"
	"fmt"

	"github.com/finn-inc/finn-server-tutorial/dip/config"
)

type registry struct {
	dbConn *sql.DB
}

func (r *registry) DBConn() *sql.DB {
	return r.dbConn
}

func NewRegistryImpl(env *config.Env) (Registry, error) {
	dbConn, err := sql.Open("postgres", env.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("postgresに接続できませんでした: %w", err)
	}

	return &registry{
		dbConn: dbConn,
	}, nil
}
