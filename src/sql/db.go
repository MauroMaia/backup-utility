package sql

import (
	"backup-utility/utils"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const create string = `
	CREATE TABLE IF NOT EXISTS servers (
		id INTEGER NOT NULL PRIMARY KEY,
		createdat DATETIME NOT NULL,
		updatedat DATETIME NOT NULL,
		name TEXT,
		ip TEXT,
		port INTEGER
	);
	CREATE TABLE IF NOT EXISTS db_version (
		version FLOAT not null 
			constraint db_version_pk primary key
	);
	insert into db_version values (1);
`

const file string = "backup-utility.db"

var db *sql.DB

func init() {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		utils.LogPanic("Can't access database: %s", err.Error())
	}
	if _, err := db.Exec(create); err != nil {
		utils.LogPanic("Can't initialize database: %s", err.Error())
	}
}
