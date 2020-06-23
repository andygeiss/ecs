package ecs_test

import (
	"fmt"
	"github.com/andygeiss/ecs"
	"math/rand"
	"testing"
)

func BenchmarkEngine_Run_Once(b *testing.B) {
	entityCounts := []int{100, 1000, 10000}
	systemCounts := []int{1, 2, 4, 8}
	for _, systemCount := range systemCounts {
		for _, entityCount := range entityCounts {
			b.Run(fmt.Sprintf("%d system(s) with %d entities", systemCount, entityCount), func(b *testing.B) {
				b.ResetTimer()
				em := ecs.NewEntityManager()
				em.Add(generateEntities(entityCount)...)
				sm := ecs.NewSystemManager()
				sm.Add(generateSystems(systemCount)...)
				engine := ecs.NewEngine(em, sm)
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

func generateEntities(count int) []*ecs.Entity {
	out := make([]*ecs.Entity, count)
	for i := 0; i < count; i++ {
		out[i] = &ecs.Entity{Id: fmt.Sprintf("e%d", rand.Uint64()), Components: []ecs.Component{&mockComponent{name: "benchmark"}}}
	}
	return out
}

func generateSystems(count int) []ecs.System {
	out := make([]ecs.System, count)
	for i := 0; i < count-1; i++ {
		out[i] = &mockupUseAllEntitiesSystem{}
	}
	out[count-1] = &mockupShouldStopSystem{}
	return out
}

// mockupUseAllEntitiesSystem works on all entities from the EntityManager which represents the worst-case scenario for performance.
type mockupUseAllEntitiesSystem struct{}

func (s *mockupUseAllEntitiesSystem) Process(entityManager *ecs.EntityManager) (state int) {
	for range entityManager.FilterBy("benchmark") {
	}
	return ecs.StateEngineContinue
}
func (s *mockupUseAllEntitiesSystem) Setup() {
}
func (s *mockupUseAllEntitiesSystem) Teardown() {
}

// mockupShouldStopSystem is the last System in the queue and should stop the engine.
type mockupShouldStopSystem struct{}

func (s *mockupShouldStopSystem) Process(entityManager *ecs.EntityManager) (state int) {
	for range entityManager.FilterBy("benchmark") {
	}
	return ecs.StateEngineStop
}
func (s *mockupShouldStopSystem) Setup() {
}
func (s *mockupShouldStopSystem) Teardown() {
}
