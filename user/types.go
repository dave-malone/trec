package user

import (
	"fmt"

	"github.com/dave-malone/trec/common"
	"github.com/dave-malone/trec/company"
)

type repository interface {
	add(user User) (err error)
	listUsers() (users []User)
	getUser(id string) (user User, err error)
}

type User struct {
	ID        int64        `json:"id"`
	Email     string       `json:"email"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Password  string       `json:"-"`
	Company   trec.Company `json:"company"`
	Verified  bool         `json:"verified"`
}

type userListResponse struct {
	Total int    `json:"total"`
	Users []User `json:"users"`
}

func (user *User) String() string {
	return fmt.Sprintf("User{Id:%v, Verified:%v, Email:%v, Name:%v %v}", user.ID, user.Verified, user.Email, user.FirstName, user.LastName)
}

func (user *User) validate() (result common.ValidationResult) {
	result = common.NewValidationResult()

	if len(user.FirstName) == 0 {
		result.AddError("first_name", "First Name is required")
	}

	if len(user.LastName) == 0 {
		result.AddError("last_name", "Last Name is required")
	}

	if len(user.Email) == 0 {
		result.AddError("email", "Email is required")
	}

	return result
}

func newUser(ID int64, FirstName string, LastName string, Email string) *User {
	return &User{
		ID:        ID,
		FirstName: FirstName,
		LastName:  LastName,
		Email:     Email,
	}
}
