package user

import (
	"errors"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/xchapter7x/lo"
)

var (
	ErrInvalidUserId   = errors.New("invalid user id")
	ErrUserDoesntExist = errors.New("This user doesnt exist")
)

func CreateUserHandler(user User, repo Repository, r render.Render) {
	errs := user.Validate()

	if errs.IsEmpty() != true {
		r.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": errs.Errors,
		})

		return
	}

	err := repo.Add(user)
	responseCode := http.StatusOK
	errMsg := ""

	if err != nil {
		lo.G.Errorf("An error occurred when saving user: %v", err)
		errMsg = err.Error()
		responseCode = http.StatusInternalServerError
	} else {
		newUserRegistrationEvent(user)
	}

	r.JSON(responseCode, map[string]interface{}{
		"user":  user,
		"error": errMsg,
	})

	return
}

func GetUsersHandler(repo Repository, r render.Render) {
	users := repo.GetUsers()

	responseCode := http.StatusOK

	r.JSON(responseCode, map[string]interface{}{
		"users": users,
	})

	return
}

func GetUserHandler(repo Repository, params martini.Params, r render.Render) {
	userId := params["id"]
	lo.G.Debugf("Getting user with id %v\n", userId)

	user, err := repo.GetUser(userId)

	responseCode := http.StatusOK
	errMsg := ""

	if err != nil {
		lo.G.Errorf("An error occurred when getting user: %v", err)
		errMsg = err.Error()
		responseCode = http.StatusInternalServerError
	}

	r.JSON(responseCode, map[string]interface{}{
		"user":  user,
		"error": errMsg,
	})

	return
}
