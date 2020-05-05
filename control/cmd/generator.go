package cmd

import "github.com/furgbol/ai/model"

// Generator - Interface to calculates velocities
type Generator interface {
	Generate(currentPose, targetPose model.Pose) StandardCommand
}