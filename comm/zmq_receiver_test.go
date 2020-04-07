package comm

import (
	"log"
	"testing"
)

var zmqReceiver Receiver

func zmqSetups() {
	socket, err := NewZMQReceiver("tcp://127.0.0.1:7000")
	if err != nil {
		log.Fatal(err.Error())
	}
	zmqReceiver = socket
}

func zmqTeardowns() {
	err := zmqReceiver.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestZMQReceiver(t *testing.T) {
	t.Skip()

	zmqSetups()

	defer zmqTeardowns()

	t.Run("Subscribe", func(t *testing.T) {
		_, err := zmqReceiver.Receive()
		if err != nil {
			t.Error(
				"[ZMQ] For subscribe receiver\ngot:", err,
				"\nwant: nil")
		}
	})
}
