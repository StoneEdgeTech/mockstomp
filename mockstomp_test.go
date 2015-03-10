package mockstomp

import (
	//"fmt"
	. "github.com/franela/goblin"
	"github.com/gmallard/stompngo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestPopulator(t *testing.T) {
	
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("stomp connection mock", func() {

		var headers stompngo.Headers
		var stompConnection = &MockStompConnection{}
		var message string

		g.BeforeEach(func() {

			// broker headers
			headers = stompngo.Headers{
				"accept-version", "1.1",
				"login", "admin",
				"passcode", "1234",
				"host", "localhost",
			}

			// message headers
			headers = stompngo.Headers{
				"persistent", "true",
				"destination", "/queue/dedupe",
				"asin", "b000159fau",
				"market", "us",
				"condition", "new",
				"triggered-at", "1252",
				"special_distribution", "true",
			}
			message = "Foo Bar"
		})

		g.It("should be successful with all headers present", func() {
			Expect(stompConnection.Send(headers, message)).To(BeNil())
		})

		g.It("should not be successful if the destination header is blank", func() {
			headers = headers.Delete("destination")
			Expect(stompConnection.Send(headers, message)).NotTo(BeNil())
		})
	})

}
