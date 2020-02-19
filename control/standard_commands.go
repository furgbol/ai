package control

// StandardCommand - Data type to store standard command data
type StandardCommand struct {
	robotID         int
	linearVelocity  float64
	angularVelocity float64
}

// RobotID returns the robot id of the command
func (sd StandardCommand) RobotID() int {
	return sd.robotID
}

// LinearVelocity returns the linear velocity of the command
func (sd StandardCommand) LinearVelocity() float64 {
	return sd.linearVelocity
}

// AngularVelocity returns the angular velocity of the command
func (sd StandardCommand) AngularVelocity() float64 {
	return sd.angularVelocity
}

// SetRobotID sets the robot id of the command
func (sd *StandardCommand) SetRobotID(id int) {
	sd.robotID = id
}

// SetLinearVelocity sets the linear velocity of the command
func (sd *StandardCommand) SetLinearVelocity(linearVel float64) {
	sd.linearVelocity = linearVel
}

// SetAngularVelocity sets the angular velocity of the command
func (sd *StandardCommand) SetAngularVelocity(angularVel float64) {
	sd.angularVelocity = angularVel
}
