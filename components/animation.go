package components

const (
	AnimationSpeedLow    = 2
	AnimationSpeedMedium = 4
	AnimationSpeedHigh   = 8
)

// Animation contains the filename of the animation which is currently active (empty = nothing).
type Animation struct {
	EventFilename map[string]string
	Filename      string
	Height        float32
	IsEnabled     bool
	Speed         int64
	SpriteCount   float32
	SpriteIndex   float32
	Width         float32
}

// Name ...
func (a *Animation) Name() string {
	return "animation"
}
