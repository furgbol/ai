package comm

import (
	zmq "github.com/pebbe/zmq4"
)

// ZMQServer - Data type that contains the zmq.Socket interface used for the connection
type ZMQServer struct {
	socket *zmq.Socket
}

// NewZMQServer create the socket at the address received as a parameter and return an 'instance' of ZMQServer
func NewZMQServer(addr string, tp zmq.Type) (*ZMQServer, error) {
	zmq.SetRetryAfterEINTR(true)

	socket, err := zmq.NewSocket(tp)
	if err != nil {
		return nil, err
	}

	err = socket.Bind(addr)
	if err != nil {
		return nil, err
	}

	return &ZMQServer{socket: socket}, nil
}

// Send sends data (bytes) to the configured address
func (z *ZMQServer) Send(data []byte) (int, error) {
	return z.socket.SendBytes(data, zmq.DONTWAIT)
}

// Receive used to receive connection data
func (z *ZMQServer) Receive() ([]byte, error) {
	bytes, err := z.socket.RecvBytes(0)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Close ends the connection
func (z *ZMQServer) Close() error {
	return z.socket.Close()
}
