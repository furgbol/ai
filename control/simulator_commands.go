package control

// SimulatorCommand - Data type to store simulator command data
type SimulatorCommand struct {
	RobotID            int
	LeftWheelVelocity  float64
	RightWheelVelocity float64
	BallPositionX			 float64
	BallPositiony			 float64
	//FinalPosition			 Pose          
}
