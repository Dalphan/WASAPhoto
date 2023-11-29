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
