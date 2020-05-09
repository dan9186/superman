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

func TestEvents(t *testing.T) {
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
				mock := georesolver.NewMock()

				mock.ExpectIP(net.ParseIP("206.81.252.6")).
					WillReturnLocation(42.4242, 42.4242, 20)

				var e Event
				json.Unmarshal([]byte(validLoginEvent), &e)

				c, err := e.ResolveLocation(mock)
				Expect(err).To(BeNil())
				Expect(c.Latitude).To(Equal(42.4242))
				Expect(c.Longitude).To(Equal(42.4242))
				Expect(c.Radius).To(Equal(20))
			})
		})

		g.Describe("Analysis", func() {
			g.It("should analyze a single event", func() {
				mockDB, mockAsserts, _ := sqlmock.New()

				mockAsserts.ExpectQuery(`SELECT uuid, timestamp, ip_address FROM logins WHERE username = (.*) AND timestamp < (.*) ORDER BY timestamp DESC LIMIT 1`).
					WithArgs("bob", 1514764800).
					WillReturnRows(
						sqlmock.NewRows([]string{"uuid", "timestamp", "ip_address"}).
							AddRow("4e837b27-2005-4dbb-8f7e-f32c6c2af699", "1514764734", "91.207.175.104"),
					)

				mockAsserts.ExpectQuery(`SELECT uuid, timestamp, ip_address FROM logins WHERE username = (.*) AND timestamp > (.*) ORDER BY timestamp ASC LIMIT 1`).
					WithArgs("bob", 1514764800).
					WillReturnRows(
						sqlmock.NewRows([]string{"uuid", "timestamp", "ip_address"}).
							AddRow("d99df0fd-3a77-4662-910f-9e4f8ecfe25b", "1588930045", "24.242.71.20"),
					)

				mockGeoDB := georesolver.NewMock()

				mockGeoDB.ExpectIP(net.ParseIP("206.81.252.6")).
					WillReturnLocation(42.4242, 42.4242, 20)

				mockGeoDB.ExpectIP(net.ParseIP("91.207.175.104")).
					WillReturnLocation(32.3242, 32.3242, 20)

				mockGeoDB.ExpectIP(net.ParseIP("24.242.71.20")).
					WillReturnLocation(22.3242, 22.3242, 20)

				var e Event
				json.Unmarshal([]byte(validLoginEvent), &e)

				a, err := e.Analyze(mockDB, mockGeoDB)
				Expect(err).To(BeNil())

				Expect(a.CurrentLocation.Latitude).To(Equal(42.4242))
				Expect(a.CurrentLocation.Longitude).To(Equal(42.4242))
				Expect(a.CurrentLocation.Radius).To(Equal(20))
			})
		})
	})
}

const (
	validLoginEvent = `{"username": "bob", "unix_timestamp": 1514764800, "event_uuid": "85ad929a-db03-4bf4-9541-8f728fa12e42", "ip_address": "206.81.252.6"}`
)
