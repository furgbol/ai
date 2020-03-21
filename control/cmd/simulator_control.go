package cmd

//Pose - Data type to model a pose
type Pose struct {
	X           float64
	Y           float64
	Orientation float64
}

//Control - Contains the data structure of the simulator control
type Control struct {
	Paused              bool
	NewBallPose         Pose
	NewRobotsBluePose   []Pose
	NewRobotsYellowPose []Pose
}
