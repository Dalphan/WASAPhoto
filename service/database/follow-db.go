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

// bid parameter is used to remove from the list the users that banned the user performing the request
func (db *appdbimpl) GetFollowings(uid int, bid int) ([]utils.User, int, error) {
	rows, err := db.c.Query(`	SELECT u.UID, u.Username
								FROM Users u, Followings f
								LEFT JOIN Bans b
									ON f.FollowedID = b.UID 
									AND b.BannedID = ?
								WHERE 	f.FollowedID = u.UID
								AND (b.UID IS NULL OR b.BannedID IS NULL)
								AND f.UID = ?`, bid, uid)
	if err != nil {
		return nil, ERROR, err
	}

	return getSelectedUsers(rows)
}

// bid parameter is used to remove from the list the users that banned the user performing the request
func (db *appdbimpl) GetFollowers(uid int, bid int) ([]utils.User, int, error) {
	rows, err := db.c.Query(`	SELECT u.UID, u.Username
								FROM Users u, Followings f
								LEFT JOIN Bans b
								ON f.UID = b.UID 
								AND b.BannedID = ?
								WHERE 	f.UID = u.UID
								AND (b.UID IS NULL OR b.BannedID IS NULL)
								AND 	f.FollowedID = ?`, bid, uid)
	if err != nil {
		return nil, ERROR, err
	}

	return getSelectedUsers(rows)
}
