/*
VIDEO_3
-------
modificamos config.go, CONFIG/config.yaml
importamos manualmente  _ "github.com/mattn/go-sqlite3"
ejecutamos go mod tidy
*/
package database

import (
	"errors"

	"github.com/CristianMarsico/TP-Go/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// NewDataBase ...
func NewDataBase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DB.Type {
	case "sqlite3":
		db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil

	default:
		return nil, errors.New("invalid db type")
	}
}
