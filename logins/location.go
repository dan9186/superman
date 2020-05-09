package logins

import (
	"math"
)

// Location represents a geographical location
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Radius    int     `json:"radius"`
}

// Radians converts the latitude and longitude of a location to radians. It
// returns the values in the order of latitude as radians, and longitude as
// radians.
func (l *Location) Radians() (float64, float64) {
	lat := l.Latitude * math.Pi / 180
	lat = math.Round(lat*10000) / 10000

	lon := l.Longitude * math.Pi / 180
	lon = math.Round(lon*10000) / 10000

	return lat, lon
}
