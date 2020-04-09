package evaluation

import "github.com/furgbol/ai/model"

// Evaluation - type that stores the evaluation of the model world
type Evaluation struct {
	world model.World
}

// DummyEvaluator - type that stores the struct evaluation 
type DummyEvaluator struct {
	eval Evaluation
}

// DummyEvaluator creates an instance of DummyEvaluator  
func NewDummyEvaluator() *DummyEvaluator {
	return &DummyEvaluator{}
}

// Evaluate is the function that return the evaluation with model world
func (duEvaluator *DummyEvaluator) Evaluate(world model.World) Evaluation {
	duEvaluator.Evaluation.World = world
	return duEvaluator.Evaluation
}