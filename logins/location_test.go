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

			latRad, lonRad := l.Radians()
			Expect(latRad).To(Equal(0.7914))
			Expect(lonRad).To(Equal(-2.1407))
		})
	})
}
