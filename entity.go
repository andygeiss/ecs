package ecs

// Entity ...
type Entity interface {
	Components() (components []Component)
	ID() (id string)
}

// EntityManager ...
type EntityManager struct {
	entities map[string]Entity
}

// NewEntityManager ...
func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: map[string]Entity{},
	}
}

// Add ...
func (m *EntityManager) Add(entities ...Entity) {
	for _, entity := range entities {
		m.entities[entity.ID()] = entity
	}
}

// Entities ...
func (m *EntityManager) Entities() (entities []Entity) {
	for _, e := range m.entities {
		entities = append(entities, e)
	}
	return
}

// Remove ...
func (m *EntityManager) Remove(entity Entity) {
	delete(m.entities, entity.ID())
}
