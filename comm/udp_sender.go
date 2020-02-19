package comm

import "net"

// UDPSender - Data type that contains the Conn interface used for the connection
type UDPSender struct {
	cnn net.Conn
}

// NewUDPSender create the socket at the address received as a parameter and return an 'instance' of UDPSender
func NewUDPSender(addr string) (*UDPSender, error) {
	cnn, err := net.Dial("udp", addr)
	if err != nil {
		return nil, err
	}
	return &UDPSender{cnn}, nil
}

// WriteDatagram sends data (bytes) to the configured address
func (s UDPSender) WriteDatagram(bytes []byte) (int, error) {
	nbytes, err := s.cnn.Write(bytes)
	return nbytes, err
}

// Close ends the connection
func (s *UDPSender) Close() error {
	err := s.cnn.Close()
	return err
}
