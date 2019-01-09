package ecs

// Entity is the composition of one or more components.
type Entity interface {
	Components() (components []Component)
	Get(name string) (component Component)
	// The ID() method must be implemented, because the EntityManager
	// uses it to get a specific entity by id.
	ID() (id string)
}

// EntityManager handles the access to each entity.
type EntityManager struct {
	entities map[string]Entity
}

// NewEntityManager creates a new EntityManager and returns its address.
func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: map[string]Entity{},
	}
}

// Add creates new map entries for the entities by id.
func (m *EntityManager) Add(entities ...Entity) {
	for _, entity := range entities {
		m.entities[entity.ID()] = entity
	}
}

// Entities returns all the mapped entities.
func (m *EntityManager) Entities() (entities []Entity) {
	for _, e := range m.entities {
		entities = append(entities, e)
	}
	return
}

// FilterBy returns the mapped entities, which components name matched.
func (m *EntityManager) FilterBy(components ...string) (entities []Entity) {
	for _, e := range m.entities {
		count := 0
		wanted := len(components)
		// Simply increase the count if the component could be found.
		for _, name := range components {
			component := e.Get(name)
			if component == nil {
				continue
			}
			count++
		}
		// Add the entity to the filter list,
		// if all components are found.
		if count == wanted {
			entities = append(entities, e)
		}
	}
	return
}

// Get a specific entity by id.
func (m *EntityManager) Get(id string) (entity Entity) {
	for _, e := range m.entities {
		if e.ID() == id {
			return e
		}
	}
	return
}

// Remove a specific entity from the map.
func (m *EntityManager) Remove(entity Entity) {
	delete(m.entities, entity.ID())
}
