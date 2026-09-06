package ecs

// DefaultSystemManager keeps the systems in the order they were added. Its
// zero value is ready to use; NewSystemManager exists for symmetry with the
// other constructors.
type DefaultSystemManager struct {
	systems []System
}

// Add appends systems to this DefaultSystemManager.
func (m *DefaultSystemManager) Add(systems ...System) {
	m.systems = append(m.systems, systems...)
}

// Systems returns the systems, which are internally stored.
func (m *DefaultSystemManager) Systems() []System {
	return m.systems
}

// NewSystemManager creates a new DefaultSystemManager and returns its address.
func NewSystemManager() *DefaultSystemManager {
	return &DefaultSystemManager{
		systems: []System{},
	}
}
