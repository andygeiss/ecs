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
- [Usage](README.md#usage)

## Goals

- Provide an **easy-to-use** framework to build a game engine from scratch.
- **No dependencies** to other modules or specific game libraries - Feel free to use what fits your needs.
- **Minimum overhead** - use only what is really needed.
- **Plugins** to offer unlimited room for improvements.
- **Interoperability** between non-Go libraries and Go via [Main/Do](https://github.com/andygeiss/ecs/blob/master/run.go). 

## Installation

**From Source**

    go get -u github.com/andygeiss/ecs

## Steps to start

We start our journey 
