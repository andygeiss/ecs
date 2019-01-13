package ecs

// System implements the behaviour of an entity by modifying the state,
// which is stored in each component of the entity.
type System interface {
	Setup()
	Process(entityManager *EntityManager)
	Teardown()
}
