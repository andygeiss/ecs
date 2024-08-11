package core

// defaultSystemManager
type defaultSystemManager struct {
	systems []System
}

// Add systems to the defaultSystemManager.
func (m *defaultSystemManager) Add(systems ...System) {
	m.systems = append(m.systems, systems...)
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
