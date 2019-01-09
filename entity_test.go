package ecs_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/ecs"
	"testing"
)

func TestEntityManager_Entities_Should_Have_No_Entity_At_Start(t *testing.T) {
	m := ecs.NewEntityManager()
	assert.That(t, len(m.Entities()), is.Equal(0))
}

type mockComponent struct {
	name string
}

func (c *mockComponent) Name() string { return c.name }

func TestEntityManager_Entities_Should_Have_One_Entity_After_Adding_One_Entity(t *testing.T) {
	m := ecs.NewEntityManager()
	m.Add(&ecs.Entity{})
	assert.That(t, len(m.Entities()), is.Equal(1))
}

func TestEntityManager_Entities_Should_Have_Two_Entities_After_Adding_Two_Entities(t *testing.T) {
	m := ecs.NewEntityManager()
	m.Add(&ecs.Entity{Id: "1"})
	m.Add(&ecs.Entity{Id: "2"})
	assert.That(t, len(m.Entities()), is.Equal(2))
}

func TestEntityManager_Entities_Should_Have_One_Entity_After_Adding_Two_Entities_With_The_Same_ID(t *testing.T) {
	m := ecs.NewEntityManager()
	m.Add(&ecs.Entity{Id: "1"})
	m.Add(&ecs.Entity{Id: "1"})
	assert.That(t, len(m.Entities()), is.Equal(1))
}

func TestEntityManager_Entities_Should_Have_One_Entity_After_Removing_One_Of_Two_Entities(t *testing.T) {
	m := ecs.NewEntityManager()
	e1 := &ecs.Entity{Id: "e1"}
	e2 := &ecs.Entity{Id: "e2"}
	m.Add(e1)
	m.Add(e2)
	m.Remove(e2)
	assert.That(t, len(m.Entities()), is.Equal(1))
	assert.That(t, m.Entities()[0].Id, is.Equal("e1"))
}

func TestEntityManager_FilterBy_Should_Return_One_Entity_Out_Of_One(t *testing.T) {
	em := ecs.NewEntityManager()
	e := &ecs.Entity{Id: "e1", Components: []ecs.Component{
		&mockComponent{name: "position"},
	}}
	em.Add(e)
	entities := em.FilterBy("position")
	assert.That(t, len(entities), is.Equal(1))
	assert.That(t, entities[0], is.Equal(e))
}

func TestEntityManager_FilterBy_Should_Return_One_Entity_Out_Of_Two(t *testing.T) {
	em := ecs.NewEntityManager()
	e1 := &ecs.Entity{Id: "e1", Components: []ecs.Component{
		&mockComponent{name: "position"},
	}}
	e2 := &ecs.Entity{Id: "e2", Components: []ecs.Component{
		&mockComponent{name: "velocity"},
	}}
	em.Add(e1, e2)
	entities := em.FilterBy("position")
	assert.That(t, len(entities), is.Equal(1))
	assert.That(t, entities[0], is.Equal(e1))
}
