# Navigation Package

This package contains all navigation and path planning algorithms used by the system.

## Contents

- [Path Type](#path): Alias to simulate the position vector.
- [Casteljau Path Planner](#casteljau): Path planning algorithm based on the Casteljau's algorithm. This algorithm uses Bézier curves.

<a name="path"></a>

 ### Path Type

 This type is an alias to a *[]Position2D*.

 <a name="casteljau"></a>

 ### Casteljau Path Planner

 This type implements the PathPlanner interface. It contains the following fields.

- **NumberOfPathPoints** *int*: Stores the number of points the path contains
- **NumberOfUsedPoints** *int*: Stores the number of points that will be used from the generated path
- **DistanceFactor** *int*: Stores the factor by which the distance between points must to be divided

To create an instance of this type, use the **NewCasteljauPathPlanner**(*numberOfPathPoints, numberOfUsedPoints, distanceFactor int*) method.