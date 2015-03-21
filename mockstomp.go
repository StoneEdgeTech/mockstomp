package mockstomp

/*
 * provides a MockStompConnection struct with accompaning functions to implement the
 * stomp interface that will record what gets called. See mock_test.go
 * for usage examples, but in general, it looks like this:
 * > mockStompConnectionInstance.Send(headers,message)
 * >
 */

import (
	"fmt"
	"github.com/gmallard/stompngo"
)

type MockStompMessage struct {
	order   int
	headers stompngo.Headers
	message string
}

type MockStompConnection struct {
	messagesSent chan MockStompMessage
}

func (m *MockStompConnection) Init() {
	m.messagesSent = make(chan MockStompMessage, 1000)
}

func (m *MockStompConnection) Send(headers stompngo.Headers, message string) (e error) {

	// check for protocol

	// check for destination header
	if headers.Value("destination") == "" {
		return fmt.Errorf("No destination header, cannot send.")
	}

	// save for later
	sentMessage := MockStompMessage{len(m.messagesSent), headers, message}
	m.messagesSent <- sentMessage

	return e
}
