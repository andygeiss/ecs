<p align="center">
<img src="https://github.com/andygeiss/ecs/blob/master/logo.png?raw=true" />
</p>

# ECS - Entity Component System

[![License](https://img.shields.io/github/license/andygeiss/ecs)](https://github.com/andygeiss/ecs/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/v/release/andygeiss/ecs)](https://github.com/andygeiss/ecs/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/ecs)](https://goreportcard.com/report/github.com/andygeiss/ecs)
[![Codacy Grade Badge](https://app.codacy.com/project/badge/Grade/b4f4c9b35f4b46d8bf19f73379864b45)](https://app.codacy.com/gh/andygeiss/ecs/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Codacy Coverage Badge](https://app.codacy.com/project/badge/Coverage/b4f4c9b35f4b46d8bf19f73379864b45)](https://app.codacy.com/gh/andygeiss/ecs/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/5a2fd230f2eae6f244f2/maintainability)](https://codeclimate.com/github/andygeiss/ecs/maintainability)

Build your own Game-Engine based on the Entity Component System concept in Golang.

## Features

- [x] Provide an **easy-to-use** framework to build a game engine from scratch.
- [x] **No dependencies** to other modules or specific game libraries - Feel free to use what fits your needs.
- [x] **Minimum overhead** - use only what is really needed.

### Example engine

See [engine-example](https://github.com/andygeiss/engine-example) for a basic 
implementation using [raylib](https://www.raylib.com)).

## Walkthrough

### Project layout

At first we create a basic project layout:

```bash
mkdir ecs-example
cd ecs-example
go mod init example
mkdir components plugins systems
```

Next we create a `main.go` with the following content:

```go
package main

import (
    ecs "github.com/andygeiss/ecs/core"
)

func main() {
    em := ecs.NewEntityManager()
    sm := ecs.NewSystemManager()
    de := ecs.NewDefaultEngine(em, sm)
    de.Setup()
    defer de.Teardown()
    de.Run()
}
```

The execution of the program leads to an endless loop, as our engine is not yet
able to react to user input.

### The movement system

A system needs to implement the methods defined by the interface
[System](https://github.com/andygeiss/ecs/blob/master/core/system.go).
So we create a new file locally at `systems/movement.go`:

```go
package systems

import (
    ecs "github.com/andygeiss/ecs/core"
)

type movementSystem struct{}

func (a *movementSystem) Process(em ecs.EntityManager) (state int) {
    // This state simply tells the engine to stop after the first call.
    return ecs.StateEngineStop
}

func (a *movementSystem) Setup() {}

func (a *movementSystem) Teardown() {}

func NewMovementSystem() ecs.System {
    return &movementSystem{}
}
```

Now we can add the following lines to `main.go`:

```go
sm := ecs.NewSystemManager()
sm.Add(systems.NewMovementSystem()) // <--
de := ecs.NewDefaultEngine(em, sm)
```

If we start our program now, it returns immediately without looping forever.

### The player entity

A game engine usually processes different types of components that represent
information about the game world itself. A component only represents the data,
and the systems are there to implement the behavior or game logic and change
these components. Entities are simply a composition of components that provide
a scalable data-oriented architecture.

A component needs to implement the methods defined by the interface
[Component](https://github.com/andygeiss/ecs/blob/master/core/entity.go).
Let's define our `Player` components by first creating a mask at
`components/components.go`:

```go
package components

const (
    MaskPosition = uint64(1 << 0)
    MaskVelocity = uint64(1 << 1)
)
```

Then create a component for `Position` and `Velocity` by creating
corresponding files such as `components/position.go`:

```go
package components

type Position struct {
    X  float32 `json:"x"`
    Y  float32 `json:"y"`
}

func (a *Position) Mask() uint64 {
    return MaskPosition
}

func (a *Position) WithX(x float32) *Position {
    a.X = x
    return a
}

func (a *Position) WithY(y float32) *Position {
    a.Y = y
    return a
}

func NewPosition() *Position {
    return &Position{}
}
```

Now we can add the following lines to `main.go`:

```go
em := ecs.NewEntityManager()
em.Add(ecs.NewEntity("player", []core.Component{ // <--
    components.NewPosition().
    WithX(10).
    WithY(10),
components.NewVelocity().
    WithX(100).
    WithY(100),
})) // -->
```

### Extend the movement system

Our final step is to add behavior to our movement system:

```go
func (a *movementSystem) Process(em ecs.EntityManager) (state int) {
    for _, e := range em.FilterByMask(components.MaskPosition | components.MaskVelocity) {
        position := e.Get(components.MaskPosition).(*components.Position)
        velocity := e.Get(components.MaskVelocity).(*components.Velocity)
        position.X += velocity.X * rl.GetFrameTime()
        position.Y += velocity.Y * rl.GetFrameTime()
    }
    return ecs.StateEngineStop
}
```

The movement system now moves every entity which has a position and velocity component.

We can replace `ecs.StateEngineStop` with `ecs.StateEngineContinue` later if we add
another system to handle user input.

A rendering system is also essential for a game, so you can use game libraries
such as [Raylib](https://pkg.go.dev/github.com/gen2brain/raylib-go/raylib) or
[SDL2](https://pkg.go.dev/github.com/veandco/go-sdl2).
This system could look like this with raylib:

```go
// ...
func (a *renderingSystem) Setup() {
    rl.InitWindow(a.width, a.height, a.title)
}

func (a *renderingSystem) Process(em core.EntityManager) (state int) {
    // First check if app should stop.
    if rl.WindowShouldClose() {
        return core.StateEngineStop
    }
    // Clear the screen
    if rl.IsWindowReady() {
        rl.BeginDrawing()
        rl.ClearBackground(rl.Black)
        rl.DrawFPS(10, 10)
        rl.EndDrawing()
    }
    return core.StateEngineContinue
}

func (a *renderingSystem) Teardown() {
    rl.CloseWindow()
}
```
