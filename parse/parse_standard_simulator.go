package parse

import "github.com/furgbol/ai/control"

// CommandsToSimulator - Data type to store the commands to simulator
type CommandsToSimulator []control.SimulatorCommand

// EncodeToCommands convert commands to the simulator 
func EncodeToCommands(cmdRepo control.StandardCommand) *control.SimulatorCommand {
	return &control.SimulatorCommand{
		LeftWheelVelocity: cmdRepo.LinearVelocity - cmdRepo.AngularVelocity,
		RightWheelVelocity: cmdRepo.LinearVelocity + cmdRepo.AngularVelocity,
	}
}

// NewCommandsToSimulator creates and returns commands to simulator  
func NewCommandsToSimulator(NumberOfRobots int) CommandsToSimulator {
	Commands := make(CommandsToSimulator,NumberOfRobots)
	return Commands
}	