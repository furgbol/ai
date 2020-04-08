package comm

// Sender - Contains methods for sending data over the network
type Sender interface {
	Send(bytes []byte) (int, error)
	Close() error
}
