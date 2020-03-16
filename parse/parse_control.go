package parse

import (
	"github.com/furgbol/ai/control/cmd"
	PBControl "github.com/furgbol/ai/parse/control_pb"
	proto "github.com/golang/protobuf/proto"
)

// ControlFromSimulator deserializes the vss control and returns the control of the game
func ControlFromSimulator(ctrl []byte) (*cmd.Control, error) {
	data := &PBControl.User_Control{}
	err := proto.Unmarshal(ctrl, data)
	if err != nil {
		return nil, err
	}

	newControl := &cmd.Control{
		Paused: data.GetPaused(),
		NewBallPose: cmd.Pose{
			X: float64(data.GetNewBallPose().GetX()),
			Y: float64(data.GetNewBallPose().GetY()),
		},
		NewRobotsBluePose:   make([]cmd.Pose, 3),
		NewRobotsYellowPose: make([]cmd.Pose, 3),
	}

	for i := 0; i < 3; i++ {
		newControl.NewRobotsBluePose[i].X = float64(data.GetNewRobotsBluePose()[i].GetX())
		newControl.NewRobotsBluePose[i].Y = float64(data.GetNewRobotsBluePose()[i].GetY())
		newControl.NewRobotsBluePose[i].Orientation = float64(data.GetNewRobotsBluePose()[i].GetYaw())

		newControl.NewRobotsYellowPose[i].X = float64(data.GetNewRobotsYellowPose()[i].GetX())
		newControl.NewRobotsYellowPose[i].Y = float64(data.GetNewRobotsYellowPose()[i].GetY())
		newControl.NewRobotsYellowPose[i].Orientation = float64(data.GetNewRobotsYellowPose()[i].GetYaw())
	}

	return newControl, nil
}
