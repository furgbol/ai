# Command Package

The command package contains modules that provide control over the properties of the robots.

## Contents

- [Standard Command Type](#stdcmd)
- [Simulator Command Type](#slrcmd)

<a name="stdcmd"></a>

### Standard Command Type

This type models the standard system command, containing the following fields:

- *RobotID()*
- *LinearVelocity()*
- *AngularVelocity()*

<a name="slrcmd"></a>

### Simulator Command Type

This type models the simulator command, having the following fields:

- *LeftWheelVelocity()*
- *RightWheelVelocity()*
- *FinalPosition()*