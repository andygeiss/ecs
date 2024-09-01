package ecs

// Component contains only the data (no behaviour at all).
type Component interface {
	Mask() uint64
}

// ComponentWithName is used by FilterByNames to enable more than 64 Components (if needed).
type ComponentWithName interface {
	Component
	Name() string
}
