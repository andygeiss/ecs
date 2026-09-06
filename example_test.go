package ecs_test

import (
	"context"
	"fmt"

	"github.com/andygeiss/ecs"
)

// Each component type owns one bit, so an entity's components fit into a
// single uint64.
const (
	maskPosition = uint64(1 << 0)
	maskVelocity = uint64(1 << 1)
)

type position struct {
	X int
	Y int
}

func (p *position) Mask() uint64 { return maskPosition }
func (p *position) Name() string { return "position" }

type velocity struct {
	X int
	Y int
}

func (v *velocity) Mask() uint64 { return maskVelocity }
func (v *velocity) Name() string { return "velocity" }

// movementSystem moves everything that has both a position and a velocity,
// and stops the engine after steps passes.
type movementSystem struct {
	steps int
}

func (s *movementSystem) Setup() {}

func (s *movementSystem) Process(em ecs.EntityManager) (state int) {
	for _, e := range em.FilterByMask(maskPosition | maskVelocity) {
		p := e.Get(maskPosition).(*position)
		v := e.Get(maskVelocity).(*velocity)
		p.X += v.X
		p.Y += v.Y
	}
	s.steps--
	if s.steps <= 0 {
		return ecs.StateEngineStop
	}
	return ecs.StateEngineContinue
}

func (s *movementSystem) Teardown() {}

// The smallest engine that does something: one entity, one system, three
// passes.
func Example() {
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("player", []ecs.Component{
		&position{},
		&velocity{X: 2, Y: 3},
	}))

	sm := ecs.NewSystemManager()
	sm.Add(&movementSystem{steps: 3})

	engine := ecs.NewDefaultEngine(em, sm)
	engine.Setup()
	defer engine.Teardown()
	engine.Run(context.Background())

	p := em.Get("player").Get(maskPosition).(*position)
	fmt.Println(p.X, p.Y)
	// Output: 6 9
}

func ExampleNewEntity() {
	e := ecs.NewEntity("player", []ecs.Component{
		&position{},
		&velocity{},
	})
	fmt.Println(e.Id, e.Mask())
	// Output: player 3
}

// Tick runs a single pass, for a game that owns its own loop and draws a
// frame between passes.
func ExampleDefaultEngine_Tick() {
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("player", []ecs.Component{
		&position{},
		&velocity{X: 2, Y: 3},
	}))

	sm := ecs.NewSystemManager()
	sm.Add(&movementSystem{steps: 100})

	engine := ecs.NewDefaultEngine(em, sm)
	engine.Setup()
	defer engine.Teardown()
	engine.Tick()

	p := em.Get("player").Get(maskPosition).(*position)
	fmt.Println(p.X, p.Y)
	// Output: 2 3
}

// FilterByMask returns every entity that carries all of the components in the
// mask. This is how a system finds its work.
func ExampleDefaultEntityManager_FilterByMask() {
	em := ecs.NewEntityManager()
	em.Add(
		ecs.NewEntity("player", []ecs.Component{&position{}, &velocity{}}),
		ecs.NewEntity("wall", []ecs.Component{&position{}}),
	)

	for _, e := range em.FilterByMask(maskPosition | maskVelocity) {
		fmt.Println(e.Id)
	}
	// Output: player
}

// FilterByNames does the same by name, for a game that outgrew the 64 bits a
// mask has. Its components implement ComponentWithName.
func ExampleDefaultEntityManager_FilterByNames() {
	em := ecs.NewEntityManager()
	em.Add(
		ecs.NewEntity("player", []ecs.Component{&position{}, &velocity{}}),
		ecs.NewEntity("wall", []ecs.Component{&position{}}),
	)

	for _, e := range em.FilterByNames("position", "velocity") {
		fmt.Println(e.Id)
	}
	// Output: player
}
