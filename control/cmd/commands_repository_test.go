package cmd

import "testing"

func TestNewCommandsRepository(t *testing.T) {
	cmdRepo := NewCommandsRepository(3)
	if len(*cmdRepo) != 3 {
		t.Errorf("Error on creating instance: Length expected: 3, got: %v", len(*cmdRepo))
	}
}
