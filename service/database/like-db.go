package database

import "github.com/Dalphan/WASAPhoto/service/utils"

func (db *appdbimpl) LikePhoto(pid int, uid int) (utils.Like, int, error) {
	var like utils.Like
	_, err := db.c.Exec(`	INSERT INTO Like
							VALUES (?, ?)`, pid, uid)

	if res := checkResults(err); res != SUCCESS {
		return like, res, err
	}
	like.PhotoID = pid
	like.UserID = uid
	return like, SUCCESS, nil
}

func (db *appdbimpl) UnlikePhoto(pid int, uid int) (int, error) {
	row, err := db.c.Exec(`	DELETE FROM Like
							WHERE PID = ? and UID = ?`, pid, uid)

	if res, err := checkRowsAffected(row); res != SUCCESS {
		return res, err
	}

	if res := checkResults(err); res != SUCCESS {
		return res, err
	}
	return SUCCESS, nil
}

// The uid is used to remove users who have banned the user receiving the list of likes
func (db *appdbimpl) GetLikesByPhoto(pid int, uid int) ([]utils.Like, int, error) {
	var likes []utils.Like
	rows, err := db.c.Query(`	SELECT l.PID, l.UID, u.Username
								FROM Like l, Users u 
								LEFT JOIN Bans b
								ON b.UID = l.UID
								AND b.BannedID = ?
								WHERE l.UID = u.UID
								AND l.PID = ?
								AND (b.BannedID IS NULL)`, uid, pid)
	if err != nil {
		return nil, ERROR, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			err = closeErr // Assign the error to the outer variable
		}
	}()

	for rows.Next() {
		var like utils.Like
		if err := rows.Scan(&like.PhotoID, &like.UserID, &like.Username); err != nil {
			return nil, ERROR, err
		}
		likes = append(likes, like)
	}
	return likes, SUCCESS, err
}
