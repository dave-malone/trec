package trec

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

type DbConn struct {
	*sql.DB
}

func newDbConn() (*DbConn, error) {
	dbUrl := flag.String("dburl", "trec:trec@localhost:3306/trec", "specify the MySQL database url to connect against")

	flag.Parse()

	var (
		data []byte
		err  error
	)

	if data, err = ioutil.ReadFile("create_db.sql"); err != nil {
		panic(fmt.Sprintf("Failed to read file: %v", err))
	}

	db, err := sql.Open("mysql", *dbUrl)

	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to connect to db %s: error: %v", *dbUrl, err))
		return nil, err
	}

	if err = db.Ping(); err != nil {
		err = errors.New(fmt.Sprintf("Unable to ping db %s: error: %v", *dbUrl, err))
		return nil, err
	}

	if _, err = db.Exec(string(data)); err != nil {
		panic(fmt.Sprintf("Failed to create db tables: %v", err))
	}

	return &DbConn{db}, nil
}
