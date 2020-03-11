package navigation

import (
	"fmt"
	"math"

	"github.com/furgbol/ai/model"
)

// CasteljauPathPlanner - This type implements the PathPlanner interface. It plans the path based on Casteljau's algorithm that uses Bézier curves.
type CasteljauPathPlanner struct {
	NumberOfPathPoints int
	NumberOfUsedPoints int
	DistanceFactor     float64
}

// NewCasteljauPathPlanner creates an instance of a Casteljau Path Planner
func NewCasteljauPathPlanner(numberOfPathPoints, numberOfUsedPoints int, distanceFactor float64) *CasteljauPathPlanner {
	return &CasteljauPathPlanner{
		NumberOfPathPoints: numberOfPathPoints,
		NumberOfUsedPoints: numberOfUsedPoints,
		DistanceFactor:     distanceFactor,
	}
}

// PlanPath is the function that calculate the path
func (pathPlanner CasteljauPathPlanner) PlanPath(initialPose, targetPose model.Pose) Path {
	distance := calculateDistanceBetweenTwoPoints(
		model.Position2D{
			X: initialPose.X,
			Y: initialPose.Y,
		},

		model.Position2D{
			X: targetPose.X,
			Y: targetPose.Y,
		},
	)

	firstControlPoint, secondControlPoint := pathPlanner.generateControlPoints(initialPose, targetPose, distance)

	path := pathPlanner.getPath(initialPose, targetPose, firstControlPoint, secondControlPoint)

	return path
}

func calculateDistanceBetweenTwoPoints(firstPoint, secondPoint model.Position2D) float64 {
	return math.Sqrt(math.Pow((secondPoint.X-firstPoint.X), 2) + math.Pow((secondPoint.Y-firstPoint.Y), 2))
}

func (pathPlanner CasteljauPathPlanner) generateControlPoints(initialPose, targetPose model.Pose, distance float64) (model.Position2D, model.Position2D) {
	firstPoint := model.Position2D{
		math.Cos(initialPose.Orientation)*(distance/pathPlanner.DistanceFactor) + initialPose.X,
		math.Sin(initialPose.Orientation)*(distance/pathPlanner.DistanceFactor) + initialPose.Y,
	}

	secondPoint := model.Position2D{
		math.Cos(targetPose.Orientation+180)*(distance/pathPlanner.DistanceFactor) + targetPose.X,
		math.Sin(targetPose.Orientation+180)*(distance/pathPlanner.DistanceFactor) + targetPose.Y,
	}

	return firstPoint, secondPoint
}

func (pathPlanner CasteljauPathPlanner) getPath(initialPose, targetPose model.Pose, firstControlPoint, secondControlPoint model.Position2D) Path {
	path := NewPath(pathPlanner.NumberOfUsedPoints)

	factor := 1.0 / float64(pathPlanner.NumberOfPathPoints)
	fmt.Printf("Factor: %v", factor)
	t := factor
	for i := 0; i < pathPlanner.NumberOfUsedPoints; i++ {
		t += factor
		path[i] = model.Position2D{
			X: float64(math.Pow((1-t), 3)*initialPose.X + 3*t*math.Pow((1-t), 2)*firstControlPoint.X + 3*math.Pow(t, 2)*(1-t)*secondControlPoint.X + math.Pow(t, 3)*targetPose.X),
			Y: float64(math.Pow((1-t), 3)*initialPose.Y + 3*t*math.Pow((1-t), 2)*firstControlPoint.Y + 3*math.Pow(t, 2)*(1-t)*secondControlPoint.Y + math.Pow(t, 3)*targetPose.Y),
		}

	}

	return path
}
