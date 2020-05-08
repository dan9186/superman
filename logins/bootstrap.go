package logins

import (
	"database/sql"
	"fmt"
)

// BootstrapLogins takes a database object to bootstrap the necessary tables for
// the login events. It returns any errors it encounters with the database.
func BootstrapLogins(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS logins (uuid TEXT PRIMARY KEY, username TEXT, timestamp INTEGER, ip_address TEXT)")
	if err != nil {
		return fmt.Errorf("failed to prepare login table creation: %v", err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("failed to execute login table creation: %v", err.Error())
	}

	return nil
}
