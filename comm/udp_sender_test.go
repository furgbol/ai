package comm

import (
	"log"
	"testing"
)

var udpSender Sender

func setups() {
	socket, err := NewUDPSender(":8001")
	if err != nil {
		log.Fatal(err.Error())
	}
	udpSender = socket
}

func teardowns() {
	err := udpSender.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestUDPSender(t *testing.T) {
	setups()

	defer teardowns()

	t.Run("Sender", func(t *testing.T) {
		msg := "testing sender"
		n, err := udpSender.Send([]byte(msg))
		if err != nil {
			t.Error(
				"[UDP] For Send\ngot:", err,
				"\nwant: nil")
		}
		if n != len(msg) {
			t.Error(
				"[UDP] For Send: bytes number\ngot:", n,
				"\nwant: ", len(msg))
		}
	})
}
