package systems

import (
	"github.com/andygeiss/ecs/core"
	"github.com/andygeiss/ecs/example/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Rendering ...
type Rendering struct {
	err           error
	title         string
	width, height int32
	plugins       []core.Plugin
}

func (a *Rendering) Error() error {
	return a.err
}

func (a *Rendering) Setup() {
	rl.InitWindow(a.width, a.height, a.title)
}

func (a *Rendering) Process(em core.EntityManager) (state int) {
	// First check if app should stop.
	if rl.WindowShouldClose() {
		return core.StateEngineStop
	}
	// Clear the screen
	if rl.IsWindowReady() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		a.renderEntities(em)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	// Dispatch work to plugins.
	for _, plugin := range a.plugins {
		plugin(em)
	}
	return core.StateEngineContinue
}

func (a *Rendering) Teardown() {
	rl.CloseWindow()
}

func (a *Rendering) WithHeight(height int) *Rendering {
	a.height = int32(height)
	return a
}

func (a *Rendering) WithPlugins(plugins ...core.Plugin) *Rendering {
	a.plugins = plugins
	return a
}

func (a *Rendering) WithTitle(title string) *Rendering {
	a.title = title
	return a
}

func (a *Rendering) WithWidth(width int) *Rendering {
	a.width = int32(width)
	return a
}

func (a *Rendering) renderEntities(em core.EntityManager) {
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskSize) {
		position := e.Get(components.MaskPosition).(*components.Position)
		size := e.Get(components.MaskSize).(*components.Size)
		rl.DrawRectangleRec(rl.Rectangle{X: position.X, Y: position.Y, Width: size.Width, Height: size.Height}, rl.Blue)
	}
}

// NewRendering ...
func NewRendering() core.System {
	return &Rendering{}
}
