package comm

import (
	"log"
	"testing"
)

var udpReceiver Receiver

func setup() {
	socket, err := NewUDPReceiver(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
	udpReceiver = socket
}

func teardown() {
	err := udpReceiver.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestUDPReceiver(t *testing.T) {
	t.Skip()

	setup()

	defer teardown()

	t.Run("Receiver", func(t *testing.T) {
		_, err := udpReceiver.Receive()
		if err != nil {
			t.Error(
				"[UDP] For receiver\ngot:", err,
				"\nwant: nil")
		}
	})
}
