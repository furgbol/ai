package comm

import (
	"log"
	"testing"
)

var udpSender *UDPSender

func setups() {
	udpsend, err := NewUDPSender(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}

	udpSender = udpsend
}

func teardowns() {
	err := udpSender.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestSender(t *testing.T) {
	setups()

	defer teardowns()

	t.Run("Sender", func(t *testing.T) {
		msg := "testing sender"
		n, err := udpSender.WriteDatagram([]byte(msg))
		if err != nil {
			t.Error(
				"[UDP] For WriteDatagram\ngot:", err,
				"\nwant: nil")
		}
		if n != len(msg) {
			t.Error(
				"[UDP] For bytes number\ngot:", n,
				"\nwant: ", len(msg))
		}
	})
}
