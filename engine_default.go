package ecs

import "context"

// DefaultEngine is simply a composition of an EntityManager and a
// SystemManager. Create one with NewDefaultEngine.
type DefaultEngine struct {
	entityManager EntityManager
	systemManager SystemManager
}

// Run calls Process on each System over and over, until a System returns
// StateEngineStop or ctx is cancelled.
//
// Cancellation is checked once per pass over the systems, so a Process that
// never returns still blocks forever.
func (e *DefaultEngine) Run(ctx context.Context) {
	for ctx.Err() == nil {
		for _, system := range e.systemManager.Systems() {
			if state := system.Process(e.entityManager); state == StateEngineStop {
				return
			}
		}
	}
}

// Tick calls Process on each System exactly once. It stops early when a
// System returns StateEngineStop.
func (e *DefaultEngine) Tick() {
	for _, system := range e.systemManager.Systems() {
		if state := system.Process(e.entityManager); state == StateEngineStop {
			break
		}
	}
}

// Setup calls Setup on each System, once, before the first Process.
func (e *DefaultEngine) Setup() {
	for _, sys := range e.systemManager.Systems() {
		sys.Setup()
	}
}

// Teardown calls Teardown on each System, once, after the last Process.
func (e *DefaultEngine) Teardown() {
	for _, sys := range e.systemManager.Systems() {
		sys.Teardown()
	}
}

// NewDefaultEngine creates a new DefaultEngine and returns its address.
func NewDefaultEngine(entityManager EntityManager, systemManager SystemManager) *DefaultEngine {
	return &DefaultEngine{
		entityManager: entityManager,
		systemManager: systemManager,
	}
}
