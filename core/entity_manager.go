package core

// EntityManager handles the access to each entity.
type EntityManager interface {
	Add(entities ...*Entity)
	Entities() (entities []*Entity)
	FilterByMask(mask uint64) (entities []*Entity)
	FilterByNames(names ...string) (entities []*Entity)
	Get(id string) (entity *Entity)
	Remove(entity *Entity)
}
