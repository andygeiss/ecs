package ecs

// engine ...
type engine struct {
	entityManager *EntityManager
	systems       []System
}

// NewEngine ...
func NewEngine() *engine {
	return &engine{
		entityManager: NewEntityManager(),
	}
}

// Run ...
func (g *engine) Run() {
	for _, system := range g.systems {
		system.Process(g.entityManager)
	}
}

// Setup ...
func (g *engine) Setup() {
	for _, sys := range g.systems {
		sys.Setup()
	}
}

// Teardown ...
func (g *engine) Teardown() {
	for _, sys := range g.systems {
		sys.Teardown()
	}
}
