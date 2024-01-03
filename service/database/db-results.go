package database

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/Dalphan/WASAPhoto/service/utils"
)

const (
	SUCCESS = iota
	NO_ROWS
	UNIQUE_FAILED
	ERROR
)

func checkRowsAffected(res sql.Result) (int, error) {
	if res != nil {
		rows, err := res.RowsAffected()

		if err != nil {
			return ERROR, err
		}

		if rows == 0 {
			return NO_ROWS, nil
		}

		return SUCCESS, nil
	}
	return ERROR, errors.New("unexpected Database error")
}

func checkResults(err error) int {
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || strings.Contains(err.Error(), "KEY constraint failed") {
			return NO_ROWS
		}
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return UNIQUE_FAILED
		}
		if strings.Contains(err.Error(), "converting NULL") {
			return SUCCESS
		}
		return ERROR
	}

	return SUCCESS
}

func getSelectedUsers(rows *sql.Rows) ([]utils.User, int, error) {
	var err error = nil
	// defer func() {
	// 	if closeErr := rows.Close(); closeErr != nil {
	// 		err = closeErr // Assign the error to the outer variable
	// 	}
	// }()

	var users []utils.User
	for rows.Next() {
		var user utils.User
		if err = rows.Scan(&user.UserID, &user.Username); err != nil {
			return nil, ERROR, err
		}
		users = append(users, user)
	}

	_ = rows.Close()

	return users, SUCCESS, err
}
