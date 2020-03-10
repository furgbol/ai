package parse

import (
	"testing"

	PBControl "github.com/furgbol/ai/parse/control_pb"
	proto "github.com/golang/protobuf/proto"
)

func TestControlFromSimulator(t *testing.T) {
	ctrl := generateControl()
	serial, err := proto.Marshal(ctrl)
	if err != nil {
		t.Error("Parse control failed: Serialize error:", err)
	}

	ctrlNew, err := ControlFromSimulator(serial)
	if err != nil {
		t.Error("Parse control failed: Deserialize error:", err)
	}

	//Paused

	if ctrlNew.Paused != true {
		t.Error("Parse control failed: Paused\nhave: ", ctrlNew.Paused,
			"\nwant: true")
	}

	//Ball

	caseThis := ctrlNew.NewBallPose
	if !compareValues(caseThis.X, 12.3) || !compareValues(caseThis.Y, 3) {
		t.Error("Parse control failed: NewBallPose\nhave: ", caseThis.X, ",", caseThis.Y,
			"\nwant: 12.3 , 3")
	}

	//Team Blue

	caseThis = ctrlNew.NewRobotsBluePose[0]
	if !compareValues(caseThis.X, 1) || !compareValues(caseThis.Y, 2) || !compareValues(caseThis.Orientation, 1) {
		t.Error("Parse control failed: NewRobotsBluePose\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 1 , 2 , 1")
	}
	caseThis = ctrlNew.NewRobotsBluePose[1]
	if !compareValues(caseThis.X, 2) || !compareValues(caseThis.Y, 3) || !compareValues(caseThis.Orientation, 2) {
		t.Error("Parse control failed: NewRobotsBluePose\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 2 , 3 , 2")
	}
	caseThis = ctrlNew.NewRobotsBluePose[2]
	if !compareValues(caseThis.X, 3) || !compareValues(caseThis.Y, 4) || !compareValues(caseThis.Orientation, 90) {
		t.Error("Parse control failed: NewRobotsBluePose\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 3 , 4 , 90")
	}

	//Team Yellow

	caseThis = ctrlNew.NewRobotsYellowPose[0]
	if !compareValues(caseThis.X, 5) || !compareValues(caseThis.Y, 6) || !compareValues(caseThis.Orientation, 130) {
		t.Error("Parse control failed: NewRobotsYellowPose\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 5 , 6 , 130")
	}
	caseThis = ctrlNew.NewRobotsYellowPose[1]
	if !compareValues(caseThis.X, 6) || !compareValues(caseThis.Y, 7) || !compareValues(caseThis.Orientation, 12.5) {
		t.Error("Parse control failed: NewRobotsYellowPose\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 6 , 7 , 12.5")
	}
	caseThis = ctrlNew.NewRobotsYellowPose[2]
	if !compareValues(caseThis.X, 7) || !compareValues(caseThis.Y, 8) || !compareValues(caseThis.Orientation, 25.3) {
		t.Error("Parse control failed: NewRobotsYellowPose\nhave: ", caseThis.X, ",", caseThis.Y, ",", caseThis.Orientation,
			"\nwant: 7 , 8 , 25.3")
	}
}

func compareValues(x, y float64) bool {
	if (x-0.1) < y && (x+0.1) > y {
		return true
	}
	return false
}

func generateControl() *PBControl.User_Control {
	paused := true
	return &PBControl.User_Control{
		Paused: &paused,
		NewBallPose: &PBControl.Pose{
			X: parseToFloat32(12.3),
			Y: parseToFloat32(3),
		},
		NewRobotsBluePose: []*PBControl.Pose{
			&PBControl.Pose{
				X:   parseToFloat32(1),
				Y:   parseToFloat32(2),
				Yaw: parseToFloat32(1),
			},
			&PBControl.Pose{
				X:   parseToFloat32(2),
				Y:   parseToFloat32(3),
				Yaw: parseToFloat32(2),
			},
			&PBControl.Pose{
				X:   parseToFloat32(3),
				Y:   parseToFloat32(4),
				Yaw: parseToFloat32(90),
			},
		},
		NewRobotsYellowPose: []*PBControl.Pose{
			&PBControl.Pose{
				X:   parseToFloat32(5),
				Y:   parseToFloat32(6),
				Yaw: parseToFloat32(130),
			},
			&PBControl.Pose{
				X:   parseToFloat32(6),
				Y:   parseToFloat32(7),
				Yaw: parseToFloat32(12.5),
			},
			&PBControl.Pose{
				X:   parseToFloat32(7),
				Y:   parseToFloat32(8),
				Yaw: parseToFloat32(25.3),
			},
		},
	}
}
