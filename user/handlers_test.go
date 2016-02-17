package user

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/dave-malone/email"
	"github.com/martini-contrib/render"
)

type fakeUserRepository struct {
	execResult     User
	execResults    []User
	execError      error
	SpyGetUserByID string
}

type fakeRender struct {
	render.Render
	SpyStatus  int
	SpyPayload map[string]interface{}
}

func (t *fakeRender) JSON(status int, v interface{}) {
	t.SpyStatus = status

	switch payload := v.(type) {
	case map[string]interface{}:
		t.SpyPayload = payload
	default:
		fmt.Printf("Payload isn't what we thought it would be and won't be assigned to SpyPayload: %v\n", payload)
	}

}

func (t *fakeUserRepository) AddUserReturns(err error) {
	t.execError = err
}

func (t *fakeUserRepository) GetUserReturns(user User, err error) {
	t.execResult = user
	t.execError = err
}

func (t *fakeUserRepository) GetUsersReturns(users []User, err error) {
	t.execResults = users
	t.execError = err
}

func (t *fakeUserRepository) Add(user User) (err error) {
	return t.execError
}

func (t *fakeUserRepository) GetUsers() (users []User) {
	return t.execResults
}

func (t *fakeUserRepository) GetUser(id string) (user User, err error) {
	t.SpyGetUserByID = id
	return t.execResult, t.execError
}

func TestValidateWithEmptyRequiredFieldsFailsWithErrors(t *testing.T) {
	user := NewUser(-1, "", "", "")

	validationErrors := user.Validate()

	if len(validationErrors.Errors) != 3 {
		t.Fatalf("Expected three errors, but there were %v errors: %v", len(validationErrors.Errors), validationErrors.Errors)
	}
}

func TestCreateUserHandler(t *testing.T) {
	email.NewSender = email.NewNoopSender

	r := new(fakeRender)
	user := NewUser(-1, "First", "Last", "Email")
	repo := new(fakeUserRepository)
	repo.AddUserReturns(nil)

	CreateUserHandler(*user, repo, r)

	if r.SpyStatus != http.StatusOK {
		t.Fatalf("Excected status %v but status was: %v", http.StatusOK, r.SpyStatus)
	}

	if len(r.SpyPayload) != 2 {
		t.Fatalf("expected exactly two values in the json payload; values %v", r.SpyPayload)
	}

	if val, ok := r.SpyPayload["err"]; val != nil || ok {
		t.Fatalf("error should have been nil in the json payload, but was %v", r.SpyPayload["err"])
	}

	if r.SpyPayload["user"] == nil {
		t.Fatal("user was nil in the json payload")
	}
}

func TestGetUserHandler(t *testing.T) {
	r := new(fakeRender)
	user := NewUser(1, "First", "Last", "Email")

	var params map[string]string
	params = make(map[string]string)
	params["id"] = "123"

	repo := new(fakeUserRepository)
	repo.GetUserReturns(*user, nil)

	GetUserHandler(repo, params, r)

	if (repo.SpyGetUserByID) != params["id"] {
		t.Fatalf("expected param value of %v; got %v", params["id"], repo.SpyGetUserByID)
	}

	if r.SpyStatus != http.StatusOK {
		t.Fatalf("Excected status %v but status was: %v", http.StatusOK, r.SpyStatus)
	}

	if len(r.SpyPayload) != 2 {
		t.Fatalf("expected exactly two values in the json payload; values %v", r.SpyPayload)
	}

	if r.SpyPayload["user"] == nil {
		t.Fatal("user was nil in the json payload")
	}
}

func TestGetUsersHandler(t *testing.T) {
	r := new(fakeRender)
	users := []User{*NewUser(1, "First", "Last", "Email")}
	repo := new(fakeUserRepository)
	repo.GetUsersReturns(users, nil)

	GetUsersHandler(repo, r)

	if r.SpyStatus != http.StatusOK {
		t.Fatalf("Excected status %v but status was: %v", http.StatusOK, r.SpyStatus)
	}

	if len(r.SpyPayload) != 1 {
		t.Fatalf("expected only one value; values %v", r.SpyPayload)
	}

	if r.SpyPayload["err"] != nil {
		t.Fatalf("error should have been nil in the response map, but was %v", r.SpyPayload["err"])
	}

	if r.SpyPayload["users"] == nil {
		t.Fatal("users was nil in the json payload")
	}
}
