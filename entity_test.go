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

type mockEntity struct {
	components []ecs.Component
	id         string
}

func (e *mockEntity) Components() []ecs.Component { return e.components }
func (e *mockEntity) Get(name string) (component ecs.Component) {
	for _, c := range e.components {
		if c.Name() == name {
			return c
		}
	}
	return
}
func (e *mockEntity) ID() string { return e.id }

func TestEntityManager_Entities_Should_Have_One_Entity_After_Adding_One_Entity(t *testing.T) {
	m := ecs.NewEntityManager()
	m.Add(&mockEntity{})
	assert.That(t, len(m.Entities()), is.Equal(1))
}

func TestEntityManager_Entities_Should_Have_Two_Entities_After_Adding_Two_Entities(t *testing.T) {
	m := ecs.NewEntityManager()
	m.Add(&mockEntity{id: "1"})
	m.Add(&mockEntity{id: "2"})
	assert.That(t, len(m.Entities()), is.Equal(2))
}

func TestEntityManager_Entities_Should_Have_One_Entity_After_Adding_Two_Entities_With_The_Same_ID(t *testing.T) {
	m := ecs.NewEntityManager()
	m.Add(&mockEntity{id: "1"})
	m.Add(&mockEntity{id: "1"})
	assert.That(t, len(m.Entities()), is.Equal(1))
}

func TestEntityManager_Entities_Should_Have_One_Entity_After_Removing_One_Of_Two_Entities(t *testing.T) {
	m := ecs.NewEntityManager()
	e1 := &mockEntity{id: "e1"}
	e2 := &mockEntity{id: "e2"}
	m.Add(e1)
	m.Add(e2)
	m.Remove(e2)
	assert.That(t, len(m.Entities()), is.Equal(1))
	assert.That(t, m.Entities()[0].ID(), is.Equal("e1"))
}

func TestEntityManager_FilterBy_Should_Return_One_Entity_Out_Of_One(t *testing.T) {
	em := ecs.NewEntityManager()
	e := &mockEntity{id: "e1", components: []ecs.Component{
		&mockComponent{name:"position"},
	}}
	em.Add(e)
	entities := em.FilterBy("position")
	assert.That(t, len(entities), is.Equal(1))
	assert.That(t, entities[0], is.Equal(e))
}

func TestEntityManager_FilterBy_Should_Return_One_Entity_Out_Of_Two(t *testing.T) {
	em := ecs.NewEntityManager()
	e1 := &mockEntity{id: "e1", components: []ecs.Component{
		&mockComponent{name:"position"},
	}}
	e2 := &mockEntity{id: "e2", components: []ecs.Component{
		&mockComponent{name:"velocity"},
	}}
	em.Add(e1, e2)
	entities := em.FilterBy("position")
	assert.That(t, len(entities), is.Equal(1))
	assert.That(t, entities[0], is.Equal(e1))
}
