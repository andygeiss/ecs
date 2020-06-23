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
