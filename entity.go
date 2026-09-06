package ecs

// Entity is simply a composition of one or more Components with an Id.
type Entity struct {
	Components []Component `json:"components"`
	Id         string      `json:"id"`
	Masked     uint64      `json:"masked"`
}

// Add appends the components the entity does not already carry, and
// recalculates its mask.
func (e *Entity) Add(cn ...Component) {
	for _, c := range cn {
		if e.Masked&c.Mask() == c.Mask() {
			continue
		}
		e.Components = append(e.Components, c)
		e.Masked = maskSlice(e.Components)
	}
}

// Get returns the component with the given mask, or nil when the entity
// does not carry one.
func (e *Entity) Get(mask uint64) Component {
	for _, c := range e.Components {
		if c.Mask() == mask {
			return c
		}
	}
	return nil
}

// Mask returns the pre-calculated bitmask of the entity's components.
func (e *Entity) Mask() uint64 {
	return e.Masked
}

// Remove deletes the component with the given mask, and recalculates the
// entity's mask. It does nothing when the entity does not carry one.
func (e *Entity) Remove(mask uint64) {
	modified := false
	for i, c := range e.Components {
		if c.Mask() == mask {
			copy(e.Components[i:], e.Components[i+1:])
			e.Components[len(e.Components)-1] = nil
			e.Components = e.Components[:len(e.Components)-1]
			modified = true
			break
		}
	}
	if modified {
		e.Masked = maskSlice(e.Components)
	}
}

// NewEntity creates a new entity and pre-calculates its component mask.
func NewEntity(id string, components []Component) *Entity {
	return &Entity{
		Components: components,
		Id:         id,
		Masked:     maskSlice(components),
	}
}

func maskSlice(components []Component) uint64 {
	mask := uint64(0)
	for _, c := range components {
		mask = mask | c.Mask()
	}
	return mask
}
