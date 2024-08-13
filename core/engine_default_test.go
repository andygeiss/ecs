package core_test

import (
	"testing"

	"github.com/andygeiss/ecs/core"
)

func TestDefaultEngine(t *testing.T) {
	em := &mockupEntityManager{}
	sm := &mockupSystemManager{}
	sys := &mockupSystem{Value: core.StateEngineStop}
	sm.Add(sys)
	engine := core.NewDefaultEngine(em, sm)
	engine.Setup()
	// app.Run()
	engine.Teardown()
	if sys.Value != core.StateEngineStop {
		t.Error("State should be correct")
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
	Value int
}

func (s *mockupSystem) Process(entityManager core.EntityManager) (state int) {
	s.Value = 1
	return core.StateEngineStop
}
func (s *mockupSystem) Setup()    {}
func (s *mockupSystem) Teardown() {}
