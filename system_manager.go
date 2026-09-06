package ecs

// SystemManager handles the access to each system.
type SystemManager interface {
	// Add appends systems to this SystemManager.
	Add(systems ...System)
	// Systems returns internally stored systems.
	Systems() []System
}
