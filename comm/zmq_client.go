package comm

import (
	zmq "github.com/pebbe/zmq4"
)

// ZMQClient - Data type that contains the zmq.Socket interface used for the connection
type ZMQClient struct {
	socket *zmq.Socket
}

// NewZMQClient connects to the specified socket and returns an 'instance' of the ZMQClient
func NewZMQClient(addr string, tp zmq.Type) (*ZMQClient, error) {
	zmq.SetRetryAfterEINTR(true)

	socket, err := zmq.NewSocket(tp)
	if err != nil {
		return nil, err
	}

	if tp == zmq.SUB {
		err = socket.SetSubscribe("")
		if err != nil {
			return nil, err
		}
		socket.SetConflate(true)
	}

	err = socket.Connect(addr)
	if err != nil {
		return nil, err
	}

	return &ZMQClient{socket: socket}, nil
}

// Send sends data (bytes) to the configured address
func (z *ZMQClient) Send(data []byte) (int, error) {
	return z.socket.SendBytes(data, zmq.DONTWAIT)
}

// Receive used to receive connection data
func (z *ZMQClient) Receive() ([]byte, error) {
	bytes, err := z.socket.RecvBytes(0)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Close ends the connection
func (z *ZMQClient) Close() error {
	return z.socket.Close()
}
