package ecs

// Component contains only the data (no behaviour at all).
// The Name() method must be implemented, because the EntityManager
// uses it to filter the entities by component names.
type Component interface {
	Name() (name string)
}

// Entity is simply a composition of one or more components with an id.
type Entity struct {
	Components []Component
	Id string
}

// Get a specific component by name.
func (e *Entity) Get(name string) Component {
	for _, c := range e.Components {
		if c.Name() == name {
			return c
		}
	}
	return nil
}
