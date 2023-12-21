package database

import (
	"github.com/Dalphan/WASAPhoto/service/utils"
)

func (db *appdbimpl) FollowUser(uid int, fid int) (int, error) {
	_, err := db.c.Exec(`	INSERT INTO Followings
							VALUES (?, ?)`, uid, fid)

	if res := checkResults(err); res != SUCCESS {
		return res, err
	} else {
		return res, nil
	}
}

func (db *appdbimpl) UnfollowUser(uid int, fid int) (int, error) {
	sql, err := db.c.Exec(`	DELETE FROM Followings
							WHERE UID = ? AND FollowedID = ?`, uid, fid)

	if res, err := checkRowsAffected(sql); res != SUCCESS {
		return res, err
	}

	if res := checkResults(err); res != SUCCESS {
		return res, err
	} else {
		return res, nil
	}
}

func (db *appdbimpl) GetFollowings(uid int) ([]utils.User, int, error) {
	rows, err := db.c.Query(`	SELECT u.UID, u.Username
								FROM Users u, Followings f
								WHERE 	f.FollowedID = u.UID
								AND 	f.UID = ?`, uid)
	if err != nil {
		return nil, ERROR, err
	}
	// defer rows.Close()

	return getSelectedUsers(rows)
}

func (db *appdbimpl) GetFollowers(uid int) ([]utils.User, int, error) {
	rows, err := db.c.Query(`	SELECT u.UID, u.Username
								FROM Users u, Followings f
								WHERE 	f.UID = u.UID
								AND 	f.FollowedID = ?`, uid)
	if err != nil {
		return nil, ERROR, err
	}
	// defer rows.Close()

	return getSelectedUsers(rows)
}
