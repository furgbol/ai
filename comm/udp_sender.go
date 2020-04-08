package comm

import "net"

// UDPSender - Data type that contains the Conn interface used for the connection
type UDPSender struct {
	cnn net.Conn
}

// NewUDPSender create the socket at the address received as a parameter
func NewUDPSender(addr string) (*UDPSender, error) {
	cnn, err := net.Dial("udp", addr)
	if err != nil {
		return nil, err
	}
	return &UDPSender{cnn: cnn}, nil
}

// Send sends data (bytes) to the configured address
func (u UDPSender) Send(bytes []byte) (int, error) {
	return u.cnn.Write(bytes)
}

// Close ends the connection
func (u *UDPSender) Close() error {
	return u.cnn.Close()
}
