package components

// Size contains the width and height of an entity.
type Size struct {
	Width  float32
	Height float32
}

// Name ...
func (s *Size) Name() string {
	return "size"
}
