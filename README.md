# ECS - Entity Component System

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/ecs)](https://goreportcard.com/report/github.com/andygeiss/ecs)

**A simple API to use an Entity Component System in Golang**

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

A component defines a specific part of an entity, 
which could be represented as a state by using a data structure. 

```go
package components
...
type Health struct { Value int }
func (c *Health) Name() string { return "health" }
```

**Entities**

An entity could be defined by adding specific components:

```go
package entities
...
player := ecs.Entity{
	Id: "player",
	Components: []ecs.Component{
		&components.Health{Value: 100},     // health
		&components.Job{Title: "warrior"},  // job
		&components.Position{X: 10, Y: 10}, // position
		&components.Velocity{X: 1, Y: 0},   // velocity
	}
}
```

**Systems**

A System must implement the behaviour of a component 
or the combination of some components:

```go
package systems

// Movement ...
type Movement struct{}

// NewMovement ...
func NewMovement() ecs.System {
    return &Movement{}
}

// Process ...
func (p *Movement) Process(entityManager *ecs.EntityManager) {
    // Select entities which have a position and velocity.
    for _, e := range entityManager.FilterBy("position", "velocity") {
    	// Access the components ...
        position := e.Get("position").(*components.Position)
        velocity := e.Get("velocity").(*components.Velocity)
        // Modify the data.
        position.X += velocity.X
        position.Y += velocity.Y
    }
}

// Setup ...
func (p *Movement) Setup() {}

// Teardown ...
func (p *Movement) Teardown() {}
```

**Engine**

A simple game engine triggers the <code>Process</code> method every frame.

```go
func (g *engine) Run() {
    for _, system := range g.systems {
        system.Process(g.entityManager)
    }
}
```
