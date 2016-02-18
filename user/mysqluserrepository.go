package user

import (
	"errors"
	"fmt"

	"github.com/SpearWind/trec/common"
	"github.com/xchapter7x/lo"
)

type mysqlRepository struct {
	dbConn *common.DbConn
}

func newMysqlRepository(dbConn *common.DbConn) *mysqlRepository {
	return &mysqlRepository{dbConn}
}

func (repo *mysqlRepository) add(user User) (err error) {
	result, err := repo.dbConn.Exec("INSERT INTO USER (first_name, last_name, email) values (?, ?, ?)", user.FirstName, user.LastName, user.Email)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to insert user; %v", err))
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to get DB generated ID; %v", err))
	}

	return nil
}
func (repo *mysqlRepository) listUsers() (users []User) {
	users = make([]User, 0)

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
		users = append(users, *user)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	return users
}
func (repo *mysqlRepository) getUser(userID string) (user User, err error) {
	var (
		id                         int64  = *new(int64)
		firstName, lastName, email string = *new(string), *new(string), *new(string)
	)

	if err := repo.dbConn.QueryRow("SELECT id, first_name, last_name, email FROM USER WHERE id = ?", userID).Scan(&id, &firstName, &lastName, &email); err == nil {
		user := newUser(id, firstName, lastName, email)

		return *user, nil
	} else {
		return User{}, err
	}
}
