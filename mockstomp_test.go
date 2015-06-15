package mockstomp

import (
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
		var stompConnection = New()
		var message string

		g.BeforeEach(func() {
			stompConnection.Clear()

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

		g.It("should be able to get messages back afterwards", func() {
			// expected behavior adding to chan
			for i := 0; i < 1000; i++ {
				Expect(stompConnection.Send(headers, message)).To(BeNil())
			}

			// should be messages in the chan
			Expect(len(stompConnection.Messages)).To(Equal(1000))

			// pop the messages off of the chan and verify
			for i := 0; i < 1000; i++ {
				msg := <-stompConnection.Messages
				expectedMessage := &MockStompMessage{
					Order: i,
					Headers: []string{
						"persistent",
						"true",
						"destination",
						"/queue/dedupe",
						"asin",
						"b000159fau",
						"market",
						"us",
						"condition",
						"new",
						"triggered-at",
						"1252",
						"special_distribution",
						"true",
					},
					Message: "Foo Bar",
				}

				Expect(msg).To(Equal(*expectedMessage))
			}
		})

		g.It("should allow for a disconnect request", func() {
			err := stompConnection.Disconnect(stompngo.Headers{})
			Expect(err).NotTo(HaveOccurred())
			Expect(stompConnection.DisconnectCalled).To(BeTrue())
		})

		g.It("should allow a subscription", func() {
			sub, err := stompConnection.Subscribe(stompngo.Headers{})
			Expect(err).NotTo(HaveOccurred())
			Expect(stompConnection.Subscription).ToNot(BeNil())
			Expect(sub).To(Equal(stompConnection.Subscription))

			msg := stompngo.MessageData{
				Message: stompngo.Message{
					Body: []uint8(message),
				},
			}
			stompConnection.PutToSubscribe(msg)
			outMsg := <-sub
			Expect(outMsg.Message.BodyString()).To(Equal(message))
			Expect(string(outMsg.Message.Body)).To(Equal(message))
		})

		g.It("should allow an unsubscribe", func() {
			err := stompConnection.Unsubscribe(stompngo.Headers{})
			Expect(err).NotTo(HaveOccurred())
			Expect(stompConnection.Subscription).To(BeNil())
		})
	})
}
