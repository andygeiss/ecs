package ecs

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

// defaultEngine is simple a composition of an defaultEntityManager and a defaultSystemManager.
type defaultEngine struct {
	entityManager EntityManager
	systemManager SystemManager
}

// Run calls the Process() method for each System
// until ShouldEngineStop is set to true.
func (e *defaultEngine) Run() {
	shouldStop := false
	for !shouldStop {
		for _, system := range e.systemManager.Systems() {
			state := system.Process(e.entityManager)
			if state == StateEngineStop {
				shouldStop = true
				break
			}
		}
	}
}

// Setup calls the Setup() method for each System
// and initializes ShouldEngineStop and ShouldEnginePause with false.
func (e *defaultEngine) Setup() {
	for _, sys := range e.systemManager.Systems() {
		sys.Setup()
	}
}

// Teardown calls the Teardown() method for each System.
func (e *defaultEngine) Teardown() {
	for _, sys := range e.systemManager.Systems() {
		sys.Teardown()
	}
}

// NewDefaultEngine creates a new Engine and returns its address.
func NewDefaultEngine(entityManager EntityManager, systemManager SystemManager) Engine {
	return &defaultEngine{
		entityManager: entityManager,
		systemManager: systemManager,
	}
}
