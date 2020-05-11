package logins

import (
	"database/sql"
	"fmt"
)

const (
	createLoginsTable = `CREATE TABLE IF NOT EXISTS logins (uuid TEXT PRIMARY KEY, username TEXT, timestamp INTEGER, ip_address TEXT)`
)

// BootstrapLogins takes a database object to bootstrap the necessary tables for
// the login events. It returns any errors it encounters with the database.
func BootstrapLogins(db *sql.DB) error {
	stmt, err := db.Prepare(createLoginsTable)
	if err != nil {
		return fmt.Errorf("failed to prepare login table creation: %v", err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("failed to execute login table creation: %v", err.Error())
	}

	return nil
}
