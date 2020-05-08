package logins

import (
	"database/sql"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
)

const (
	insertLoginEvent = `INSERT INTO logins (uuid, username, timestamp, ip_address) VALUES($1, $2, $3, $4)`
)

// Event represents a singular login event for a given user at a given time and
// IP address.
type Event struct {
	Username      string    `json:"username"`
	UnixTimestamp int64     `json:"unix_timestamp"`
	ID            uuid.UUID `json:"event_uuid"`
	IPAddress     net.IP    `json:"ip_address"`
}

// Analyze looks up comparative details of a login event and provides an
// Analysis of the comparative details.
func (e *Event) Analyze() (*Analysis, error) {
	loc, err := e.ResolveLocation()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve location: %v", err.Error())
	}

	a := &Analysis{
		CurrentLocation: loc,
	}

	return a, nil
}

// ResolveLocation uses an event's IP address to determine a geolocation and
// returns the details as a Location object.
func (e *Event) ResolveLocation() *Location {
	l := &Location{}

	return l
}

// Timestamp retuns the unix timestamp of the event as a golang Time object. The
// time will always be returned in UTC.
func (e *Event) Timestamp() *time.Time {
	t := time.Unix(e.UnixTimestamp, 0).UTC()

	return &t
}

// Store takes a database object and stores the designated event in the
// database. It returns an errors it encounters with the database.
func (e *Event) Store(db *sql.DB) error {
	_, err := db.Exec(insertLoginEvent, e.ID, e.Username, e.UnixTimestamp, e.IPAddress.String())
	if err != nil {
		return fmt.Errorf("logins: failed inserting login event: %v", err.Error())
	}

	return nil
}
