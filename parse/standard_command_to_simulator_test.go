package parse

import (
	"testing"

	"github.com/furgbol/ai/control/cmd"
	PBCommand "github.com/furgbol/ai/parse/command_pb"
	proto "github.com/golang/protobuf/proto"
)

func TestEncodeForCommand(t *testing.T) {
	var command Encoder

	command = StandardCommandToProtobuf{
		Commands: []cmd.StandardCommand{
			cmd.StandardCommand{
				LinearVelocity:  5,
				AngularVelocity: 2,
			},
			cmd.StandardCommand{
				LinearVelocity:  12.3,
				AngularVelocity: 4,
			},
			cmd.StandardCommand{
				LinearVelocity:  4.5,
				AngularVelocity: 0,
			},
		},
	}

	serial, err := command.Encode()
	if err != nil {
		t.Error("Parse command failed: Serialize error:", err)
	}

	commandTest := &PBCommand.Global_Commands{}
	err = proto.Unmarshal(serial, commandTest)
	if err != nil {
		t.Error("Parse command failed: Deserialize error:", err)
	}

	if commandTest.GetRobotCommands()[0].GetLeftVel() != 3 {
		t.Error("Parse command failed: left wheel speed\nhave: ", commandTest.GetRobotCommands()[0].GetLeftVel(),
			"\nwant: 3")
	}
	if commandTest.GetRobotCommands()[0].GetRightVel() != 7 {
		t.Error("Parse command failed: right wheel speed\nhave: ", commandTest.GetRobotCommands()[0].GetRightVel(),
			"\nwant: 7")
	}
	if commandTest.GetRobotCommands()[1].GetLeftVel() != 8.3 {
		t.Error("Parse command failed: left wheel speed\nhave: ", commandTest.GetRobotCommands()[1].GetLeftVel(),
			"\nwant: 8.3")
	}
	if commandTest.GetRobotCommands()[1].GetRightVel() != 16.3 {
		t.Error("Parse command failed: right wheel speed\nhave: ", commandTest.GetRobotCommands()[1].GetRightVel(),
			"\nwant: 16.3")
	}
	if commandTest.GetRobotCommands()[2].GetLeftVel() != 4.5 {
		t.Error("Parse command failed: left wheel speed\nhave: ", commandTest.GetRobotCommands()[2].GetLeftVel(),
			"\nwant: 4.5")
	}
	if commandTest.GetRobotCommands()[2].GetRightVel() != 4.5 {
		t.Error("Parse command failed: right wheel speed\nhave: ", commandTest.GetRobotCommands()[2].GetRightVel(),
			"\nwant: 4.5")
	}
}
