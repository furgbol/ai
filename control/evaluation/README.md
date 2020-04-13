# Evaluation Package

This package contains the types and methods needed to evaluate the game environment.

## Contents

- [Evaluator Interface](#evaluator)
- [Evaluation Type](#evaluation)
- [Dummy Evaluator Type](#DummyEvaluator)

<a name="evaluator"></a>

### Evaluator Interface

This interface returns the evaluator containing the model word, with the following method:

- *Evaluator(world model.World)*

<a name="evaluation"></a>

### Evaluation Type

- **world** *model.World*: stores the actual instance of the world.

<a name="DummyEvaluator"></a>

### Dummy Evaluator Type

This type implements the Evaluator interface. It contains the following instace.

- **evaluation** *Evaluation*: stores the actual intance of the evaluator, consequently the world.