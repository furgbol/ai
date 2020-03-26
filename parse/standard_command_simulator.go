package parse

import "github.com/furgbol/ai/control/cmd"

// CommandsToSimulator - Data type to store the commands to simulator
type CommandsToSimulator []cmd.SimulatorCommand

// Encode convert commands to the simulator 
func Encode(cmdRepo cmd.StandardCommand) *cmd.SimulatorCommand {
	return &cmd.SimulatorCommand{
		LeftWheelVelocity: cmdRepo.LinearVelocity - cmdRepo.AngularVelocity,
		RightWheelVelocity: cmdRepo.LinearVelocity + cmdRepo.AngularVelocity,
	}
}

// NewCommandsToSimulator creates and returns commands to simulator  
func NewCommandsToSimulator(NumberOfRobots int) CommandsToSimulator {
	Commands := make(CommandsToSimulator,NumberOfRobots)
	return Commands
}	