package core

// SystemManager handles the access to each system.
type SystemManager interface {
	Add(systems ...System)
	Systems() []System
}
