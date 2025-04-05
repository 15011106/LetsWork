package infrastructure

import "database/sql"

type Database struct {
	Rds *sql.DB
}

func NewDatabase() (*Database, error) {

	rds, err := NewRdsConnection()
	if err != nil {
		return nil, err
	}

	return &Database{Rds: rds}, nil
}
