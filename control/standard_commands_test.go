package control

import (
	"testing"
)

func TestRobotID(t *testing.T) {
	var stdCmd StandardCommand
	stdCmd.SetRobotID(2)
	robotID := stdCmd.RobotID()
	if robotID != 2 {
		t.Errorf("Error on getting value: expecting: 2, got: %v", robotID)
	}
}

func TestLinearVelocity(t *testing.T) {
	var stdCmd StandardCommand
	stdCmd.SetLinearVelocity(12.5)
	linearVel := stdCmd.LinearVelocity()
	if linearVel != 12.5 {
		t.Errorf("Error on getting value: expecting: 12.5, got: %v", linearVel)
	}
}

func TestAngularVelocity(t *testing.T) {
	var stdCmd StandardCommand
	stdCmd.SetAngularVelocity(2.5)
	angularVel := stdCmd.AngularVelocity()
	if angularVel != 2.5 {
		t.Errorf("Error on getting value: expecting: 2.5, got: %v", angularVel)
	}
}

func TestSetRobotID(t *testing.T) {
	var stdCmd StandardCommand
	stdCmd.SetRobotID(2)
	robotID := stdCmd.RobotID()
	if robotID != 2 {
		t.Errorf("Error on setting value: expecting: 2, got: %v", robotID)
	}
}

func TestSetLinearVelocity(t *testing.T) {
	var stdCmd StandardCommand
	stdCmd.SetLinearVelocity(12.5)
	linearVel := stdCmd.LinearVelocity()
	if linearVel != 12.5 {
		t.Errorf("Error on setting value: expecting: 12.5, got: %v", linearVel)
	}
}

func TestSetAngularVelocity(t *testing.T) {
	var stdCmd StandardCommand
	stdCmd.SetAngularVelocity(2.5)
	angularVel := stdCmd.AngularVelocity()
	if angularVel != 2.5 {
		t.Errorf("Error on setting value: expecting: 2.5, got: %v", angularVel)
	}
}
