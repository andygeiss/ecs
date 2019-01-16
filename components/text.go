package components

import "github.com/gen2brain/raylib-go/raylib"

const (
	TextAlignBottom = 0
	TextAlignCenter = 1
	TextAlignLeft   = 2
	TextAlignRight  = 3
	TextAlignTop    = 4
)

// Text ...
type Text struct {
	Align     int
	Color     rl.Color
	Content   string
	FontSize  int32
	IsEnabled bool
}

// Name ...
func (t *Text) Name() string {
	return "text"
}
