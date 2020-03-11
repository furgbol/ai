package parse

import "github.com/furgbol/ai/control"

// CommandsToSimulator - Data type to store the commands to simulator
type CommandsToSimulator []control.SimulatorCommand

// Encode creates a list with commands to simulator
func Encode(cmdRepo control.StandardCommand) *control.SimulatorCommand {
	return &control.SimulatorCommand{
		LeftWheelVelocity: cmdRepo.LinearVelocity - cmdRepo.AngularVelocity,
		RightWheelVelocity: cmdRepo.LinearVelocity + cmdRepo.AngularVelocity,
	}
}

func NewCommandsToSimulator(NumberOfRobots int) CommandsToSimulator {
	Commands := make(CommandsToSimulator,NumberOfRobots)
	return Commands
}