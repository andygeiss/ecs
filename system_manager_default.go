package ecs

var _ SystemManager = (*DefaultSystemManager)(nil)

// DefaultSystemManager keeps the systems in the order they were added. Its
// zero value is ready to use; NewSystemManager exists for symmetry with the
// other constructors.
type DefaultSystemManager struct {
	systems []System
}

// Add appends systems to this DefaultSystemManager. A nil system is a
// programmer error, and it panics later — in Setup or Process — not here.
func (m *DefaultSystemManager) Add(systems ...System) {
	m.systems = append(m.systems, systems...)
}

// Systems returns the systems in the order they were added. The slice is the
// manager's own, not a copy, so changing it changes the manager.
func (m *DefaultSystemManager) Systems() []System {
	return m.systems
}

// NewSystemManager creates a new DefaultSystemManager and returns its address.
func NewSystemManager() *DefaultSystemManager {
	return &DefaultSystemManager{
		systems: []System{},
	}
}
