package cmd

import (
	"github.com/furgbol/ai/model"
	"github.com/golang/math"
)

// GeneratorCommand - Data type to model the commands to be generated
type GeneratorCommand struct {
	robotID						 int
	ConstAngularMult	 float64
	ConstLinearMult		 float64
	MaxAngularVelocity float64
	MaxLinearVelocity  float64
}

// NewGeneratorCommands creates an instance of a struct GeneratorCommand
func NewGeneratorCommands(maxLinearVelocity, maxAngularVelocity, constAngularMult, constLinearMult float64) *GeneratorCommand {
	return &GeneratorCommand{
		ConstAngularMult: constAngularMult,
		ConstLinearMult: constLinearMult, 
		MaxAngularVelocity: maxAngularVelocity,
		MaxLinearVelocity: maxLinearVelocity,
	}
}

// Generate is the function that calculates and ruturns the velocities of the robots
func (commands GeneratorCommand) Generate(currentPose, targetPose model.Pose) StandardCommand {
	DistanceTargetCurrent := math.Sqrt(math.Pow(currentPose.Position2D.X - targetPose.Position2D.X, 2) + math.Pow(currentPose.Position2D.Y - targetPose.Position2D.Y, 2))
	linearVelocity :=  DistanceTargetCurrent * commands.ConstLinearMult
	angularVelocity := DistanceTargetCurrent * commands.ConstAngularMult

	if linearVelocity > commands.MaxLinearVelocity {
		linearVelocity = commands.MaxLinearVelocity 
	}
	if angularVelocity > commands.MaxAngularVelocity {
		angularVelocity = commands.maxAngularVelocity
	}

	return StandardCommand{
		RobotID: commands.robotID,
		LinearVelocity: linearVelocity,
		AngularVelocity: angularVelocity,
	}
}