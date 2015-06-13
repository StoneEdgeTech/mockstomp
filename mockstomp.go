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
	Order   int
	Headers stompngo.Headers
	Message string
}

type MockStompConnection struct {
	Messages         chan MockStompMessage
	NumMessages      int
	DisconnectCalled bool
}

func (m *MockStompConnection) Clear() {
	m.Messages = make(chan MockStompMessage, 1000)
	m.DisconnectCalled = false
}

func (m *MockStompConnection) Disconnect(stompngo.Headers) error {
	m.DisconnectCalled = true
	return nil
}

func (m MockStompConnection) Connected() bool {
	return true
}


	// initialize if chan not created yet:
	if cap(m.MessagesSent) < 1000 {
		m.MessagesSent = make(chan MockStompMessage, 1000)
	}

func (m *MockStompConnection) Send(headers stompngo.Headers, message string) error {
	// check for protocol

	// check for destination header
	if headers.Value("destination") == "" {
		return fmt.Errorf("No destination header, cannot send.")
	}

	// save for later
	msg := MockStompMessage{len(m.Messages), headers, message}
	m.Messages <- msg

	m.NumMessages++

	return nil
}
}
