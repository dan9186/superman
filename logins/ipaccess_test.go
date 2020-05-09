package logins

import (
	"testing"
	"time"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestIPAccess(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("IP Access", func() {
		g.It("should calculate the speed betweeen two locations", func() {
			t := time.Now().Unix()

			// West Linn, OR
			l := &Location{
				Latitude:  45.345,
				Longitude: -122.6533,
			}

			// Bend, OR
			ipa := &IPAccess{
				Latitude:      44.0591,
				Longitude:     -121.3057,
				UnixTimestamp: t - 3600,
			}

			ipa.CalculateSpeed(l, t)
			Expect(ipa.Speed).To(Equal(110))
		})

		g.It("should use the radius to reduce the possible distance to travel", func() {
			t := time.Now().Unix()

			// West Linn, OR
			l := &Location{
				Latitude:  45.345,
				Longitude: -122.6533,
				Radius:    20,
			}

			// Bend, OR
			ipa := &IPAccess{
				Latitude:      44.0591,
				Longitude:     -121.3057,
				Radius:        20,
				UnixTimestamp: t - 3600,
			}

			ipa.CalculateSpeed(l, t)
			Expect(ipa.Speed).To(Equal(70))
		})

		g.It("should not go below 0 distance", func() {
			t := time.Now().Unix()

			// West Linn, OR
			l := &Location{
				Latitude:  45.345,
				Longitude: -122.6533,
				Radius:    70,
			}

			// Bend, OR
			ipa := &IPAccess{
				Latitude:      44.0591,
				Longitude:     -121.3057,
				Radius:        80,
				UnixTimestamp: t - 3600,
			}

			ipa.CalculateSpeed(l, t)
			Expect(ipa.Speed).To(Equal(0))
		})
	})
}
