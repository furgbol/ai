package comm

import (
	"log"
	"testing"
)

var udpReceiver *UDPReceiver

func setup() {
	udprec, err := NewUDPReceiver(":8000")

	if err != nil {
		log.Fatal(err.Error())
	}

	udpReceiver = udprec
}

func teardown() {
	err := udpReceiver.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestReceiver(t *testing.T) {
	t.Skip()

	setup()

	defer teardown()

	t.Run("Receiver", func(t *testing.T) {
		_, err := udpReceiver.ReceiveDatagram()
		if err != nil {
			t.Error(
				"For receiver\n got:", err,
				"\nwant: nil")
		}
	})
}
