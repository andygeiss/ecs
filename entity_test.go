package ecs_test

import (
	"testing"

	"github.com/andygeiss/ecs"
)

func TestEntity_NewEntity_Should_Create_A_Correct_Mask(t *testing.T) {
	t.Parallel()
	entity := ecs.NewEntity("e", []ecs.Component{
		&mockComponent{name: "position", mask: 1},
	})
	if entity.Mask() != 1 {
		t.Errorf("Entity mask should be 1, but got %d", entity.Mask())
	}
}

func TestEntity_Add_Should_Work_With_Multiple_Components(t *testing.T) {
	t.Parallel()
	entity := ecs.NewEntity("e", []ecs.Component{
		&mockComponent{name: "position", mask: 1},
	})
	entity.Add(&mockComponent{name: "velocity", mask: 2})
	if entity.Mask() != 3 {
		t.Errorf("Entity mask should be 3, but got %d", entity.Mask())
	}
}

func TestEntity_Add_Should_Not_Add_Existing_Component(t *testing.T) {
	t.Parallel()
	entity := ecs.NewEntity("e", []ecs.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "velocity", mask: 2},
	})
	entity.Add(&mockComponent{name: "velocity", mask: 2})
	if len(entity.Components) != 2 {
		t.Errorf("Component count should be 2, but got %d", len(entity.Components))
	}
}

func TestEntity_Get_Should_Return_Component(t *testing.T) {
	t.Parallel()
	entity := ecs.NewEntity("e", []ecs.Component{
		&mockComponent{name: "position", mask: 1},
	})
	component := entity.Get(1)
	if component == nil {
		t.Error("Component should not be nil")
	}
}

func TestEntity_Remove_Should_Work_With_Multiple_Components(t *testing.T) {
	t.Parallel()
	entity := ecs.NewEntity("e", []ecs.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
		&mockComponent{name: "velocity", mask: 4},
	})
	entity.Remove(4)
	if entity.Mask() != 3 {
		t.Errorf("Entity mask should be 1, but got %d", entity.Mask())
	}
}

func TestEntity_NewEntity_Without_Components_Should_Have_An_Empty_Mask(t *testing.T) {
	t.Parallel()
	entity := ecs.NewEntity("e", nil)
	if entity.Mask() != 0 {
		t.Errorf("Entity mask should be 0, but got %d", entity.Mask())
	}
}

func TestEntity_Get_Should_Return_Nil_For_An_Unknown_Mask(t *testing.T) {
	t.Parallel()
	entity := ecs.NewEntity("e", []ecs.Component{
		&mockComponent{name: "position", mask: 1},
	})
	if component := entity.Get(2); component != nil {
		t.Errorf("Component should be nil, but got %v", component)
	}
}

func TestEntity_Remove_Should_Change_Nothing_For_An_Unknown_Mask(t *testing.T) {
	t.Parallel()
	entity := ecs.NewEntity("e", []ecs.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	entity.Remove(4)
	if len(entity.Components) != 2 {
		t.Errorf("Component count should be 2, but got %d", len(entity.Components))
	}
	if entity.Mask() != 3 {
		t.Errorf("Entity mask should be 3, but got %d", entity.Mask())
	}
}

func TestEntity_Remove_Should_Empty_The_Mask_After_The_Last_Component(t *testing.T) {
	t.Parallel()
	entity := ecs.NewEntity("e", []ecs.Component{
		&mockComponent{name: "position", mask: 1},
	})
	entity.Remove(1)
	if len(entity.Components) != 0 {
		t.Errorf("Component count should be 0, but got %d", len(entity.Components))
	}
	if entity.Mask() != 0 {
		t.Errorf("Entity mask should be 0, but got %d", entity.Mask())
	}
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

// maskOnlyComponent is a plain Component, so FilterByNames must never match it.
type maskOnlyComponent struct {
	mask uint64
}

func (c *maskOnlyComponent) Mask() uint64 { return c.mask }
