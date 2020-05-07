package logins

import (
	"net"
	"time"

	"github.com/google/uuid"
)

// Event represents a singular login event for a given user at a given time and
// IP address.
type Event struct {
	Username      string    `json:"username"`
	UnixTimestamp int64     `json:"unix_timestamp"`
	ID            uuid.UUID `json:"event_uuid"`
	IPAddress     net.IP    `json:"ip_address"`
}

// Timestamp retuns the unix timestamp of the event as a golang Time object.
func (e *Event) Timestamp() *time.Time {
	t := time.Unix(e.UnixTimestamp, 0)

	return &t
}
