package ecs

// EntityManager handles the access to each entity.
type EntityManager struct {
	entities []*Entity
	index []int
}

// NewEntityManager creates a new EntityManager and returns its address.
func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: []*Entity{},
	}
}

// Add entries to the manager.
func (m *EntityManager) Add(entities ...*Entity) {
	for _, entity := range entities {
		m.entities = append(m.entities, entity)
	}
}

// Entities returns all the entities.
func (m *EntityManager) Entities() (entities []*Entity) {
	return m.entities
}

// FilterBy returns the mapped entities, which components name matched.
func (m *EntityManager) FilterBy(components ...string) (entities []*Entity) {
	// Allocate the worst-case amount of memory (all entities needed).
	entities = make([]*Entity, len(m.entities))
	index := 0
	for _, e := range m.entities {
		count := 0
		wanted := len(components)
		// Simply increase the count if the component could be found.
		for _, name := range components {
			for _, c := range e.Components {
				if c.Name() == name {
					count++
				}
			}
		}
		// Add the entity to the filter list, if all components are found.
		if count == wanted {
			// Direct access
			entities[index] = e
			// entities = append(entities, e)
		}
	}
	// Return only the needed slice.
	return entities[:index]
	// return
}

// Get a specific entity by id.
func (m *EntityManager) Get(id string) (entity *Entity) {
	for _, e := range m.entities {
		if e.Id == id {
			return e
		}
	}
	return
}

// Remove a specific entity.
func (m *EntityManager) Remove(entity *Entity) {
	for i, e := range m.entities {
		if e.Id == entity.Id {
			copy(m.entities[i:], m.entities[i+1:])
			m.entities[len(m.entities)-1] = nil
			m.entities = m.entities[:len(m.entities)-1]
			break
		}
	}
}
