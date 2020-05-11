package logins

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestLocation(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Location", func() {
		g.It("should convert a location to radians", func() {
			l := &Location{
				Latitude:  45.345,
				Longitude: -122.6533,
			}

			latRad, lonRad := l.radians()
			Expect(latRad).To(Equal(0.7914))
			Expect(lonRad).To(Equal(-2.1407))
		})

		g.It("should calculate the distance between two locations", func() {
			// West Linn, OR
			l1 := &Location{
				Latitude:  45.345,
				Longitude: -122.6533,
			}

			// San Diego Zoo
			l2 := &Location{
				Latitude:  32.7352,
				Longitude: -117.149,
			}

			d := l1.Distance(l2)
			Expect(d).To(Equal(919.489))

			// Bend, OR
			l1 = &Location{
				Latitude:  44.0591,
				Longitude: -121.3057,
			}

			// Santa Fe, Argentina
			l2 = &Location{
				Latitude:  -31.6796,
				Longitude: -60.6422,
			}

			d = l1.Distance(l2)
			Expect(d).To(Equal(6478.3509))

			// Melbourne, Australia
			l1 = &Location{
				Latitude:  -37.8086,
				Longitude: 144.9166,
			}

			// London, UK
			l2 = &Location{
				Latitude:  51.4884,
				Longitude: -0.116,
			}

			d = l1.Distance(l2)
			Expect(d).To(Equal(10501.3591))
		})
	})
}
