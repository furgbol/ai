# Model Package

The model package contains the modeling of the game world.

## Contents

### Types

- [Ball](#ball)
- [Pose](#pose)
- [Position2D](#position2d)
- [Robot](#robot)
- [Team](#team)
- [World](#world)

<a name="ball"></a>

#### Ball

The Ball type models the game ball, containing a **Position2D** embbeded.

<a name="pose"></a>

#### Pose

The Pose type models the robot pose into the game field, containg a **Position2D** embbeded and an **Orientation** that is a *float64*.

<a name="position2d"></a>

#### Position2D

The Position2D type models a position into the game field, containing the following fields:

- **X**: *float64*
- **Y**: *float64*

<a name="robot"></a>

#### Robot

The Robot type models a robot into the game environment, containing the following fields:

- **ID**: *int*
- **Role**: *string*
- **CurrentPosition**: *Pose*
- **TargetPosition**: *Pose*

<a name="team"></a>

#### Team

The Team type models the team. This type is an alias to a slice *[]Robot*.

To create an instance of this type, the **NewTeam**(*numberOfRobots*: int) function is provided.

<a name="world"></a>

#### World

The World type models the game environment itself. This type contains a **Ball** and a **Team**.

To create an instance of this type, the **NewWorld**(*ball*: Ball, *team*: Team) function is provided.
