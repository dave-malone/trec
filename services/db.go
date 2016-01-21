package trec

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DbConn struct {
	*sql.DB
}

func newDbConn(dbUrl string) (*DbConn, error) {
	db, err := sql.Open("mysql", dbUrl)

	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to connect to db %s: error: %v", dbUrl, err))
		return nil, err
	}

	if err = db.Ping(); err != nil {
		err = errors.New(fmt.Sprintf("Unable to ping db %s: error: %v", dbUrl, err))
		return nil, err
	}

	return &DbConn{db}, nil
}
