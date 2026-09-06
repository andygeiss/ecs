package ecs

// EntityManager handles the access to each entity.
type EntityManager interface {
	// Add appends entities to the manager.
	Add(entities ...*Entity)
	// Entities returns all the entities.
	Entities() (entities []*Entity)
	// FilterByMask returns every entity that carries all the components in mask.
	FilterByMask(mask uint64) (entities []*Entity)
	// FilterByNames returns every entity that carries a component for
	// each of the given names.
	FilterByNames(names ...string) (entities []*Entity)
	// Get returns the entity with the given Id, or nil when there is none.
	Get(id string) (entity *Entity)
	// Remove deletes the entity with the same Id.
	Remove(entity *Entity)
}
