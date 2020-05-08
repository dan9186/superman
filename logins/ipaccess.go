package logins

import (
	"net"
)

// IPAccess represents the derived details of an comparison between two login
// events.
type IPAccess struct {
	IPAddress     net.IP  `json:"ip_address"`
	Speed         int     `json:"speed"`
	Latitude      float64 `json:"lat"`
	Longitude     float64 `json:"lon"`
	Radius        int     `json:"radius"`
	UnixTimestamp int64   `json:"timestamp"`
}
