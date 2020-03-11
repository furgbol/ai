package navigation

import "github.com/furgbol/ai/model"

// PathPlanner - interface to path planning
type PathPlanner interface {
	PlanPath(model.Pose, model.Pose) Path
}
