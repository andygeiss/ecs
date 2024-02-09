package core

// SystemManager handles the access to each system.
type SystemManager interface {
	// Add systems to the this SystemManager.
	Add(systems ...System)
	// Systems returns internally stored systems.
	Systems() []System
}
