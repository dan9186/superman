package logins

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Login Events", func() {
		g.It("should parse a login event", func() {
			var e Event
			err := json.Unmarshal([]byte(validLoginEvent), &e)
			Expect(err).To(BeNil())

			Expect(e.Username).To(Equal("bob"))
			Expect(e.ID.String()).To(Equal("85ad929a-db03-4bf4-9541-8f728fa12e42"))
			Expect(e.IPAddress.String()).To(Equal("206.81.252.6"))

			t := e.Timestamp()
			Expect(t.Year()).To(Equal(2018))
			Expect(t.Month()).To(Equal(time.January))
			Expect(t.Day()).To(Equal(1))
			Expect(t.Hour()).To(Equal(0))
			Expect(t.Minute()).To(Equal(0))
		})
	})
}

const (
	validLoginEvent = `{"username": "bob", "unix_timestamp": 1514764800, "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e42", "ip_address": "206.81.252.6"}`
)
