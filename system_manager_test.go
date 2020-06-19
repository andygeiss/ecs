package ecs_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/ecs"
	"testing"
)

func TestSystemManager_Systems_Should_Have_No_System_At_Start(t *testing.T) {
	m := ecs.NewSystemManager()
	assert.That("manager should have no system at start", t, len(m.Systems()), 0)
}

type MockSystem struct{}

func (s *MockSystem) Process(entityManager *ecs.EntityManager) (state int) {
	return ecs.StateEngineContinue
}
func (s *MockSystem) Setup()    {}
func (s *MockSystem) Teardown() {}

func TestSystemManager_Systems_Should_Have_One_System_After_Adding_One_System(t *testing.T) {
	m := ecs.NewSystemManager()
	s := &MockSystem{}
	m.Add(s)
	assert.That("manager should have one system after adding one", t, len(m.Systems()), 1)
}

func TestSystemManager_Systems_Should_Have_Two_System_After_Adding_Two_System(t *testing.T) {
	m := ecs.NewSystemManager()
	s1 := &MockSystem{}
	s2 := &MockSystem{}
	m.Add(s1, s2)
	assert.That("manager should have two systems", t, len(m.Systems()), 2)
}
