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
// provided details. The radius of the IPAccess and the radius of the provided
// location are used to reduce to reduce the total possible distance traveled
// before calculating the speed. This is to reduce the possible false negatives
// of a person who could have legitimately been on the inside edges of the two
// locations. The resulting speed is updated back to the IPAccess speed field
// and will be in miles per hour.
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
