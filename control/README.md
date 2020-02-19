# Control Package

The control package contains modules that provide control of the environment, such as: commands, strategies, controllers, etc.

## Contents

- [Standard Command Type](#stdcmd)
- [Commands Repository Type](#cmdrepo)

<a name="stdcmd"></a>

### Standard Command Type

This type models the standard system command, containing the following fields:

- *RobotID()*
- *LinearVelocity()*
- *AngularVelocity()*

<a name="cmdrepo"></a>

### Commands Repository Type

This type is an alias to a slice: *[]StandardCommand*
