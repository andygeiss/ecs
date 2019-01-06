package ecs

import "strings"

// System ...
type System interface {
	Setup()
	Process(entityManager *EntityManager)
	Teardown()
}

// SystemManager ...
type SystemManager struct {
	systems []System
}

// NewSystemManager ...
func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems: []System{},
	}
}

// Add ...
func (m *SystemManager) Add(systems ...System) {
	for _, system := range systems {
		m.systems = append(m.systems, system)
	}
}

// Systems ...
func (m *SystemManager) Systems() []System {
	return m.systems
}

// Filter ...
func Filter(entity Entity, names ...string) (filtered map[string]Component, isComplete bool) {
	filtered = map[string]Component{}
	pattern := strings.Join(names, " ")
	for _, component := range entity.Components() {
		name := component.Name()
		if strings.Contains(pattern, name) {
			filtered[name] = component
		}
	}
	isComplete = false
	if len(filtered) == len(names) {
		isComplete = true
	}
	return
}
