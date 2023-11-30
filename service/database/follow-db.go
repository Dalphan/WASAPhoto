package database

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
