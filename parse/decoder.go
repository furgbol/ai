package parse

import "github.com/furgbol/ai/model"

// Decoder - Contains methods for deserializing data.
type Decoder interface {
	Decode() (*model.World, error)
}
