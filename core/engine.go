package core

const (
	StateEngineContinue = 0
	StateEngineStop     = 1
)

// Engine handles the stages Setup(), Run() and Teardown() for all the systems.
type Engine interface {
	Run()
	Setup()
	Teardown()
}
