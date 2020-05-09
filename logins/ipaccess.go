package logins

import (
	"math"
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

// CalculateSpeed takes a location and corresponding timestamp for the location
// to calculate the speed to travel between the location of the IPAccess and the
// provided details. It updates the IPAccess record with the result in miles per
// hour.
func (ipa *IPAccess) CalculateSpeed(l *Location, timestamp int64) {
	l2 := &Location{
		Latitude:  ipa.Latitude,
		Longitude: ipa.Longitude,
	}

	d := l.Distance(l2)
	d = math.Max(d-float64(ipa.Radius)-float64(l.Radius), 0.0)

	t := math.Abs(float64(ipa.UnixTimestamp-timestamp)) / 3600.0

	ipa.Speed = int(d / t)
}
