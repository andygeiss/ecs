package ecs

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

// Entity is simply a composition of one or more Components with an Id.
type Entity struct {
	Components []Component `json:"components"`
	Id         string      `json:"id"`
	Masked     uint64      `json:"masked"`
}

// Add a component.
func (e *Entity) Add(cn ...Component) {
	for _, c := range cn {
		if e.Masked&c.Mask() == c.Mask() {
			continue
		}
		e.Components = append(e.Components, c)
		e.Masked = maskSlice(e.Components)
	}
}

// Get a component by its bitmask.
func (e *Entity) Get(mask uint64) Component {
	for _, c := range e.Components {
		if c.Mask() == mask {
			return c
		}
	}
	return nil
}

// ID ...
func (e *Entity) ID() string {
	return e.Id
}

// Mask returns a pre-calculated maskSlice to identify the Components.
func (e *Entity) Mask() uint64 {
	return e.Masked
}

// Remove a component by using its maskSlice.
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

// NewEntity creates a new entity and pre-calculates the component maskSlice.
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
