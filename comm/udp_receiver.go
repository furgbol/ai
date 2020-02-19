package comm

import "net"

// UDPReceiver - Data type that contains the PacketConn interface used for the connection
type UDPReceiver struct {
	bufferSize int64
	cnnUDP     net.PacketConn //UDP connection variable
}

// NewUDPReceiver create the socket at the address received as a parameter and return an 'instance' of UDPReceiver
func NewUDPReceiver(addr string) (*UDPReceiver, error) {
	cnn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, err
	}
	return &UDPReceiver{1024, cnn}, nil
}

// ReceiveDatagram used to receive connection data
func (r UDPReceiver) ReceiveDatagram() ([]byte, error) {
	buffer := make([]byte, r.bufferSize)
	n, _, err := r.cnnUDP.ReadFrom(buffer)
	return buffer[0:n], err
}

// Close ends the connection
func (r *UDPReceiver) Close() error {
	err := r.cnnUDP.Close()
	return err
}
