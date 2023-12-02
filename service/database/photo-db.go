package database

import "github.com/Dalphan/WASAPhoto/service/utils"

func (db *appdbimpl) CreatePhoto(photo utils.Photo) (int, int, error) {
	var PID int
	err := db.c.QueryRow(`	INSERT INTO Photo (UID, image, date)
							VALUES (?, ?, ?)
							RETURNING PID`, photo.UserID, photo.Image, photo.Timestamp).Scan(&PID)

	return PID, checkResults(err), err
}

func (db *appdbimpl) DeletePhoto(pid int) (int, error) {
	row, err := db.c.Exec(`	DELETE FROM Photo
							WHERE PID = ?`, pid)

	if err != nil {
		return ERROR, err
	}

	if res, err := checkRowsAffected(row); res != SUCCESS {
		return res, err
	}
	return SUCCESS, nil
}

func (db *appdbimpl) GetPhotoById(pid int) (utils.Photo, int, error) {
	var photo utils.Photo
	err := db.c.QueryRow(`	SELECT *
							FROM Photo
							WHERE PID = ?`, pid).Scan(&photo.PhotoID, &photo.UserID, &photo.Image, &photo.Timestamp)

	if res := checkResults(err); res != SUCCESS {
		return *new(utils.Photo), res, err
	}
	return photo, SUCCESS, nil
}
