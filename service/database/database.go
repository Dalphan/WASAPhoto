/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Dalphan/WASAPhoto/service/utils"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	FindUserByUsername(string) (utils.User, int, error)
	FindUserByID(int) (utils.User, int, error)
	CreateUser(string) (int, int, error)
	UpdateUser(utils.User) (utils.User, int, error)
	UpdateUsername(int, string) (utils.User, int, error)
	GetUserStream(int, int) ([]utils.Photo, int, error)
	GetUserPhotos(int, int) ([]utils.Photo, int, error)
	CheckIfBanned(int, int) (bool, error)
	BanUser(int, int) (int, error)
	UnbanUser(int, int) (int, error)
	FollowUser(int, int) (int, error)
	UnfollowUser(int, int) (int, error)
	GetFollowings(int) ([]utils.User, int, error)
	GetFollowers(int) ([]utils.User, int, error)
	CreatePhoto(utils.Photo) (int, int, error)
	DeletePhoto(int) (int, error)
	GetPhotoById(int) (utils.Photo, int, error)
	LikePhoto(int, int) (utils.Like, int, error)
	UnlikePhoto(int, int) (int, error)
	GetLikesByPhoto(int) ([]utils.Like, int, error)
	CommentPhoto(utils.Comment) (int, int, error)
	UncommentPhoto(int, int) (int, error)
	GetCommentById(int) (utils.Comment, int, error)
	GetCommentsByPhoto(int, int) ([]utils.Comment, int, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='Users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		// _, err = db.Exec(sqlStmt)
		// if err != nil {
		// 	return nil, fmt.Errorf("error creating database structure: %w", err)
		// }
		sqlStmt := `PRAGMA foreign_keys = ON`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		sqlStmt = `CREATE TABLE Users (
			UID INTEGER NOT NULL PRIMARY KEY,
			Username TEXT NOT NULL UNIQUE CHECK(length(Username) >= 3 AND length(Username <= 16)),
			name TEXT,
			surname TEXT
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		sqlStmt = `CREATE TABLE Photo (
			PID INTEGER NOT NULL PRIMARY KEY,
			UID INTEGER NOT NULL,
			image BLOB,
			date TEXT,
			caption TEXT,
			FOREIGN KEY (UID) references Users(UID)
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		sqlStmt = `CREATE TABLE Comment (
			CID INTEGER NOT NULL PRIMARY KEY,
			PID INTEGER NOT NULL,
			UID INTEGER NOT NULL,
			text TEXT,
			date TEXT,
			FOREIGN KEY (UID) references Users(UID)
			FOREIGN KEY (PID) references Photo(PID) ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		sqlStmt = `CREATE TABLE Like (
			PID INTEGER NOT NULL,
			UID INTEGER NOT NULL,
			PRIMARY KEY (PID, UID),
			FOREIGN KEY (UID) references Users(UID)
			FOREIGN KEY (PID) references Photo(PID) ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		sqlStmt = `CREATE TABLE Bans (
			UID INTEGER NOT NULL,
			BannedID INTEGER NOT NULL CHECK(BannedID != UID),
			PRIMARY KEY (UID, BannedID),
			FOREIGN KEY (UID) references Users(UID)
			FOREIGN KEY (BannedID) references Users(UID)
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		sqlStmt = `CREATE TABLE Followings (
			UID INTEGER NOT NULL,
			FollowedID INTEGER NOT NULL,
			PRIMARY KEY (UID, followedID),
			FOREIGN KEY (UID) references Users(UID)
			FOREIGN KEY (followedID) references Users(UID)
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
