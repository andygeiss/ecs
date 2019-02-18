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

See a real-world example at [ecs-pong](https://github.com/andygeiss/ecs-pong/tree/master/internal/app).
