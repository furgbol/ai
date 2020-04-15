package debug

import (
	"github.com/furgbol/ai/control/navigation"
	"github.com/furgbol/ai/model"
)

// Debug - Type that contains the positions of the robots for debugging
type Debug struct {
	InitialPoses []model.Pose
	FinalPoses   []model.Pose
	PathRobots   []navigation.Path
}
