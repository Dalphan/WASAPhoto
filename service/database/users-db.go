package database

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/Dalphan/WASAPhoto/service/utils"
)

func checkUpdateUser(user utils.User, err error) (utils.User, int, error) {
	if errors.Is(err, sql.ErrNoRows) {
		return *new(utils.User), NO_ROWS, err
	}
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return *new(utils.User), UNIQUE_FAILED, err
		}
		return *new(utils.User), ERROR, err
	}

	return user, SUCCESS, nil
}

func (db *appdbimpl) UpdateUsername(uid int, username string) (utils.User, int, error) {
	var user utils.User
	var name, surname sql.NullString

	err := db.c.QueryRow(`	UPDATE Users
							SET
								Username = ?
							WHERE UID = ?
							RETURNING *`, username, uid).Scan(&user.UserID, &user.Username, &name, &surname)

	user.Name = name.String
	user.Surname = surname.String
	return checkUpdateUser(user, err)
}

func (db *appdbimpl) UpdateUser(user utils.User) (utils.User, int, error) {
	err := db.c.QueryRow(`	UPDATE Users
							SET
								Username = ?,
								name = ?,
								surname = ?
							WHERE UID = ?
							RETURNING *`, user.Username, user.Name, user.Surname, user.UserID).Scan(&user.UserID, &user.Username, &user.Name, &user.Surname)

	return checkUpdateUser(user, err)
}

func (db *appdbimpl) FindUserByUsername(username string) (utils.User, int, error) {
	var user utils.User

	err := db.c.QueryRow(`	SELECT 
								UID,
								Username,
								name,
								surname
							FROM Users
							Where Username = ?`, username).Scan(&user.UserID, &user.Username, &user.Name, &user.Surname)

	// La SELECT non ritorna niente
	if errors.Is(err, sql.ErrNoRows) {
		return *new(utils.User), NO_ROWS, nil
	} else if user.UserID != 0 { // La SELECT ha ritornato un utente che ha solo username al momento
		return user, SUCCESS, nil
	} else if err != nil { //La SELECT ha ritornato un errore imprevisto
		return user, ERROR, err
	}
	//La SELECT ha ritornato l'utente completo di informazioni
	return user, SUCCESS, nil
}

func (db *appdbimpl) CreateUser(username string) (int, int, error) {
	var UID int
	err := db.c.QueryRow(`	INSERT INTO Users (Username)
							VALUES (?)
							RETURNING UID`, username).Scan(&UID)
	return UID, SUCCESS, err
}
