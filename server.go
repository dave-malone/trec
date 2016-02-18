package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/dave-malone/email"
	"github.com/dave-malone/trec/user"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	initFactories()

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)

	n.UseHandler(router)
	// m.Use(render.Renderer())

	return n
}

func initFactories() {
	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if awsEndpoint != "" && awsAccessKeyID != "" && awsSecretAccessKey != "" {
		email.NewSender = email.NewAmazonSESSender(awsEndpoint, awsAccessKeyID, awsSecretAccessKey)
	} else {
		email.NewSender = email.NewNoopSender
	}
}

func initRoutes(router *mux.Router, formatter *render.Render) {
	user.InitRoutes(router, formatter)
	router.HandleFunc("/", homeHandler(formatter)).Methods("GET")
}

func homeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Message string }{"trec api home; nothing to see here"})
	}
}

// m.Get("/", func() string {
// 	return "trec api home; nothing to see here"
// })
//
// m.Post("/login", auth.LoginHandler)
// m.Get("/validate", auth.ValidationHandler)
//
// m.Group("/company", func(r martini.Router) {
// 	r.Get("/info", func() string {
// 		return "An API that allows you to work with Companies"
// 	})
// })
