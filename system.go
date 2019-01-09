package ecs

// System implements the behaviour of an entity by modifying the state,
// which is stored in each component of the entity.
type System interface {
	Setup()
	Process(entityManager *EntityManager)
	Teardown()
}

// SystemManager handles the access to each system.
type SystemManager struct {
	systems []System
}

// NewSystemManager creates a new SystemManager and returns its address.
func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems: []System{},
	}
}

// Add systems to the SystemManager.
func (m *SystemManager) Add(systems ...System) {
	for _, system := range systems {
		m.systems = append(m.systems, system)
	}
}

// Systems returns the system, which are internally stored.
func (m *SystemManager) Systems() []System {
	return m.systems
}
