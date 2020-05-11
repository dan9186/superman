package logins

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCleanup(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Cleanup", func() {
		g.It("should cleanup only cuketest users", func() {
			mockDB, mockAsserts, _ := sqlmock.New()
			mockAsserts.ExpectExec(`DELETE FROM logins WHERE username = 'cuketest'`).
				WillReturnResult(sqlmock.NewResult(0, 1))

			err := Cleanup(mockDB)
			Expect(err).To(BeNil())
		})
	})
}
