package comm

import (
	"time"

	zmq "github.com/pebbe/zmq4"
)

// ZMQReceiver - Data type that contains the zmq.Socket interface used for the connection
type ZMQReceiver struct {
	socket *zmq.Socket
}

// NewZMQReceiver connects to the specified socket for receiving data
func NewZMQReceiver(addr string) (*ZMQReceiver, error) {
	zmq.SetRetryAfterEINTR(true)

	socket, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		return nil, err
	}

	err = socket.SetSubscribe("")
	if err != nil {
		return nil, err
	}

	socket.SetConflate(true)
	err = socket.Connect(addr)
	if err != nil {
		return nil, err
	}
	return &ZMQReceiver{socket: socket}, nil
}

// Receive used to receive connection data
func (z *ZMQReceiver) Receive() ([]byte, error) {
	bytes, err := z.socket.RecvBytes(0)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Close ends the connection
func (z *ZMQReceiver) Close() error {
	time.Sleep(time.Second)
	return z.socket.Close()
}
