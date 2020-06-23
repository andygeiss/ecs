package ecs_test

import (
	"fmt"
	"github.com/andygeiss/ecs"
	"math/rand"
	"testing"
)

func BenchmarkEngine_Run(b *testing.B) {
	entityCounts := []int{100, 1000, 10000}
	systemCounts := []int{1, 2, 4}
	for _, systemCount := range systemCounts {
		for _, entityCount := range entityCounts {
			b.Run(fmt.Sprintf("%d system(s) with %d entities", systemCount, entityCount), func(b *testing.B) {
				b.ResetTimer()
				em := ecs.NewEntityManager()
				em.Add(generateEntities(entityCount)...)
				sm := ecs.NewSystemManager()
				sm.Add(generateUseAllEntitiesSystems(systemCount)...)
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

func BenchmarkEngine_Run_Cached(b *testing.B) {
	entityCounts := []int{100, 1000, 10000}
	systemCounts := []int{1, 2, 4}
	for _, systemCount := range systemCounts {
		for _, entityCount := range entityCounts {
			b.Run(fmt.Sprintf("%d system(s) with %d entities", systemCount, entityCount), func(b *testing.B) {
				b.ResetTimer()
				em := ecs.NewEntityManager()
				entities := generateEntities(entityCount)
				em.Add(entities...)
				sm := ecs.NewSystemManager()
				systems := generateUseCacheSystems(systemCount)
				for _, system := range systems {
					switch sys := system.(type) {
					case *mockupUseACacheSystem:
						sys.cache = entities
					}
				}
				sm.Add(systems...)
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

func generateUseAllEntitiesSystems(count int) []ecs.System {
	out := make([]ecs.System, count)
	for i := 0; i < count-1; i++ {
		out[i] = &mockupUseAllEntitiesSystem{}
	}
	out[count-1] = &mockupShouldStopSystem{}
	return out
}

func generateUseCacheSystems(count int) []ecs.System {
	out := make([]ecs.System, count)
	for i := 0; i < count-1; i++ {
		out[i] = &mockupUseACacheSystem{}
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

// mockupUseACacheSystem uses a cache to get a faster access to the entities.
type mockupUseACacheSystem struct {
	cache []*ecs.Entity
}

func (s *mockupUseACacheSystem) Process(entityManager *ecs.EntityManager) (state int) {
	// Don't use the EntityManager - instead use the local cache.
	// Important: This only works for Entities with an infinity lifespan (at best).
	// You need to somehow update the cache if new entities were added to the manager. ;-)
	for range s.cache {
	}
	return ecs.StateEngineContinue
}
func (s *mockupUseACacheSystem) Setup() {
}
func (s *mockupUseACacheSystem) Teardown() {
}
