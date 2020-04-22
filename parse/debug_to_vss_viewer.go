package parse

import (
	"errors"

	"github.com/furgbol/ai/debug"
	PBDebug "github.com/furgbol/ai/parse/debug_pb"
	proto "github.com/golang/protobuf/proto"
)

// DebugToProtobuf - Type that contains attributes for sending the debug
type DebugToProtobuf struct {
	Debug debug.Debug
}

// Encode receives a debug type and serializes with the protobuf
func (d DebugToProtobuf) Encode() ([]byte, error) {
	check := checkSizes(d)
	if check != nil {
		return nil, errors.New(string("[Parse Error] Insufficient arguments in Encode Debug.\nhave: " + string(check) + "\nwant: >= 3"))
	}

	globalDebug := pathToPBDebug(d)

	serial, err := proto.Marshal(&globalDebug)
	if err != nil {
		return nil, err
	}

	return serial, nil
}

func pathToPBDebug(d DebugToProtobuf) PBDebug.Global_Debug {
	stepPoses := make([](*PBDebug.Pose), len(d.Debug.InitialPoses))
	finalPoses := make([](*PBDebug.Pose), len(d.Debug.FinalPoses))

	for i := 0; i < len(d.Debug.InitialPoses); i++ {
		stepPoses[i] = &PBDebug.Pose{
			X:   parseToFloat32(d.Debug.InitialPoses[i].X),
			Y:   parseToFloat32(d.Debug.InitialPoses[i].Y),
			Yaw: parseToFloat32(d.Debug.InitialPoses[i].Orientation),
		}

		finalPoses[i] = &PBDebug.Pose{
			X:   parseToFloat32(d.Debug.FinalPoses[i].X),
			Y:   parseToFloat32(d.Debug.FinalPoses[i].Y),
			Yaw: parseToFloat32(d.Debug.FinalPoses[i].Orientation),
		}
	}

	paths := make([](*PBDebug.Path), len(d.Debug.PathRobots))

	for i := 0; i < len(d.Debug.PathRobots); i++ {
		path := make([](*PBDebug.Pose), len(d.Debug.PathRobots[i]))
		for j := 0; j < len(d.Debug.PathRobots[i]); j++ {
			path[j] = &PBDebug.Pose{
				X:   parseToFloat32(d.Debug.PathRobots[i][j].X),
				Y:   parseToFloat32(d.Debug.PathRobots[i][j].Y),
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
	if len(d.Debug.InitialPoses) >= 3 {
		if len(d.Debug.InitialPoses) == len(d.Debug.FinalPoses) && len(d.Debug.InitialPoses) == len(d.Debug.PathRobots) {
			return nil
		}
	}

	message := "\nInitialPoses: " + string(len(d.Debug.InitialPoses))
	message += "\nFinalPoses: " + string(len(d.Debug.FinalPoses))
	message += "\nPathRobots: " + string(len(d.Debug.PathRobots))
	message += "\n"

	return []byte(message)
}
