package ecs

// Run simplifies the engine usage by calling the Setup(), Run() and Teardown() internally.
func Run(em *EntityManager, sm *SystemManager) {
	engine := NewEngine(em, sm)
	engine.Setup()
	defer engine.Teardown()
	engine.Run()
}
