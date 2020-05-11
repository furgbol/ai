package parse

import (
	"errors"

	"github.com/furgbol/ai/control/navigation"
	"github.com/furgbol/ai/model"
	PBDebug "github.com/furgbol/ai/parse/debug_pb"
	proto "github.com/golang/protobuf/proto"
)

// DebugToProtobuf - Type that contains attributes for sending the debug
type DebugToProtobuf struct {
	InitialPoses []model.Pose
	FinalPoses   []model.Pose
	PathRobots   []navigation.Path
}

// Encode receives a debug type and serializes with the protobuf
func (d DebugToProtobuf) Encode() ([]byte, error) {
	check := checkSizes(d)
	if check != nil {
		return nil, errors.New(string("[Parse Error] Insufficient arguments in Encode Debug. \nhave: " + string(check) + "\nwant: >= 3"))
	}

	globalDebug := pathToPBDebug(d)

	serial, err := proto.Marshal(&globalDebug)
	if err != nil {
		return nil, err
	}

	return serial, nil
}

func pathToPBDebug(d DebugToProtobuf) PBDebug.Global_Debug {
	stepPoses := make([](*PBDebug.Pose), len(d.InitialPoses))
	finalPoses := make([](*PBDebug.Pose), len(d.FinalPoses))

	for i := 0; i < len(d.InitialPoses); i++ {
		stepPoses[i] = &PBDebug.Pose{
			X:   parseToFloat32(d.InitialPoses[i].X),
			Y:   parseToFloat32(d.InitialPoses[i].Y),
			Yaw: parseToFloat32(d.InitialPoses[i].Orientation),
		}

		finalPoses[i] = &PBDebug.Pose{
			X:   parseToFloat32(d.FinalPoses[i].X),
			Y:   parseToFloat32(d.FinalPoses[i].Y),
			Yaw: parseToFloat32(d.FinalPoses[i].Orientation),
		}
	}

	paths := make([](*PBDebug.Path), len(d.PathRobots))

	for i := 0; i < len(d.PathRobots); i++ {
		path := make([](*PBDebug.Pose), len(d.PathRobots[i]))
		for j := 0; j < len(d.PathRobots[i]); j++ {
			path[j] = &PBDebug.Pose{
				X:   parseToFloat32(d.PathRobots[i][j].X),
				Y:   parseToFloat32(d.PathRobots[i][j].Y),
				Yaw: parseToFloat32(0),
			}
		}
		paths[i] = &PBDebug.Path{Poses: path}
	}

	return PBDebug.Global_Debug{
		StepPoses:  stepPoses,
		FinalPoses: finalPoses,
		Paths:      paths,
	}
}

func checkSizes(d DebugToProtobuf) []byte {
	if len(d.InitialPoses) >= 3 {
		if len(d.InitialPoses) == len(d.FinalPoses) && len(d.InitialPoses) == len(d.PathRobots) {
			return nil
		}
	}

	message := "\nInitialPoses: " + string(len(d.InitialPoses))
	message += "\nFinalPoses: " + string(len(d.FinalPoses))
	message += "\nPathRobots: " + string(len(d.PathRobots))
	message += "\n"

	return []byte(message)
}
