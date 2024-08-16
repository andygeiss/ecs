package core

// Component contains only the data (no behaviour at all).
// The Name() method must be implemented, because the defaultEntityManager
// uses it to filter the entities by component names.
type Component interface {
	Mask() uint64
}

// ComponentWithName is used by FilterByNames to enable more than 64 Components (if needed).
type ComponentWithName interface {
	Component
	Name() string
}
