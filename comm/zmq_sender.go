package comm

import (
	"time"

	zmq "github.com/pebbe/zmq4"
)

// ZMQSender - Data type that contains the zmq.Socket interface used for the connection
type ZMQSender struct {
	socket *zmq.Socket
}

// NewZMQSender connect to the specified socket for sending data
func NewZMQSender(addr string) (*ZMQSender, error) {
	zmq.SetRetryAfterEINTR(true)

	socket, err := zmq.NewSocket(zmq.PAIR)
	if err != nil {
		return nil, err
	}

	err = socket.Connect(addr)
	if err != nil {
		return nil, err
	}

	return &ZMQSender{socket: socket}, nil
}

// Send sends data (bytes) to the configured address
func (z *ZMQSender) Send(bytes []byte) (int, error) {
	return z.socket.SendBytes(bytes, zmq.DONTWAIT)
}

// Close ends the connection
func (z *ZMQSender) Close() error {
	time.Sleep(time.Second)
	return z.socket.Close()
}
