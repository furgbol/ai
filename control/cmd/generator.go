package cmd

// Generator - Interface to calculates velocities
type Generator interface {
	Generate(currentPose, targetPose model.Pose) StandardCommand
}