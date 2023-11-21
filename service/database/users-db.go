package database

import (
	"database/sql"
	"errors"

	"github.com/Dalphan/WASAPhoto/service/utils"
)

func (db *appdbimpl) FindUserByUsername(username string) (utils.User, int, error) {
	var UID int
	var name, surname string
	err := db.c.QueryRow(`	SELECT 
								UID,
								Username,
								name,
								surname
							FROM Users
							Where Username = ?`, username).Scan(&UID, &name, &surname)
	if errors.Is(err, sql.ErrNoRows) {
		return *new(utils.User), NO_ROWS, nil
	} else if err != nil {
		return *new(utils.User), NO_ROWS, nil
	}

	return *new(utils.User), SUCCESS, nil
}

func (db *appdbimpl) CreateUser(username string) (int, int, error) {
	var UID int
	err := db.c.QueryRow(`	INSERT INTO Users (Username)
							VALUES (?)
							RETURNING UID`, username).Scan(&UID)
	return UID, SUCCESS, err
}
