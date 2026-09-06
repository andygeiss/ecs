<p align="center">
<img src="https://github.com/andygeiss/ecs/blob/master/logo.png?raw=true" />
</p>

# ECS - Entity Component System

[![License](https://img.shields.io/github/license/andygeiss/ecs)](https://github.com/andygeiss/ecs/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/v/release/andygeiss/ecs)](https://github.com/andygeiss/ecs/releases)

Build a game engine in Go out of entities, components and systems. It is for people who
want that core and nothing else: no rendering, no input, no physics, and no dependencies
— you bring the game library you like.

- **Data and behaviour stay apart.** A component is data. A system is behaviour. An
  entity is just an id and a bag of components.
- **Nothing else is in the box.** The module has zero dependencies, so it adds none to
  your game.
- **Small enough to read in one sitting.** About 200 lines of code.

## Install

```bash
go get github.com/andygeiss/ecs
```

## 30 seconds

```go
package main

import (
	"context"
	"fmt"

	"github.com/andygeiss/ecs"
)

const maskPosition = uint64(1 << 0)

type position struct{ X, Y int }

func (p *position) Mask() uint64 { return maskPosition }

// stepSystem moves every position by one, three times, then stops the engine.
type stepSystem struct{ steps int }

func (s *stepSystem) Setup() { s.steps = 3 }

func (s *stepSystem) Process(em ecs.EntityManager) (state int) {
	for _, e := range em.FilterByMask(maskPosition) {
		e.Get(maskPosition).(*position).X++
	}
	s.steps--
	if s.steps == 0 {
		return ecs.StateEngineStop
	}
	return ecs.StateEngineContinue
}

func (s *stepSystem) Teardown() {}

func main() {
	em := ecs.NewEntityManager()
	em.Add(ecs.NewEntity("player", []ecs.Component{&position{}}))

	sm := ecs.NewSystemManager()
	sm.Add(&stepSystem{})

	engine := ecs.NewDefaultEngine(em, sm)
	engine.Setup()
	defer engine.Teardown()
	engine.Run(context.Background())

	fmt.Println(em.Get("player").Get(maskPosition).(*position).X) // 3
}
```

### Example engine

See [engine-example](https://github.com/andygeiss/engine-example) for a real
implementation using [raylib](https://www.raylib.com).

## Walkthrough

### Project layout

At first we create a basic project layout:

```bash
mkdir ecs-example
cd ecs-example
go mod init example
mkdir components systems
```

Next we create a `main.go` with the following content:

```go
package main

import (
    "context"
    "os"
    "os/signal"

    "github.com/andygeiss/ecs"
)

func main() {
    // Ctrl-C cancels the context, and the engine stops after the current pass.
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
    defer stop()

    em := ecs.NewEntityManager()
    sm := ecs.NewSystemManager()
    de := ecs.NewDefaultEngine(em, sm)
    de.Setup()
    defer de.Teardown()
    de.Run(ctx)
}
```

The program loops until you press Ctrl-C, because no system tells it to stop yet.
It spins a core flat out while it does: `Run` never sleeps, so pacing the loop is a
system's job — a renderer waiting for vsync, say.

### The movement system

A system needs to implement the methods defined by the interface
[System](https://github.com/andygeiss/ecs/blob/master/system.go).
So we create a new file locally at `systems/movement.go`:

```go
package systems

import (
    "github.com/andygeiss/ecs"
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
[Component](https://github.com/andygeiss/ecs/blob/master/component.go).
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
em.Add(ecs.NewEntity("player", []ecs.Component{ // <--
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
such as [raylib](https://www.raylib.com) or
[SDL](https://github.com/libsdl-org/SDL).
This system could look like this with raylib:

```go
// ...
func (a *renderingSystem) Setup() {
    rl.InitWindow(a.width, a.height, a.title)
}

func (a *renderingSystem) Process(em ecs.EntityManager) (state int) {
    // First check if app should stop.
    if rl.WindowShouldClose() {
        return ecs.StateEngineStop
    }
    // Clear the screen
    if rl.IsWindowReady() {
        rl.BeginDrawing()
        rl.ClearBackground(rl.Black)
        rl.DrawFPS(10, 10)
        rl.EndDrawing()
    }
    return ecs.StateEngineContinue
}

func (a *renderingSystem) Teardown() {
    rl.CloseWindow()
}
```

## Upgrading to v0.4.0

Two changes, and only the first one touches your code.

1. **`Run` takes a `context.Context`.** `de.Run()` becomes `de.Run(ctx)`, and the engine
   now stops when that context is cancelled — checked once per pass over the systems.
   Pass `context.Background()` if you do not want that. `System.Process` is unchanged, so
   your systems stay as they are.
2. **The constructors return structs rather than interfaces**: `*DefaultEngine`,
   `*DefaultEntityManager`, `*DefaultSystemManager`. Existing call sites keep compiling —
   the interfaces `Engine`, `EntityManager` and `SystemManager` are still there.

One bug went with them: `FilterByNames` counted a name once per matching component, so an
entity carrying two components with the same name matched a name it did not carry at all.

## Working on this repository

```bash
make        # every gate against the working tree; run before a commit
make ci     # the same gates against the commit; run before a push
make fmt    # apply goimports and go fix
make test   # the inner loop
```

Built to the [engineering baseline](https://github.com/andygeiss/baseline) — its
`checklists/library.md` is the definition of done here. No rule is waived.
