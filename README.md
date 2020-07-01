# ECS - Entity Component System

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/ecs)](https://goreportcard.com/report/github.com/andygeiss/ecs)
[![BCH compliance](https://bettercodehub.com/edge/badge/andygeiss/ecs?branch=master)](https://bettercodehub.com/)

**Build scalable solutions based on the Entity Component System concept in Golang**

The architectural pattern of an ECS is mostly used in game development,
to provide long-term maintainability and extendability of large, dynamic systems.

[![Overview](ecs.svg)](ecs.svg)

An **Entity** is basically a composition of different components and has an ID.  
A **Component** contains only the state or data of one specific aspect like health, position, velocity etc.  
A **System** handles the behaviour or logic of the components. A movement system uses the position and velocity to implement an entities movement. 

**Table of Contents**

- [Goals](README.md#goals)
- [Installation](README.md#installation)
- [Usage](README.md#usage)

## Goals

- Provide an **easy-to-use module** to create a ECS architecture.
- **No dependencies** to other modules.
- **Minimum overhead** - use only what is really needed.
- **Plugins** to offer unlimited room for improvements.
- **Interoperability** between Cgo and Go via [Prepare](https://github.com/andygeiss/ecs/blob/master/run.go) 

## Installation

**From Source**

    go get -u github.com/andygeiss/ecs

## Usage

An example scenario that distributes the load across multiple processor cores could look like this:

```go
func main() {
    // Create a new EntityManager
    em := ecs.NewEntityManager()
    em.Add(
        ecs.NewEntity("background", []ecs.Component{
            components.NewSize(width, height),
        }))
    em.Add(en...)
    
    // Create a new SystemManager which should only handle collisions.
    systems1 := ecs.NewSystemManager()
    systems1.Add(
        systems.NewCollision(),
    )
    
    // Create another SystemManager which should handle movement.
    systems2 := ecs.NewSystemManager()
    systems2.Add(
        systems.NewMovement(
            topDown.TranslateInputs2Velocity(),
            topDown.MoveEntities(),
        ),
    )
    
    // Create a SystemManager which should handle Cgo calls.
    systemsCgo := ecs.NewSystemManager()
    systemsCgo.Add(
        systems.NewCamera(
            topDown.UpdateCamera(width, height, 4.0),
        ),
        systems.NewInput(
            topDown.ReadFromKeyboard(),
        ),
        systems.NewTexture(
            topDown.LoadTextures(filepath.Join("assets", "sprites")),
        ),
        systems.NewRendering(width, height, "Demo",
            topDown.RenderEntities(width, height),
        ),
    )
    
    ecs.Prepare(func() {
        // systems1 will be handled by CPU 1
        go ecs.Run(em, systems1)
        // systems1 will be handled by CPU 2
        go ecs.Run(em, systems2)
        // systemsCgo will be handled by CPU 3 (Cgo calls needs to be locked).
        ecs.RunAsMain(func() {
            ecs.Run(em, systemsCgo)
        })
    })
}
```


See a real-world example of ECS in action at [ecs-pong](https://github.com/andygeiss/ecs-pong).

![ecs-pong](https://github.com/andygeiss/ecs-pong/blob/master/assets/pong.png)

## Benchmarks

Run the benchmarks on your current machine by using the following command:

    go test -bench ./...

**Important Notice**

The version **v0.0.51** got a major **Performance Upgrade** by replacing looped <code>append</code> calls to a single <code>make</code> to allocate the entity slice only once.

Performance before:
    
    BenchmarkEngine_Run/1_system(s)_with_10000_entities-4       	    3384	    334659 ns/op
    BenchmarkEngine_Run/2_system(s)_with_10000_entities-4       	    1687	    667612 ns/op
    BenchmarkEngine_Run/4_system(s)_with_10000_entities-4       	     868	   1355581 ns/op   
    
Performance after:

    BenchmarkEngine_Run/1_system(s)_with_10000_entities-4               7812            143810 ns/op
    BenchmarkEngine_Run/2_system(s)_with_10000_entities-4               4216            281840 ns/op
    BenchmarkEngine_Run/4_system(s)_with_10000_entities-4               2044            559839 ns/op

**UPDATE**

The version **v0.0.54** got another **Performance Upgrade** by replacing the String-compare <code>if c.Name() == name</code> with a bitmask.
Thus, the for-loop checking for each component's name could be removed, too.

Performance after:

    BenchmarkEngine_Run/1_system(s)_with_10000_entities-4          	   30087	     38109 ns/op
    BenchmarkEngine_Run/2_system(s)_with_10000_entities-4          	   15320	     76884 ns/op
    BenchmarkEngine_Run/4_system(s)_with_10000_entities-4          	    7609	    154035 ns/op

We finally end up running one loop through all the systems and entities **9 times faster** than **v0.0.50** ;-)
