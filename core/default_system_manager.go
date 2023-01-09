package core

// defaultSystemManager
type defaultSystemManager struct {
	systems []System
}

// Add systems to the defaultSystemManager.
func (m *defaultSystemManager) Add(systems ...System) {
	for _, system := range systems {
		m.systems = append(m.systems, system)
	}
}

// Systems returns the system, which are internally stored.
func (m *defaultSystemManager) Systems() []System {
	return m.systems
}

// NewSystemManager creates a new defaultSystemManager and returns its address.
func NewSystemManager() SystemManager {
	return &defaultSystemManager{
		systems: []System{},
	}
}
