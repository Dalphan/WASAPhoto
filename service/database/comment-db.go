package database

import "github.com/Dalphan/WASAPhoto/service/utils"

func (db *appdbimpl) CommentPhoto(c utils.Comment) (int, error) {
	err := db.c.QueryRow(` 	INSERT INTO Comment (PID, UID, text, date)
							VALUES (?,?,?,?)`, c.PhotoID, c.UserID, c.Text, c.Date).Scan(&c.CommentID)

	if res := checkResults(err); res != SUCCESS {
		return res, err
	} else {
		return SUCCESS, nil
	}
}

func (db *appdbimpl) UncommentPhoto(cid int) (int, error) {
	row, err := db.c.Exec(`	DELETE FROM Comment
							WHERE CID = ?`, cid)

	if res, err := checkRowsAffected(row); res != SUCCESS {
		return res, err
	}

	if res := checkResults(err); res != SUCCESS {
		return res, err
	} else {
		return res, nil
	}
}
