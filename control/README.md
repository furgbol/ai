# Control Package

The control package contains modules that provide control of the environment, such as: commands, strategies, controllers, etc.

## Contents

- [Standard Command Type](#stdcmd)
- [Commands Repository Type](#cmdrepo)

<a name="stdcmd"></a>

### Standard Command Type

This type models the standard system command, containing the following methods:

- *RobotID()*: Returns the robot id of the command.
- *LinearVelocity()*: Returns the linear velocity of the command.
- *AngularVelocity()*: Returns the angular velocity of the command.
- *SetRobotID()*: Sets the robot id of the command.
- *SetLinearVelocity()*: Sets the linear velocity of the command.
- *SetAngularVelocity()*: Sets the angular velocity of the command.

<a name="cmdrepo"></a>

### Commands Repository Type

This type is an alias to a slice: *[]StandardCommand*
