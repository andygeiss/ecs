package components

// Sound contains the filename of the sound which is currently playing (empty = nothing).
type Sound struct {
	EventFilename map[string]string
	Filename      string
}

// Name ...
func (s *Sound) Name() string {
	return "sound"
}
