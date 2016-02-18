package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/SpearWind/trec/common"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/xchapter7x/lo"
)

var (
	ErrInvalidUserId   = errors.New("invalid user id")
	ErrUserDoesntExist = errors.New("This user doesnt exist")
)

func repo() repository {
	profile := os.Getenv("PROFILE")

	var repo repository

	if profile == "mysql" {
		db, err := common.NewDbConn()
		if err != nil {
			repo = newMysqlRepository(db)
		}
	} else {
		lo.G.Info("Using in-memory repositories")
		repo = newInMemoryRepository()
	}

	return repo
}

func InitRoutes(router *mux.Router, formatter *render.Render) {
	repo := repo()
	router.HandleFunc("/user", createUserHandler(formatter, repo)).Methods("POST")
	router.HandleFunc("/user", getUserListHandler(formatter, repo)).Methods("GET")
	router.HandleFunc("/user/{id}", getUserHandler(formatter, repo)).Methods("GET")
}

func createUserHandler(formatter *render.Render, repo repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		var user User

		err := json.Unmarshal(payload, &user)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse create user request")
			return
		}

		if result := user.validate(); result.HasErrors() {
			formatter.JSON(w, http.StatusBadRequest, map[string]interface{}{
				"errors": result.Errors,
			})
			return
		}

		if err := repo.add(user); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, map[string]interface{}{
				"user":  user,
				"error": err.Error(),
			})
		} else {
			w.Header().Add("Location", fmt.Sprintf("/user/%d", user.ID))
			formatter.JSON(w, http.StatusCreated, user)
			newUserRegistrationEvent(user)
		}
	}
}

func getUserListHandler(formatter *render.Render, repo repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		users := repo.listUsers()

		formatter.JSON(w, http.StatusOK, map[string]interface{}{
			"users": users,
			"total": len(users),
		})
	}
}

func getUserHandler(formatter *render.Render, repo repository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		userID := vars["id"]

		if user, err := repo.getUser(userID); err != nil {
			formatter.JSON(w, http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
		} else {
			formatter.JSON(w, http.StatusOK, user)
		}
	}
}
