package app_test

import (
	"github.com/andygeiss/ecs/example/app"
	"github.com/andygeiss/utils/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	cfg := app.Config{NumberOfEntities: 1000, Width: 1024, Height: 768}
	out := app.Generate(&cfg)
	assert.That("entitiy count should be 1000", t, len(out), 1000)
}
