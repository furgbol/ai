package comm

import (
	"log"
	"testing"
)

var zmqClientSub *ZMQClient
var zmqClientPair *ZMQClient

func zmqSetup() {
	zmqSub, err := NewZMQClient("tcp://127.0.0.1:7000", SUB)
	if err != nil {
		log.Fatal(err.Error())
	}

	zmqPair, err := NewZMQClient("tcp://127.0.0.1:7001", PAIR)
	if err != nil {
		log.Fatal(err.Error())
	}

	zmqClientSub = zmqSub
	zmqClientPair = zmqPair
}

func zmqTeardown() {
	err := zmqClientSub.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = zmqClientPair.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestClient(t *testing.T) {
	t.Skip()

	zmqSetup()

	defer zmqTeardown()

	t.Run("Subscribe", func(t *testing.T) {
		_, err := zmqClientSub.Receive()
		if err != nil {
			t.Error(
				"[ZMQ] For subscribe receiver\ngot:", err,
				"\nwant: nil")
		}
	})
	t.Run("RecPair", func(t *testing.T) {
		_, err := zmqClientPair.Receive()
		if err != nil {
			t.Error(
				"[ZMQ] For pair receiver\ngot:", err,
				"\nwant: nil")
		}
	})
}
