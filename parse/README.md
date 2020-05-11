# Parse Package

This package provides types and methods for converting standard data to simulator data, it also serializes and deserializes that data.

# Contents

Data type of StandardCommand to protobuf and simulator state to world model.

1. [StandardCommandToProtobuf](#command)
2. [ProtobufToWorldModel](#state)
3. [DebugToProtobuf](#debug)

StandardCommand to protobuf and simulator state to world model use the Encoder/Decoder interfaces.

1. [Encoder](#encoder)
2. [Decoder](#decoder)

<a name="command"></a>
## StandardCommandToProtobuf Type

This type contains the following fields:

- *Commands()*

<a name="state"></a>
## ProtobufToWorldModel Type

Contains the following fields:

- *Team()*
- *ReceivedData()*

Constants for the *Team* attribute:

- *TeamYellow*
- *TeamBlue*

<a name="debug"></a>
## DebugToProtobuf Type

This type contains the following fields:

- *InitialPoses()*
- *FinalPoses()*
- *PathRobots()*

<a name="encoder"></a>
## Encoder

The encoder interface contains methods for serializing data.

- **Encode**() (*[]byte, error*)

<a name="decoder"></a>
## Decoder

The decoder interface contains methods for deserializing data.

- **Decode**() (**model.World, error*)

