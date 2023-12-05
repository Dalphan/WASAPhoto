package database

import "github.com/Dalphan/WASAPhoto/service/utils"

func (db *appdbimpl) CommentPhoto(c utils.Comment) (int, error) {
	err := db.c.QueryRow(` 	INSERT INTO Comment (PID, UID, text, date)
							VALUES (?,?,?,?)`, c.PhotoID, c.UserID, c.Text, c.Date).Scan(&c.CommentID)

	if err != nil {
		return ERROR, err
	}
	return SUCCESS, nil
}
