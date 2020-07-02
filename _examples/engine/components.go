package engine

/*
  ____                                             _
 / ___|___  _ __ ___  _ __   ___  _ __   ___ _ __ | |_ ___
| |   / _ \| '_ ` _ \| '_ \ / _ \| '_ \ / _ \ '_ \| __/ __|
| |__| (_) | | | | | | |_) | (_) | | | |  __/ | | | |_\__ \
 \____\___/|_| |_| |_| .__/ \___/|_| |_|\___|_| |_|\__|___/
                     |_|
*/

const (
	MaskPosition = uint64(1 << 0)
	MaskRenderer = uint64(1 << 1)
	MaskSize     = uint64(1 << 2)
	MaskVelocity = uint64(1 << 3)
	MaskWindow   = uint64(1 << 4)
)

// Position ...
type Position struct {
	ID string `json:"id"`
	X  int32  `json:"x"`
	Y  int32  `json:"y"`
}

func (p *Position) Mask() uint64 {
	return MaskPosition
}

// NewPosition ...
func NewPosition(x, y int32) *Position {
	return &Position{
		ID: "position",
		X:  x,
		Y:  y,
	}
}

// Renderer ...
type Renderer struct {
	ID   string `json:"id"`
	Addr interface{}
}

func (m *Renderer) Mask() uint64 {
	return MaskWindow
}

// NewRenderer ...
func NewRenderer(addr interface{}) *Renderer {
	return &Renderer{
		ID: "renderer",
	}
}

// Size ...
type Size struct {
	ID     string `json:"id"`
	Height int32  `json:"height"`
	Width  int32  `json:"width"`
}

func (s *Size) Mask() uint64 {
	return MaskSize
}

// NewSize ...
func NewSize(width, height int32) *Size {
	return &Size{
		ID:     "size",
		Width:  width,
		Height: height,
	}
}

// Velocity ...
type Velocity struct {
	ID string `json:"id"`
	X  int32  `json:"x"`
	Y  int32  `json:"y"`
}

func (s *Velocity) Mask() uint64 {
	return MaskVelocity
}

// NewVelocity ...
func NewVelocity(x, y int32) *Velocity {
	return &Velocity{
		ID: "velocity",
		X:  x,
		Y:  y,
	}
}

// Window ...
type Window struct {
	ID   string `json:"id"`
	Addr interface{}
}

func (m *Window) Mask() uint64 {
	return MaskWindow
}

// NewWindow ...
func NewWindow(addr interface{}) *Window {
	return &Window{
		ID: "window",
	}
}
