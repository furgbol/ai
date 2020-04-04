# Coach Package

The coach package contains the function of distributing the roles of the players.

## Contents

- [Coach Interface](#cch)
- [Half Attack Coach Type](#hact)

<a name="cch"></a>

### Coach Interface

This interface take the roles of robots, containing the following method:

- *GetRoles(evaluation evaluation.Evaluation)*

<a name="hact"></a>

### Half Attack Coach Type

This type models implements the Coach interface. It contains the following fields:

- **Index** *map[int]string*: Stores the dictionary of robots roles to robots id.

