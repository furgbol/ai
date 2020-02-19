package comm

import (
	"log"
	"testing"
)

var udpreceiver *UDPReceiver

func setup() {
	var err error
	udpreceiver, err = NewUDPReceiver(":8000")

	if err != nil {
		log.Fatal(err.Error())
	}
}

func teardown() {
	err := udpreceiver.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestReceiver(t *testing.T) {
	t.Skip()

	setup()

	defer teardown()

	t.Run("Receiver", func(t *testing.T) {
		_, err := udpreceiver.ReceiveDatagram()
		if err != nil {
			t.Error(
				"For receiver\n got:", err,
				"\nwant: nil")
		}
	})
}
