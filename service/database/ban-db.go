package database

import "github.com/Dalphan/WASAPhoto/service/utils"

// Banning a user involves removing their follow status from both users
// and deleting likes and comments made by the banned user on the photos
// of the user who is implementing the ban.
func (db *appdbimpl) BanUser(uid int, bid int) (int, error) {
	tx, err := db.c.Begin()
	if err != nil {
		return ERROR, err
	}

	_, err = tx.Exec(`	INSERT INTO Bans
						VALUES (?, ?)`, uid, bid)

	res := checkResults(err)
	if res != SUCCESS {
		err2 := tx.Rollback()
		if err2 != nil {
			err = err2
		}
		return res, err
	}
	_, err = tx.Exec(`	DELETE FROM Followings
						WHERE (UID = ? AND FollowedID)
						OR (UID = ? and FollowedID = ?)`, uid, bid, bid, uid)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = err2
		}
		return ERROR, err
	}

	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = err2
		}
		return ERROR, err
	}
	return SUCCESS, err
}

func (db *appdbimpl) UnbanUser(uid int, bid int) (int, error) {
	sql, err := db.c.Exec(`	DELETE FROM Bans
							WHERE UID = ? AND BannedID = ?`, uid, bid)

	if res, err := checkRowsAffected(sql); res != SUCCESS {
		return res, err
	}

	if res := checkResults(err); res != SUCCESS {
		return res, err
	} else {
		return res, nil
	}
}

func (db *appdbimpl) GetBanned(uid int) ([]utils.User, int, error) {
	rows, err := db.c.Query(`	SELECT u.UID, u.Username
								FROM 	Users u, Bans b
								WHERE 	b.BannedID = u.UID
								AND 	b.UID = ?`, uid)
	if err != nil {
		return nil, ERROR, err
	}

	return getSelectedUsers(rows)
}
