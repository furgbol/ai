package navigation

import "github.com/furgbol/ai/model"

// Path is the data type to store a path
type Path []model.Position2D

// NewPath creates a path instance
func NewPath(numberOfPoints int) Path {
	path := make(Path, numberOfPoints, numberOfPoints)
	return path
}
