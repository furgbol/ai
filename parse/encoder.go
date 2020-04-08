package parse

// Encoder - Contains methods for serializing data.
type Encoder interface {
	Encode() ([]byte, error)
}
