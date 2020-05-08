package logins

import (
	"net"
)

// Analysis represents a resulting analysis of a singular login event. It
// contains the current geo locational resolution to the IP address of a login
// event, comparative details of the preceding login event if it exists,
// comparative details of the subsequent login event if it exists, and whether
// or not the comparison of the events result in suspicious activity.
type Analysis struct {
	CurrentLocation             Location  `json:"currentGeo"`
	PrecedingAccess             *IPAccess `json:"precedingIpAccess,omitempty"`
	SuspiciousPrecedingAccess   bool      `json:"travelToCurrentGeoSuspicious"`
	SubsequentAccess            *IPAccess `json:"subsequentIpAccess,omitempty"`
	SuspiciiousSubsequentAccess bool      `json:"travelFromCurrentGeoSuspicious"`
}

// Location represents a geographical location
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Radius    int     `json:"radius"`
}

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
