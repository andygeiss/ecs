package ecs

// SystemManager handles the access to each system.
type SystemManager interface {
	// Add appends systems to this SystemManager.
	Add(systems ...System)
	// Systems returns the systems in the order they were added. The slice may
	// be the manager's own, so changing it may change the manager.
	Systems() []System
}
