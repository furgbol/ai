# Parse Package

This package provides types and methods for converting standard data to simulator data, it also serializes and deserializes that data.

# Contents

1. [Parse Command](#command)
2. [Parse Control](#control)
3. [Parse State](#state)
4. [Parse Standard Simulator](#psdsmr)

<a name="command"></a>
## Parse Command

Serializes the simulator commands.

### Parse Command Methods

- **SerializeToSimulatorCommand**(*commands []control.SimulatorCommand*) (*[]byte, error*)


<a name="control"></a>
## Parse Control

Deserializes the simulator control data.

### Parse Control Methods

- **ControlFromSimulator**(*ctrl []byte*) (**control.Control, error*)


<a name="state"></a>
## Parse State

Deserializes the simulator state data and returns world model.

### Parse State Constants

- *TeamYellow*
- *TeamBlue*

### Parse State Methods

- **StateFromSimulator**(*state []byte, team int*) (**model.World, error*)


<a name="#psdsmr"></a>
## Parse Standard Simulator

This module transform the commands passed into commands for the simulator.

### Parse Standard Simulator Methods

- **EncodeToCommands**(cmdRepo control.StandardCommand) (*control.SimulatorCommand)
- **NewCommandsToSimulator**(NumberOfRobots int) (CommandsToSimulator)