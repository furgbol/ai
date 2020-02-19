package comm

import (
	"log"
	"testing"
)

var zmqServerPub *ZMQServer
var zmqServerPair *ZMQServer

func zmqSetups() {
	zmqPub, err := NewZMQServer("tcp://*:7002", PUB)
	if err != nil {
		log.Fatal(err.Error())
	}

	zmqPair, err := NewZMQServer("tcp://*:7003", PAIR)
	if err != nil {
		log.Fatal(err.Error())
	}

	zmqServerPub = zmqPub
	zmqServerPair = zmqPair
}

func zmqTeardowns() {
	err := zmqServerPub.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = zmqServerPair.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestServer(t *testing.T) {
	zmqSetups()

	defer zmqTeardowns()

	t.Run("Publisher", func(t *testing.T) {
		msg := "testing publisher"
		n, err := zmqServerPub.Send([]byte(msg))
		if err != nil {
			t.Error(
				"[ZMQ] For Publisher\ngot:", err,
				"\nwant: nil")
		}
		if n != len(msg) {
			t.Error(
				"[ZMQ] For bytes number\ngot:", n,
				"\nwant: ", len(msg))
		}
	})
	t.Run("SendPair", func(t *testing.T) {
		t.SkipNow()

		msg := "testing sender pair"
		n, err := zmqServerPair.Send([]byte(msg))
		if err != nil {
			t.Error(
				"[ZMQ] For Sender Pair\ngot:", err,
				"\nwant: nil")
		}
		if n != len(msg) {
			t.Error(
				"[ZMQ] For bytes number\ngot:", n,
				"\nwant: ", len(msg))
		}
	})
}
