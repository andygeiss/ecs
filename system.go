package ecs

// System ...
type System interface {
	Setup()
	Process(entityManager *EntityManager)
	Teardown()
}

// SystemManager ...
type SystemManager struct {
	systems []System
}

// NewSystemManager ...
func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems: []System{},
	}
}

// Add ...
func (m *SystemManager) Add(systems ...System) {
	for _, system := range systems {
		m.systems = append(m.systems, system)
	}
}

// Systems ...
func (m *SystemManager) Systems() []System {
	return m.systems
}
