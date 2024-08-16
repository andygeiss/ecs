package ecs

// EntityManager handles the access to each entity.
type EntityManager interface {
	// Add entries to the manager.
	Add(entities ...*Entity)
	// Entities returns all the entities.
	Entities() (entities []*Entity)
	// FilterByMask returns the mapped entities, which Components mask matched.
	FilterByMask(mask uint64) (entities []*Entity)
	// FilterByNames returns the mapped entities, which Components names matched.
	FilterByNames(names ...string) (entities []*Entity)
	// Get a specific entity by Id.
	Get(id string) (entity *Entity)
	// Remove a specific entity.
	Remove(entity *Entity)
}
