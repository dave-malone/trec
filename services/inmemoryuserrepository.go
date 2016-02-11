package trec

import (
	"errors"
	"strconv"
)

type inMemoryUserRepository struct {
	users []User
}

func newInMemoryUserRepository() userRepository {
	repo := inMemoryUserRepository{}
	repo.users = []User{}
	return repo
}

func (repo inMemoryUserRepository) addUser(user User) (err error) {
	repo.users = append(repo.users, user)
	return err
}
func (repo inMemoryUserRepository) getUsers() (users []User) {
	return repo.users
}
func (repo inMemoryUserRepository) getUser(id string) (user User, err error) {
	found := false

	for _, target := range repo.users {
		if userId, err := strconv.ParseInt(id, 10, 64); err == nil && userId == target.Id {
			user = target
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find user in repository")
	}
	return user, err
}
