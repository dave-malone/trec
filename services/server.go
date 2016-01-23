package trec

import (
	"os"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/xchapter7x/lo"
)

// NewServer configures and returns a Server.
func NewServer() *martini.ClassicMartini {
	m := martini.Classic()
	initRoutes(m)
	initMappings(m)
	m.Use(render.Renderer())

	return m
}

func initMappings(m *martini.ClassicMartini) {
	profile := os.Getenv("PROFILE")

	m.Map(newNoopEmailSender())

	if profile == "mysql" {
		db, err := newDbConn()
		if err != nil {
			userRepo := newMysqlUserRepository(db)
			m.Map(userRepo)
		}
	} else {
		lo.G.Info("Using in-memory repositories")

		userRepo := newInMemoryUserRepository()
		m.Map(userRepo)
	}
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
