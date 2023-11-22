package database

import (
	"database/sql"
	"errors"

	"github.com/Dalphan/WASAPhoto/service/utils"
)

func (db *appdbimpl) FindUserByUsername(username string) (utils.User, int, error) {
	var user utils.User

	err := db.c.QueryRow(`	SELECT 
								UID,
								name,
								surname
							FROM Users
							Where Username = ?`, username).Scan(&user.UserID, &user.Name, &user.Surname)

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
