package evaluation

import "github.com/furgbol/ai/model"

// DummyEvaluator - type that stores the struct evaluation 
type DummyEvaluator struct {
	evaluation Evaluation
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