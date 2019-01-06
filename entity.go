package ecs

// Entity ...
type Entity interface {
	Components() (components []Component)
	ID() (id string)
}
