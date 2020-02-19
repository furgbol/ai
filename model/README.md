# Model Package

The model package contains the modeling of the game world.

## Contents

- [Ball Type](#ball)
- [Pose Type](#pose)
- [Position2D Type](#position2d)
- [Robot Type](#robot)
- [Team Type](#team)
- [World Type](#world)

<a name="ball"></a>

### Ball Type

The Ball type models the game ball, containing a **Position2D** embbeded.

<a name="pose"></a>

### Pose Type

The Pose type models the robot pose into the game field, containg a **Position2D** embbeded and an **Orientation** that is a *float64*.

<a name="position2d"></a>

### Position2D Type

The Position2D type models a position into the game field, containing the following fields:

- **X**: *float64*
- **Y**: *float64*

<a name="robot"></a>

### Robot Type

The Robot type models a robot into the game environment, containing the following fields:

- **ID**: *int*
- **Role**: *string*
- **CurrentPosition**: *Pose*
- **TargetPosition**: *Pose*

<a name="team"></a>

### Team Type

The Team type models the team. This type is an alias to a slice *[]Robot*.

To create an instance of this type, the **NewTeam**(*numberOfRobots*: int) function is provided.

<a name="world"></a>

### World Type

The World type models the game environment itself. This type contains a **Ball** and a **Team**.

To create an instance of this type, the **NewWorld**(*ball*: Ball, *team*: Team) function is provided.
