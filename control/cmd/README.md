# Commands Package

This package contains the manipulation and modeling of all fundamental system commands.

## Contents

- [Standard Command Type](#stdcmd)
- [Commands Repository Type](#cmdrepo)
- [Simulator Command Type](#simcmd)

<a name="stdcmd"></a>

### Standard Command Type

This type models the standard system command, containing the following fields:

- *RobotID()*
- *LinearVelocity()*
- *AngularVelocity()*

<a name="cmdrepo"></a>

### Commands Repository Type

This type is an alias to a slice: *[]StandardCommand*

To create a new Commands Repository, the **NewCommandsRepository**(*numberOfRobots* int) function is provided.

<a name="simcmd"></a>

### Simulator Command Type

This type models the simulator command, having the following fields:

- *RobotID()*
- *LeftWheelVelocity()*
- *RightWheelVelocity()*
- *FinalPosition()*