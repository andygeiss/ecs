package components

import "github.com/andygeiss/ecs/core"

// Size ...
type Size struct {
	ID     string  `json:"id"`
	Height float32 `json:"height"`
	Width  float32 `json:"width"`
}

func (a *Size) Mask() uint64 {
	return MaskSize
}

func (a *Size) WithHeight(height float32) *Size {
	a.Height = height
	return a
}

func (a *Size) WithWidth(width float32) *Size {
	a.Width = width
	return a
}

// NewSize ...
func NewSize() core.Component {
	return &Size{
		ID: "size",
	}
}
