package database

import (
	"database/sql"
	"errors"

	"github.com/Dalphan/WASAPhoto/service/utils"
)

const LIMIT_STREAM = 10

func checkUpdateUser(user utils.User, err error) (utils.User, int, error) {
	res := checkResults(err)

	if res != SUCCESS {
		return *new(utils.User), res, err
	}

	return user, SUCCESS, nil
}

func (db *appdbimpl) fillUser(user utils.User) (utils.User, int, error) {
	err := db.c.QueryRow(`	SELECT COUNT(*)
							FROM Followings
							WHERE UID = ?`, user.UserID).Scan(&user.FollowingCount)
	if err != nil {
		return user, ERROR, err
	}

	err = db.c.QueryRow(`	SELECT COUNT(*)
							FROM Followings
							WHERE FollowedID = ?`, user.UserID).Scan(&user.FollowersCount)
	if err != nil {
		return user, ERROR, err
	}

	err = db.c.QueryRow(`	SELECT COUNT(*)
							FROM Photo
							WHERE UID = ?`, user.UserID).Scan(&user.PhotoCount)
	if err != nil {
		return user, ERROR, err
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

func (db *appdbimpl) FindUserByID(uid int) (utils.User, int, error) {
	var user utils.User

	err := db.c.QueryRow(`	SELECT *
							FROM Users
							WHERE UID = ?`, uid).Scan(&user.UserID, &user.Username, &user.Name, &user.Surname)

	if res := checkResults(err); res != SUCCESS {
		return *new(utils.User), res, err
	} else {
		return db.fillUser(user)
	}
}

func (db *appdbimpl) FindUserByUsername(username string) (utils.User, int, error) {
	var user utils.User

	err := db.c.QueryRow(`	SELECT *
							FROM Users
							Where Username = ?`, username).Scan(&user.UserID, &user.Username, &user.Name, &user.Surname)

	// La SELECT non ritorna niente
	if errors.Is(err, sql.ErrNoRows) {
		return *new(utils.User), NO_ROWS, nil
	} else if user.UserID != 0 { // La SELECT ha ritornato un utente che ha solo username al momento
		return user, SUCCESS, nil
	} else if err != nil { // La SELECT ha ritornato un errore imprevisto
		return user, ERROR, err
	}
	// La SELECT ha ritornato l'utente completo di informazioni
	return db.fillUser(user)
}

func (db *appdbimpl) CreateUser(username string) (int, int, error) {
	var UID int
	err := db.c.QueryRow(`	INSERT INTO Users (Username)
							VALUES (?)
							RETURNING UID`, username).Scan(&UID)
	return UID, SUCCESS, err
}

func (db *appdbimpl) GetUserStream(uid int, page int) ([]utils.Photo, int, error) {
	rows, err := db.c.Query(`	SELECT p.*
								FROM Followings f, Photo p
								WHERE p.UID = f.FollowedID
								AND f.UID = ?
								ORDER BY p.date DESC
								LIMIT ? OFFSET ?`, uid, LIMIT_STREAM, page*LIMIT_STREAM)

	var photos []utils.Photo
	if err != nil {
		return nil, ERROR, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			err = closeErr // Assign the error to the outer variable
		}
	}()

	for rows.Next() {
		var photo utils.Photo
		var nullCaption sql.NullString
		if err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.Image, &photo.Timestamp, &nullCaption); err != nil {
			return nil, ERROR, err
		}
		photo.Caption = nullCaption.String
		photo, res, err := db.fillPhoto(photo)
		if res != SUCCESS {
			return nil, ERROR, err
		}
		photos = append(photos, photo)
	}
	return photos, SUCCESS, nil
}

func (db *appdbimpl) CheckIfBanned(uid int, bid int) (bool, error) {
	var ret bool
	err := db.c.QueryRow(`SELECT EXISTS(SELECT *
									FROM BANS
									WHERE UID = ?
									AND BannedID = ?)`, uid, bid).Scan(&ret)

	if err != nil {
		return ret, err
	}

	return ret, nil
}
