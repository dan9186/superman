package logins

import (
	"database/sql"
)

// BootstrapLogins takes a database object to bootstrap the necessary tables for
// the login events. It returns any errors it encounters with the database.
func BootstrapLogins(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS logins (uuid TEXT PRIMARY KEY, username TEXT, timestamp INTEGER, ip_address TEXT)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
