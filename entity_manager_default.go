package ecs

// DefaultEntityManager keeps the entities in a slice and scans it on every
// lookup. Its zero value is ready to use; NewEntityManager exists for
// symmetry with the other constructors.
type DefaultEntityManager struct {
	entities []*Entity
}

// Add appends entities to the manager.
func (m *DefaultEntityManager) Add(entities ...*Entity) {
	m.entities = append(m.entities, entities...)
}

// Entities returns all the entities.
func (m *DefaultEntityManager) Entities() (entities []*Entity) {
	return m.entities
}

// FilterByMask returns every entity that carries all the components in mask.
func (m *DefaultEntityManager) FilterByMask(mask uint64) (entities []*Entity) {
	// Allocate the worst-case amount of memory (all entities needed).
	entities = make([]*Entity, len(m.entities))
	index := 0
	for _, e := range m.entities {
		// Use the pre-calculated component mask.
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

// FilterByNames returns every entity that carries a component for each
// of the given names.
func (m *DefaultEntityManager) FilterByNames(names ...string) (entities []*Entity) {
	// Allocate the worst-case amount of memory (all entities needed).
	entities = make([]*Entity, len(m.entities))
	index := 0
	for _, e := range m.entities {
		// Count each name at most once. Two components sharing a name would
		// otherwise stand in for a name the entity does not carry at all.
		matched := 0
		for _, name := range names {
			for _, c := range e.Components {
				if v, ok := c.(ComponentWithName); ok && v.Name() == name {
					matched++
					break
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

// Get returns the entity with the given Id, or nil when there is none.
func (m *DefaultEntityManager) Get(id string) (entity *Entity) {
	for _, e := range m.entities {
		if e.Id == id {
			return e
		}
	}
	return
}

// Remove deletes the entity with the same Id. It does nothing when the
// manager does not hold one.
func (m *DefaultEntityManager) Remove(entity *Entity) {
	for i, e := range m.entities {
		if e.Id == entity.Id {
			copy(m.entities[i:], m.entities[i+1:])
			m.entities[len(m.entities)-1] = nil
			m.entities = m.entities[:len(m.entities)-1]
			break
		}
	}
}

// NewEntityManager creates a new DefaultEntityManager and returns its address.
func NewEntityManager() *DefaultEntityManager {
	return &DefaultEntityManager{
		entities: []*Entity{},
	}
}
