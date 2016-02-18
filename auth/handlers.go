package auth

import (
	"fmt"
	"net/http"
)

func ValidationHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("the future home of the ValidationHandler")
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("the future home of the LoginHandler")
}
