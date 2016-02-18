package user

import (
	"fmt"

	"github.com/SpearWind/trec/common"
	"github.com/SpearWind/trec/company"
)

type (
	Repository interface {
		Add(user User) (err error)
		GetUsers() (users []User)
		GetUser(id string) (user User, err error)
	}

	RepositoryFactory func() Repository
)

var (
	NewRepository RepositoryFactory
)

type User struct {
	ID        int64        `json:"id"`
	Email     string       `json:"email"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Password  string       `json:"-"`
	Company   trec.Company `json:"company"`
	Verified  bool         `json:"verified"`
}

func (user *User) String() string {
	return fmt.Sprintf("User{Id:%v, Verified:%v, Email:%v, Name:%v %v}", user.ID, user.Verified, user.Email, user.FirstName, user.LastName)
}

func (user *User) Validate() (errs common.ValidationErrors) {
	errs = common.NewValidationErrors()

	if len(user.FirstName) == 0 {
		errs.Add("first_name", "First Name is required")
	}

	if len(user.LastName) == 0 {
		errs.Add("last_name", "Last Name is required")
	}

	if len(user.Email) == 0 {
		errs.Add("email", "Email is required")
	}

	return errs
}

func NewUser(ID int64, FirstName string, LastName string, Email string) *User {
	return &User{
		ID:        ID,
		FirstName: FirstName,
		LastName:  LastName,
		Email:     Email,
	}
}
