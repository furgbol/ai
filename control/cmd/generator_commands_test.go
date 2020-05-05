package cmd

import (
	"testing"
	"github.com/furgbol/ai/model"
)

type GeneratorTester struct {
	generator Generator
}

func TestGenerate(t *testing.T) {
	currentPose := model.Pose {
		Position2D: model.Position2D{
			X: 3.0,
			Y: 1.0,
		},
		Orientation: 67.3,
	}

	targetPose := model.Pose {
		Position2D: model.Position2D{
			X: 6.0,
			Y: 5.0,
		},
		Orientation: 53.2,
	}

	commands := GeneratorTester{generator: NewGeneratorCommands(1,30,30,5,5)}
	standardCommands := commands.generator.Generate(currentPose, targetPose)

	if standardCommands.RobotID != 1 {
		t.Errorf("Error on testing RobotID expected: 1, returned: %v", standardCommands.RobotID)
	}
	if standardCommands.LinearVelocity != 25 {
		t.Errorf("Error on testing LinearVelocity expected: 25, returned: %v", standardCommands.LinearVelocity)
	}
	if standardCommands.LinearVelocity != 25 {
		t.Errorf("Error on testing AngularVelocity expected: 25, returned: %v", standardCommands.AngularVelocity)
	} 
}
