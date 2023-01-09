package core_test

import (
	"github.com/andygeiss/ecs/core"
	"testing"

	"github.com/andygeiss/utils/assert"
)

func TestSystemManager_Systems_Should_Have_No_System_At_Start(t *testing.T) {
	m := core.NewSystemManager()
	assert.That("manager should have no system at start", t, len(m.Systems()), 0)
}

func TestSystemManager_Systems_Should_Have_One_System_After_Adding_One_System(t *testing.T) {
	m := core.NewSystemManager()
	s := &mockupDedicatedSystem{}
	m.Add(s)
	assert.That("manager should have one system after adding one", t, len(m.Systems()), 1)
}

func TestSystemManager_Systems_Should_Have_Two_System_After_Adding_Two_System(t *testing.T) {
	m := core.NewSystemManager()
	s1 := &mockupDedicatedSystem{}
	s2 := &mockupDedicatedSystem{}
	m.Add(s1, s2)
	assert.That("manager should have two systems", t, len(m.Systems()), 2)
}

/*
       _   _ _
 _   _| |_(_) |___
| | | | __| | / __|
| |_| | |_| | \__ \
 \__,_|\__|_|_|___/
*/

// mockupDedicatedSystem is used without an defaultEngine to test the defaultSystemManager behaviour.
type mockupDedicatedSystem struct{}

func (s *mockupDedicatedSystem) Process(entityManager core.EntityManager) (state int) {
	return core.StateEngineContinue
}
func (s *mockupDedicatedSystem) Setup()    {}
func (s *mockupDedicatedSystem) Teardown() {}
