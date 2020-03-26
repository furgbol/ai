package parse

import (
	"testing"
	"github.com/furgbol/ai/control/cmd"
)	

func TestNewCommandsToSimulator(t *testing.T){
	Commands := NewCommandsToSimulator(3)
	if len(Commands) != 3 {
		t.Errorf("Team creation failed: expected team size: 3, got: %v", len(Commands))
	}
}

func TestEncode(t *testing.T){
	cmdRepo := cmd.NewCommandsRepository(3)
	(*cmdRepo)[0] = cmd.StandardCommand{
		RobotID: 1,
		LinearVelocity: 10,
		AngularVelocity: 5,
	}
	commands := Encode((*cmdRepo)[0])
	if commands.LeftWheelVelocity != 5{
		t.Errorf("Encode failed: expected left velocity: 5, got: %v", commands.LeftWheelVelocity)
	}
	if commands.RightWheelVelocity != 15{
		t.Errorf("Encode failed: expected right velocity: 15, got: %v", commands.RightWheelVelocity)
	}
}
