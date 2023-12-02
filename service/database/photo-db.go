package database

import "github.com/Dalphan/WASAPhoto/service/utils"

func (db *appdbimpl) CreatePhoto(photo utils.Photo) (int, int, error) {
	var PID int
	err := db.c.QueryRow(`	INSERT INTO Photo (UID, image, date)
							VALUES (?, ?, ?)
							RETURNING PID`, photo.UserID, photo.Image, photo.Timestamp).Scan(&PID)

	return PID, checkResults(err), err
}
