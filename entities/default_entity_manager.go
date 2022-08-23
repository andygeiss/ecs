package entities

import (
	"github.com/andygeiss/ecs/core"
)

type defaultEntityManager struct {
	entities []*core.Entity
}

// Add entries to the manager.
func (m *defaultEntityManager) Add(entities ...*core.Entity) {
	m.entities = append(m.entities, entities...)
}

// Entities returns all the entities.
func (m *defaultEntityManager) Entities() (entities []*core.Entity) {
	return m.entities
}

// FilterByMask returns the mapped entities, which Components mask matched.
func (m *defaultEntityManager) FilterByMask(mask uint64) (entities []*core.Entity) {
	// Allocate the worst-case amount of memory (all entities needed).
	entities = make([]*core.Entity, len(m.entities))
	index := 0
	for _, e := range m.entities {
		// Use the pre-calculated Components maskSlice.
		observed := e.Mask()
		// Add the entity to the filter list, if all Components are found.
		if observed&mask == mask {
			// Direct access
			entities[index] = e
			index++
		}
	}
	// Return only the needed slice.
	return entities[:index]
}

// FilterByNames returns the mapped entities, which Components names matched.
func (m *defaultEntityManager) FilterByNames(names ...string) (entities []*core.Entity) {
	// Allocate the worst-case amount of memory (all entities needed).
	entities = make([]*core.Entity, len(m.entities))
	index := 0
	for _, e := range m.entities {
		// Each component should match
		matched := 0
		for _, name := range names {
			for _, c := range e.Components {
				switch v := c.(type) {
				case core.ComponentWithName:
					if v.Name() == name {
						matched++
					}
				}
			}
		}
		// Add the entity to the filter list, if all Components are found.
		if matched == len(names) {
			// Direct access
			entities[index] = e
			index++
		}
	}
	// Return only the needed slice.
	return entities[:index]
}

// Get a specific entity by Id.
func (m *defaultEntityManager) Get(id string) (entity *core.Entity) {
	for _, e := range m.entities {
		if e.ID() == id {
			return e
		}
	}
	return
}

// Remove a specific entity.
func (m *defaultEntityManager) Remove(entity *core.Entity) {
	for i, e := range m.entities {
		if e.Id == entity.Id {
			copy(m.entities[i:], m.entities[i+1:])
			m.entities[len(m.entities)-1] = nil
			m.entities = m.entities[:len(m.entities)-1]
			break
		}
	}
}

// NewEntityManager creates a new defaultEntityManager and returns its address.
func NewEntityManager() core.EntityManager {
	return &defaultEntityManager{
		entities: []*core.Entity{},
	}
}
