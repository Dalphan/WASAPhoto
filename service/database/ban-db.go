package database

import "fmt"

func (db *appdbimpl) BanUser(uid int, bid int) (int, error) {
	_, err := db.c.Exec(`	INSERT INTO Bans
							VALUES (?, ?)`, uid, bid)

	if res := checkResults(err); res != SUCCESS {
		return res, err
	} else {
		return res, nil
	}
}

func (db *appdbimpl) UnbanUser(uid int, bid int) (int, error) {
	a, err := db.c.Exec(`	DELETE FROM Bans
							WHERE UID = ? AND BannedID = ?`, uid, bid)

	fmt.Println("ERRORE SE ESITSTE ", err)
	fmt.Println("SQL RESULT: ", a)

	if res := checkResults(err); res != SUCCESS {
		return res, err
	} else {
		return res, nil
	}
}
