package parse

import (
	"testing"

	PBState "github.com/furgbol/ai/parse/state_pb"
	proto "github.com/golang/protobuf/proto"
)

func TestStateFromSimulator(t *testing.T) {
	state := generateState()
	serial, err := proto.Marshal(state)
	if err != nil {
		t.Error("Parse state failed: Serialize error:", err)
	}

	var receivedState Decoder

	receivedState = ProtobufToWorldModel{
		Team:         TeamBlue,
		ReceivedData: serial,
	}

	stateNew, err := receivedState.Decode()
	if err != nil {
		t.Error("Parse state failed: Deserialize error:", err)
	}

	//Ball
	caseBall := stateNew.Ball
	if !compareValue(caseBall.X, 85) || !compareValue(caseBall.Y, 65) {
		t.Error("Parse state failed: Ball\nhave: ", caseBall.X, ",", caseBall.Y,
			"\nwant: 85 , 65")
	}

	//Team Blue
	caseThis := stateNew.Team[0].CurrentPosition
	if !compareValue(caseThis.X, 10) || !compareValue(caseThis.Y, 11) || !compareValue(caseThis.Orientation, 100) {
		t.Error("Parse state failed: TeamBlue\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 10 , 11, 100")
	}
	caseThis = stateNew.Team[1].CurrentPosition
	if !compareValue(caseThis.X, 12.8) || !compareValue(caseThis.Y, 13) || !compareValue(caseThis.Orientation, 130.3) {
		t.Error("Parse state failed: TeamBlue\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 12.8 , 13, 130.3")
	}
	caseThis = stateNew.Team[2].CurrentPosition
	if !compareValue(caseThis.X, 14) || !compareValue(caseThis.Y, 15) || !compareValue(caseThis.Orientation, 30) {
		t.Error("Parse state failed: TeamBlue\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 14 , 15, 30")
	}

	//Team Yellow

	receivedState = ProtobufToWorldModel{
		Team:         TeamYellow,
		ReceivedData: serial,
	}

	stateNew, err = receivedState.Decode()
	if err != nil {
		t.Error("Parse state failed: Deserialize error:", err)
	}

	caseThis = stateNew.Team[0].CurrentPosition
	if !compareValue(caseThis.X, 15) || !compareValue(caseThis.Y, 16) || !compareValue(caseThis.Orientation, 22.7) {
		t.Error("Parse state failed: TeamYellow\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 15 , 16, 22.7")
	}
	caseThis = stateNew.Team[1].CurrentPosition
	if !compareValue(caseThis.X, 16) || !compareValue(caseThis.Y, 17) || !compareValue(caseThis.Orientation, 54.2) {
		t.Error("Parse state failed: TeamYellow\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 16 , 17, 54.2")
	}
	caseThis = stateNew.Team[2].CurrentPosition
	if !compareValue(caseThis.X, 17) || !compareValue(caseThis.Y, 18) || !compareValue(caseThis.Orientation, 101.1) {
		t.Error("Parse state failed: TeamYellow\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 17 , 18, 101.1")
	}
}

func compareValue(x, y float64) bool {
	if (x-0.1) < y && (x+0.1) > y {
		return true
	}
	return false
}

func generateState() *PBState.Global_State {
	return &PBState.Global_State{
		Balls: []*PBState.Ball_State{
			&PBState.Ball_State{
				Pose: &PBState.Pose{
					X: parseToFloat32(85),
					Y: parseToFloat32(65),
				},
			},
		},
		RobotsBlue: []*PBState.Robot_State{
			&PBState.Robot_State{
				Pose: &PBState.Pose{
					X:   parseToFloat32(10),
					Y:   parseToFloat32(11),
					Yaw: parseToFloat32(100),
				},
			},
			&PBState.Robot_State{
				Pose: &PBState.Pose{
					X:   parseToFloat32(12.8),
					Y:   parseToFloat32(13),
					Yaw: parseToFloat32(130.3),
				},
			},
			&PBState.Robot_State{
				Pose: &PBState.Pose{
					X:   parseToFloat32(14),
					Y:   parseToFloat32(15),
					Yaw: parseToFloat32(30),
				},
			},
		},
		RobotsYellow: []*PBState.Robot_State{
			&PBState.Robot_State{
				Pose: &PBState.Pose{
					X:   parseToFloat32(15),
					Y:   parseToFloat32(16),
					Yaw: parseToFloat32(22.7),
				},
			},
			&PBState.Robot_State{
				Pose: &PBState.Pose{
					X:   parseToFloat32(16),
					Y:   parseToFloat32(17),
					Yaw: parseToFloat32(54.2),
				},
			},
			&PBState.Robot_State{
				Pose: &PBState.Pose{
					X:   parseToFloat32(17),
					Y:   parseToFloat32(18),
					Yaw: parseToFloat32(101.1),
				},
			},
		},
	}
}
