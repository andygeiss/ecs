package components

// Texture contains the filename of the current texture.
type Texture struct {
	Filename  string
	IsEnabled bool
}

// Name ...
func (t *Texture) Name() string {
	return "texture"
}
