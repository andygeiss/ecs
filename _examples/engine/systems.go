package engine

import (
	"github.com/andygeiss/ecs"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

/*
 ____            _
/ ___| _   _ ___| |_ ___ _ __ ___  ___
\___ \| | | / __| __/ _ \ '_ ` _ \/ __|
 ___) | |_| \__ \ ||  __/ | | | | \__ \
|____/ \__, |___/\__\___|_| |_| |_|___/
       |___/
*/

// Collision ...
type Collision struct {
	height int32
	width  int32
}

func (m *Collision) Setup() {}

func (m *Collision) Process(em *ecs.EntityManager) (state int) {
	for _, entity := range em.FilterByMask(MaskPosition | MaskVelocity) {
		position := entity.Get(MaskPosition).(*Position)
		velocity := entity.Get(MaskVelocity).(*Velocity)
		if position.X >= m.width || position.X <= 0 {
			velocity.X = -velocity.X
		}
		if position.Y >= m.height || position.Y <= 0 {
			velocity.Y = -velocity.Y
		}
	}
	return ecs.StateEngineContinue
}

// NewCollision ...
func NewCollision(width, height int32) ecs.System {
	return &Collision{
		height: height,
		width:  width,
	}
}

func (m *Collision) Teardown() {}

// Movement ...
type Movement struct{}

func (m *Movement) Setup() {}

func (m *Movement) Process(em *ecs.EntityManager) (state int) {
	for _, entity := range em.FilterByMask(MaskPosition | MaskVelocity) {
		position := entity.Get(MaskPosition).(*Position)
		velocity := entity.Get(MaskVelocity).(*Velocity)
		position.X += velocity.X
		position.Y += velocity.Y
	}
	return ecs.StateEngineContinue
}

func (m *Movement) Teardown() {

}

// NewMovement ...
func NewMovement() ecs.System {
	return &Movement{}
}

// Rendering ...
type Rendering struct {
	renderer      *sdl.Renderer
	plugins       []ecs.Plugin
	title         string
	window        *sdl.Window
	width, height int32
}

func (r *Rendering) Setup() {
	ecs.Do(func() {
		r.window, _ = sdl.CreateWindow(r.title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, r.width, r.height, sdl.WINDOW_SHOWN)
		r.renderer, _ = sdl.CreateRenderer(r.window, -1, sdl.RENDERER_SOFTWARE)
	})
}

func (r *Rendering) Process(em *ecs.EntityManager) (state int) {
	// First check if engine should stop.
	shouldStop := false
	ecs.Do(func() {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				shouldStop = true
			}
		}
	})
	if shouldStop {
		return ecs.StateEngineStop
	}
	// Clear the screen
	ecs.Do(func() {
		_ = r.renderer.Clear()
		_ = r.renderer.SetDrawColor(0, 0, 0, 0x20)
		_ = r.renderer.FillRect(&sdl.Rect{0, 0, r.width, r.height})
	})
	// Render entities
	ecs.Do(func() {
		_ = r.renderer.SetDrawColor(255, 255, 255, 255)
		for _, e := range em.FilterByMask(MaskPosition | MaskSize) {
			position := e.Get(MaskPosition).(*Position)
			size := e.Get(MaskSize).(*Size)
			_ = r.renderer.FillRect(&sdl.Rect{X: position.X, Y: position.Y, W: size.Width, H: size.Height})
		}
	})
	// Set FPS to 60
	ecs.Do(func() {
		r.renderer.Present()
		time.Sleep(time.Millisecond * 1000 / 60)
	})
	// Save the window and renderer address for further usage in a plugin context.
	for _, e := range em.FilterByMask(MaskRenderer) {
		renderer := e.Get(MaskRenderer).(*Renderer)
		renderer.Addr = r.renderer
	}
	for _, e := range em.FilterByMask(MaskWindow) {
		window := e.Get(MaskWindow).(*Window)
		window.Addr = r.window
	}
	// Dispatch work to plugins.
	for _, plugin := range r.plugins {
		plugin(em)
	}
	return ecs.StateEngineContinue
}

func (r *Rendering) Teardown() {
	_ = r.window.Destroy()
	_ = r.renderer.Destroy()
	sdl.Quit()
}

// NewRendering ...
func NewRendering(width, height int32, title string, plugins ...ecs.Plugin) ecs.System {
	return &Rendering{
		height:  height,
		plugins: plugins,
		title:   title,
		width:   width,
	}
}
