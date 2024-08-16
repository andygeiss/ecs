package ecs

const (
	StateEngineContinue = 0
	StateEngineStop     = 1
)

// Engine handles the stages Setup(), Run() and Teardown() for all the systems.
type Engine interface {
	// Run calls the Process() method for each System
	// until ShouldEngineStop is set to true.
	Run()
	// Setup calls the Setup() method for each System
	// and initializes ShouldEngineStop and ShouldEnginePause with false.
	Setup()
	// Teardown calls the Teardown() method for each System.
	Teardown()
	// Tick calls the Process() method for each System exactly once
	Tick()
}
