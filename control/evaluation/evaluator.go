package evaluation

import "github.com/furgbol/ai/model"

// Evaluator - Inteface to Evaluate
type Evaluator interface {
	Evaluate(world model.World) *Evaluation
}