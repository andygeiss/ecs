package ecs

// EntityManager handles the access to each entity.
type EntityManager struct {
	entities []*Entity
	index    []int
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
func (m *EntityManager) FilterByMask(mask uint64) (entities []*Entity) {
	// Allocate the worst-case amount of memory (all entities needed).
	entities = make([]*Entity, len(m.entities))
	index := 0
	for _, e := range m.entities {
		// Use the pre-calculated components maskSlice.
		observed := e.Mask()
		// Add the entity to the filter list, if all components are found.
		if observed&mask == mask {
			// Direct access
			entities[index] = e
			index++
		}
	}
	// Return only the needed slice.
	return entities[:index]
}

// Get a specific entity by id.
func (m *EntityManager) Get(id string) (entity *Entity) {
	for _, e := range m.entities {
		if e.id == id {
			return e
		}
	}
	return
}

// Remove a specific entity.
func (m *EntityManager) Remove(entity *Entity) {
	for i, e := range m.entities {
		if e.id == entity.id {
			copy(m.entities[i:], m.entities[i+1:])
			m.entities[len(m.entities)-1] = nil
			m.entities = m.entities[:len(m.entities)-1]
			break
		}
	}
}
