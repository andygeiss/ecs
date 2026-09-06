package ecs

var _ EntityManager = (*DefaultEntityManager)(nil)

// DefaultEntityManager keeps the entities in a slice and scans it on every
// lookup. Its zero value is ready to use; NewEntityManager exists for
// symmetry with the other constructors.
//
// It is not safe for concurrent use. A system that fans work out to goroutines
// guards the manager itself.
type DefaultEntityManager struct {
	entities []*Entity
}

// Add appends entities to the manager. A nil entity is a programmer error, and
// it panics later — in Get or one of the filters — not here.
func (m *DefaultEntityManager) Add(entities ...*Entity) {
	m.entities = append(m.entities, entities...)
}

// Entities returns all the entities. The slice is the manager's own, not a
// copy, so changing it changes the manager.
func (m *DefaultEntityManager) Entities() (entities []*Entity) {
	return m.entities
}

// FilterByMask returns every entity that carries all the components in mask.
func (m *DefaultEntityManager) FilterByMask(mask uint64) (entities []*Entity) {
	// Allocate the worst-case amount of memory (all entities needed).
	entities = make([]*Entity, len(m.entities))
	index := 0
	for _, e := range m.entities {
		observed := e.Mask()
		if observed&mask == mask {
			entities[index] = e
			index++
		}
	}
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
		if matched == len(names) {
			entities[index] = e
			index++
		}
	}
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

// Remove deletes the entity with the same Id. It does nothing when entity is
// nil, or when the manager does not hold one — so Remove(Get(id)) is safe for
// an id that is not there.
func (m *DefaultEntityManager) Remove(entity *Entity) {
	if entity == nil {
		return
	}
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
