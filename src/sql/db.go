package sql

import (
	"backup-utility/utils"
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

const file string = "backup-utility.db"

var db *sql.DB

func init() {

	utils.LogDebug("Initialize database")

	var err error
	db, err = sql.Open("sqlite3", file)
	if err != nil {
		utils.LogPanic("Can't access database: %s", err.Error())
	}

	lastMigrationVersion, _ := GetLasMigration()
	// apply migrations
	for i, migration := range migrations {
		if i <= lastMigrationVersion {
			// skip version lower that last applied
			continue
		}
		if _, err := db.Exec(migration); err != nil {
			utils.LogWarning("Failed to apply migration %d. Reason: %s", i, err.Error())
		}

	}

}

func GetLasMigration() (int, error) {

	strSql := `SELECT max(version) from db_version;`

	result, err := db.Query(strSql)
	if err != nil {
		return -1, errors.New("Failed to execute query to read migration version. Reason: " + err.Error())
	}

	var lastMigrationVersion int
	result.Next()
	if err = result.Scan(&lastMigrationVersion); err != nil {
		return -1, errors.New("No max migration version found. Reason: " + err.Error())
	}

	return lastMigrationVersion, nil
}

var migrations = [...]string{
	`
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
	INSERT INTO db_version values (1);
`,
}
