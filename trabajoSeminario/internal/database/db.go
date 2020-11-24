package database

import (
	"errors"
	"seminarioGo/trabajoSeminario/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // import driver sqlite3
)

//...
func NewDatabase(c *config.Config) (*sqlx.DB, error) {
	switch c.DB.Type {
	case "sqlite3":
		db, err := sqlx.Open(c.DB.Driver, c.DB.Conn)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil
	default:
		return nil, errors.New("tipo de db invalido")

	}

}
