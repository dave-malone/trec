package trec

import (
	"errors"
	"fmt"

	"github.com/xchapter7x/lo"
)

type sqlUserRepository struct {
	dbConn *DbConn
}

func newSqlUserRepository(dbConn *DbConn) *sqlUserRepository {
	return &sqlUserRepository{dbConn}
}

func (repo *sqlUserRepository) addUser(user User) (err error) {
	result, err := repo.dbConn.Exec("INSERT INTO USER (first_name, last_name, email) values (?, ?, ?)", user.FirstName, user.LastName, user.Email)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to insert user; %v", err))
	}

	user.Id, err = result.LastInsertId()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to get DB generated ID; %v", err))
	}

	return nil
}
func (repo *sqlUserRepository) getUsers() (users []*User) {
	users = make([]*User, 0)

	rows, err := repo.dbConn.Query("SELECT id, first_name, last_name, email FROM USER")

	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id                         int64  = *new(int64)
			firstName, lastName, email string = *new(string), *new(string), *new(string)
		)

		if err := rows.Scan(&id, &firstName, &lastName, &email); err != nil {
			lo.G.Fatal(err)
		}

		user := newUser(id, firstName, lastName, email)
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	return users
}
func (repo *sqlUserRepository) getUser(userId string) (user User, err error) {
	var (
		id                         int64  = *new(int64)
		firstName, lastName, email string = *new(string), *new(string), *new(string)
	)

	if err := repo.dbConn.QueryRow("SELECT id, first_name, last_name, email FROM USER WHERE id = ?", userId).Scan(&id, &firstName, &lastName, &email); err == nil {
		user := newUser(id, firstName, lastName, email)

		return *user, nil
	} else {
		return *new(User), err
	}
}
