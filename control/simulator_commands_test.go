package control

import "testing"

// Commands - Data type to store the commands to be send to the Simulator
type Commands []SimulatorCommand

func NewSimulatorCommands(NumberRobots int) *Commands {
	sltcmd := make(Commands, NumberRobots)
	return &sltcmd
}

func testNewSimulatorCommands(t *testing.T) {
	sltcmd := NewSimulatorCommands(2)
	if len(*sltcmd) != 2 {
		t.Errorf("Error on creating instance: Length expected: 2, got: %v", len(*sltcmd))
	}
}