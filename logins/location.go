package logins

import (
	"math"
)

const (
	earthRadius = 3958.8 // miles
)

// Location represents a geographical location
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Radius    int     `json:"radius"`
}

// Distance calculates the distance from the calling location to the provided
// location in miles.
func (l *Location) Distance(l2 *Location) float64 {
	lat1, lon1 := l.Radians()
	lat2, lon2 := l2.Radians()

	latD := (lat2 - lat1) / 2
	latDSqrd := math.Sin(latD) * math.Sin(latD)

	lonD := (lon2 - lon1) / 2
	lonDSqrd := math.Sin(lonD) * math.Sin(lonD)

	sqrtH := math.Sqrt(latDSqrd + math.Cos(lat1)*math.Cos(lat2)*lonDSqrd)
	d := 2 * earthRadius * math.Asin(sqrtH)

	return math.Round(d*10000) / 10000
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
