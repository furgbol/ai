package comm

// Receiver - Contains methods for receiving data over the network
type Receiver interface {
	Receive() ([]byte, error)
	Close() error
}
