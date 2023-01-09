package app

import (
	"github.com/andygeiss/ecs/core"
	"github.com/andygeiss/ecs/example/plugins"
	"github.com/andygeiss/ecs/example/systems"
)

// Entrypoint ...
func Entrypoint(cfg *Config) {
	// set up entities
	em := core.NewEntityManager()
	em.Add(Generate(cfg)...)
	// set up the systems
	sm := core.NewSystemManager()
	sm.Add(
		systems.NewMovement(em),
		systems.NewCollision().(*systems.Collision).
			WithWidth(cfg.Width).
			WithHeight(cfg.Height),
		systems.NewRendering().(*systems.Rendering).
			WithWidth(cfg.Width).
			WithHeight(cfg.Height).
			WithTitle(cfg.Title).
			WithPlugins(plugins.ShowEngineStats()),
	)
	// set up the engine
	e := core.NewDefaultEngine(em, sm)
	e.Setup()
	defer e.Teardown()
	e.Run()
}
