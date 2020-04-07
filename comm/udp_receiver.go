package comm

import "net"

// UDPReceiver - Data type that contains the PacketConn interface used for the connection
type UDPReceiver struct {
	cnn net.PacketConn //UDP connection variable
}

// NewUDPReceiver create the socket at the address received as a parameter
func NewUDPReceiver(addr string) (*UDPReceiver, error) {
	cnn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, err
	}
	return &UDPReceiver{cnn: cnn}, nil
}

// Receive used to receive connection data
func (u UDPReceiver) Receive() ([]byte, error) {
	buffer := make([]byte, 1024)
	n, _, err := u.cnn.ReadFrom(buffer)
	return buffer[0:n], err
}

// Close ends the connection
func (u *UDPReceiver) Close() error {
	return u.cnn.Close()
}
