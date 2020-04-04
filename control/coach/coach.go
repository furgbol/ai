package coach

import "github.com/furgbol/ai/control/evaluation"

// Coach - Inteface to get roles
type Coach interface {
	GetRoles(evaluation evaluation.Evaluation) map[int]string
}