package main

import (
	"github.com/andygeiss/ecs/example/app"
)

func main() {
	cfg := &app.Config{
		NumberOfEntities: 200000,
		Width:            1366,
		Height:           768,
		Title:            "ECS Example Benchmark",
	}
	app.Entrypoint(cfg)
}
