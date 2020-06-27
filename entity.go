package ecs

// Component contains only the data (no behaviour at all).
// The Name() method must be implemented, because the EntityManager
// uses it to filter the entities by component names.
type Component interface {
	Mask() uint64
}

// Entity is simply a composition of one or more components with an id.
type Entity struct {
	components []Component
	id         string
	mask       uint64
}

// Add a component.
func (e *Entity) Add(cn ...Component) {
	for _, c := range cn {
		if e.mask&c.Mask() == c.Mask() {
			continue
		}
		e.components = append(e.components, c)
		e.mask = maskSlice(e.components)
	}
}

// Get a component by its bitmask.
func (e *Entity) Get(mask uint64) Component {
	for _, c := range e.components {
		if c.Mask() == mask {
			return c
		}
	}
	return nil
}

// ID ...
func (e *Entity) ID() string {
	return e.id
}

// Mask returns a pre-calculated maskSlice to identify the components.
func (e *Entity) Mask() uint64 {
	return e.mask
}

// Remove a component by using its maskSlice.
func (e *Entity) Remove(mask uint64) {
	modified := false
	for i, c := range e.components {
		if c.Mask() == mask {
			copy(e.components[i:], e.components[i+1:])
			e.components[len(e.components)-1] = nil
			e.components = e.components[:len(e.components)-1]
			modified = true
			break
		}
	}
	if modified {
		e.mask = maskSlice(e.components)
	}
}

// NewEntity creates a new entity and pre-calculates the component maskSlice.
func NewEntity(id string, components []Component) *Entity {
	return &Entity{
		components: components,
		id:         id,
		mask:       maskSlice(components),
	}
}

func maskSlice(components []Component) uint64 {
	mask := uint64(0)
	for _, c := range components {
		mask = mask | c.Mask()
	}
	return mask
}
