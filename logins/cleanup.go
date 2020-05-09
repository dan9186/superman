package logins

import (
	"database/sql"
	"fmt"
)

const (
	deleteTestData = `DELETE FROM logins WHERE username = 'cuketest'`
)

// Cleanup takes a database object to remove test data from. It is hard coded to
// delete events associated to test data so no important data can be touched. It
// returns any errors it encounters with the database.
func Cleanup(db *sql.DB) error {
	_, err := db.Exec(deleteTestData)
	if err != nil {
		return fmt.Errorf("failed to execute test data delete query: %v", err.Error())
	}

	return nil
}
