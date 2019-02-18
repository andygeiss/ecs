package ecs

// engine is simple a composition of an EntityManager and a SystemManager.
// It handles the stages Setup(), Run() and Teardown() for all the systems.
type engine struct {
	entityManager *EntityManager
	systemManager *SystemManager
}

// NewEngine creates a new Engine and returns its address.
func NewEngine(entityManager *EntityManager, systemManager *SystemManager) *engine {
	return &engine{
		entityManager: entityManager,
		systemManager: systemManager,
	}
}

// Run calls the Process() method for each System
// until ShouldEngineStop is set to true.
func (e *engine) Run() {
	for _, system := range e.systemManager.Systems() {
		system.Process(e.entityManager)
	}
}

// Setup calls the Setup() method for each System
// and initializes ShouldEngineStop and ShouldEnginePause with false.
func (e *engine) Setup() {
	for _, sys := range e.systemManager.Systems() {
		sys.Setup()
	}
}

// Teardown calls the Teardown(= method for each System.
func (e *engine) Teardown() {
	for _, sys := range e.systemManager.Systems() {
		sys.Teardown()
	}
}
