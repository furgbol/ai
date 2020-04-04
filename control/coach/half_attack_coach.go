package coach

import "github.com/furgbol/ai/control/evaluation"

const (
	// Position of robot in goal
	Goalkeeper = "inGoal"
	// Position of robot in attack
	Attacker = "inAttack"
	// Position of robot in defense
	Defender = "inDefense"
)

// HalfAttackCoach - This type implement the Coach interface.
type HalfAttackCoach struct {
	Index map[int]string
}

// NewHalfAttackCoach creates an instance of a Half Attack Coach  
func NewHalfAttackCoach(index map[int]string) *HalfAttackCoach {
	return &HalfAttackCoach{
		Index: index,
	} 
}

// GetRoles is the function that return the role of robot
func (coach HalfAttackCoach) GetRoles(evaluation evaluation.Evaluation) map[int]string {
	var NewCoach map[int]string
	Robots := evaluation.word.Team 
	for i := 0; i < len(Robots); i+++ {
		robot := Robots[i]
		NewCoach[robots.ID], _ := coach[i]  
	}
	return NewCoach
}