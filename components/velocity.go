package components

// Velocity contains the speed in pixels on the X and Y axis.
type Velocity struct {
	X float32
	Y float32
}

// Name ...
func (v *Velocity) Name() string {
	return "velocity"
}
