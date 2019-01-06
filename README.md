# ECS - Entity Component System

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/ecs)](https://goreportcard.com/report/github.com/andygeiss/ecs)

**A simple API to use an Entity Component System in Golang ;-)**

The architectural pattern of an ECS is mostly used in game development,
to provide long-term maintainability and extendability of large, dynamic systems.

An **Entity** is basically a composition of different components and has an ID.  
A **Component** contains only the state or data of one specific aspect like health, position, velocity etc.  
A **System** handles the behaviour or logic of the components. A movement system uses the position and velocity to implement an entities movement. 

    Entity  +-------has------>  Component  <-----handles----+  System
    
    Compositon of               State                          Behaviour
    Components                  Data                           Logic

**Table of Contents**

- [Goals](README.md#goals)
- [Installation](README.md#installation)
- [Usage](README.md#usage)

## Goals

- Provide an **easy-to-use module** to create a ECS architecture.
- **No dependencies** to other modules.
- **Minimum overhead** - use only what is really needed.

## Installation

**From Source**

    go get -u github.com/andygeiss/ecs

## Usage

**Components**

A component must implement a <code>Name</code> method,
which will be used later to filter entities with specific components.

```go
type HealthComponent int
func (c *HealthComponent) Name() string { return "health" }
```

**Entities**

An entity must implement a <code>Components</code> and <code>ID</code> method.

```go
type PlayerEntity struct {
	components []ecs.Component
	id string
}
func (e *PlayerEntity) Components() string { return e.components }
func (e *PlayerEntity) ID() string { return e.id }
```

**Systems**

A System must implement a <code>Setup</code>, <code>Process</code> and <code>Teardown</code> method.

```go
type MovementSystem struct {}
func (s *MovementSystem) Process(entities []ecs.Entities) { 
    for _, e := range entities {
        // Filter entity components, but skip further action
        // if the filter could not be applied completely.
        filtered, complete := ecs.Filter(e, "position", "velocity")
        if !complete {
            continue
        }
        // Map the filter to concrete components.
        position := filtered["position"].(*components.Position)
        velocity := filtered["velocity"].(*components.Velocity)
        // Apply the system logic.
        position.X += velocity.X
        position.Y += velocity.Y
    }
}
func (s *MovementSystem) Setup() {}
func (s *MovementSystem) Teardown() {}
```

**Games**

A game loop for each frame should now run each <code>Process</code> method for all active entities by the existing systems.

```go
func (game *Game) Run() {
	for !game.Over {
		for _, system := range g.systems {
			system.Process(g.entities)
		}
	}
}
```
