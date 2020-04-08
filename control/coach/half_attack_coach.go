package coach

import "github.com/furgbol/ai/control/evaluation"

// HalfAttackCoach - This type implement the Coach interface.
type HalfAttackCoach struct {
	Index map[int]string
}

// NewHalfAttackCoach creates an instance of a Half Attack Coach  
func NewHalfAttackCoach(index map[int]string) *HalfAttackCoach {
	index[1] = "Goalkeeper"
	index[2] = "Centerback"
	index[3] = "striker"
	return &HalfAttackCoach{
		Index: index,
	} 
}

// GetRoles is the function that return the role of robot
func (coach HalfAttackCoach) GetRoles(evaluation evaluation.Evaluation) map[int]string {
	return coach.Index
}