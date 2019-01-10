package components

// Animation contains the filename of the animation which is currently active (empty = nothing).
type Animation struct {
	Count         float32
	EventFilename map[string]string
	Filename      string
	Height        float32
	Index         float32
	Width         float32
}

// Name ...
func (a *Animation) Name() string {
	return "animation"
}
