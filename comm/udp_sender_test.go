package comm

import (
	"log"
	"testing"
)

var udpsender *UDPSender

func setups() {
	var err error
	udpsender, err = NewUDPSender(":8000")

	if err != nil {
		log.Fatal(err.Error())
	}
}

func teardowns() {
	err := udpsender.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestSender(t *testing.T) {
	setups()

	defer teardowns()

	t.Run("Sender", func(t *testing.T) {
		msg := "testing sender"
		n, err := udpsender.WriteDatagram([]byte(msg))
		if err != nil {
			t.Error(
				"For WriteDatagram\n got:", err,
				"\nwant: nil")
		}
		if n != len(msg) {
			t.Error(
				"For bytes number\n got:", n,
				"\nwant: ", len(msg))
		}
	})
}
