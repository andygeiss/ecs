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
func (e *PlayerEntity) Components() []ecs.Component { return e.components }
func (e *PlayerEntity) ID() string { return e.id }
```

**Systems**

A System must implement a <code>Setup</code>, <code>Process</code> and <code>Teardown</code> method.

```go

type rendering struct {}

func NewRendering() ecs.System {
    return &rendering{}
}

// Process ...
func (s *rendering) Process(entityManager *ecs.EntityManager) {
    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)
    for _, e := range entityManager.Entities() {
        // Filter entity components, but skip further action
        // if the filter could not be applied completely.
        filtered, complete := ecs.Filter(e, "position", "size")
        if !complete {
            continue
        }
        // Map the filter to concrete components.
        position := filtered["position"].(*components.Position)
        size := filtered["size"].(*components.Size)
        // Apply the system logic.
        rl.DrawRectangleLines(
            int32(position.X),
            int32(position.Y),
            int32(size.Width),
            int32(size.Height),
            rl.RayWhite,
        )
    }
    rl.EndDrawing()
}

// Setup ...
func (s *rendering) Setup() {
    rl.InitWindow(1024, 576, "Demo")
    rl.SetTargetFPS(60)
}

// Teardown ...
func (s *rendering) Teardown() {
    rl.CloseWindow()
}
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
