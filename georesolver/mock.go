package georesolver

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
)

// MockGeoDB is a basic representation of all the data necessary to handle a
// mock lookup and response to a geo locational db.
type MockGeoDB struct {
	expectations []*Expectation
}

// NewMock returns a properly initialized MockGeoDB.
func NewMock() *MockGeoDB {
	db := &MockGeoDB{
		expectations: make([]*Expectation, 0),
	}

	return db
}

// City meets the GeoResolver interface as a lookup for a City by IP Address
func (mgdb *MockGeoDB) City(ip net.IP) (*geoip2.City, error) {
	exp, r := mgdb.expectations[0], mgdb.expectations[1:]
	mgdb.expectations = r

	if exp.ip.String() != ip.String() {
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
			AccuracyRadius: uint16(exp.radius),
			Latitude:       exp.latitude,
			Longitude:      exp.longitude,
		},
	}

	return c, nil
}

// ExpectIP takes an IP address to add to the expectation stack and returns the
// expectation so the expected returns can be added to the expectation.
func (mgdb *MockGeoDB) ExpectIP(ip net.IP) *Expectation {
	e := &Expectation{
		ip: ip,
	}

	mgdb.expectations = append(mgdb.expectations, e)
	return e
}

// Expectation is a matching of an IP address to the return values it should
// give.
type Expectation struct {
	ip        net.IP
	latitude  float64
	longitude float64
	radius    int
}

// WillReturnLocation collects the additional information for an appropriate
// return on an IP lookup.
func (e *Expectation) WillReturnLocation(lat, lon float64, rad int) {
	e.latitude = lat
	e.longitude = lon
	e.radius = rad
}
