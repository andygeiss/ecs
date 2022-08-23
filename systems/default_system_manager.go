package systems

import (
	"github.com/andygeiss/ecs/core"
)

// defaultSystemManager
type defaultSystemManager struct {
	systems []core.System
}

// Add systems to the defaultSystemManager.
func (m *defaultSystemManager) Add(systems ...core.System) {
	for _, system := range systems {
		m.systems = append(m.systems, system)
	}
}

// Systems returns the system, which are internally stored.
func (m *defaultSystemManager) Systems() []core.System {
	return m.systems
}

// NewSystemManager creates a new defaultSystemManager and returns its address.
func NewSystemManager() core.SystemManager {
	return &defaultSystemManager{
		systems: []core.System{},
	}
}
