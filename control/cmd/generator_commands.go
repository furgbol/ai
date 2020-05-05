package cmd

import (
	"github.com/furgbol/ai/model"
	"math"
)

// GeneratorCommand - Data type to model the commands to be generated
type GeneratorCommand struct {
	RobotID						 int
	ConstAngularMult	 float64
	ConstLinearMult		 float64
	MaxAngularVelocity float64
	MaxLinearVelocity  float64
}

// NewGeneratorCommands creates an instance of a struct GeneratorCommand
func NewGeneratorCommands(robotID int, maxLinearVelocity, maxAngularVelocity, constAngularMult, constLinearMult float64) *GeneratorCommand {
	return &GeneratorCommand{
		RobotID: robotID,
		ConstAngularMult: constAngularMult,
		ConstLinearMult: constLinearMult, 
		MaxAngularVelocity: maxAngularVelocity,
		MaxLinearVelocity: maxLinearVelocity,
	}
}

func calculateDistanceBetweenTwoPoints(firstPoint, secondPoint model.Position2D) float64 {
	return math.Sqrt(math.Pow((secondPoint.X - firstPoint.X), 2) + math.Pow((secondPoint.Y - firstPoint.Y), 2))
}

// Generate is the function that calculates and ruturns the velocities of the robots
func (commands GeneratorCommand) Generate(currentPose, targetPose model.Pose) StandardCommand {
	Distance := calculateDistanceBetweenTwoPoints(currentPose.Position2D, targetPose.Position2D)
	linearVelocity :=  Distance * commands.ConstLinearMult
	angularVelocity := Distance * commands.ConstAngularMult

	if linearVelocity > commands.MaxLinearVelocity {
		linearVelocity = commands.MaxLinearVelocity 
	}
	if angularVelocity > commands.MaxAngularVelocity {
		angularVelocity = commands.MaxAngularVelocity
	}

	return StandardCommand{
		RobotID: commands.RobotID,
		LinearVelocity: linearVelocity,
		AngularVelocity: angularVelocity,
	}
}