package ecs

// System implements the behaviour of an entity by modifying the state,
// which is stored in each component of the entity.
type System interface {
	// Setup is called once before the first call to Process.
	// It is used to initialize the state of the system.
	Setup()
	// Process is called once per pass of the engine. The system reads the
	// entities it cares about from entityManager, changes their components,
	// and returns StateEngineContinue or StateEngineStop.
	Process(entityManager EntityManager) (state int)
	// Teardown is called once after the last call to Process.
	// It is used to clean up the state of the system.
	Teardown()
}
