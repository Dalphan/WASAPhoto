package database

import (
	"database/sql"
	"errors"
	"strings"
)

const (
	SUCCESS = iota
	NO_ROWS
	UNIQUE_FAILED
	ERROR
)

func checkResults(err error) int {
	if errors.Is(err, sql.ErrNoRows) {
		return NO_ROWS
	}
	if err != nil {
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
