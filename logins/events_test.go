package logins

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/dan9186/superman/georesolver"
)

func TestConfig(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Login Events", func() {
		g.Describe("JSON", func() {
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

		g.Describe("Database", func() {
			var mockDB *sql.DB
			var mockAsserts sqlmock.Sqlmock

			g.BeforeEach(func() {
				mockDB, mockAsserts, _ = sqlmock.New()
			})

			g.It("should store a login event", func() {
				mockAsserts.ExpectExec(`INSERT INTO logins \(uuid, username, timestamp, ip_address\) VALUES(.*)`).
					WithArgs("85ad929a-db03-4bf4-9541-8f728fa12e42", "bob", 1514764800, "206.81.252.6").
					WillReturnResult(sqlmock.NewResult(0, 1))

				var e Event
				json.Unmarshal([]byte(validLoginEvent), &e)

				err := e.Store(mockDB)
				Expect(err).To(BeNil())
			})

			g.It("should return an error if the db returns an error", func() {
				mockAsserts.ExpectExec(`INSERT INTO logins \(uuid, username, timestamp, ip_address\) VALUES(.*)`).
					WithArgs("85ad929a-db03-4bf4-9541-8f728fa12e42", "bob", 1514764800, "206.81.252.6").
					WillReturnError(fmt.Errorf("some bad thing with the DB"))

				var e Event
				json.Unmarshal([]byte(validLoginEvent), &e)

				err := e.Store(mockDB)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("some bad thing with the DB"))
			})
		})

		g.Describe("Geo Location", func() {
			g.It("should resolve an events location", func() {
				mock := georesolver.MockGeoDB{
					ExpectedIP: net.ParseIP("206.81.252.6"),
					Radius:     20,
					Latitude:   42.4242,
					Longitude:  42.4242,
				}

				var e Event
				json.Unmarshal([]byte(validLoginEvent), &e)

				c, err := e.ResolveLocation(mock)
				Expect(err).To(BeNil())
				Expect(c.Latitude).To(Equal(42.4242))
				Expect(c.Longitude).To(Equal(42.4242))
				Expect(c.Radius).To(Equal(20))
			})
		})
	})
}

const (
	validLoginEvent = `{"username": "bob", "unix_timestamp": 1514764800, "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e42", "ip_address": "206.81.252.6"}`
)
