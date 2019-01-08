package ecs

// Entity ...
type Entity interface {
	Components() (components []Component)
	Get(name string) (component Component)
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

func (m *EntityManager) FilterBy(components ...string) (entities []Entity) {
	for _, e := range m.entities {
		count := 0
		wanted := len(components)
		for _, name := range components {
			component := e.Get(name)
			if component == nil {
				continue
			}
			count++
		}
		if count == wanted {
			entities = append(entities, e)
		}
	}
	return
}

// Get ...
func (m *EntityManager) Get(id string) (entity Entity) {
	for _, e := range m.entities {
		if e.ID() == id {
			return e
		}
	}
	return
}

// Remove ...
func (m *EntityManager) Remove(entity Entity) {
	delete(m.entities, entity.ID())
}
