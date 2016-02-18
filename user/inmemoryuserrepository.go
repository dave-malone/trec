package user

import (
	"errors"
	"strconv"
)

type inMemoryRepository struct {
	users []User
}

func newInMemoryRepository() *inMemoryRepository {
	repo := &inMemoryRepository{}
	repo.users = []User{}
	return repo
}

func (repo *inMemoryRepository) add(user User) (err error) {
	repo.users = append(repo.users, user)
	user.ID = int64(len(repo.users))
	return err
}
func (repo *inMemoryRepository) listUsers() (users []User) {
	return repo.users
}
func (repo *inMemoryRepository) getUser(id string) (user User, err error) {
	found := false

	for _, target := range repo.users {
		if userID, err := strconv.ParseInt(id, 10, 64); err == nil && userID == target.ID {
			user = target
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find user in repository")
	}
	return user, err
}
