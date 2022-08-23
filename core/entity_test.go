package core_test

import (
	"github.com/andygeiss/ecs/core"
	"testing"

	"github.com/andygeiss/utils/assert"
)

func TestEntity_NewEntity_Should_Create_A_Correct_Mask(t *testing.T) {
	entity := core.NewEntity("e", []core.Component{
		&mockComponent{name: "position", mask: 1},
	})
	assert.That("mask should be 1", t, entity.Mask(), 1)
}

func TestEntity_Add_Should_Work_With_Multiple_Components(t *testing.T) {
	entity := core.NewEntity("e", []core.Component{
		&mockComponent{name: "position", mask: 1},
	})
	entity.Add(&mockComponent{name: "velocity", mask: 2})
	assert.That("mask should be 3", t, entity.Mask(), 3)
}

func TestEntity_Remove_Should_Work_With_Multiple_Components(t *testing.T) {
	entity := core.NewEntity("e", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
		&mockComponent{name: "velocity", mask: 4},
	})
	entity.Remove(4)
	assert.That("mask should be 3", t, entity.Mask(), 3)
}

/*
       _   _ _
 _   _| |_(_) |___
| | | | __| | / __|
| |_| | |_| | \__ \
 \__,_|\__|_|_|___/
*/

type mockComponent struct {
	mask uint64
	name string
}

func (c *mockComponent) Mask() uint64 { return c.mask }

func (c *mockComponent) Name() string { return c.name }
