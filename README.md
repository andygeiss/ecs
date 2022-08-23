# ECS - Entity Component System

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/ecs)](https://goreportcard.com/report/github.com/andygeiss/ecs)
[![BCH compliance](https://bettercodehub.com/edge/badge/andygeiss/ecs?branch=master)](https://bettercodehub.com/)

**Build your own Game-Engine based on the Entity Component System concept in Golang**

The architectural pattern of an ECS is mostly used in game development,
to provide long-term maintainability and extendability of large, dynamic systems.

[![Overview](ecs.svg)](ecs.svg)

An **Entity** is basically a composition of different components and has an ID.  
A **Component** contains only the state or data of one specific aspect like health, position, velocity etc.  
A **System** handles the behaviour or logic of the components. A movement system uses the position and velocity to implement an entities movement. 

**Table of Contents**

- [Goals](README.md#goals)
- [Installation](README.md#installation)
- [Steps to start](README.md#steps-to-start)

## Goals

- Provide an **easy-to-use** framework to build a game engine from scratch.
- **No dependencies** to other modules or specific game libraries - Feel free to use what fits your needs.
- **Minimum overhead** - use only what is really needed.
- **Plugins** to offer unlimited room for improvements.

## Installation

**From Source**

    go get -u github.com/andygeiss/ecs

## Steps to start

In the first step we have to be clear about what our game engine should do.
The main task is to make sure that we have all the essential components that are necessary for the technical and logical aspects 
are responsible, are combined with each other.

An Entity Component System (ECS) helps us to do just that, as the logical components (data) such as entities and their components
can be separated from the actual logic. 
One of the advantages of this is that we can implement and test the game mechanics independently of the rest.

So let's start...

We decide to use 2D and define the **three most important components**:
* Position
* Size
* Velocity

We store these as [components.go](https://github.com/andygeiss/ecs-example/blob/main/engine/components.go).

In the next step, the **three most important systems** implement
* Collision
* Movement
* Rendering

We store these as [systems.go](https://github.com/andygeiss/ecs-example/blob/main/engine/systems.go).

The rendering system uses a specific game library like [Raylib](https://www.raylib.com/index.html).

Finally we create [main.go](https://github.com/andygeiss/ecs-example/blob/main/main.go) and link the systems together:

## I want more than 64 Components !

Yes! You can do that by adding a `Name()` function to your component and using `FilterByNames` instead of `FilterByMask`.
However if you want the fastest possible solution you need to know that `FilterByNames` is `40` times slower than `FilterByMask`.
This will maybe not impact your Game if there are only a few thousands of entities. ;-)

![bench](bench.png)
