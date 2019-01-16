package components

// Sound contains the filename of the sound which is currently playing (empty = nothing).
type Sound struct {
	EventFilename map[string]string
	Filename      string
	IsEnabled     bool
	Volume        float32
}

// Name ...
func (s *Sound) Name() string {
	return "sound"
}
