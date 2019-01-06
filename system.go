package ecs

import "strings"

// System ...
type System interface {
	Setup()
	Process(entityManager *EntityManager)
	Teardown()
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
