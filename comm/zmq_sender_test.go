package comm

import (
	"log"
	"testing"
)

var zmqSender Sender

func zmqSetup() {
	socket, err := NewZMQSender("tcp://127.0.0.1:7001")
	if err != nil {
		log.Fatal(err.Error())
	}
	zmqSender = socket
}

func zmqTeardown() {
	err := zmqSender.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestZMQSender(t *testing.T) {
	t.Skip()

	zmqSetup()

	defer zmqTeardown()

	t.Run("ZMQ Sender", func(t *testing.T) {
		msg := "testing sender"
		n, err := zmqSender.Send([]byte(msg))
		if err != nil {
			t.Error(
				"[ZMQ] For Send\ngot:", err,
				"\nwant: nil")
		}
		if n != len(msg) {
			t.Error(
				"[ZMQ] For Send: bytes number\ngot:", n,
				"\nwant: ", len(msg))
		}
	})
}
