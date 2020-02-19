package model

// Position2D - Data type to model a two-dimensional position
type Position2D struct {
	X float64
	Y float64
}

// Pose - Data type to model a pose with orientation
type Pose struct {
	Position2D
	Orientation float64
}

// Ball - Data type to model the ball
type Ball struct {
	Position2D
}

// Robot - Data type to model the robots
type Robot struct {
	ID              int
	Role            string
	CurrentPosition Pose
	TargetPosition  Pose
}

// Team - Data type to model the team
type Team []Robot

// World - Data type to model the world
type World struct {
	Ball Ball
	Team Team
}

// NewTeam creates an instance of the team
func NewTeam(numberOfRobots int) Team {
	team := make(Team, numberOfRobots)
	return team
}

// NewWorld creates an instance of the world
func NewWorld(ball Ball, team Team) *World {
	return &World{
		Ball: ball,
		Team: team,
	}
}
