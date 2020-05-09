package georesolver

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

// GeoResolver represents a geo locational resolution resource. The normal DB
// meets this interface so that mocks can be made.
type GeoResolver interface {
	City(net.IP) (*geoip2.City, error)
}
