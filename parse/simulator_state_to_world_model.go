package parse

import (
	"github.com/furgbol/ai/model"
	PBState "github.com/furgbol/ai/parse/state_pb"
	"github.com/golang/protobuf/proto"
)

// ProtobufToWorldModel - Contains attributes to deserialize the simulator data
type ProtobufToWorldModel struct {
	Team         int
	ReceivedData []byte
}

const (
	// TeamYellow - Used to specify the team (yellow)
	TeamYellow = 0
	// TeamBlue - Used to specify the team (blue)
	TeamBlue = 1
)

// Decode deserializes the vss state and returns the state of the game
func (state ProtobufToWorldModel) Decode() (*model.World, error) {
	data := &PBState.Global_State{}
	err := proto.Unmarshal(state.ReceivedData, data)
	if err != nil {
		return nil, err
	}

	world := model.NewWorld(model.Ball{}, model.NewTeam(3))

	world.Ball.Position2D.X = float64(data.GetBalls()[0].GetPose().GetX())
	world.Ball.Position2D.Y = float64(data.GetBalls()[0].GetPose().GetY())

	if state.Team == TeamYellow {
		for i := 0; i < 3; i++ {
			world.Team[i].ID = i
			world.Team[i].CurrentPosition.Position2D.X = float64(data.GetRobotsYellow()[i].GetPose().GetX())
			world.Team[i].CurrentPosition.Position2D.Y = float64(data.GetRobotsYellow()[i].GetPose().GetY())
			world.Team[i].CurrentPosition.Orientation = float64(data.GetRobotsYellow()[i].GetPose().GetYaw())
		}
	} else {
		for i := 0; i < 3; i++ {
			world.Team[i].ID = i
			world.Team[i].CurrentPosition.Position2D.X = float64(data.GetRobotsBlue()[i].GetPose().GetX())
			world.Team[i].CurrentPosition.Position2D.Y = float64(data.GetRobotsBlue()[i].GetPose().GetY())
			world.Team[i].CurrentPosition.Orientation = float64(data.GetRobotsBlue()[i].GetPose().GetYaw())
		}
	}

	return world, nil
}
