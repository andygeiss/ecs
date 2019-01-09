package ecs

// Component contains only the data (no behaviour at all).
// The Name() method must be implemented, because the EntityManager
// uses it to filter the entities by component names.
type Component interface {
	Name() (name string)
}
