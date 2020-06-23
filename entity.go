package ecs

// Component contains only the data (no behaviour at all).
// The Name() method must be implemented, because the EntityManager
// uses it to filter the entities by component names.
type Component interface {
	Mask() uint64
}

// Entity is simply a composition of one or more components with an id.
type Entity struct {
	Components []Component
	Id         string
	mask       uint64
}

// Get a specific component by a bitmask.
func (e *Entity) Get(mask uint64) Component {
	for _, c := range e.Components {
		if c.Mask() == mask {
			return c
		}
	}
	return nil
}

// Mask returns a pre-calculated mask to identify the components.
func (e *Entity) Mask() uint64 {
	return e.mask
}

// NewEntity creates a new entity and pre-calculates the component mask.
func NewEntity(id string, components []Component) *Entity {
	mask := uint64(0)
	for _, c := range components {
		mask = mask | c.Mask()
	}
	return &Entity{
		Components: components,
		Id:         id,
		mask:       mask,
	}
}
