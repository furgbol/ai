package parse

import (
	"errors"

	"github.com/furgbol/ai/control/cmd"
	PBCommand "github.com/furgbol/ai/parse/command_pb"
	proto "github.com/golang/protobuf/proto"
)

// SerializeToSimulatorCommand receives a list of simulator commands and serializes with the protobuf
func SerializeToSimulatorCommand(commands []cmd.SimulatorCommand) ([]byte, error) {
	if len(commands) < 3 {
		return nil, errors.New(string("[Parse Error] Insufficient arguments in SerializeToSimulatorCommand.\nhave: " + string(len(commands)) + "\nwant: 3"))
	}

	robots := make([](*PBCommand.Robot_Command), 3)

	for i := 0; i < 3; i++ {
		robots[i] = &PBCommand.Robot_Command{
			LeftVel:  parseToFloat32(commands[i].LeftWheelVelocity),
			RightVel: parseToFloat32(commands[i].RightWheelVelocity),
		}
	}

	command := PBCommand.Global_Commands{RobotCommands: robots}

	serial, err := proto.Marshal(&command)
	if err != nil {
		return nil, err
	}

	return serial, nil
}

func parseToFloat32(value float64) *float32 {
	valueReturn := float32(value)
	return &valueReturn
}
