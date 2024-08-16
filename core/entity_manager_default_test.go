package core_test

import (
	"testing"

	"github.com/andygeiss/ecs/core"
)

func TestEntityManager_Entities_Should_Have_No_Entity_At_Start(t *testing.T) {
	m := core.NewEntityManager()
	if len(m.Entities()) != 0 {
		t.Errorf("EntityManager should have no entity at start, but got %d", len(m.Entities()))
	}
}

func TestEntityManager_Entities_Should_Have_One_Entity_After_Adding_One_Entity(t *testing.T) {
	m := core.NewEntityManager()
	m.Add(&core.Entity{})
	if len(m.Entities()) != 1 {
		t.Errorf("EntityManager should have one entity, but got %d", len(m.Entities()))
	}
}

func TestEntityManager_Entities_Should_Have_Two_Entities_After_Adding_Two_Entities(t *testing.T) {
	m := core.NewEntityManager()
	m.Add(core.NewEntity("e1", nil))
	m.Add(core.NewEntity("e2", nil))
	if len(m.Entities()) != 2 {
		t.Errorf("EntityManager should have two entities, but got %d", len(m.Entities()))
	}
}

func TestEntityManager_Entities_Should_Have_One_Entity_After_Removing_One_Of_Two_Entities(t *testing.T) {
	m := core.NewEntityManager()
	e1 := core.NewEntity("e1", nil)
	e2 := core.NewEntity("e2", nil)
	m.Add(e1)
	m.Add(e2)
	m.Remove(e2)
	if len(m.Entities()) != 1 {
		t.Errorf("EntityManager should have one entity after removing one, but got %d", len(m.Entities()))
	}
	if m.Entities()[0].Id != "e1" {
		t.Errorf("Entity should have correct Id, but got %s", m.Entities()[0].Id)
	}
}

func TestEntityManager_FilterByMask_Should_Return_No_Entity_Out_Of_One(t *testing.T) {
	em := core.NewEntityManager()
	e := core.NewEntity("e", []core.Component{
		&mockComponent{name: "position", mask: 1},
	})
	em.Add(e)
	filtered := em.FilterByMask(2)
	if len(filtered) != 0 {
		t.Errorf("EntityManager should return no entity, but got %d", len(filtered))
	}
}

func TestEntityManager_FilterByMask_Should_Return_One_Entity_Out_Of_One(t *testing.T) {
	em := core.NewEntityManager()
	e := core.NewEntity("e", []core.Component{
		&mockComponent{name: "position", mask: 1},
	})
	em.Add(e)
	filtered := em.FilterByMask(1)
	if len(filtered) != 1 {
		t.Errorf("EntityManager should return one entity, but got %d", len(filtered))
	}
}

func TestEntityManager_FilterByMask_Should_Return_One_Entity_Out_Of_Two(t *testing.T) {
	em := core.NewEntityManager()
	e1 := core.NewEntity("e1", []core.Component{
		&mockComponent{name: "position", mask: 1},
	})
	e2 := core.NewEntity("e2", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	em.Add(e1, e2)
	filtered := em.FilterByMask(2)
	if len(filtered) != 1 {
		t.Errorf("EntityManager should return one entity, but got %d", len(filtered))
	}
	if filtered[0].Id != "e2" {
		t.Errorf("Entity should have correct Id, but got %s", filtered[0].Id)
	}
}

func TestEntityManager_FilterByMask_Should_Return_Two_Entities_Out_Of_Three(t *testing.T) {
	em := core.NewEntityManager()
	e1 := core.NewEntity("e1", []core.Component{
		&mockComponent{name: "position", mask: 1},
	})
	e2 := core.NewEntity("e2", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	e3 := core.NewEntity("e3", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	em.Add(e1, e2, e3)
	filtered := em.FilterByMask(2)
	if len(filtered) != 2 {
		t.Errorf("EntityManager should return two entities, but got %d", len(filtered))
	}
	if filtered[0].Id != "e2" {
		t.Errorf("Entity should have correct Id, but got %s", filtered[0].Id)
	}
	if filtered[1].Id != "e3" {
		t.Errorf("Entity should have correct Id, but got %s", filtered[1].Id)
	}
}

func TestEntityManager_FilterByMask_Should_Return_Three_Entities_Out_Of_Three(t *testing.T) {
	em := core.NewEntityManager()
	e1 := core.NewEntity("e1", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	e2 := core.NewEntity("e2", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	e3 := core.NewEntity("e3", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	em.Add(e1, e2, e3)
	filtered := em.FilterByMask(1 | 2)
	if len(filtered) != 3 {
		t.Errorf("EntityManager should return three entities, but got %d", len(filtered))
	}
	if filtered[0].Id != "e1" {
		t.Errorf("Entity should have correct Id, but got %s", filtered[0].Id)
	}
	if filtered[1].Id != "e2" {
		t.Errorf("Entity should have correct Id, but got %s", filtered[1].Id)
	}
	if filtered[2].Id != "e3" {
		t.Errorf("Entity should have correct Id, but got %s", filtered[2].Id)
	}
}

func TestEntityManager_FilterByNames_Should_Return_Three_Entities_Out_Of_Three(t *testing.T) {
	em := core.NewEntityManager()
	e1 := core.NewEntity("e1", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	e2 := core.NewEntity("e2", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	e3 := core.NewEntity("e3", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	em.Add(e1, e2, e3)
	filtered := em.FilterByNames("position", "size")
	if len(filtered) != 3 {
		t.Errorf("EntityManager should return three entities, but got %d", len(filtered))
	}
	if filtered[0].Id != "e1" {
		t.Errorf("Entity should have correct Id, but got %s", filtered[0].Id)
	}
	if filtered[1].Id != "e2" {
		t.Errorf("Entity should have correct Id, but got %s", filtered[1].Id)
	}
	if filtered[2].Id != "e3" {
		t.Errorf("Entity should have correct Id, but got %s", filtered[2].Id)
	}
}

func TestEntityManager_Get_Should_Return_Entity(t *testing.T) {
	em := core.NewEntityManager()
	e1 := core.NewEntity("e1", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	e2 := core.NewEntity("e2", []core.Component{
		&mockComponent{name: "position", mask: 1},
		&mockComponent{name: "size", mask: 2},
	})
	em.Add(e1, e2)
	if e := em.Get("e1"); e == nil {
		t.Error("Entity should not be nil")
	}
	if e := em.Get("e2"); e == nil {
		t.Error("Entity should not be nil")
	}
}
