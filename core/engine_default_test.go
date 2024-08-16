package core_test

import (
	"testing"

	"github.com/andygeiss/ecs/core"
)

func prepare() (e core.Engine, sys *mockupSystem) {
	em := &mockupEntityManager{}
	sm := &mockupSystemManager{}
	system := &mockupSystem{State: core.StateEngineStop}
	sm.Add(system)
	engine := core.NewDefaultEngine(em, sm)
	return engine, system
}

func TestDefaultEngine_Teardown_After_Setup_Should_Set_StateEngineStop(t *testing.T) {
	engine, system := prepare()
	engine.Setup()
	// app.Run()
	engine.Teardown()
	if system.State != core.StateEngineStop {
		t.Error("State should be correct")
	}
}

func TestDefaultEngine_Run_Twice_Should_Increase_Counter_By_Two(t *testing.T) {
	engine, system := prepare()
	engine.Run()
	if system.Counter != 1 {
		t.Errorf("Counter should be 1, but got %d", system.Counter)
	}
	engine.Run()
	if system.Counter != 2 {
		t.Errorf("Counter should be 2, but got %d", system.Counter)
	}
}

func TestDefaultEngine_Tick_Twice_Should_Increase_Counter_By_Two(t *testing.T) {
	engine, system := prepare()
	engine.Tick()
	if system.Counter != 1 {
		t.Errorf("Counter should be 1, but got %d", system.Counter)
	}
	engine.Run()
	if system.Counter != 2 {
		t.Errorf("Counter should be 2, but got %d", system.Counter)
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

func (m *mockupEntityManager) Add(entities ...*core.Entity) {}

func (m *mockupEntityManager) Entities() (entities []*core.Entity) { return nil }

func (m *mockupEntityManager) FilterByMask(mask uint64) (entities []*core.Entity) { return nil }

func (m *mockupEntityManager) FilterByNames(names ...string) (entities []*core.Entity) { return nil }

func (m *mockupEntityManager) Get(id string) (entity *core.Entity) { return nil }

func (m *mockupEntityManager) Remove(entity *core.Entity) {}

type mockupSystemManager struct {
	systems []core.System
}

func (m *mockupSystemManager) Add(systems ...core.System) {
	m.systems = append(m.systems, systems...)
}

func (m *mockupSystemManager) Systems() []core.System {
	return m.systems
}

type mockupSystem struct {
	Counter int
	State   int
}

func (s *mockupSystem) Process(entityManager core.EntityManager) (state int) {
	s.Counter++
	return core.StateEngineStop
}
func (s *mockupSystem) Setup()    {}
func (s *mockupSystem) Teardown() {}
