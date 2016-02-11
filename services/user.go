package trec

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/xchapter7x/lo"
)

type (
	userRepository interface {
		addUser(user User) (err error)
		getUsers() (users []User)
		getUser(id string) (user User, err error)
	}

	userRepositoryFactory func() userRepository
)

var (
	newUserRepository  userRepositoryFactory
	ErrInvalidUserId   = errors.New("invalid user id")
	ErrUserDoesntExist = errors.New("This user doesnt exist")
)

type User struct {
	Id        int64   `json:"id"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Password  string  `json:"-"`
	Company   Company `json:"company"`
	Verified  bool    `json:"verified"`
}

func (user *User) String() string {
	return fmt.Sprintf("User{Id:%v, Verified:%v, Email:%v, Name:%v %v}", user.Id, user.Verified, user.Email, user.FirstName, user.LastName)
}

func (user *User) validate() (errs ValidationErrors) {
	errs = newValidationErrors()

	if len(user.FirstName) == 0 {
		errs.add("first_name", "First Name is required")
	}

	if len(user.LastName) == 0 {
		errs.add("last_name", "Last Name is required")
	}

	if len(user.Email) == 0 {
		errs.add("email", "Email is required")
	}

	return errs
}

func newUser(Id int64, FirstName string, LastName string, Email string) *User {
	return &User{
		Id:        Id,
		FirstName: FirstName,
		LastName:  LastName,
		Email:     Email,
	}
}

func createUserHandler(user User, repo userRepository, r render.Render) {
	errs := user.validate()

	if errs.isEmpty() != true {
		r.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": errs.Errors,
		})

		return
	}

	err := repo.addUser(user)
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

func getUsersHandler(repo userRepository, r render.Render) {
	users := repo.getUsers()

	responseCode := http.StatusOK

	r.JSON(responseCode, map[string]interface{}{
		"users": users,
	})

	return
}

func getUserHandler(repo userRepository, params martini.Params, r render.Render) {
	userId := params["id"]
	lo.G.Debugf("Getting user with id %v\n", userId)

	user, err := repo.getUser(userId)

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
