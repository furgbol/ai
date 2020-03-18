# Command Package

The command package contains modules that provide control over the properties of the robots.

## Contents

- [Standard Command Type](#stdcmd)
- [Commands Repository Type](#cmdrepo)
- [Simulator Command Type](#slrcmd)
- [Simulator Control Type](#ctrl)

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

<a name="slrcmd"></a>

### Simulator Command Type

This type models the simulator command, having the following fields:

- *RobotID()*
- *LeftWheelVelocity()*
- *RightWheelVelocity()*
- *FinalPosition()*

<a name="ctrl"></a>

### Simulator Control Type

This type contains the data structure of the simulator control, these are:

- *Paused()*
- *NewBallPose()*
- *NewRobotsBluePose()*
- *NewRobotsYellowPose()*
