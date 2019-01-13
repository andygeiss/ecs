package ecs

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
