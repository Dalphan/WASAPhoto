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

func (db *appdbimpl) UncommentPhoto(cid int, pid int) (int, error) {
	row, err := db.c.Exec(`	DELETE FROM Comment
							WHERE CID = ? AND PID = ?`, cid, pid)

	if res, err := checkRowsAffected(row); res != SUCCESS {
		return res, err
	}

	if res := checkResults(err); res != SUCCESS {
		return res, err
	} else {
		return res, nil
	}
}

func (db *appdbimpl) GetCommentsByPhoto(pid int) ([]utils.Comment, int, error) {
	var c []utils.Comment
	rows, err := db.c.Query(`	SELECT *
								FROM Comment
								WHERE PID = ?`, pid)
	if err != nil {
		return nil, ERROR, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment utils.Comment
		if err := rows.Scan(&comment.CommentID, &comment.PhotoID, &comment.UserID, &comment.Text, &comment.Date); err != nil {
			return nil, ERROR, err
		}
		c = append(c, comment)
	}
	return c, SUCCESS, nil
}

func (db *appdbimpl) GetCommentById(cid int) (utils.Comment, int, error) {
	var c utils.Comment
	err := db.c.QueryRow(`	SELECT *
							FROM Comment
							WHERE CID = ?`, cid).Scan(&c.CommentID, &c.PhotoID, &c.UserID, &c.Text, &c.Date)

	return c, checkResults(err), err
}
