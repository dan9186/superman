package georesolver

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
)

// GeoResolver represents a geo locational resolution resource. The normal DB
// meets this interface so that mocks can be made.
type GeoResolver interface {
	City(net.IP) (*geoip2.City, error)
}

// MockGeoDB is a basic representation of all the data necessary to handle a
// mock lookup and response to a geo locational db.
type MockGeoDB struct {
	ExpectedIP net.IP
	Latitude   float64
	Longitude  float64
	Radius     int
}

// City meets the GeoResolver interface as a lookup for a City by IP Address
func (mgdb MockGeoDB) City(ip net.IP) (*geoip2.City, error) {
	if mgdb.ExpectedIP.String() != ip.String() {
		return nil, fmt.Errorf("incorrect mock IP requested")
	}

	c := &geoip2.City{
		Location: struct {
			AccuracyRadius uint16  `maxminddb:"accuracy_radius"`
			Latitude       float64 `maxminddb:"latitude"`
			Longitude      float64 `maxminddb:"longitude"`
			MetroCode      uint    `maxminddb:"metro_code"`
			TimeZone       string  `maxminddb:"time_zone"`
		}{
			AccuracyRadius: uint16(mgdb.Radius),
			Latitude:       mgdb.Latitude,
			Longitude:      mgdb.Longitude,
		},
	}

	return c, nil
}
