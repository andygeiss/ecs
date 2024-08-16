package ecs

// System implements the behaviour of an entity by modifying the state,
// which is stored in each component of the entity.
type System interface {
	// Setup is called once before the first call to Process.
	// It is used to initialize the state of the system.
	Setup()
	// Process is called for each entity in the system.
	// It is used to modify the state of the entity.
	Process(entityManager EntityManager) (state int)
	// Teardown is called once after the last call to Process.
	// It is used to clean up the state of the system.
	Teardown()
}
