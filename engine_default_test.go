package ecs_test

import (
	"context"
	"testing"

	"github.com/andygeiss/ecs"
)

func prepare() (e *ecs.DefaultEngine, sys *mockupSystem) {
	em := &mockupEntityManager{}
	sm := &mockupSystemManager{}
	system := &mockupSystem{State: ecs.StateEngineStop}
	sm.Add(system)
	engine := ecs.NewDefaultEngine(em, sm)
	return engine, system
}

func TestDefaultEngine_Teardown_After_Setup_Should_Set_StateEngineStop(t *testing.T) {
	t.Parallel()
	engine, system := prepare()
	engine.Setup()
	engine.Teardown()
	if system.State != ecs.StateEngineStop {
		t.Error("State should be correct")
	}
}

func TestDefaultEngine_Run_Twice_Should_Increase_Counter_By_Two(t *testing.T) {
	t.Parallel()
	engine, system := prepare()
	engine.Run(t.Context())
	if system.Counter != 1 {
		t.Errorf("Counter should be 1, but got %d", system.Counter)
	}
	engine.Run(t.Context())
	if system.Counter != 2 {
		t.Errorf("Counter should be 2, but got %d", system.Counter)
	}
}

func TestDefaultEngine_Tick_Twice_Should_Increase_Counter_By_Two(t *testing.T) {
	t.Parallel()
	engine, system := prepare()
	engine.Tick()
	if system.Counter != 1 {
		t.Errorf("Counter should be 1, but got %d", system.Counter)
	}
	engine.Tick()
	if system.Counter != 2 {
		t.Errorf("Counter should be 2, but got %d", system.Counter)
	}
}

func TestDefaultEngine_Run_Should_Return_When_Context_Is_Cancelled(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	system := &mockupNeverStoppingSystem{cancelAfter: 3, cancel: cancel}
	sm := &mockupSystemManager{}
	sm.Add(system)
	engine := ecs.NewDefaultEngine(&mockupEntityManager{}, sm)
	engine.Run(ctx)
	if system.Counter != 3 {
		t.Errorf("Counter should be 3, but got %d", system.Counter)
	}
}

func TestDefaultEngine_Run_Should_Not_Process_With_An_Already_Cancelled_Context(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	engine, system := prepare()
	engine.Run(ctx)
	if system.Counter != 0 {
		t.Errorf("Counter should be 0, but got %d", system.Counter)
	}
}

/*
       _   _ _
 _   _| |_(_) |___
| | | | __| | / __|
| |_| | |_| | \__ \
 \__,_|\__|_|_|___/
*/

type mockupEntityManager struct{}

func (m *mockupEntityManager) Add(entities ...*ecs.Entity) {}

func (m *mockupEntityManager) Entities() (entities []*ecs.Entity) { return nil }

func (m *mockupEntityManager) FilterByMask(mask uint64) (entities []*ecs.Entity) { return nil }

func (m *mockupEntityManager) FilterByNames(names ...string) (entities []*ecs.Entity) { return nil }

func (m *mockupEntityManager) Get(id string) (entity *ecs.Entity) { return nil }

func (m *mockupEntityManager) Remove(entity *ecs.Entity) {}

type mockupSystemManager struct {
	systems []ecs.System
}

func (m *mockupSystemManager) Add(systems ...ecs.System) {
	m.systems = append(m.systems, systems...)
}

func (m *mockupSystemManager) Systems() []ecs.System {
	return m.systems
}

type mockupSystem struct {
	Counter int
	State   int
}

func (s *mockupSystem) Process(entityManager ecs.EntityManager) (state int) {
	s.Counter++
	return ecs.StateEngineStop
}
func (s *mockupSystem) Setup()    {}
func (s *mockupSystem) Teardown() {}

// mockupNeverStoppingSystem never returns StateEngineStop, so only the
// cancelled context can end the loop it is driven by.
type mockupNeverStoppingSystem struct {
	cancel      context.CancelFunc
	cancelAfter int
	Counter     int
}

func (s *mockupNeverStoppingSystem) Process(entityManager ecs.EntityManager) (state int) {
	s.Counter++
	if s.Counter == s.cancelAfter {
		s.cancel()
	}
	return ecs.StateEngineContinue
}
func (s *mockupNeverStoppingSystem) Setup()    {}
func (s *mockupNeverStoppingSystem) Teardown() {}
