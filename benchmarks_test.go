package ecs_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/andygeiss/ecs/core"
)

func BenchmarkEntityManager_Get_With_1_Entity_Id_Found(b *testing.B) {
	m := core.NewEntityManager()
	m.Add(core.NewEntity("foo", nil))
	for i := 0; i < b.N; i++ {
		m.Get("foo")
	}
}

func BenchmarkEntityManager_Get_With_1000_Entities_Id_Not_Found(b *testing.B) {
	m := core.NewEntityManager()
	for i := 0; i < 1000; i++ {
		m.Add(core.NewEntity("foo", nil))
	}
	for i := 0; i < b.N; i++ {
		m.Get("1000")
	}
}

func BenchmarkEntityManager_FilterByMask_With_1000_Entities(b *testing.B) {
	m := core.NewEntityManager()
	for i := 0; i < 1000; i++ {
		m.Add(core.NewEntity(fmt.Sprintf("%d", i), []core.Component{
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
			&mockComponent{name: "velocity", mask: 3},
		}))
	}
	for i := 0; i < b.N; i++ {
		m.FilterByMask(1 | 2 | 3)
	}
}

func BenchmarkEntityManager_FilterByNames_With_1000_Entities(b *testing.B) {
	m := core.NewEntityManager()
	for i := 0; i < 1000; i++ {
		m.Add(core.NewEntity(fmt.Sprintf("%d", i), []core.Component{
			&mockComponent{name: "position", mask: 1},
			&mockComponent{name: "size", mask: 2},
			&mockComponent{name: "velocity", mask: 3},
		}))
	}
	for i := 0; i < b.N; i++ {
		m.FilterByNames("position", "size", "velocity")
	}
}

func BenchmarkEngine_Run(b *testing.B) {
	entityCounts := []int{100, 1000, 10000}
	systemCounts := []int{1, 2, 4}
	for _, systemCount := range systemCounts {
		for _, entityCount := range entityCounts {
			b.Run(fmt.Sprintf("%d system(s) with %d entities", systemCount, entityCount), func(b *testing.B) {
				b.ResetTimer()
				em := core.NewEntityManager()
				em.Add(generateEntities(entityCount)...)
				sm := core.NewSystemManager()
				sm.Add(generateUseAllEntitiesSystems(systemCount)...)
				engine := core.NewDefaultEngine(em, sm)
				engine.Setup()
				defer engine.Teardown()
				for i := 0; i < b.N; i++ {
					engine.Run()
				}
			})
		}
	}
}

/*
       _   _ _
 _   _| |_(_) |___
| | | | __| | / __|
| |_| | |_| | \__ \
 \__,_|\__|_|_|___/
*/

func generateEntities(count int) []*core.Entity {
	out := make([]*core.Entity, count)
	for i := 0; i < count; i++ {
		out[i] = core.NewEntity(
			fmt.Sprintf("e%d", rand.Uint64()),
			[]core.Component{
				&mockComponent{mask: 1},
			},
		)
	}
	return out
}

func generateUseAllEntitiesSystems(count int) []core.System {
	out := make([]core.System, count)
	for i := 0; i < count-1; i++ {
		out[i] = &mockupUseAllEntitiesSystem{}
	}
	out[count-1] = &mockupShouldStopSystem{}
	return out
}

// mockupUseAllEntitiesSystem works on all entities from the defaultEntityManager which represents the worst-case scenario for performance.
type mockupUseAllEntitiesSystem struct{}

func (s *mockupUseAllEntitiesSystem) Process(entityManager core.EntityManager) (state int) {
	for range entityManager.FilterByMask(1) {
	}
	return core.StateEngineContinue
}
func (s *mockupUseAllEntitiesSystem) Setup() {
}
func (s *mockupUseAllEntitiesSystem) Teardown() {
}

// mockupShouldStopSystem is the last System in the queue and should stop the defaultEngine.
type mockupShouldStopSystem struct{}

func (s *mockupShouldStopSystem) Process(entityManager core.EntityManager) (state int) {
	for range entityManager.FilterByMask(1) {
	}
	return core.StateEngineStop
}
func (s *mockupShouldStopSystem) Setup() {
}
func (s *mockupShouldStopSystem) Teardown() {
}

type mockComponent struct {
	mask uint64
	name string
}

func (c *mockComponent) Mask() uint64 { return c.mask }

func (c *mockComponent) Name() string { return c.name }
