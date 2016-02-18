package main

import (
	"net/http"
	"os"

	"github.com/SpearWind/trec/user"
	"github.com/codegangsta/negroni"
	"github.com/dave-malone/email"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if awsEndpoint != "" && awsAccessKeyID != "" && awsSecretAccessKey != "" {
		email.NewSender = email.NewAmazonSESSender(awsEndpoint, awsAccessKeyID, awsSecretAccessKey)
	} else {
		email.NewSender = email.NewNoopSender
	}

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)

	n.UseHandler(router)
	return n
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
