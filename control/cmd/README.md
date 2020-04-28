# Command Package

The command package contains modules that provide control over the properties of the robots.

## Contents

- [Standard Command Type](#stdcmd)
- [Simulator Command Type](#slrcmd)
- [Simulator Control Type](#ctrl)
- [Generator Command Type](#gnrcmd)
- [Generator Interface](#gnri)

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

<a name="ctrl"></a>

### Simulator Control Type

This type contains the data structure of the simulator control, these are:

- *Paused()*
- *NewBallPose()*
- *NewRobotsBluePose()*
- *NewRobotsYellowPose()*

<a name="gnrcmd"></a>

### Generator Command Type

This type models the commands to be generated, owning the following fields:

- *robotID()*
- *ConstAngularMult()*
- *ConstLinearMult()*
- *MaxAngularVelocity()*
- *MaxLinearVelocity()*

<a name="gnri"></a>

### Generator Interface

This interface calculates velocities to the robots, containing the following method:

- *Generate(currentPose, targetPose model.Pose)*