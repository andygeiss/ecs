package ecs

// ShouldEngineStop ...
var ShouldEngineStop bool

// engine ...
type engine struct {
	entityManager *EntityManager
	systemManager *SystemManager
}

// NewEngine ...
func NewEngine(entityManager *EntityManager, systemManager *SystemManager) *engine {
	return &engine{
		entityManager: entityManager,
		systemManager: systemManager,
	}
}

// Run ...
func (g *engine) Run() {
	for !ShouldEngineStop {
		for _, system := range g.systemManager.Systems() {
			system.Process(g.entityManager)
		}
	}
}

// Setup ...
func (g *engine) Setup() {
	ShouldEngineStop = false
	for _, sys := range g.systemManager.Systems() {
		sys.Setup()
	}
}

// Teardown ...
func (g *engine) Teardown() {
	for _, sys := range g.systemManager.Systems() {
		sys.Teardown()
	}
}
