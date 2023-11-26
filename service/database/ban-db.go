package database

func (db *appdbimpl) BanUser(uid int, bid int) (int, error) {
	_, err := db.c.Exec(`	INSERT INTO Bans
							VALUES (?, ?)`, uid, bid, uid)

	if res := checkResults(err); res != SUCCESS {
		return res, err
	} else {
		return res, nil
	}
}
