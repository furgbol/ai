package parse

import (
	"testing"

	"github.com/furgbol/ai/control/cmd"
	PBCommand "github.com/furgbol/ai/parse/command_pb"
	proto "github.com/golang/protobuf/proto"
)

func TestSerializeToSimulatorCommand(t *testing.T) {
	command := []cmd.SimulatorCommand{
		cmd.SimulatorCommand{
			LeftWheelVelocity:  2,
			RightWheelVelocity: 3,
		},
		cmd.SimulatorCommand{
			LeftWheelVelocity:  -12.3,
			RightWheelVelocity: -13,
		},
		cmd.SimulatorCommand{
			LeftWheelVelocity:  -4.5,
			RightWheelVelocity: 0,
		},
	}

	serial, err := SerializeToSimulatorCommand(command)
	if err != nil {
		t.Error("Parse command failed: Serialize error:", err)
	}

	commandTest := &PBCommand.Global_Commands{}
	err = proto.Unmarshal(serial, commandTest)
	if err != nil {
		t.Error("Parse command failed: Deserialize error:", err)
	}

	if commandTest.GetRobotCommands()[0].GetLeftVel() != 2 {
		t.Error("Parse command failed: left wheel speed\nhave: ", commandTest.GetRobotCommands()[0].GetLeftVel(),
			"\nwant: 2")
	}
	if commandTest.GetRobotCommands()[0].GetRightVel() != 3 {
		t.Error("Parse command failed: right wheel speed\nhave: ", commandTest.GetRobotCommands()[0].GetRightVel(),
			"\nwant: 3")
	}
	if commandTest.GetRobotCommands()[1].GetLeftVel() != -12.3 {
		t.Error("Parse command failed: left wheel speed\nhave: ", commandTest.GetRobotCommands()[1].GetLeftVel(),
			"\nwant: -12.3")
	}
	if commandTest.GetRobotCommands()[1].GetRightVel() != -13 {
		t.Error("Parse command failed: right wheel speed\nhave: ", commandTest.GetRobotCommands()[1].GetRightVel(),
			"\nwant: -13")
	}
	if commandTest.GetRobotCommands()[2].GetLeftVel() != -4.5 {
		t.Error("Parse command failed: left wheel speed\nhave: ", commandTest.GetRobotCommands()[2].GetLeftVel(),
			"\nwant: -4.5")
	}
	if commandTest.GetRobotCommands()[2].GetRightVel() != 0 {
		t.Error("Parse command failed: right wheel speed\nhave: ", commandTest.GetRobotCommands()[2].GetRightVel(),
			"\nwant: 0")
	}
}
