package components

// Position contains the 2D X and Y coordinate.
type Position struct {
	X float32
	Y float32
}

// Name ...
func (p *Position) Name() string {
	return "position"
}
