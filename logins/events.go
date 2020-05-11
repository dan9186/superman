package logins

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/google/uuid"

	"github.com/dan9186/superman/georesolver"
)

const (
	insertLoginEvent      = `INSERT INTO logins (uuid, username, timestamp, ip_address) VALUES($1, $2, $3, $4)`
	selectPrecedingEvent  = `SELECT uuid, timestamp, ip_address FROM logins WHERE username = $1 AND timestamp < $2 ORDER BY timestamp DESC LIMIT 1`
	selectSubsequentEvent = `SELECT uuid, timestamp, ip_address FROM logins WHERE username = $1 AND timestamp > $2 ORDER BY timestamp ASC LIMIT 1`
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
func (e *Event) Analyze(db *sql.DB, geodb georesolver.GeoResolver) (*Analysis, error) {
	loc, err := e.ResolveLocation(geodb)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve location: %v", err.Error())
	}

	analysis := &Analysis{
		CurrentLocation: loc,
	}

	pe, err := e.getPreceding(db)
	if err != nil {
		return nil, fmt.Errorf("failed to get preceding event info: %v", err.Error())
	}

	if pe != nil {
		pAccess, err := pe.toIPAccess(geodb)
		if err != nil {
			return nil, fmt.Errorf("failed to convert preceding event to ip access: %v", err.Error())
		}

		pAccess.CalculateSpeed(loc, e.UnixTimestamp)

		analysis.PrecedingAccess = pAccess
		analysis.SuspiciousPrecedingAccess = pAccess.Speed > 500
	}

	se, err := e.getSubsequent(db)
	if err != nil {
		return nil, fmt.Errorf("failed to get subsequent event info: %v", err.Error())
	}

	if se != nil {
		sAccess, err := se.toIPAccess(geodb)
		if err != nil {
			return nil, fmt.Errorf("failed to convert subsequent event to ip access: %v", err.Error())
		}

		sAccess.CalculateSpeed(loc, e.UnixTimestamp)

		analysis.SubsequentAccess = sAccess
		analysis.SuspiciiousSubsequentAccess = sAccess.Speed > 500
	}

	return analysis, nil
}

func (e *Event) getPreceding(db *sql.DB) (*Event, error) {
	return e.getEvent(selectPrecedingEvent, db)
}

func (e *Event) getSubsequent(db *sql.DB) (*Event, error) {
	return e.getEvent(selectSubsequentEvent, db)
}

func (e *Event) getEvent(query string, db *sql.DB) (*Event, error) {
	var id uuid.UUID
	var ipStr string
	var unix int64

	err := db.QueryRow(query, e.Username, e.UnixTimestamp).Scan(&id, &unix, &ipStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to query for subsequent event: %v", err.Error())
	}

	se := &Event{
		Username:      e.Username,
		ID:            id,
		UnixTimestamp: unix,
		IPAddress:     net.ParseIP(ipStr),
	}

	return se, nil
}

// ResolveLocation uses an event's IP address to determine a geolocation and
// returns the details as a Location object.
func (e *Event) ResolveLocation(geodb georesolver.GeoResolver) (*Location, error) {
	r, err := geodb.City(e.IPAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup city for IP: %v", err.Error())
	}

	l := &Location{
		Latitude:  r.Location.Latitude,
		Longitude: r.Location.Longitude,
		Radius:    int(r.Location.AccuracyRadius),
	}

	return l, nil
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

func (e *Event) toIPAccess(geodb georesolver.GeoResolver) (*IPAccess, error) {
	loc, err := e.ResolveLocation(geodb)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve location for preceding event: %v", err.Error())
	}

	a := &IPAccess{
		IPAddress:     e.IPAddress,
		UnixTimestamp: e.UnixTimestamp,
		Latitude:      loc.Latitude,
		Longitude:     loc.Longitude,
		Radius:        loc.Radius,
	}

	return a, nil
}
