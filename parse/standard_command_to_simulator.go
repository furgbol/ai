package parse

import (
	"errors"

	"github.com/furgbol/ai/control/cmd"
	PBCommand "github.com/furgbol/ai/parse/command_pb"
	proto "github.com/golang/protobuf/proto"
)

// StandardCommandToProtobuf - Contains a list of three commands for serializing data
type StandardCommandToProtobuf struct {
	Commands []cmd.StandardCommand
}

// Encode receives a list of simulator commands and serializes with the protobuf
func (commands StandardCommandToProtobuf) Encode() ([]byte, error) {
	if len(commands.Commands) < 3 {
		return nil, errors.New(string("[Parse Error] Insufficient arguments in SerializeToSimulatorCommand.\nhave: " + string(len(commands.Commands)) + "\nwant: 3"))
	}

	simulatorCmd := commands.standardCmdToSimulatorCmd()

	robots := make([](*PBCommand.Robot_Command), 3)

	for i := 0; i < 3; i++ {
		robots[i] = &PBCommand.Robot_Command{
			LeftVel:  parseToFloat32(simulatorCmd[i].LeftWheelVelocity),
			RightVel: parseToFloat32(simulatorCmd[i].RightWheelVelocity),
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

func (commands StandardCommandToProtobuf) standardCmdToSimulatorCmd() []cmd.SimulatorCommand {
	simulatorCommands := make([]cmd.SimulatorCommand, 3)

	for i, v := range commands.Commands {
		simulatorCommands[i] = cmd.SimulatorCommand{
			LeftWheelVelocity:  (v.LinearVelocity - v.AngularVelocity),
			RightWheelVelocity: (v.LinearVelocity + v.AngularVelocity),
		}
	}

	return simulatorCommands
}
