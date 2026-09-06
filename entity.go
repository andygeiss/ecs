package ecs

// Entity is simply a composition of one or more Components with an Id.
//
// Components and Masked have to agree. Change Components by hand and Masked
// goes stale, which makes FilterByMask skip the entity. Add and Remove keep the
// two in step.
//
// An Entity marshals to JSON, but it does not unmarshal back: Components holds
// an interface, so the decoder cannot know which concrete type to build. A
// consumer that persists entities writes its own UnmarshalJSON, which knows its
// own component types.
type Entity struct {
	Components []Component `json:"components"`
	Id         string      `json:"id"`
	Masked     uint64      `json:"masked"`
}

// Add appends the components the entity does not already carry, and extends its
// mask. A component is identified by its mask, the way Get and Remove address
// one, so a second component with a mask the entity already carries is skipped.
func (e *Entity) Add(cn ...Component) {
	for _, c := range cn {
		if e.Get(c.Mask()) != nil {
			continue
		}
		e.Components = append(e.Components, c)
		e.Masked |= c.Mask()
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

// NewEntity creates a new entity and pre-calculates its component mask. It
// takes the slice as given, so unlike Add it does not skip a second component
// with a mask another one already carries.
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
