package model

import "testing"

func TestNewTeam(t *testing.T) {
	team := NewTeam(3)
	if len(team) != 3 {
		t.Errorf("Team creation failed: expected team size: 3, got: %v", len(team))
	}
}

func TestNewWorld(t *testing.T) {
	ball := Ball{
		Position2D: Position2D{
			X: 1.4,
			Y: 2.1,
		},
	}

	team := NewTeam(2)
	team[0] = Robot{
		ID: 1,
		CurrentPosition: Pose{
			Position2D: Position2D{
				X: 2,
				Y: 3,
			},
			Orientation: 3.5,
		},
		TargetPosition: Pose{
			Position2D: Position2D{
				X: 3,
				Y: 2,
			},
			Orientation: 5.3,
		},
	}

	world := NewWorld(ball, team)

	if world.Ball.X != 1.4 {
		t.Errorf("World creation failed: expected ball x: 1.4, got: %v", world.Ball.X)
	}
	if world.Ball.Y != 2.1 {
		t.Errorf("World creation failed: expected ball y: 2.1, got: %v", world.Ball.Y)
	}
	if world.Team[0].ID != 1 {
		t.Errorf("World creation failed: expected robot id: 1, got: %v", world.Team[0].ID)
	}
	if world.Team[0].CurrentPosition.X != 2 {
		t.Errorf("World creation failed: expected robot current position x: 2, got: %v", world.Team[0].CurrentPosition.X)
	}
	if world.Team[0].CurrentPosition.Y != 3 {
		t.Errorf("World creation failed: expected robot current position y: 3, got: %v", world.Team[0].CurrentPosition.Y)
	}
	if world.Team[0].CurrentPosition.Orientation != 3.5 {
		t.Errorf("World creation failed: expected robot current orientation: 3.5, got: %v", world.Team[0].CurrentPosition.Orientation)
	}
	if world.Team[0].TargetPosition.X != 3 {
		t.Errorf("World creation failed: expected robot target position y: 3, got: %v", world.Team[0].TargetPosition.X)
	}
	if world.Team[0].TargetPosition.Y != 2 {
		t.Errorf("World creation failed: expected robot target position y: 2, got: %v", world.Team[0].TargetPosition.Y)
	}
	if world.Team[0].TargetPosition.Orientation != 5.3 {
		t.Errorf("World creation failed: expected robot target orientation: 5.3, got: %v", world.Team[0].TargetPosition.Orientation)
	}
}
