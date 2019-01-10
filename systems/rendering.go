package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs/components"
	"github.com/gen2brain/raylib-go/raylib"
)

type rendering struct {
	background   rl.Color
	images       map[string]*rl.Image
	textures     map[string]rl.Texture2D
	title        string
	windowHeight int32
	windowWidth  int32
}

// NewRendering ...
func NewRendering(width, height int32, title string, background rl.Color) ecs.System {
	return &rendering{
		background:   background,
		images:       map[string]*rl.Image{},
		textures:     map[string]rl.Texture2D{},
		title:        title,
		windowHeight: height,
		windowWidth:  width,
	}
}

// Process ...
func (s *rendering) Process(entityManages *ecs.EntityManager) {
	if rl.WindowShouldClose() {
		ecs.ShouldEngineStop = true
		return
	}
	rl.BeginDrawing()
	rl.ClearBackground(s.background)
	for _, e := range entityManages.FilterBy("position", "size") {
		isAnimationPresent := s.renderAnimationIfPresent(e)
		isTexturePresent := s.renderTextureIfPresent(e)
		// Render a colored rectangle if no texture and animation is present.
		if !isAnimationPresent && !isTexturePresent {
			s.renderBoundingBox(e)
		}
	}
	rl.EndDrawing()
}

// Setup ...
func (s *rendering) Setup() {
	rl.InitWindow(s.windowWidth, s.windowHeight, s.title)
	rl.SetTargetFPS(60)
}

// Teardown ...
func (s *rendering) Teardown() {
	for _, img := range s.images {
		rl.UnloadImage(img)
	}
	for _, tx := range s.textures {
		rl.UnloadTexture(tx)
	}
	rl.CloseWindow()
}

func (s *rendering) renderAnimationIfPresent(entity *ecs.Entity) (present bool) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	// Return if animation is not present.
	anim := entity.Get("animation")
	if anim == nil {
		return false
	}
	// Get texture from cache or load it from the filesystem into the cache.
	animation := anim.(*components.Animation)
	filename := animation.Filename
	tx, exists := s.textures[filename]
	if !exists {
		img := rl.LoadImage(filename)
		rl.ImageResize(img, int32(animation.Width), int32(animation.Height))
		s.images[filename] = img
		s.textures[filename] = rl.LoadTextureFromImage(img)
		tx = s.textures[filename]
	}
	rl.DrawTextureRec(
		tx,
		rl.NewRectangle(animation.Index*size.Width, 0, size.Width, size.Height),
		rl.NewVector2(position.X, position.Y),
		rl.RayWhite,
	)
	return true
}

func (s *rendering) renderBoundingBox(entity *ecs.Entity) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	rl.DrawRectangleLines(
		int32(position.X),
		int32(position.Y),
		int32(size.Width),
		int32(size.Height),
		rl.RayWhite,
	)
}

func (s *rendering) renderTextureIfPresent(entity *ecs.Entity) (present bool) {
	position := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	// Return if texture is not present.
	texture := entity.Get("texture")
	if texture == nil {
		return false
	}
	// Get texture from cache or load it from the filesystem into the cache.
	fileName := texture.(*components.Texture).Filename
	tx, exists := s.textures[fileName]
	if !exists {
		s.textures[fileName] = rl.LoadTexture(fileName)
		tx = s.textures[fileName]
	}
	rl.DrawTextureRec(
		tx,
		rl.NewRectangle(0, 0, size.Width, size.Height),
		rl.NewVector2(position.X, position.Y),
		rl.RayWhite,
	)
	return true
}