package trec

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// NewServer configures and returns a Server.
func NewServer() *martini.ClassicMartini {
	db := initDb()
	userRepo := newSqlUserRepository(db)

	m := martini.Classic()
	m.Map(userRepo)
	m.Use(render.Renderer())

	initRoutes(m)

	return m
}

func initDb() (db *DbConn) {
	dbUrl := flag.String("dburl", "trec:trec@localhost:3306/trec", "specify the MySQL database url to connect against")

	flag.Parse()

	var (
		data []byte
		err  error
	)

	if data, err = ioutil.ReadFile("create_db.sql"); err != nil {
		panic(fmt.Sprintf("Failed to read file: %v", err))
	}

	if db, err = newDbConn(*dbUrl); err != nil {
		panic(fmt.Sprintf("Failed to connect to db: %v", err))
	}

	if _, err = db.Exec(string(data)); err != nil {
		panic(fmt.Sprintf("Failed to create db tables: %v", err))
	}

	return db
}

func initRoutes(m *martini.ClassicMartini) {
	m.Get("/", func() string {
		return "trec api home; nothing to see here"
	})

	m.Group("/user", func(r martini.Router) {
		r.Get("/info", func() string {
			return "An API that allows you to work with Users"
		})
		r.Get("/", getUsersHandler)
		r.Get("/:id", getUserHandler)
		r.Post("/", binding.Json(User{}), createUserHandler)
	})

}
