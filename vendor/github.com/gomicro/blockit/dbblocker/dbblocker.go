package dbblocker

import (
	"database/sql"
	"time"
)

// Blocker represents a SQL database connection to monitor for connectivity
// and block until a connection is established.
type Blocker struct {
	db *sql.DB
}

// New takes a SQL database object and returns a newly instantiated Blocker
func New(db *sql.DB) *Blocker {
	return &Blocker{db: db}
}

// Blockit meets the blocker interface. It returns a read only channel that will
// receive true when the database is connected.
func (d *Blocker) Blockit() <-chan bool {
	connected := make(chan bool)

	go func() {
		for {
			err := d.db.Ping()
			if err != nil {
				<-time.After(1 * time.Second)
				continue
			}

			connected <- true
			close(connected)

			break
		}
	}()

	return connected
}
