// Package ecs builds a game engine out of entities, components and systems.
//
// A [Component] is data and nothing else. An [Entity] is a bag of components
// with an id. A [System] is the behaviour: it reads and writes the components
// of the entities it cares about. That split is the whole idea — you bring
// the rendering, the input and the physics.
//
// [NewDefaultEngine] is the entry point. It takes an [EntityManager] and a
// [SystemManager], and drives them:
//
//	em := ecs.NewEntityManager()
//	em.Add(ecs.NewEntity("player", []ecs.Component{&position{}, &velocity{}}))
//
//	sm := ecs.NewSystemManager()
//	sm.Add(&movementSystem{})
//
//	engine := ecs.NewDefaultEngine(em, sm)
//	engine.Setup()
//	defer engine.Teardown()
//	engine.Run(ctx)
//
// [DefaultEngine.Run] calls every system over and over until one of them
// returns [StateEngineStop] or ctx is cancelled. [DefaultEngine.Tick] runs a
// single pass instead, for a caller that owns its own loop.
//
// A system finds its entities through the manager it is handed. Use
// [DefaultEntityManager.FilterByMask] while 64 component types are enough,
// and [DefaultEntityManager.FilterByNames] when they are not.
//
// Nothing here is safe for concurrent use. The engine calls one system after
// another on one goroutine, so systems need no locking between themselves. A
// system that fans its own work out to goroutines owns the locking for that.
package ecs
