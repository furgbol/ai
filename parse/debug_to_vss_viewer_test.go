package parse

import (
	"testing"

	"github.com/furgbol/ai/control/navigation"
	"github.com/furgbol/ai/debug"
	"github.com/furgbol/ai/model"
	PBDebug "github.com/furgbol/ai/parse/debug_pb"
	"github.com/golang/protobuf/proto"
)

func TestDebugParser(t *testing.T) {
	var Encoder Encoder

	ini, fim := generatePos()
	caminhos := generatePath()

	debug := DebugToProtobuf{
		Debug: debug.Debug{
			InitialPoses: ini,
			FinalPoses:   fim,
			PathRobots:   caminhos,
		},
	}

	Encoder = debug
	serial, err := Encoder.Encode()
	if err != nil {
		t.Error("Parse debug failed: Serialize error:", err)
	}

	debugTest := &PBDebug.Global_Debug{}
	err = proto.Unmarshal(serial, debugTest)
	if err != nil {
		t.Error("Parse debug failed: Deserialize error:", err)
	}

	for i := 0; i < len(debugTest.GetStepPoses()); i++ {
		mid := debugTest.GetStepPoses()[i]
		if !compareValues(mid.GetX(), ini[i].X) || !compareValues(mid.GetY(), ini[i].Y) {
			t.Error("Parse debug failed: InitialPoses\nhave: ", mid.GetX(), ",", mid.GetY(),
				"\nwant: ", ini[i].X, ",", ini[i].Y)
		}
	}

	for i := 0; i < len(debugTest.GetFinalPoses()); i++ {
		mid := debugTest.GetFinalPoses()[i]
		if !compareValues(mid.GetX(), fim[i].X) || !compareValues(mid.GetY(), fim[i].Y) {
			t.Error("Parse debug failed: FinalPoses\nhave: ", mid.GetX(), ",", mid.GetY(),
				"\nwant: ", fim[i].X, ",", fim[i].Y)
		}
	}
	for i := 0; i < len(debugTest.GetPaths()); i++ {
		for j := 0; j < len(debugTest.GetPaths()[i].GetPoses()); j++ {
			mid := debugTest.GetPaths()[i].GetPoses()[j]
			if !compareValues(mid.GetX(), caminhos[i][j].X) || !compareValues(mid.GetY(), caminhos[i][j].Y) {
				t.Error("Parse debug failed: Paths\nhave: ", mid.GetX(), ",", mid.GetY(),
					"\nwant: ", caminhos[i][j].X, ",", caminhos[i][j].Y)
			}
		}
	}
}

func compareValues(x float32, y float64) bool {
	if (x-0.1) < float32(y) && (x+0.1) > float32(y) {
		return true
	}
	return false
}

func generatePos() ([]model.Pose, []model.Pose) {
	pos := make([]model.Pose, 3)
	pos[0] = model.Pose{
		Position2D: model.Position2D{
			X: 1,
			Y: 2,
		},
		Orientation: 0,
	}
	pos[1] = model.Pose{
		Position2D: model.Position2D{
			X: 2,
			Y: 3,
		},
		Orientation: 0,
	}
	pos[2] = model.Pose{
		Position2D: model.Position2D{
			X: 3,
			Y: 4,
		},
		Orientation: 0,
	}

	pos2 := make([]model.Pose, 3)
	pos2[0] = model.Pose{
		Position2D: model.Position2D{
			X: 70,
			Y: 20,
		},
		Orientation: 30,
	}
	pos2[1] = model.Pose{
		Position2D: model.Position2D{
			X: 30,
			Y: 120,
		},
		Orientation: 100,
	}
	pos2[2] = model.Pose{
		Position2D: model.Position2D{
			X: 40,
			Y: 50,
		},
		Orientation: 0,
	}
	return pos, pos2
}

func generatePath() []navigation.Path {
	caminho := make([]navigation.Path, 3)
	caminho[0] = navigation.Path{
		model.Position2D{
			X: 157,
			Y: 65,
		},
		model.Position2D{
			X: 130,
			Y: 20,
		},
		model.Position2D{
			X: 120,
			Y: 15,
		},
		model.Position2D{
			X: 100,
			Y: 10,
		},
		model.Position2D{
			X: 70,
			Y: 20,
		},
	}

	caminho[1] = navigation.Path{
		model.Position2D{
			X: 95,
			Y: 55,
		},
		model.Position2D{
			X: 110,
			Y: 100,
		},
		model.Position2D{
			X: 105,
			Y: 120,
		},
		model.Position2D{
			X: 30,
			Y: 120,
		},
	}

	caminho[2] = navigation.Path{
		model.Position2D{
			X: 95,
			Y: 97,
		},
		model.Position2D{
			X: 80,
			Y: 60,
		},
		model.Position2D{
			X: 40,
			Y: 50,
		},
	}
	return caminho
}
